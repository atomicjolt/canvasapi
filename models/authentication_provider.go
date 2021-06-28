package models

type AuthenticationProvider struct {
	IDentifierFormat       string                     `json:"identifier_format" url:"identifier_format,omitempty"`             // Valid for SAML providers..Example: urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress
	AuthType               string                     `json:"auth_type" url:"auth_type,omitempty"`                             // Valid for all providers..Example: saml
	ID                     int64                      `json:"id" url:"id,omitempty"`                                           // Valid for all providers..Example: 1649
	LogOutUrl              string                     `json:"log_out_url" url:"log_out_url,omitempty"`                         // Valid for SAML providers..Example: http://example.com/saml1/slo
	LogInUrl               string                     `json:"log_in_url" url:"log_in_url,omitempty"`                           // Valid for SAML and CAS providers..Example: http://example.com/saml1/sli
	CertificateFingerprint string                     `json:"certificate_fingerprint" url:"certificate_fingerprint,omitempty"` // Valid for SAML providers..Example: 111222
	RequestedAuthnContext  string                     `json:"requested_authn_context" url:"requested_authn_context,omitempty"` // Valid for SAML providers..
	AuthHost               string                     `json:"auth_host" url:"auth_host,omitempty"`                             // Valid for LDAP providers..Example: 127.0.0.1
	AuthFilter             string                     `json:"auth_filter" url:"auth_filter,omitempty"`                         // Valid for LDAP providers..Example: filter1
	AuthOverTls            int64                      `json:"auth_over_tls" url:"auth_over_tls,omitempty"`                     // Valid for LDAP providers..
	AuthBase               string                     `json:"auth_base" url:"auth_base,omitempty"`                             // Valid for LDAP and CAS providers..
	AuthUsername           string                     `json:"auth_username" url:"auth_username,omitempty"`                     // Valid for LDAP providers..Example: username1
	AuthPort               int64                      `json:"auth_port" url:"auth_port,omitempty"`                             // Valid for LDAP providers..
	Position               int64                      `json:"position" url:"position,omitempty"`                               // Valid for all providers..Example: 1
	IDpEntityID            string                     `json:"idp_entity_id" url:"idp_entity_id,omitempty"`                     // Valid for SAML providers..Example: http://example.com/saml1
	LoginAttribute         string                     `json:"login_attribute" url:"login_attribute,omitempty"`                 // Valid for SAML providers..Example: nameid
	SigAlg                 string                     `json:"sig_alg" url:"sig_alg,omitempty"`                                 // Valid for SAML providers..Example: http://www.w3.org/2001/04/xmldsig-more#rsa-sha256
	JitProvisioning        bool                       `json:"jit_provisioning" url:"jit_provisioning,omitempty"`               // Just In Time provisioning. Valid for all providers except Canvas (which has the similar in concept self_registration setting)..
	FederatedAttributes    *FederatedAttributesConfig `json:"federated_attributes" url:"federated_attributes,omitempty"`       //
	MfaRequired            bool                       `json:"mfa_required" url:"mfa_required,omitempty"`                       // If multi-factor authentication is required when logging in with this authentication provider. The account must not have MFA disabled..
}

func (t *AuthenticationProvider) HasError() error {
	return nil
}
