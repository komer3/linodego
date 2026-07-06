package linodego

import (
	"context"
	"encoding/json"
	"time"
)

type NFSLDAPConfigStatus string

const (
	NFSLDAPConfigStatusActive      NFSLDAPConfigStatus = "active"
	NFSLDAPConfigStatusUpdating    NFSLDAPConfigStatus = "updating"
	NFSLDAPConfigStatusUnreachable NFSLDAPConfigStatus = "unreachable"
	NFSLDAPConfigStatusError       NFSLDAPConfigStatus = "error"
)

type NFSLDAPConfig struct {
	SpaceID         int                 `json:"space_id"`
	URL             string              `json:"url"`
	BaseDN          string              `json:"base_dn"`
	BindDN          string              `json:"bind_dn"`
	TLSCACert       *string             `json:"tls_ca_cert"`
	UserSearchBase  string              `json:"user_search_base"`
	GroupSearchBase string              `json:"group_search_base"`
	UIDAttribute    string              `json:"uid_attribute"`
	GIDAttribute    string              `json:"gid_attribute"`
	Status          NFSLDAPConfigStatus `json:"status"`
	Created         *time.Time          `json:"-"`
	Updated         *time.Time          `json:"-"`
	Verified        *time.Time          `json:"-"`
}

type NFSLDAPConfigUpsertOptions struct {
	URL             *string  `json:"url,omitzero"`
	BaseDN          *string  `json:"base_dn,omitzero"`
	BindDN          *string  `json:"bind_dn,omitzero"`
	BindPassword    *string  `json:"bind_password,omitzero"`
	TLSCACert       **string `json:"tls_ca_cert,omitzero"`
	UserSearchBase  *string  `json:"user_search_base,omitzero"`
	GroupSearchBase *string  `json:"group_search_base,omitzero"`
	UIDAttribute    *string  `json:"uid_attribute,omitzero"`
	GIDAttribute    *string  `json:"gid_attribute,omitzero"`
}

type NFSLDAPTestResult struct {
	Reachable   bool       `json:"reachable"`
	BindSuccess bool       `json:"bind_success"`
	LatencyMS   int        `json:"latency_ms"`
	Message     string     `json:"message"`
	Verified    *time.Time `json:"-"`
}

func (n *NFSLDAPConfig) UnmarshalJSON(b []byte) error {
	type Mask NFSLDAPConfig

	p := struct {
		*Mask

		Created  *time.Time `json:"created"`
		Updated  *time.Time `json:"updated"`
		Verified *time.Time `json:"verified"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Created = p.Created
	n.Updated = p.Updated
	n.Verified = p.Verified

	return nil
}

func (n *NFSLDAPTestResult) UnmarshalJSON(b []byte) error {
	type Mask NFSLDAPTestResult

	p := struct {
		*Mask

		Verified *time.Time `json:"verified"`
	}{
		Mask: (*Mask)(n),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	n.Verified = p.Verified

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
