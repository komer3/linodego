package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
)

type NFSSpaceStatus string

const (
	NFSSpaceStatusCreating NFSSpaceStatus = "creating"
	NFSSpaceStatusActive   NFSSpaceStatus = "active"
	NFSSpaceStatusDeleting NFSSpaceStatus = "deleting"
	NFSSpaceStatusError    NFSSpaceStatus = "error"
)

type NFSSpace struct {
	ID          string         `json:"id"`
	Label       string         `json:"label"`
	Description *string        `json:"description"`
	Default     bool           `json:"default"`
	Status      NFSSpaceStatus `json:"status"`
	Created     *time.Time     `json:"-"`
	Updated     *time.Time     `json:"-"`
	Tags        []string       `json:"tags"`
}

type NFSSpaceCreateOptions struct {
	Label       string   `json:"label"`
	Description *string  `json:"description,omitzero"`
	Tags        []string `json:"tags,omitzero"`
}

type NFSSpaceUpdateOptions struct {
	Label       string   `json:"label,omitzero"`
	Description *string  `json:"description,omitzero"`
	Tags        []string `json:"tags,omitzero"`
}

func (n *NFSSpace) UnmarshalJSON(b []byte) error {
	type Mask NFSSpace

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

func (n NFSSpace) GetCreateOptions() NFSSpaceCreateOptions {
	return NFSSpaceCreateOptions{
		Label:       n.Label,
		Description: n.Description,
		Tags:        n.Tags,
	}
}

func (n NFSSpace) GetUpdateOptions() NFSSpaceUpdateOptions {
	return NFSSpaceUpdateOptions{
		Label:       n.Label,
		Description: n.Description,
		Tags:        n.Tags,
	}
}

func (c *Client) ListNFSSpaces(ctx context.Context, opts *ListOptions) ([]NFSSpace, error) {
	return getPaginatedResults[NFSSpace](ctx, c, "nfs/spaces", opts)
}

func (c *Client) GetNFSSpace(ctx context.Context, spaceID string) (*NFSSpace, error) {
	return doGETRequest[NFSSpace](ctx, c, formatAPIPath("nfs/spaces/%s", spaceID))
}

func (c *Client) CreateNFSSpace(ctx context.Context, opts NFSSpaceCreateOptions) (*NFSSpace, error) {
	return doPOSTRequest[NFSSpace](ctx, c, "nfs/spaces", opts)
}

func (c *Client) UpdateNFSSpace(ctx context.Context, spaceID string, opts NFSSpaceUpdateOptions) (*NFSSpace, error) {
	return doPUTRequest[NFSSpace](ctx, c, formatAPIPath("nfs/spaces/%s", spaceID), opts)
}

func (c *Client) DeleteNFSSpace(ctx context.Context, spaceID string) error {
	return doDELETERequest(ctx, c, formatAPIPath("nfs/spaces/%s", spaceID))
}
