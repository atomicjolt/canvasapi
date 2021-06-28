package models

import (
	"time"
)

type DeveloperKey struct {
	ID               float64   `json:"id" url:"id,omitempty"`                                 // The ID should match the Developer Key ID in canvas.Example: 1000000000040
	IsLtiKey         bool      `json:"is_lti_key" url:"is_lti_key,omitempty"`                 // true the tool is a lti key, null is not a lti key.Example: true
	Visible          bool      `json:"visible" url:"visible,omitempty"`                       // Controls if the tool is visable.Example: true
	AccountName      string    `json:"account_name" url:"account_name,omitempty"`             // The name of the account associated with the tool.Example: The Academy
	PublicJwk        string    `json:"public_jwk" url:"public_jwk,omitempty"`                 // The public key in jwk format.Example: { 	'kty':'RSA', 	'e':'AQAB', 	'n':'ufmgt156hs168mgdhy168jrsydt168ju816rtahesuvdbmnrtd87t7h8ser', 	'alg':'RS256', 	'use':'sig', 	'kid':'Se68gr16s6tj_87sdr98g489dsfjy-547a6eht1', }
	VendorCode       string    `json:"vendor_code" url:"vendor_code,omitempty"`               // The code of the vendor managing the tool.Example: fi5689s9avewr68
	LastUsedAt       time.Time `json:"last_used_at" url:"last_used_at,omitempty"`             // The date and time the tool was last used.Example: 2019-06-07T20:34:33Z
	AccessTokenCount float64   `json:"access_token_count" url:"access_token_count,omitempty"` // The number of active access tokens associated with the tool.Example: 0
	RedirectUris     string    `json:"redirect_uris" url:"redirect_uris,omitempty"`           // redirect uris description.Example: https://redirect.to.here.com
	RedirectUri      string    `json:"redirect_uri" url:"redirect_uri,omitempty"`             // redirect uri description.Example: https://redirect.to.here.com
	ApiKey           string    `json:"api_key" url:"api_key,omitempty"`                       // Api key for api access for the tool.Example: sd45fg648sr546tgh15S15df5se56r4xdf45asef456
	Notes            string    `json:"notes" url:"notes,omitempty"`                           // Notes for use specifications for the tool.Example: Used for sorting graded assignments
	Name             string    `json:"name" url:"name,omitempty"`                             // Display name of the tool.Example: Tool 1
	UserID           string    `json:"user_id" url:"user_id,omitempty"`                       // ID of the user associated with the tool.Example: tu816dnrs6zdsg148918dmu
	CreatedAt        time.Time `json:"created_at" url:"created_at,omitempty"`                 // The time the jwk was created.Example: 2019-06-07T20:34:33Z
	UserName         string    `json:"user_name" url:"user_name,omitempty"`                   // The user name of the tool creator.Example: johnsmith
	Email            string    `json:"email" url:"email,omitempty"`                           // Email associated with the tool owner.Example: johnsmith@instructure.com
	RequireScopes    bool      `json:"require_scopes" url:"require_scopes,omitempty"`         // True if the tool has required permissions, null if there are no needed permissions.Example: true
	IconUrl          string    `json:"icon_url" url:"icon_url,omitempty"`                     // Icon to be displayed with the name of the tool.Example: null
	Scopes           string    `json:"scopes" url:"scopes,omitempty"`                         // Specified permissions for the tool.Example: https://canvas.instructure.com/lti/public_jwk/scope/update
	WorkflowState    string    `json:"workflow_state" url:"workflow_state,omitempty"`         // The current state of the tool.Example: active
}

func (t *DeveloperKey) HasError() error {
	return nil
}
