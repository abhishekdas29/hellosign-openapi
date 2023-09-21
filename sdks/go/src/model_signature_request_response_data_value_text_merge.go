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

// SignatureRequestResponseDataValueTextMerge struct for SignatureRequestResponseDataValueTextMerge
type SignatureRequestResponseDataValueTextMerge struct {
	// A text field that has default text set by the api
	Type *string `json:"type,omitempty"`
	// The value of the form field.
	Value *string `json:"value,omitempty"`
}

// NewSignatureRequestResponseDataValueTextMerge instantiates a new SignatureRequestResponseDataValueTextMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestResponseDataValueTextMerge() *SignatureRequestResponseDataValueTextMerge {
	this := SignatureRequestResponseDataValueTextMerge{}
	var type_ string = "text-merge"
	this.Type = &type_
	return &this
}

// NewSignatureRequestResponseDataValueTextMergeWithDefaults instantiates a new SignatureRequestResponseDataValueTextMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestResponseDataValueTextMergeWithDefaults() *SignatureRequestResponseDataValueTextMerge {
	this := SignatureRequestResponseDataValueTextMerge{}
	var type_ string = "text-merge"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SignatureRequestResponseDataValueTextMerge) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseDataValueTextMerge) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SignatureRequestResponseDataValueTextMerge) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SignatureRequestResponseDataValueTextMerge) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SignatureRequestResponseDataValueTextMerge) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseDataValueTextMerge) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SignatureRequestResponseDataValueTextMerge) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SignatureRequestResponseDataValueTextMerge) SetValue(v string) {
	o.Value = &v
}

func (o SignatureRequestResponseDataValueTextMerge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureRequestResponseDataValueTextMerge struct {
	value *SignatureRequestResponseDataValueTextMerge
	isSet bool
}

func (v NullableSignatureRequestResponseDataValueTextMerge) Get() *SignatureRequestResponseDataValueTextMerge {
	return v.value
}

func (v *NullableSignatureRequestResponseDataValueTextMerge) Set(val *SignatureRequestResponseDataValueTextMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseDataValueTextMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseDataValueTextMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseDataValueTextMerge(val *SignatureRequestResponseDataValueTextMerge) *NullableSignatureRequestResponseDataValueTextMerge {
	return &NullableSignatureRequestResponseDataValueTextMerge{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseDataValueTextMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseDataValueTextMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


