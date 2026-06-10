package linodego

import "context"

type NFSHomeDirectoryConfig struct {
	FilesystemID   string `json:"filesystem_id"`
	PathTemplate   string `json:"path_template"`
	AutoCreateDirs bool   `json:"auto_create_dirs"`
	AutofsMap      string `json:"autofs_map"`
}

type NFSHomeDirectoryConfigUpdateOptions struct {
	PathTemplate   string `json:"path_template"`
	AutoCreateDirs bool   `json:"auto_create_dirs"`
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
