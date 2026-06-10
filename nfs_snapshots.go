package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
)

type NFSSnapshotStatus string

const (
	NFSSnapshotStatusCreating NFSSnapshotStatus = "creating"
	NFSSnapshotStatusActive   NFSSnapshotStatus = "active"
	NFSSnapshotStatusFailed   NFSSnapshotStatus = "failed"
	NFSSnapshotStatusDeleting NFSSnapshotStatus = "deleting"
)

type NFSSnapshotSource string

const (
	NFSSnapshotSourceManual NFSSnapshotSource = "manual"
	NFSSnapshotSourcePolicy NFSSnapshotSource = "policy"
)

type NFSSnapshot struct {
	ID                  string            `json:"id"`
	FilesystemID        string            `json:"filesystem_id"`
	SpaceID             string            `json:"space_id"`
	Label               string            `json:"label"`
	Status              NFSSnapshotStatus `json:"status"`
	Created             *time.Time        `json:"-"`
	Expiration          *time.Time        `json:"-"`
	Locked              bool              `json:"locked"`
	Source              NFSSnapshotSource `json:"source"`
	PolicyID            *string           `json:"policy_id"`
	SizeBytes           int64             `json:"size_bytes"`
	AggregatedSizeBytes int64             `json:"aggregated_size_bytes"`
}

type NFSSnapshotCreateOptions struct {
	Label      string  `json:"label"`
	Expiration *string `json:"expiration,omitzero"`
	Locked     *bool   `json:"locked,omitzero"`
}

type NFSSnapshotCloneOptions struct {
	Label   string   `json:"label"`
	Region  string   `json:"region"`
	SpaceID string   `json:"space_id,omitzero"`
	SizeGib *int     `json:"size_gib,omitzero"`
	Tags    []string `json:"tags,omitzero"`
}

func (n *NFSSnapshot) UnmarshalJSON(b []byte) error {
	type Mask NFSSnapshot

	p := struct {
		*Mask

		Created    *parseabletime.ParseableTime `json:"created"`
		Expiration *parseabletime.ParseableTime `json:"expiration"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Created = (*time.Time)(p.Created)
	n.Expiration = (*time.Time)(p.Expiration)

	return nil
}

func (c *Client) ListNFSSnapshots(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts *ListOptions,
) ([]NFSSnapshot, error) {
	return getPaginatedResults[NFSSnapshot](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) GetNFSSnapshot(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	snapshotID string,
) (*NFSSnapshot, error) {
	return doGETRequest[NFSSnapshot](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots/%s", spaceID, filesystemID, snapshotID),
	)
}

func (c *Client) CreateNFSSnapshot(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSSnapshotCreateOptions,
) (*NFSSnapshot, error) {
	return doPOSTRequest[NFSSnapshot](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSSnapshot(ctx context.Context, spaceID string, filesystemID string, snapshotID string) error {
	return doDELETERequest(
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots/%s", spaceID, filesystemID, snapshotID),
	)
}

func (c *Client) CloneNFSSnapshot(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	snapshotID string,
	opts NFSSnapshotCloneOptions,
) (*NFSFilesystem, error) {
	return doPOSTRequest[NFSFilesystem](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots/%s/clone", spaceID, filesystemID, snapshotID),
		opts,
	)
}
