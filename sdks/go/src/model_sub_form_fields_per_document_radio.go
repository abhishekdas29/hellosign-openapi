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

// SubFormFieldsPerDocumentRadio This class extends `SubFormFieldsPerDocumentBase`.
type SubFormFieldsPerDocumentRadio struct {
	SubFormFieldsPerDocumentBase
	// An input field for radios. Use the `SubFormFieldsPerDocumentRadio` class.
	Type string `json:"type"`
	// String referencing group defined in `form_field_groups` parameter.
	Group string `json:"group"`
	// `true` for checking the radio field by default, otherwise `false`. Only one radio field per group can be `true`.
	IsChecked bool `json:"is_checked"`
}

// NewSubFormFieldsPerDocumentRadio instantiates a new SubFormFieldsPerDocumentRadio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentRadio(type_ string, group string, isChecked bool, documentIndex int32, apiId string, height int32, required bool, signer string, width int32, x int32, y int32) *SubFormFieldsPerDocumentRadio {
	this := SubFormFieldsPerDocumentRadio{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	this.Group = group
	this.IsChecked = isChecked
	return &this
}

// NewSubFormFieldsPerDocumentRadioWithDefaults instantiates a new SubFormFieldsPerDocumentRadio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentRadioWithDefaults() *SubFormFieldsPerDocumentRadio {
	this := SubFormFieldsPerDocumentRadio{}
	var type_ string = "radio"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentRadio) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentRadio) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentRadio) SetType(v string) {
	o.Type = v
}

// GetGroup returns the Group field value
func (o *SubFormFieldsPerDocumentRadio) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentRadio) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SubFormFieldsPerDocumentRadio) SetGroup(v string) {
	o.Group = v
}

// GetIsChecked returns the IsChecked field value
func (o *SubFormFieldsPerDocumentRadio) GetIsChecked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsChecked
}

// GetIsCheckedOk returns a tuple with the IsChecked field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentRadio) GetIsCheckedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsChecked, true
}

// SetIsChecked sets field value
func (o *SubFormFieldsPerDocumentRadio) SetIsChecked(v bool) {
	o.IsChecked = v
}

func (o SubFormFieldsPerDocumentRadio) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["is_checked"] = o.IsChecked
	}
	return json.Marshal(toSerialize)
}

type NullableSubFormFieldsPerDocumentRadio struct {
	value *SubFormFieldsPerDocumentRadio
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentRadio) Get() *SubFormFieldsPerDocumentRadio {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentRadio) Set(val *SubFormFieldsPerDocumentRadio) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentRadio) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentRadio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentRadio(val *SubFormFieldsPerDocumentRadio) *NullableSubFormFieldsPerDocumentRadio {
	return &NullableSubFormFieldsPerDocumentRadio{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentRadio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentRadio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


