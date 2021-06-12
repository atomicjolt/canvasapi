package models

type AuthenticationProvider struct {
	IDentifierFormat       string                     `json:"identifier_format"`       // Valid for SAML providers..Example: urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress
	AuthType               string                     `json:"auth_type"`               // Valid for all providers..Example: saml
	ID                     int64                      `json:"id"`                      // Valid for all providers..Example: 1649
	LogOutUrl              string                     `json:"log_out_url"`             // Valid for SAML providers..Example: http://example.com/saml1/slo
	LogInUrl               string                     `json:"log_in_url"`              // Valid for SAML and CAS providers..Example: http://example.com/saml1/sli
	CertificateFingerprint string                     `json:"certificate_fingerprint"` // Valid for SAML providers..Example: 111222
	RequestedAuthnContext  string                     `json:"requested_authn_context"` // Valid for SAML providers..
	AuthHost               string                     `json:"auth_host"`               // Valid for LDAP providers..Example: 127.0.0.1
	AuthFilter             string                     `json:"auth_filter"`             // Valid for LDAP providers..Example: filter1
	AuthOverTls            int64                      `json:"auth_over_tls"`           // Valid for LDAP providers..
	AuthBase               string                     `json:"auth_base"`               // Valid for LDAP and CAS providers..
	AuthUsername           string                     `json:"auth_username"`           // Valid for LDAP providers..Example: username1
	AuthPort               int64                      `json:"auth_port"`               // Valid for LDAP providers..
	Position               int64                      `json:"position"`                // Valid for all providers..Example: 1
	IDpEntityID            string                     `json:"idp_entity_id"`           // Valid for SAML providers..Example: http://example.com/saml1
	LoginAttribute         string                     `json:"login_attribute"`         // Valid for SAML providers..Example: nameid
	SigAlg                 string                     `json:"sig_alg"`                 // Valid for SAML providers..Example: http://www.w3.org/2001/04/xmldsig-more#rsa-sha256
	JitProvisioning        bool                       `json:"jit_provisioning"`        // Just In Time provisioning. Valid for all providers except Canvas (which has the similar in concept self_registration setting)..
	FederatedAttributes    *FederatedAttributesConfig `json:"federated_attributes"`    //
	MfaRequired            bool                       `json:"mfa_required"`            // If multi-factor authentication is required when logging in with this authentication provider. The account must not have MFA disabled..
}

func (t *AuthenticationProvider) HasError() error {
	return nil
}
