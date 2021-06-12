package models

type FederatedAttributeConfig struct {
	Attribute        string `json:"attribute"`         // The name of the attribute as it will be sent from the authentication provider.Example: mail
	ProvisioningOnly bool   `json:"provisioning_only"` // If the attribute should be applied only when provisioning a new user, rather than all logins.
}

func (t *FederatedAttributeConfig) HasError() error {
	return nil
}
