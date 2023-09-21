/*
Dropbox Sign API

Dropbox Sign v3 API

API version: 3.0.0
Contact: apisupport@hellosign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dropboxsign

import (
	"encoding/json"
)

// AccountCreateRequest struct for AccountCreateRequest
type AccountCreateRequest struct {
	// The email address which will be associated with the new Account.
	EmailAddress string `json:"email_address"`
	// Used when creating a new account with OAuth authorization.  See [OAuth 2.0 Authorization](https://app.hellosign.com/api/oauthWalkthrough#OAuthAuthorization)
	ClientId *string `json:"client_id,omitempty"`
	// Used when creating a new account with OAuth authorization.  See [OAuth 2.0 Authorization](https://app.hellosign.com/api/oauthWalkthrough#OAuthAuthorization)
	ClientSecret *string `json:"client_secret,omitempty"`
	// The locale used in this Account. Check out the list of [supported locales](/api/reference/constants/#supported-locales) to learn more about the possible values.
	Locale *string `json:"locale,omitempty"`
}

// NewAccountCreateRequest instantiates a new AccountCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateRequest(emailAddress string) *AccountCreateRequest {
	this := AccountCreateRequest{}
	this.EmailAddress = emailAddress
	return &this
}

// NewAccountCreateRequestWithDefaults instantiates a new AccountCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateRequestWithDefaults() *AccountCreateRequest {
	this := AccountCreateRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *AccountCreateRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *AccountCreateRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AccountCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AccountCreateRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AccountCreateRequest) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateRequest) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AccountCreateRequest) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AccountCreateRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o AccountCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreateRequest struct {
	value *AccountCreateRequest
	isSet bool
}

func (v NullableAccountCreateRequest) Get() *AccountCreateRequest {
	return v.value
}

func (v *NullableAccountCreateRequest) Set(val *AccountCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateRequest(val *AccountCreateRequest) *NullableAccountCreateRequest {
	return &NullableAccountCreateRequest{value: val, isSet: true}
}

func (v NullableAccountCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


