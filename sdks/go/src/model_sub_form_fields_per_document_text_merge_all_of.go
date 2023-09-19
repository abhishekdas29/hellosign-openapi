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

// SubFormFieldsPerDocumentTextMergeAllOf struct for SubFormFieldsPerDocumentTextMergeAllOf
type SubFormFieldsPerDocumentTextMergeAllOf struct {
	// A text field that has default text set using pre-filled data. Use the `SubFormFieldsPerDocumentTextMerge` class.
	Type string `json:"type"`
	// Font family for the field.
	FontFamily *string `json:"font_family,omitempty"`
	// The initial px font size for the field contents. Can be any integer value between `7` and `49`.  **NOTE**: Font size may be reduced during processing in order to fit the contents within the dimensions of the field.
	FontSize *int32 `json:"font_size,omitempty"`
}

// NewSubFormFieldsPerDocumentTextMergeAllOf instantiates a new SubFormFieldsPerDocumentTextMergeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentTextMergeAllOf(type_ string) *SubFormFieldsPerDocumentTextMergeAllOf {
	this := SubFormFieldsPerDocumentTextMergeAllOf{}
	this.Type = type_
	return &this
}

// NewSubFormFieldsPerDocumentTextMergeAllOfWithDefaults instantiates a new SubFormFieldsPerDocumentTextMergeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentTextMergeAllOfWithDefaults() *SubFormFieldsPerDocumentTextMergeAllOf {
	this := SubFormFieldsPerDocumentTextMergeAllOf{}
	var type_ string = "text-merge"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentTextMergeAllOf) SetType(v string) {
	o.Type = v
}

// GetFontFamily returns the FontFamily field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetFontFamily() string {
	if o == nil || o.FontFamily == nil {
		var ret string
		return ret
	}
	return *o.FontFamily
}

// GetFontFamilyOk returns a tuple with the FontFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetFontFamilyOk() (*string, bool) {
	if o == nil || o.FontFamily == nil {
		return nil, false
	}
	return o.FontFamily, true
}

// HasFontFamily returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) HasFontFamily() bool {
	if o != nil && o.FontFamily != nil {
		return true
	}

	return false
}

// SetFontFamily gets a reference to the given string and assigns it to the FontFamily field.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) SetFontFamily(v string) {
	o.FontFamily = &v
}

// GetFontSize returns the FontSize field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetFontSize() int32 {
	if o == nil || o.FontSize == nil {
		var ret int32
		return ret
	}
	return *o.FontSize
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) GetFontSizeOk() (*int32, bool) {
	if o == nil || o.FontSize == nil {
		return nil, false
	}
	return o.FontSize, true
}

// HasFontSize returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) HasFontSize() bool {
	if o != nil && o.FontSize != nil {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given int32 and assigns it to the FontSize field.
func (o *SubFormFieldsPerDocumentTextMergeAllOf) SetFontSize(v int32) {
	o.FontSize = &v
}

func (o SubFormFieldsPerDocumentTextMergeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FontFamily != nil {
		toSerialize["font_family"] = o.FontFamily
	}
	if o.FontSize != nil {
		toSerialize["font_size"] = o.FontSize
	}
	return json.Marshal(toSerialize)
}

type NullableSubFormFieldsPerDocumentTextMergeAllOf struct {
	value *SubFormFieldsPerDocumentTextMergeAllOf
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentTextMergeAllOf) Get() *SubFormFieldsPerDocumentTextMergeAllOf {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentTextMergeAllOf) Set(val *SubFormFieldsPerDocumentTextMergeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentTextMergeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentTextMergeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentTextMergeAllOf(val *SubFormFieldsPerDocumentTextMergeAllOf) *NullableSubFormFieldsPerDocumentTextMergeAllOf {
	return &NullableSubFormFieldsPerDocumentTextMergeAllOf{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentTextMergeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentTextMergeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

