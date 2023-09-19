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

// TemplateResponseDocumentStaticFieldHyperlinkAllOf struct for TemplateResponseDocumentStaticFieldHyperlinkAllOf
type TemplateResponseDocumentStaticFieldHyperlinkAllOf struct {
	// The type of this static field. See [field types](/api/reference/constants/#field-types).  * Text Field uses `TemplateResponseDocumentStaticFieldText` * Dropdown Field uses `TemplateResponseDocumentStaticFieldDropdown` * Hyperlink Field uses `TemplateResponseDocumentStaticFieldHyperlink` * Checkbox Field uses `TemplateResponseDocumentStaticFieldCheckbox` * Radio Field uses `TemplateResponseDocumentStaticFieldRadio` * Signature Field uses `TemplateResponseDocumentStaticFieldSignature` * Date Signed Field uses `TemplateResponseDocumentStaticFieldDateSigned` * Initials Field uses `TemplateResponseDocumentStaticFieldInitials`
	Type *string `json:"type,omitempty"`
}

// NewTemplateResponseDocumentStaticFieldHyperlinkAllOf instantiates a new TemplateResponseDocumentStaticFieldHyperlinkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentStaticFieldHyperlinkAllOf() *TemplateResponseDocumentStaticFieldHyperlinkAllOf {
	this := TemplateResponseDocumentStaticFieldHyperlinkAllOf{}
	var type_ string = "hyperlink"
	this.Type = &type_
	return &this
}

// NewTemplateResponseDocumentStaticFieldHyperlinkAllOfWithDefaults instantiates a new TemplateResponseDocumentStaticFieldHyperlinkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentStaticFieldHyperlinkAllOfWithDefaults() *TemplateResponseDocumentStaticFieldHyperlinkAllOf {
	this := TemplateResponseDocumentStaticFieldHyperlinkAllOf{}
	var type_ string = "hyperlink"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldHyperlinkAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldHyperlinkAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldHyperlinkAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplateResponseDocumentStaticFieldHyperlinkAllOf) SetType(v string) {
	o.Type = &v
}

func (o TemplateResponseDocumentStaticFieldHyperlinkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf struct {
	value *TemplateResponseDocumentStaticFieldHyperlinkAllOf
	isSet bool
}

func (v NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) Get() *TemplateResponseDocumentStaticFieldHyperlinkAllOf {
	return v.value
}

func (v *NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) Set(val *TemplateResponseDocumentStaticFieldHyperlinkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentStaticFieldHyperlinkAllOf(val *TemplateResponseDocumentStaticFieldHyperlinkAllOf) *NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf {
	return &NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentStaticFieldHyperlinkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

