package linodego

import (
	"context"
	"encoding/json"
	"time"
)

type NFSSpaceStatus string

const (
	NFSSpaceStatusCreating NFSSpaceStatus = "creating"
	NFSSpaceStatusActive   NFSSpaceStatus = "active"
	NFSSpaceStatusDeleting NFSSpaceStatus = "deleting"
	NFSSpaceStatusError    NFSSpaceStatus = "error"
)

type NFSSpace struct {
	ID          int            `json:"id"`
	Label       string         `json:"label"`
	Description *string        `json:"description"`
	Status      NFSSpaceStatus `json:"status"`
	Created     *time.Time     `json:"-"`
	Updated     *time.Time     `json:"-"`
	Tags        []string       `json:"tags"`
}

type NFSSpaceCreateOptions struct {
	Label       string    `json:"label"`
	Description **string  `json:"description,omitzero"`
	Tags        *[]string `json:"tags,omitzero"`
}

type NFSSpaceUpdateOptions struct {
	Label       *string   `json:"label,omitzero"`
	Description **string  `json:"description,omitzero"`
	Tags        *[]string `json:"tags,omitzero"`
}

func (n *NFSSpace) UnmarshalJSON(b []byte) error {
	type Mask NFSSpace

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

func (n NFSSpace) GetCreateOptions() NFSSpaceCreateOptions {
	result := NFSSpaceCreateOptions{
		Label:       n.Label,
		Description: Pointer(n.Description),
	}

	if n.Tags != nil {
		result.Tags = Pointer(n.Tags)
	}

	return result
}

func (n NFSSpace) GetUpdateOptions() NFSSpaceUpdateOptions {
	result := NFSSpaceUpdateOptions{
		Label:       Pointer(n.Label),
		Description: Pointer(n.Description),
	}

	if n.Tags != nil {
		result.Tags = Pointer(n.Tags)
	}

	return result
}

func (c *Client) ListNFSSpaces(ctx context.Context, opts *ListOptions) ([]NFSSpace, error) {
	return getPaginatedResults[NFSSpace](ctx, c, "nfs/spaces", opts)
}

func (c *Client) GetNFSSpace(ctx context.Context, spaceID int) (*NFSSpace, error) {
	return doGETRequest[NFSSpace](ctx, c, formatAPIPath("nfs/spaces/%d", spaceID))
}

func (c *Client) CreateNFSSpace(ctx context.Context, opts NFSSpaceCreateOptions) (*NFSSpace, error) {
	return doPOSTRequest[NFSSpace](ctx, c, "nfs/spaces", opts)
}

func (c *Client) UpdateNFSSpace(ctx context.Context, spaceID int, opts NFSSpaceUpdateOptions) (*NFSSpace, error) {
	return doPUTRequest[NFSSpace](ctx, c, formatAPIPath("nfs/spaces/%d", spaceID), opts)
}

func (c *Client) DeleteNFSSpace(ctx context.Context, spaceID int) error {
	return doDELETERequest(ctx, c, formatAPIPath("nfs/spaces/%d", spaceID))
}
