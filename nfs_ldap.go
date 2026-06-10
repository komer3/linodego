package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/v2/internal/parseabletime"
)

type NFSLDAPConfigStatus string

const (
	NFSLDAPConfigStatusActive      NFSLDAPConfigStatus = "active"
	NFSLDAPConfigStatusUnreachable NFSLDAPConfigStatus = "unreachable"
	NFSLDAPConfigStatusError       NFSLDAPConfigStatus = "error"
)

type NFSLDAPConfig struct {
	SpaceID         string              `json:"space_id"`
	URL             string              `json:"url"`
	BaseDN          string              `json:"base_dn"`
	BindDN          string              `json:"bind_dn"`
	TLSCACert       *string             `json:"tls_ca_cert"`
	UserSearchBase  string              `json:"user_search_base"`
	GroupSearchBase string              `json:"group_search_base"`
	UIDAttribute    string              `json:"uid_attribute"`
	GIDAttribute    string              `json:"gid_attribute"`
	Status          NFSLDAPConfigStatus `json:"status"`
}

type NFSLDAPConfigUpsertOptions struct {
	URL             string  `json:"url"`
	BaseDN          string  `json:"base_dn"`
	BindDN          string  `json:"bind_dn"`
	BindPassword    string  `json:"bind_password"`
	TLSCACert       *string `json:"tls_ca_cert,omitzero"`
	UserSearchBase  string  `json:"user_search_base"`
	GroupSearchBase string  `json:"group_search_base"`
	UIDAttribute    string  `json:"uid_attribute,omitzero"`
	GIDAttribute    string  `json:"gid_attribute,omitzero"`
}

type NFSLDAPTestResult struct {
	Reachable   bool       `json:"reachable"`
	BindSuccess bool       `json:"bind_success"`
	LatencyMS   int        `json:"latency_ms"`
	Message     string     `json:"message"`
	TestedAt    *time.Time `json:"-"`
}

func (n *NFSLDAPTestResult) UnmarshalJSON(b []byte) error {
	type Mask NFSLDAPTestResult

	p := struct {
		*Mask

		TestedAt *parseabletime.ParseableTime `json:"tested_at"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.TestedAt = (*time.Time)(p.TestedAt)

	return nil
}

func (c *Client) GetNFSLDAPConfig(ctx context.Context, spaceID string) (*NFSLDAPConfig, error) {
	return doGETRequest[NFSLDAPConfig](ctx, c, formatAPIPath("nfs/spaces/%s/ldap-config", spaceID))
}

func (c *Client) UpsertNFSLDAPConfig(
	ctx context.Context,
	spaceID string,
	opts NFSLDAPConfigUpsertOptions,
) (*NFSLDAPConfig, error) {
	return doPUTRequest[NFSLDAPConfig](ctx, c, formatAPIPath("nfs/spaces/%s/ldap-config", spaceID), opts)
}

func (c *Client) DeleteNFSLDAPConfig(ctx context.Context, spaceID string) error {
	return doDELETERequest(ctx, c, formatAPIPath("nfs/spaces/%s/ldap-config", spaceID))
}

func (c *Client) TestNFSLDAPConfig(ctx context.Context, spaceID string) (*NFSLDAPTestResult, error) {
	return doPOSTRequestNoRequestBody[NFSLDAPTestResult](ctx, c, formatAPIPath("nfs/spaces/%s/ldap-config/test", spaceID))
}
