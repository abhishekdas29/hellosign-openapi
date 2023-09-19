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

// OAuthTokenGenerateRequest struct for OAuthTokenGenerateRequest
type OAuthTokenGenerateRequest struct {
	// The client id of the app requesting authorization.
	ClientId string `json:"client_id"`
	// The secret token of your app.
	ClientSecret string `json:"client_secret"`
	// The code passed to your callback when the user granted access.
	Code string `json:"code"`
	// When generating a new token use `authorization_code`.
	GrantType string `json:"grant_type"`
	// Same as the state you specified earlier.
	State string `json:"state"`
}

// NewOAuthTokenGenerateRequest instantiates a new OAuthTokenGenerateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthTokenGenerateRequest(clientId string, clientSecret string, code string, grantType string, state string) *OAuthTokenGenerateRequest {
	this := OAuthTokenGenerateRequest{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Code = code
	this.GrantType = grantType
	this.State = state
	return &this
}

// NewOAuthTokenGenerateRequestWithDefaults instantiates a new OAuthTokenGenerateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthTokenGenerateRequestWithDefaults() *OAuthTokenGenerateRequest {
	this := OAuthTokenGenerateRequest{}
	var grantType string = "authorization_code"
	this.GrantType = grantType
	return &this
}

// GetClientId returns the ClientId field value
func (o *OAuthTokenGenerateRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenGenerateRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuthTokenGenerateRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *OAuthTokenGenerateRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenGenerateRequest) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *OAuthTokenGenerateRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCode returns the Code field value
func (o *OAuthTokenGenerateRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenGenerateRequest) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OAuthTokenGenerateRequest) SetCode(v string) {
	o.Code = v
}

// GetGrantType returns the GrantType field value
func (o *OAuthTokenGenerateRequest) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenGenerateRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *OAuthTokenGenerateRequest) SetGrantType(v string) {
	o.GrantType = v
}

// GetState returns the State field value
func (o *OAuthTokenGenerateRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenGenerateRequest) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OAuthTokenGenerateRequest) SetState(v string) {
	o.State = v
}

func (o OAuthTokenGenerateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["grant_type"] = o.GrantType
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthTokenGenerateRequest struct {
	value *OAuthTokenGenerateRequest
	isSet bool
}

func (v NullableOAuthTokenGenerateRequest) Get() *OAuthTokenGenerateRequest {
	return v.value
}

func (v *NullableOAuthTokenGenerateRequest) Set(val *OAuthTokenGenerateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthTokenGenerateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthTokenGenerateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthTokenGenerateRequest(val *OAuthTokenGenerateRequest) *NullableOAuthTokenGenerateRequest {
	return &NullableOAuthTokenGenerateRequest{value: val, isSet: true}
}

func (v NullableOAuthTokenGenerateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthTokenGenerateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

