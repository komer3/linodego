package linodego

import "context"

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
	FilesystemID      string         `json:"filesystem_id"`
	UsedCapacityBytes int64          `json:"used_capacity_bytes"`
	PathQuotas        []NFSPathQuota `json:"path_quotas"`
	Status            NFSQuotaStatus `json:"status"`
}

type NFSPathQuota struct {
	Path                   string              `json:"path"`
	Managed                bool                `json:"managed"`
	CapacitySoftLimitBytes *int64              `json:"capacity_soft_limit_bytes"`
	CapacityHardLimitBytes *int64              `json:"capacity_hard_limit_bytes"`
	FilesSoftLimit         *int64              `json:"files_soft_limit"`
	FilesHardLimit         *int64              `json:"files_hard_limit"`
	GracePeriod            *string             `json:"grace_period"`
	UserGroupQuotas        *NFSUserGroupQuotas `json:"user_group_quotas"`
}

type NFSUserGroupQuotas struct {
	Enabled          bool                     `json:"enabled"`
	DefaultUserRule  *NFSQuotaRule            `json:"default_user_rule"`
	DefaultGroupRule *NFSQuotaRule            `json:"default_group_rule"`
	UserRules        []NFSIdentifiedQuotaRule `json:"user_rules"`
	GroupRules       []NFSIdentifiedQuotaRule `json:"group_rules"`
}

type NFSQuotaRule struct {
	CapacitySoftLimitBytes *int64  `json:"capacity_soft_limit_bytes"`
	CapacityHardLimitBytes *int64  `json:"capacity_hard_limit_bytes"`
	FilesSoftLimit         *int64  `json:"files_soft_limit"`
	FilesHardLimit         *int64  `json:"files_hard_limit"`
	GracePeriod            *string `json:"grace_period"`
}

type NFSIdentifiedQuotaRule struct {
	IdentifierType         NFSQuotaRuleIdentifierType `json:"identifier_type"`
	Identifier             string                     `json:"identifier"`
	CapacitySoftLimitBytes *int64                     `json:"capacity_soft_limit_bytes"`
	CapacityHardLimitBytes *int64                     `json:"capacity_hard_limit_bytes"`
	FilesSoftLimit         *int64                     `json:"files_soft_limit"`
	FilesHardLimit         *int64                     `json:"files_hard_limit"`
	GracePeriod            *string                    `json:"grace_period"`
}

type NFSQuotaUpdateOptions struct {
	PathQuotas []NFSPathQuota `json:"path_quotas"`
}

func (c *Client) GetNFSQuota(ctx context.Context, spaceID string, filesystemID string) (*NFSQuota, error) {
	return doGETRequest[NFSQuota](ctx, c, formatAPIPath("nfs/spaces/%s/filesystems/%s/quota", spaceID, filesystemID))
}

func (c *Client) UpdateNFSQuota(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSQuotaUpdateOptions,
) (*NFSQuota, error) {
	return doPUTRequest[NFSQuota](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/quota", spaceID, filesystemID),
		opts,
	)
}
