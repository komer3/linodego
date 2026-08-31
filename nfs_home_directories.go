package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
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

func (c *Client) GetNFSHomeDirectoryConfig(
	ctx context.Context,
	spaceID int,
	filesystemID int,
) (*NFSHomeDirectoryConfig, error) {
	return doGETRequest[NFSHomeDirectoryConfig](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/home-directory", spaceID, filesystemID),
	)
}

func (c *Client) UpsertNFSHomeDirectoryConfig(
	ctx context.Context,
	spaceID int,
	filesystemID int,
	opts NFSHomeDirectoryConfigUpdateOptions,
) (*NFSHomeDirectoryConfig, error) {
	return doPUTRequest[NFSHomeDirectoryConfig](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/home-directory", spaceID, filesystemID),
		opts,
	)
}

func (c *Client) DeleteNFSHomeDirectoryConfig(ctx context.Context, spaceID int, filesystemID int) error {
	return doDELETERequest(
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/home-directory", spaceID, filesystemID),
	)
}
