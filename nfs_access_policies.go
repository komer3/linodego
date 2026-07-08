package linodego

import (
	"context"
	"encoding/json"
	"time"
)

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
	NFSAccessPolicyStatusError    NFSAccessPolicyStatus = "error"
)

type NFSSquashPolicy string

const (
	NFSSquashPolicyNone       NFSSquashPolicy = "none"
	NFSSquashPolicyRootSquash NFSSquashPolicy = "root_squash"
	NFSSquashPolicyAllSquash  NFSSquashPolicy = "all_squash"
)

type NFSSpaceAccessPolicyVPC struct {
	ID      int                             `json:"id"`
	Label   *string                         `json:"label"`
	URL     *string                         `json:"url"`
	Range   *string                         `json:"range"`
	Subnets []NFSSpaceAccessPolicyVPCSubnet `json:"subnets"`
}

type NFSSpaceAccessPolicyVPCSubnet struct {
	ID    int     `json:"id"`
	Label *string `json:"label"`
	URL   *string `json:"url"`
	Range *string `json:"range"`
}

type NFSSpaceAccessPolicyVPCOptions struct {
	ID      int   `json:"id"`
	Subnets []int `json:"subnets,omitzero"`
}

type NFSFilesystemAccessPolicyLinode struct {
	ID    int     `json:"id"`
	Label *string `json:"label"`
	URL   *string `json:"url"`
	IPv6  *string `json:"ipv6"`
}

type NFSSpaceAccessPolicy struct {
	SpaceID    int                       `json:"space_id"`
	Label      string                    `json:"label"`
	Enabled    bool                      `json:"enabled"`
	VPCACL     []NFSSpaceAccessPolicyVPC `json:"vpc_acl"`
	MTLSCACert *string                   `json:"mtls_ca_cert"`
	MTLSMode   NFSMTLSMode               `json:"mtls_mode"`
	Status     NFSAccessPolicyStatus     `json:"status"`
	Created    *time.Time                `json:"-"`
	Updated    *time.Time                `json:"-"`
}

type NFSSpaceAccessPolicyUpdateOptions struct {
	Label      *string                           `json:"label,omitzero"`
	Enabled    *bool                             `json:"enabled,omitzero"`
	VPCs       *[]NFSSpaceAccessPolicyVPCOptions `json:"vpcs,omitzero"`
	MTLSCACert *string                           `json:"mtls_ca_cert,omitzero"`
	MTLSMode   *NFSMTLSMode                      `json:"mtls_mode,omitzero"`
}

type NFSFilesystemAccessPolicy struct {
	FilesystemID int                               `json:"filesystem_id"`
	Label        string                            `json:"label"`
	Enabled      bool                              `json:"enabled"`
	LinodeACL    []NFSFilesystemAccessPolicyLinode `json:"linode_acl"`
	SquashPolicy NFSSquashPolicy                   `json:"squash_policy"`
	Protocols    []NFSProtocolVersion              `json:"protocols"`
	Status       NFSAccessPolicyStatus             `json:"status"`
	Created      *time.Time                        `json:"-"`
	Updated      *time.Time                        `json:"-"`
}

type NFSFilesystemAccessPolicyUpdateOptions struct {
	Label        *string               `json:"label,omitzero"`
	Enabled      *bool                 `json:"enabled,omitzero"`
	LinodeIDs    *[]int                `json:"linode_ids,omitzero"`
	SquashPolicy *NFSSquashPolicy      `json:"squash_policy,omitzero"`
	Protocols    *[]NFSProtocolVersion `json:"protocols,omitzero"`
}

func (n *NFSSpaceAccessPolicy) UnmarshalJSON(b []byte) error {
	type Mask NFSSpaceAccessPolicy

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

func (n *NFSFilesystemAccessPolicy) UnmarshalJSON(b []byte) error {
	type Mask NFSFilesystemAccessPolicy

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

func (c *Client) GetNFSSpaceAccessPolicy(ctx context.Context, spaceID int) (*NFSSpaceAccessPolicy, error) {
	return doGETRequest[NFSSpaceAccessPolicy](ctx, c, formatAPIPath("nfs/spaces/%d/access-policy", spaceID))
}

func (c *Client) UpdateNFSSpaceAccessPolicy(
	ctx context.Context,
	spaceID int,
	opts NFSSpaceAccessPolicyUpdateOptions,
) (*NFSSpaceAccessPolicy, error) {
	return doPUTRequest[NFSSpaceAccessPolicy](ctx, c, formatAPIPath("nfs/spaces/%d/access-policy", spaceID), opts)
}

func (c *Client) GetNFSFilesystemAccessPolicy(
	ctx context.Context,
	spaceID int,
	filesystemID int,
) (*NFSFilesystemAccessPolicy, error) {
	return doGETRequest[NFSFilesystemAccessPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/access-policy", spaceID, filesystemID),
	)
}

func (c *Client) UpdateNFSFilesystemAccessPolicy(
	ctx context.Context,
	spaceID int,
	filesystemID int,
	opts NFSFilesystemAccessPolicyUpdateOptions,
) (*NFSFilesystemAccessPolicy, error) {
	return doPUTRequest[NFSFilesystemAccessPolicy](
		ctx,
		c,
		formatAPIPath("nfs/spaces/%d/filesystems/%d/access-policy", spaceID, filesystemID),
		opts,
	)
}
