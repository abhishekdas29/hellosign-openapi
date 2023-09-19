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

// EmbeddedSignUrlResponseEmbedded An object that contains necessary information to set up embedded signing.
type EmbeddedSignUrlResponseEmbedded struct {
	// A signature url that can be opened in an iFrame.
	SignUrl *string `json:"sign_url,omitempty"`
	// The specific time that the the `sign_url` link expires, in epoch.
	ExpiresAt *int32 `json:"expires_at,omitempty"`
}

// NewEmbeddedSignUrlResponseEmbedded instantiates a new EmbeddedSignUrlResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedSignUrlResponseEmbedded() *EmbeddedSignUrlResponseEmbedded {
	this := EmbeddedSignUrlResponseEmbedded{}
	return &this
}

// NewEmbeddedSignUrlResponseEmbeddedWithDefaults instantiates a new EmbeddedSignUrlResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedSignUrlResponseEmbeddedWithDefaults() *EmbeddedSignUrlResponseEmbedded {
	this := EmbeddedSignUrlResponseEmbedded{}
	return &this
}

// GetSignUrl returns the SignUrl field value if set, zero value otherwise.
func (o *EmbeddedSignUrlResponseEmbedded) GetSignUrl() string {
	if o == nil || o.SignUrl == nil {
		var ret string
		return ret
	}
	return *o.SignUrl
}

// GetSignUrlOk returns a tuple with the SignUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedSignUrlResponseEmbedded) GetSignUrlOk() (*string, bool) {
	if o == nil || o.SignUrl == nil {
		return nil, false
	}
	return o.SignUrl, true
}

// HasSignUrl returns a boolean if a field has been set.
func (o *EmbeddedSignUrlResponseEmbedded) HasSignUrl() bool {
	if o != nil && o.SignUrl != nil {
		return true
	}

	return false
}

// SetSignUrl gets a reference to the given string and assigns it to the SignUrl field.
func (o *EmbeddedSignUrlResponseEmbedded) SetSignUrl(v string) {
	o.SignUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *EmbeddedSignUrlResponseEmbedded) GetExpiresAt() int32 {
	if o == nil || o.ExpiresAt == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedSignUrlResponseEmbedded) GetExpiresAtOk() (*int32, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *EmbeddedSignUrlResponseEmbedded) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int32 and assigns it to the ExpiresAt field.
func (o *EmbeddedSignUrlResponseEmbedded) SetExpiresAt(v int32) {
	o.ExpiresAt = &v
}

func (o EmbeddedSignUrlResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignUrl != nil {
		toSerialize["sign_url"] = o.SignUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedSignUrlResponseEmbedded struct {
	value *EmbeddedSignUrlResponseEmbedded
	isSet bool
}

func (v NullableEmbeddedSignUrlResponseEmbedded) Get() *EmbeddedSignUrlResponseEmbedded {
	return v.value
}

func (v *NullableEmbeddedSignUrlResponseEmbedded) Set(val *EmbeddedSignUrlResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedSignUrlResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedSignUrlResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedSignUrlResponseEmbedded(val *EmbeddedSignUrlResponseEmbedded) *NullableEmbeddedSignUrlResponseEmbedded {
	return &NullableEmbeddedSignUrlResponseEmbedded{value: val, isSet: true}
}

func (v NullableEmbeddedSignUrlResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedSignUrlResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

