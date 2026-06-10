package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
)

type NFSProtocolVersion string

const (
	NFSProtocolVersionV4 NFSProtocolVersion = "nfsv4"
)

type NFSFilesystemStatus string

const (
	NFSFilesystemStatusCreating NFSFilesystemStatus = "creating"
	NFSFilesystemStatusActive   NFSFilesystemStatus = "active"
	NFSFilesystemStatusDeleting NFSFilesystemStatus = "deleting"
	NFSFilesystemStatusFailed   NFSFilesystemStatus = "failed"
)

type NFSFilesystem struct {
	ID                 string               `json:"id"`
	SpaceID            string               `json:"space_id"`
	Label              string               `json:"label"`
	Region             string               `json:"region"`
	ProtocolVersions   []NFSProtocolVersion `json:"protocol_versions"`
	Status             NFSFilesystemStatus  `json:"status"`
	MountTarget        string               `json:"mount_target"`
	SnapshotUsageBytes *int64               `json:"snapshot_usage_bytes"`
	LDAPConfigID       *string              `json:"ldap_config_id"`
	SourceSnapshotID   *string              `json:"source_snapshot_id"`
	Created            *time.Time           `json:"-"`
	Updated            *time.Time           `json:"-"`
	Tags               []string             `json:"tags"`
}

type NFSFilesystemCreateOptions struct {
	Label            string               `json:"label"`
	Region           string               `json:"region"`
	ProtocolVersions []NFSProtocolVersion `json:"protocol_versions,omitzero"`
	Tags             []string             `json:"tags,omitzero"`
}

type NFSFilesystemUpdateOptions struct {
	Label string   `json:"label,omitzero"`
	Tags  []string `json:"tags,omitzero"`
}

type NFSFilesystemStats struct {
	FilesystemID               string     `json:"filesystem_id"`
	ReadThroughputBytesPerSec  int64      `json:"read_throughput_bytes_per_sec"`
	WriteThroughputBytesPerSec int64      `json:"write_throughput_bytes_per_sec"`
	ReadIOPS                   int        `json:"read_iops"`
	WriteIOPS                  int        `json:"write_iops"`
	UsedCapacityBytes          int64      `json:"used_capacity_bytes"`
	FreeCapacityBytes          int64      `json:"free_capacity_bytes"`
	TotalInodes                int64      `json:"total_inodes"`
	AvailableInodes            int64      `json:"available_inodes"`
	CollectedAt                *time.Time `json:"-"`
}

type NFSFilesystemListQueryOptions struct {
	Label  string `query:"label"`
	Region string `query:"region"`
}

func (n *NFSFilesystem) UnmarshalJSON(b []byte) error {
	type Mask NFSFilesystem

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

func (n *NFSFilesystemStats) UnmarshalJSON(b []byte) error {
	type Mask NFSFilesystemStats

	p := struct {
		*Mask

		CollectedAt *parseabletime.ParseableTime `json:"collected_at"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.CollectedAt = (*time.Time)(p.CollectedAt)

	return nil
}

func (n NFSFilesystem) GetCreateOptions() NFSFilesystemCreateOptions {
	return NFSFilesystemCreateOptions{
		Label:            n.Label,
		Region:           n.Region,
		ProtocolVersions: n.ProtocolVersions,
		Tags:             n.Tags,
	}
}

func (n NFSFilesystem) GetUpdateOptions() NFSFilesystemUpdateOptions {
	return NFSFilesystemUpdateOptions{
		Label: n.Label,
		Tags:  n.Tags,
	}
}

func (c *Client) ListNFSFilesystems(ctx context.Context, spaceID string, opts *ListOptions) ([]NFSFilesystem, error) {
	return getPaginatedResults[NFSFilesystem](ctx, c, formatAPIPath("nfs/spaces/%s/filesystems", spaceID), opts)
}

func (c *Client) GetNFSFilesystem(ctx context.Context, spaceID string, filesystemID string) (*NFSFilesystem, error) {
	return doGETRequest[NFSFilesystem](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s", spaceID, filesystemID),
	)
}

func (c *Client) GetNFSFilesystemByID(ctx context.Context, filesystemID string) (*NFSFilesystem, error) {
	return doGETRequest[NFSFilesystem](ctx, c, formatAPIPath("nfs/filesystems/%s", filesystemID))
}

func (c *Client) CreateNFSFilesystem(
	ctx context.Context,
	spaceID string,
	opts NFSFilesystemCreateOptions,
) (*NFSFilesystem, error) {
	return doPOSTRequest[NFSFilesystem](ctx, c, formatAPIPath("nfs/spaces/%s/filesystems", spaceID), opts)
}

func (c *Client) UpdateNFSFilesystem(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSFilesystemUpdateOptions,
) (*NFSFilesystem, error) {
	return doPUTRequest[NFSFilesystem](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSFilesystem(ctx context.Context, spaceID string, filesystemID string) error {
	return doDELETERequest(ctx, c, formatAPIPath("nfs/spaces/%s/filesystems/%s", spaceID, filesystemID))
}

func (c *Client) GetNFSFilesystemStats(
	ctx context.Context,
	spaceID string,
	filesystemID string,
) (*NFSFilesystemStats, error) {
	return doGETRequest[NFSFilesystemStats](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/stats", spaceID, filesystemID),
	)
}
