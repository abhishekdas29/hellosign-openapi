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

// SignatureRequestUpdateRequest struct for SignatureRequestUpdateRequest
type SignatureRequestUpdateRequest struct {
	// The signature ID for the recipient.
	SignatureId string `json:"signature_id"`
	// The new email address for the recipient.  This will generate a new `signature_id` value.  **NOTE**: Optional if `name` is provided.
	EmailAddress *string `json:"email_address,omitempty"`
	// The new name for the recipient.  **NOTE**: Optional if `email_address` is provided.
	Name *string `json:"name,omitempty"`
	// The new time when the signature request will expire. Unsigned signatures will be moved to the expired status, and no longer signable. See [Signature Request Expiration Date](https://developers.hellosign.com/docs/signature-request/expiration/) for details.
	ExpiresAt NullableInt32 `json:"expires_at,omitempty"`
}

// NewSignatureRequestUpdateRequest instantiates a new SignatureRequestUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestUpdateRequest(signatureId string) *SignatureRequestUpdateRequest {
	this := SignatureRequestUpdateRequest{}
	this.SignatureId = signatureId
	return &this
}

// NewSignatureRequestUpdateRequestWithDefaults instantiates a new SignatureRequestUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestUpdateRequestWithDefaults() *SignatureRequestUpdateRequest {
	this := SignatureRequestUpdateRequest{}
	return &this
}

// GetSignatureId returns the SignatureId field value
func (o *SignatureRequestUpdateRequest) GetSignatureId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureId
}

// GetSignatureIdOk returns a tuple with the SignatureId field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestUpdateRequest) GetSignatureIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignatureId, true
}

// SetSignatureId sets field value
func (o *SignatureRequestUpdateRequest) SetSignatureId(v string) {
	o.SignatureId = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *SignatureRequestUpdateRequest) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestUpdateRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *SignatureRequestUpdateRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *SignatureRequestUpdateRequest) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SignatureRequestUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SignatureRequestUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SignatureRequestUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignatureRequestUpdateRequest) GetExpiresAt() int32 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignatureRequestUpdateRequest) GetExpiresAtOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SignatureRequestUpdateRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt32 and assigns it to the ExpiresAt field.
func (o *SignatureRequestUpdateRequest) SetExpiresAt(v int32) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *SignatureRequestUpdateRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *SignatureRequestUpdateRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o SignatureRequestUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signature_id"] = o.SignatureId
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureRequestUpdateRequest struct {
	value *SignatureRequestUpdateRequest
	isSet bool
}

func (v NullableSignatureRequestUpdateRequest) Get() *SignatureRequestUpdateRequest {
	return v.value
}

func (v *NullableSignatureRequestUpdateRequest) Set(val *SignatureRequestUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestUpdateRequest(val *SignatureRequestUpdateRequest) *NullableSignatureRequestUpdateRequest {
	return &NullableSignatureRequestUpdateRequest{value: val, isSet: true}
}

func (v NullableSignatureRequestUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


