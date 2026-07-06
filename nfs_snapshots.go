package linodego

import (
	"context"
	"encoding/json"
	"time"
)

type NFSSnapshotStatus string

const (
	NFSSnapshotStatusCreating NFSSnapshotStatus = "creating"
	NFSSnapshotStatusActive   NFSSnapshotStatus = "active"
	NFSSnapshotStatusUpdating NFSSnapshotStatus = "updating"
	NFSSnapshotStatusDeleting NFSSnapshotStatus = "deleting"
	NFSSnapshotStatusError    NFSSnapshotStatus = "error"
)

type NFSSnapshotSource string

const (
	NFSSnapshotSourceManual NFSSnapshotSource = "manual"
	NFSSnapshotSourcePolicy NFSSnapshotSource = "policy"
)

type NFSSnapshot struct {
	ID                  int               `json:"id"`
	FilesystemID        int               `json:"filesystem_id"`
	SpaceID             int               `json:"space_id"`
	Label               string            `json:"label"`
	Status              NFSSnapshotStatus `json:"status"`
	Created             *time.Time        `json:"-"`
	Expiration          *time.Time        `json:"-"`
	Locked              bool              `json:"locked"`
	Source              NFSSnapshotSource `json:"source"`
	PolicyID            *int              `json:"policy_id"`
	SizeBytes           int64             `json:"size_bytes"`
	AggregatedSizeBytes int64             `json:"aggregated_size_bytes"`
	Tags                []string          `json:"tags"`
}

type NFSSnapshotCreateOptions struct {
	Label      string    `json:"label"`
	Expiration **string  `json:"expiration,omitzero"`
	Locked     *bool     `json:"locked,omitzero"`
	Tags       *[]string `json:"tags,omitzero"`
}

type NFSSnapshotUpdateOptions struct {
	Expiration **string  `json:"expiration,omitzero"`
	Locked     *bool     `json:"locked,omitzero"`
	Tags       *[]string `json:"tags,omitzero"`
}

type NFSSnapshotCloneOptions struct {
	Label   string    `json:"label"`
	Region  string    `json:"region"`
	SpaceID **int     `json:"space_id,omitzero"`
	SizeGib **int     `json:"size_gib,omitzero"`
	Tags    *[]string `json:"tags,omitzero"`
}

func (n *NFSSnapshot) UnmarshalJSON(b []byte) error {
	type Mask NFSSnapshot

	p := struct {
		*Mask

		Created    *time.Time `json:"created"`
		Expiration *time.Time `json:"expiration"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Created = p.Created
	n.Expiration = p.Expiration

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

func (c *Client) UpdateNFSSnapshot(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	snapshotID string,
	opts NFSSnapshotUpdateOptions,
) (*NFSSnapshot, error) {
	return doPUTRequest[NFSSnapshot](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/snapshots/%s", spaceID, filesystemID, snapshotID),
		opts,
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
