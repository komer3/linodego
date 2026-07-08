package linodego

import (
	"context"
	"encoding/json"
	"time"
)

type NFSProtocolVersion string

const (
	NFSProtocolVersionV4 NFSProtocolVersion = "nfsv4"
)

type NFSFilesystemStatus string

const (
	NFSFilesystemStatusCreating NFSFilesystemStatus = "creating"
	NFSFilesystemStatusActive   NFSFilesystemStatus = "active"
	NFSFilesystemStatusUpdating NFSFilesystemStatus = "updating"
	NFSFilesystemStatusDeleting NFSFilesystemStatus = "deleting"
	NFSFilesystemStatusError    NFSFilesystemStatus = "error"
)

type NFSFilesystem struct {
	ID                 int                  `json:"id"`
	SpaceID            int                  `json:"space_id"`
	Label              string               `json:"label"`
	Region             string               `json:"region"`
	ProtocolVersions   []NFSProtocolVersion `json:"protocol_versions"`
	Status             NFSFilesystemStatus  `json:"status"`
	MountTargetIPs     []string             `json:"mount_target_ips"`
	MountTargetFQDN    *string              `json:"mount_target_fqdn"`
	SnapshotUsageBytes *int64               `json:"snapshot_usage_bytes"`
	LDAPConfigID       *string              `json:"ldap_config_id"`
	SourceSnapshotID   *int                 `json:"source_snapshot_id"`
	Created            *time.Time           `json:"-"`
	Updated            *time.Time           `json:"-"`
	Tags               []string             `json:"tags"`
	Stats              NFSFilesystemStats   `json:"stats"`
}

type NFSFilesystemCreateOptions struct {
	Label            string                `json:"label"`
	Region           string                `json:"region"`
	ProtocolVersions *[]NFSProtocolVersion `json:"protocol_versions,omitzero"`
	Tags             *[]string             `json:"tags,omitzero"`
}

type NFSFilesystemUpdateOptions struct {
	Label *string   `json:"label,omitzero"`
	Tags  *[]string `json:"tags,omitzero"`
}

type NFSFilesystemStats struct {
	UsedCapacityBytes *int64     `json:"used_capacity_bytes"`
	MaxCapacityBytes  *int64     `json:"max_capacity_bytes"`
	CollectedAt       *time.Time `json:"-"`
}

type NFSFilesystemListQueryOptions struct {
	Label  string `query:"label"`
	Region string `query:"region"`
}

func (n *NFSFilesystem) UnmarshalJSON(b []byte) error {
	type Mask NFSFilesystem

	p := struct {
		*Mask

		Created *time.Time `json:"created"`
		Updated *time.Time `json:"updated"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Created = p.Created
	n.Updated = p.Updated

	return nil
}

func (n *NFSFilesystemStats) UnmarshalJSON(b []byte) error {
	type Mask NFSFilesystemStats

	p := struct {
		*Mask

		CollectedAt *time.Time `json:"collected_at"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.CollectedAt = p.CollectedAt

	return nil
}

func (n NFSFilesystem) GetCreateOptions() NFSFilesystemCreateOptions {
	result := NFSFilesystemCreateOptions{
		Label:  n.Label,
		Region: n.Region,
	}

	if n.ProtocolVersions != nil {
		result.ProtocolVersions = Pointer(n.ProtocolVersions)
	}

	if n.Tags != nil {
		result.Tags = Pointer(n.Tags)
	}

	return result
}

func (n NFSFilesystem) GetUpdateOptions() NFSFilesystemUpdateOptions {
	result := NFSFilesystemUpdateOptions{
		Label: Pointer(n.Label),
	}

	if n.Tags != nil {
		result.Tags = Pointer(n.Tags)
	}

	return result
}

func (c *Client) ListNFSFilesystems(ctx context.Context, spaceID int, opts *ListOptions) ([]NFSFilesystem, error) {
	return getPaginatedResults[NFSFilesystem](ctx, c, formatAPIPath("nfs/spaces/%d/filesystems", spaceID), opts)
}

func (c *Client) GetNFSFilesystem(ctx context.Context, spaceID int, filesystemID int) (*NFSFilesystem, error) {
	return doGETRequest[NFSFilesystem](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d", spaceID, filesystemID),
	)
}

func (c *Client) GetNFSFilesystemByID(ctx context.Context, filesystemID int) (*NFSFilesystem, error) {
	return doGETRequest[NFSFilesystem](ctx, c, formatAPIPath("nfs/filesystems/%d", filesystemID))
}

func (c *Client) CreateNFSFilesystem(
	ctx context.Context,
	spaceID int,
	opts NFSFilesystemCreateOptions,
) (*NFSFilesystem, error) {
	return doPOSTRequest[NFSFilesystem](ctx, c, formatAPIPath("nfs/spaces/%d/filesystems", spaceID), opts)
}

func (c *Client) UpdateNFSFilesystem(
	ctx context.Context,
	spaceID int,
	filesystemID int,
	opts NFSFilesystemUpdateOptions,
) (*NFSFilesystem, error) {
	return doPUTRequest[NFSFilesystem](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSFilesystem(ctx context.Context, spaceID int, filesystemID int) error {
	return doDELETERequest(ctx, c, formatAPIPath("nfs/spaces/%d/filesystems/%d", spaceID, filesystemID))
}
