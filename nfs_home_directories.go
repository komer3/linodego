package linodego

import (
	"context"
	"encoding/json"
	"time"
)

type NFSHomeDirectoryConfig struct {
	FilesystemID   int        `json:"filesystem_id"`
	PathTemplate   string     `json:"path_template"`
	AutoCreateDirs bool       `json:"auto_create_dirs"`
	AutofsMap      *string    `json:"autofs_map"`
	Created        *time.Time `json:"-"`
	Updated        *time.Time `json:"-"`
}

type NFSHomeDirectoryConfigUpdateOptions struct {
	PathTemplate   *string `json:"path_template,omitzero"`
	AutoCreateDirs *bool   `json:"auto_create_dirs,omitzero"`
}

func (n *NFSHomeDirectoryConfig) UnmarshalJSON(b []byte) error {
	type Mask NFSHomeDirectoryConfig

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

func (c *Client) GetNFSHomeDirectoryConfig(
	ctx context.Context,
	spaceID string,
	filesystemID string,
) (*NFSHomeDirectoryConfig, error) {
	return doGETRequest[NFSHomeDirectoryConfig](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/home-directory", spaceID, filesystemID),
	)
}

func (c *Client) UpsertNFSHomeDirectoryConfig(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSHomeDirectoryConfigUpdateOptions,
) (*NFSHomeDirectoryConfig, error) {
	return doPUTRequest[NFSHomeDirectoryConfig](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/home-directory", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSHomeDirectoryConfig(ctx context.Context, spaceID string, filesystemID string) error {
	return doDELETERequest(
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/home-directory", spaceID, filesystemID),
	)
}
