package models

type FederatedAttributeConfig struct {
	Attribute        string `json:"attribute" url:"attribute,omitempty"`                 // The name of the attribute as it will be sent from the authentication provider.Example: mail
	ProvisioningOnly bool   `json:"provisioning_only" url:"provisioning_only,omitempty"` // If the attribute should be applied only when provisioning a new user, rather than all logins.
}

func (t *FederatedAttributeConfig) HasErrors() error {
	return nil
}
