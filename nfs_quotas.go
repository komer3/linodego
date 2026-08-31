package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
)

type NFSQuotaStatus string

const (
	NFSQuotaStatusActive   NFSQuotaStatus = "active"
	NFSQuotaStatusUpdating NFSQuotaStatus = "updating"
	NFSQuotaStatusError    NFSQuotaStatus = "error"
)

type NFSQuotaRuleIdentifierType string

const (
	NFSQuotaRuleIdentifierTypeUID       NFSQuotaRuleIdentifierType = "uid"
	NFSQuotaRuleIdentifierTypeUsername  NFSQuotaRuleIdentifierType = "username"
	NFSQuotaRuleIdentifierTypeGID       NFSQuotaRuleIdentifierType = "gid"
	NFSQuotaRuleIdentifierTypeGroupname NFSQuotaRuleIdentifierType = "groupname"
)

type NFSQuota struct {
	FilesystemID      int            `json:"filesystem_id"`
	UsedCapacityBytes int64          `json:"used_capacity_bytes"`
	PathLimits        []NFSPathLimit `json:"path_limits"`
	Status            NFSQuotaStatus `json:"status"`
	Created           *time.Time     `json:"-"`
	Updated           *time.Time     `json:"-"`
}

type NFSPathLimit struct {
	Path             string              `json:"path"`
	MaxCapacityBytes *int64              `json:"max_capacity_bytes"`
	MaxFileCount     *int64              `json:"max_file_count"`
	GracePeriod      *string             `json:"grace_period"`
	UserGroupConfig  *NFSUserGroupConfig `json:"user_group_config"`
}

type NFSUserGroupConfig struct {
	Enabled           bool                        `json:"enabled"`
	DefaultUserLimit  *NFSCapacityLimit           `json:"default_user_limit"`
	DefaultGroupLimit *NFSCapacityLimit           `json:"default_group_limit"`
	UserLimits        []NFSIdentifiedLimit        `json:"user_limits"`
	GroupLimits       []NFSIdentifiedLimit        `json:"group_limits"`
	UserGroupLimits   []NFSUserGroupCombinedLimit `json:"user_group_limits"`
}

type NFSCapacityLimit struct {
	MaxCapacityBytes *int64 `json:"max_capacity_bytes"`
	MaxFileCount     *int64 `json:"max_file_count"`
}

type NFSIdentifiedLimit struct {
	IdentifierType   NFSQuotaRuleIdentifierType `json:"identifier_type"`
	Identifier       string                     `json:"identifier"`
	MaxCapacityBytes *int64                     `json:"max_capacity_bytes"`
	MaxFileCount     *int64                     `json:"max_file_count"`
}

type NFSUserGroupCombinedLimit struct {
	User             NFSUserGroupLimitIdentity `json:"user"`
	Group            NFSUserGroupLimitIdentity `json:"group"`
	MaxCapacityBytes *int64                    `json:"max_capacity_bytes"`
	MaxFileCount     *int64                    `json:"max_file_count"`
}

type NFSUserGroupLimitIdentity struct {
	IdentifierType NFSQuotaRuleIdentifierType `json:"identifier_type"`
	Identifier     string                     `json:"identifier"`
}

type NFSQuotaUpdateOptions struct {
	PathLimits []NFSPathLimitUpdateOptions `json:"path_limits"`
}

type NFSPathLimitUpdateOptions struct {
	Path             string                            `json:"path"`
	MaxCapacityBytes **int64                           `json:"max_capacity_bytes,omitzero"`
	MaxFileCount     **int64                           `json:"max_file_count,omitzero"`
	UserGroupConfig  **NFSUserGroupConfigUpdateOptions `json:"user_group_config,omitzero"`
}

type NFSUserGroupConfigUpdateOptions struct {
	Enabled           bool                                      `json:"enabled"`
	DefaultUserLimit  **NFSCapacityLimitUpdateOptions           `json:"default_user_limit,omitzero"`
	DefaultGroupLimit **NFSCapacityLimitUpdateOptions           `json:"default_group_limit,omitzero"`
	UserLimits        *[]NFSIdentifiedLimitUpdateOptions        `json:"user_limits,omitzero"`
	GroupLimits       *[]NFSIdentifiedLimitUpdateOptions        `json:"group_limits,omitzero"`
	UserGroupLimits   *[]NFSUserGroupCombinedLimitUpdateOptions `json:"user_group_limits,omitzero"`
}

type NFSCapacityLimitUpdateOptions struct {
	MaxCapacityBytes **int64 `json:"max_capacity_bytes,omitzero"`
	MaxFileCount     **int64 `json:"max_file_count,omitzero"`
}

type NFSIdentifiedLimitUpdateOptions struct {
	IdentifierType   NFSQuotaRuleIdentifierType `json:"identifier_type"`
	Identifier       string                     `json:"identifier"`
	MaxCapacityBytes **int64                    `json:"max_capacity_bytes,omitzero"`
	MaxFileCount     **int64                    `json:"max_file_count,omitzero"`
}

type NFSUserGroupCombinedLimitUpdateOptions struct {
	User             NFSUserGroupLimitIdentity `json:"user"`
	Group            NFSUserGroupLimitIdentity `json:"group"`
	MaxCapacityBytes **int64                   `json:"max_capacity_bytes,omitzero"`
	MaxFileCount     **int64                   `json:"max_file_count,omitzero"`
}

func (n *NFSQuota) UnmarshalJSON(b []byte) error {
	type Mask NFSQuota

	p := struct {
		*Mask

		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Created = (*time.Time)(p.Created)
	n.Updated = (*time.Time)(p.Updated)

	return nil
}

func (c *Client) GetNFSQuota(ctx context.Context, spaceID int, filesystemID int) (*NFSQuota, error) {
	return doGETRequest[NFSQuota](ctx, c, formatAPIPath("nfs/spaces/%d/filesystems/%d/quota", spaceID, filesystemID))
}

func (c *Client) UpdateNFSQuota(
	ctx context.Context,
	spaceID int,
	filesystemID int,
	opts NFSQuotaUpdateOptions,
) (*NFSQuota, error) {
	return doPUTRequest[NFSQuota](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/quota", spaceID, filesystemID),
		opts,
	)
}
