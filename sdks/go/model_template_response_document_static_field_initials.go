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

// TemplateResponseDocumentStaticFieldInitials This class extends `TemplateResponseDocumentStaticFieldBase`
type TemplateResponseDocumentStaticFieldInitials struct {
	// The type of this static field. See [field types](/api/reference/constants/#field-types).  * Text Field uses `TemplateResponseDocumentStaticFieldText` * Dropdown Field uses `TemplateResponseDocumentStaticFieldDropdown` * Hyperlink Field uses `TemplateResponseDocumentStaticFieldHyperlink` * Checkbox Field uses `TemplateResponseDocumentStaticFieldCheckbox` * Radio Field uses `TemplateResponseDocumentStaticFieldRadio` * Signature Field uses `TemplateResponseDocumentStaticFieldSignature` * Date Signed Field uses `TemplateResponseDocumentStaticFieldDateSigned` * Initials Field uses `TemplateResponseDocumentStaticFieldInitials`
	Type string `json:"type"`
}

// NewTemplateResponseDocumentStaticFieldInitials instantiates a new TemplateResponseDocumentStaticFieldInitials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentStaticFieldInitials(type_ string) *TemplateResponseDocumentStaticFieldInitials {
	this := TemplateResponseDocumentStaticFieldInitials{}
	this.Type = type_
	var signer string = "me_now"
	this.Signer = &signer
	return &this
}

// NewTemplateResponseDocumentStaticFieldInitialsWithDefaults instantiates a new TemplateResponseDocumentStaticFieldInitials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentStaticFieldInitialsWithDefaults() *TemplateResponseDocumentStaticFieldInitials {
	this := TemplateResponseDocumentStaticFieldInitials{}
	var type_ string = "initials"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentStaticFieldInitials) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldInitials) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentStaticFieldInitials) SetType(v string) {
	o.Type = v
}

func (o TemplateResponseDocumentStaticFieldInitials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateResponseDocumentStaticFieldInitials struct {
	value *TemplateResponseDocumentStaticFieldInitials
	isSet bool
}

func (v NullableTemplateResponseDocumentStaticFieldInitials) Get() *TemplateResponseDocumentStaticFieldInitials {
	return v.value
}

func (v *NullableTemplateResponseDocumentStaticFieldInitials) Set(val *TemplateResponseDocumentStaticFieldInitials) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentStaticFieldInitials) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentStaticFieldInitials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentStaticFieldInitials(val *TemplateResponseDocumentStaticFieldInitials) *NullableTemplateResponseDocumentStaticFieldInitials {
	return &NullableTemplateResponseDocumentStaticFieldInitials{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentStaticFieldInitials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentStaticFieldInitials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


