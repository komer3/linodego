package linodego

import "context"

type NFSMTLSMode string

const (
	NFSMTLSModeRequired NFSMTLSMode = "required"
	NFSMTLSModeOptional NFSMTLSMode = "optional"
	NFSMTLSModeDisabled NFSMTLSMode = "disabled"
)

type NFSAccessPolicyStatus string

const (
	NFSAccessPolicyStatusCreating NFSAccessPolicyStatus = "creating"
	NFSAccessPolicyStatusActive   NFSAccessPolicyStatus = "active"
	NFSAccessPolicyStatusUpdating NFSAccessPolicyStatus = "updating"
	NFSAccessPolicyStatusDeleting NFSAccessPolicyStatus = "deleting"
)

type NFSRootSquashMode string

const (
	NFSRootSquashModeNone       NFSRootSquashMode = "none"
	NFSRootSquashModeRootSquash NFSRootSquashMode = "root_squash"
	NFSRootSquashModeAllSquash  NFSRootSquashMode = "all_squash"
)

type NFSPosixOverride struct {
	UID int `json:"uid"`
	GID int `json:"gid"`
}

type NFSSpaceAccessPolicy struct {
	SpaceID      string                `json:"space_id"`
	Label        string                `json:"label"`
	Enabled      bool                  `json:"enabled"`
	VPCIDs       []string              `json:"vpc_ids"`
	AllowedCIDRs []string              `json:"allowed_cidrs"`
	MTLSMode     NFSMTLSMode           `json:"mtls_mode"`
	Status       NFSAccessPolicyStatus `json:"status"`
}

type NFSSpaceAccessPolicyUpdateOptions struct {
	Label        string      `json:"label,omitzero"`
	Enabled      *bool       `json:"enabled,omitzero"`
	VPCIDs       []string    `json:"vpc_ids,omitzero"`
	AllowedCIDRs []string    `json:"allowed_cidrs,omitzero"`
	MTLSCACert   *string     `json:"mtls_ca_cert,omitzero"`
	MTLSMode     NFSMTLSMode `json:"mtls_mode,omitzero"`
}

type NFSFilesystemAccessPolicy struct {
	FilesystemID  string                `json:"filesystem_id"`
	Label         string                `json:"label"`
	Enabled       bool                  `json:"enabled"`
	LinodeIDs     []int                 `json:"linode_ids"`
	LinodeIPs     []string              `json:"linode_ips"`
	RootSquash    NFSRootSquashMode     `json:"root_squash"`
	Protocols     []NFSProtocolVersion  `json:"protocols"`
	PosixOverride *NFSPosixOverride     `json:"posix_override"`
	Status        NFSAccessPolicyStatus `json:"status"`
}

type NFSFilesystemAccessPolicyUpdateOptions struct {
	Label         string               `json:"label,omitzero"`
	Enabled       *bool                `json:"enabled,omitzero"`
	LinodeIDs     []int                `json:"linode_ids,omitzero"`
	RootSquash    NFSRootSquashMode    `json:"root_squash,omitzero"`
	Protocols     []NFSProtocolVersion `json:"protocols,omitzero"`
	PosixOverride *NFSPosixOverride    `json:"posix_override,omitzero"`
}

func (c *Client) GetNFSSpaceAccessPolicy(ctx context.Context, spaceID string) (*NFSSpaceAccessPolicy, error) {
	return doGETRequest[NFSSpaceAccessPolicy](ctx, c, formatAPIPath("nfs/spaces/%s/access-policy", spaceID))
}

func (c *Client) UpdateNFSSpaceAccessPolicy(
	ctx context.Context,
	spaceID string,
	opts NFSSpaceAccessPolicyUpdateOptions,
) (*NFSSpaceAccessPolicy, error) {
	return doPUTRequest[NFSSpaceAccessPolicy](ctx, c, formatAPIPath("nfs/spaces/%s/access-policy", spaceID), opts)
}

func (c *Client) GetNFSFilesystemAccessPolicy(
	ctx context.Context,
	spaceID string,
	filesystemID string,
) (*NFSFilesystemAccessPolicy, error) {
	return doGETRequest[NFSFilesystemAccessPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/access-policy", spaceID, filesystemID),
	)
}

func (c *Client) UpdateNFSFilesystemAccessPolicy(
	ctx context.Context,
	spaceID string,
	filesystemID string,
	opts NFSFilesystemAccessPolicyUpdateOptions,
) (*NFSFilesystemAccessPolicy, error) {
	return doPUTRequest[NFSFilesystemAccessPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%s/filesystems/%s/access-policy", spaceID, filesystemID),
		opts,
	)
}
