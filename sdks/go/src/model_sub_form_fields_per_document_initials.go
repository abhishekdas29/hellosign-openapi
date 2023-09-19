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

// SubFormFieldsPerDocumentInitials This class extends `SubFormFieldsPerDocumentBase`.
type SubFormFieldsPerDocumentInitials struct {
	SubFormFieldsPerDocumentBase
	// An input field for initials. Use the `SubFormFieldsPerDocumentInitials` class.
	Type string `json:"type"`
}

// NewSubFormFieldsPerDocumentInitials instantiates a new SubFormFieldsPerDocumentInitials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentInitials(type_ string, documentIndex int32, apiId string, height int32, required bool, signer string, width int32, x int32, y int32) *SubFormFieldsPerDocumentInitials {
	this := SubFormFieldsPerDocumentInitials{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewSubFormFieldsPerDocumentInitialsWithDefaults instantiates a new SubFormFieldsPerDocumentInitials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentInitialsWithDefaults() *SubFormFieldsPerDocumentInitials {
	this := SubFormFieldsPerDocumentInitials{}
	var type_ string = "initials"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentInitials) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentInitials) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentInitials) SetType(v string) {
	o.Type = v
}

func (o SubFormFieldsPerDocumentInitials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSubFormFieldsPerDocumentBase, errSubFormFieldsPerDocumentBase := json.Marshal(o.SubFormFieldsPerDocumentBase)
	if errSubFormFieldsPerDocumentBase != nil {
		return []byte{}, errSubFormFieldsPerDocumentBase
	}
	errSubFormFieldsPerDocumentBase = json.Unmarshal([]byte(serializedSubFormFieldsPerDocumentBase), &toSerialize)
	if errSubFormFieldsPerDocumentBase != nil {
		return []byte{}, errSubFormFieldsPerDocumentBase
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSubFormFieldsPerDocumentInitials struct {
	value *SubFormFieldsPerDocumentInitials
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentInitials) Get() *SubFormFieldsPerDocumentInitials {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentInitials) Set(val *SubFormFieldsPerDocumentInitials) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentInitials) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentInitials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentInitials(val *SubFormFieldsPerDocumentInitials) *NullableSubFormFieldsPerDocumentInitials {
	return &NullableSubFormFieldsPerDocumentInitials{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentInitials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentInitials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

