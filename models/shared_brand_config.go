package models

import (
	"time"
)

type SharedBrandConfig struct {
	ID             int64     `json:"id" url:"id,omitempty"`                             // The shared_brand_config identifier..Example: 987
	AccountID      string    `json:"account_id" url:"account_id,omitempty"`             // The id of the account it should be shared within..
	BrandConfigMd5 string    `json:"brand_config_md5" url:"brand_config_md5,omitempty"` // The md5 (since BrandConfigs are identified by MD5 and not numeric id) of the BrandConfig to share..Example: 1d31002c95842f8fe16da7dfcc0d1f39
	Name           string    `json:"name" url:"name,omitempty"`                         // The name to share this theme as.Example: Crimson and Gold Verson 1
	CreatedAt      time.Time `json:"created_at" url:"created_at,omitempty"`             // When this was created.Example: 2012-07-13T10:55:20-06:00
	UpdatedAt      time.Time `json:"updated_at" url:"updated_at,omitempty"`             // When this was last updated.Example: 2012-07-13T10:55:20-06:00
}

func (t *SharedBrandConfig) HasError() error {
	return nil
}
