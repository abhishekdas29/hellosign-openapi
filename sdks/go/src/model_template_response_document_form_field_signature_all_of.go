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

// TemplateResponseDocumentFormFieldSignatureAllOf struct for TemplateResponseDocumentFormFieldSignatureAllOf
type TemplateResponseDocumentFormFieldSignatureAllOf struct {
	// The type of this form field. See [field types](/api/reference/constants/#field-types).  * Text Field uses `TemplateResponseDocumentFormFieldText` * Dropdown Field uses `TemplateResponseDocumentFormFieldDropdown` * Hyperlink Field uses `TemplateResponseDocumentFormFieldHyperlink` * Checkbox Field uses `TemplateResponseDocumentFormFieldCheckbox` * Radio Field uses `TemplateResponseDocumentFormFieldRadio` * Signature Field uses `TemplateResponseDocumentFormFieldSignature` * Date Signed Field uses `TemplateResponseDocumentFormFieldDateSigned` * Initials Field uses `TemplateResponseDocumentFormFieldInitials`
	Type *string `json:"type,omitempty"`
}

// NewTemplateResponseDocumentFormFieldSignatureAllOf instantiates a new TemplateResponseDocumentFormFieldSignatureAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentFormFieldSignatureAllOf() *TemplateResponseDocumentFormFieldSignatureAllOf {
	this := TemplateResponseDocumentFormFieldSignatureAllOf{}
	var type_ string = "signature"
	this.Type = &type_
	return &this
}

// NewTemplateResponseDocumentFormFieldSignatureAllOfWithDefaults instantiates a new TemplateResponseDocumentFormFieldSignatureAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentFormFieldSignatureAllOfWithDefaults() *TemplateResponseDocumentFormFieldSignatureAllOf {
	this := TemplateResponseDocumentFormFieldSignatureAllOf{}
	var type_ string = "signature"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFormFieldSignatureAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldSignatureAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFormFieldSignatureAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplateResponseDocumentFormFieldSignatureAllOf) SetType(v string) {
	o.Type = &v
}

func (o TemplateResponseDocumentFormFieldSignatureAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateResponseDocumentFormFieldSignatureAllOf struct {
	value *TemplateResponseDocumentFormFieldSignatureAllOf
	isSet bool
}

func (v NullableTemplateResponseDocumentFormFieldSignatureAllOf) Get() *TemplateResponseDocumentFormFieldSignatureAllOf {
	return v.value
}

func (v *NullableTemplateResponseDocumentFormFieldSignatureAllOf) Set(val *TemplateResponseDocumentFormFieldSignatureAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentFormFieldSignatureAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentFormFieldSignatureAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentFormFieldSignatureAllOf(val *TemplateResponseDocumentFormFieldSignatureAllOf) *NullableTemplateResponseDocumentFormFieldSignatureAllOf {
	return &NullableTemplateResponseDocumentFormFieldSignatureAllOf{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentFormFieldSignatureAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentFormFieldSignatureAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


