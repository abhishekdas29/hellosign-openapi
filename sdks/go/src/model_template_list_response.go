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

// TemplateListResponse struct for TemplateListResponse
type TemplateListResponse struct {
	// List of templates that the API caller has access to.
	Templates *[]TemplateResponse `json:"templates,omitempty"`
	ListInfo *ListInfoResponse `json:"list_info,omitempty"`
	// A list of warnings.
	Warnings *[]WarningResponse `json:"warnings,omitempty"`
}

// NewTemplateListResponse instantiates a new TemplateListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateListResponse() *TemplateListResponse {
	this := TemplateListResponse{}
	return &this
}

// NewTemplateListResponseWithDefaults instantiates a new TemplateListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateListResponseWithDefaults() *TemplateListResponse {
	this := TemplateListResponse{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *TemplateListResponse) GetTemplates() []TemplateResponse {
	if o == nil || o.Templates == nil {
		var ret []TemplateResponse
		return ret
	}
	return *o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateListResponse) GetTemplatesOk() (*[]TemplateResponse, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *TemplateListResponse) HasTemplates() bool {
	if o != nil && o.Templates != nil {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []TemplateResponse and assigns it to the Templates field.
func (o *TemplateListResponse) SetTemplates(v []TemplateResponse) {
	o.Templates = &v
}

// GetListInfo returns the ListInfo field value if set, zero value otherwise.
func (o *TemplateListResponse) GetListInfo() ListInfoResponse {
	if o == nil || o.ListInfo == nil {
		var ret ListInfoResponse
		return ret
	}
	return *o.ListInfo
}

// GetListInfoOk returns a tuple with the ListInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateListResponse) GetListInfoOk() (*ListInfoResponse, bool) {
	if o == nil || o.ListInfo == nil {
		return nil, false
	}
	return o.ListInfo, true
}

// HasListInfo returns a boolean if a field has been set.
func (o *TemplateListResponse) HasListInfo() bool {
	if o != nil && o.ListInfo != nil {
		return true
	}

	return false
}

// SetListInfo gets a reference to the given ListInfoResponse and assigns it to the ListInfo field.
func (o *TemplateListResponse) SetListInfo(v ListInfoResponse) {
	o.ListInfo = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TemplateListResponse) GetWarnings() []WarningResponse {
	if o == nil || o.Warnings == nil {
		var ret []WarningResponse
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateListResponse) GetWarningsOk() (*[]WarningResponse, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TemplateListResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
func (o *TemplateListResponse) SetWarnings(v []WarningResponse) {
	o.Warnings = &v
}

func (o TemplateListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.ListInfo != nil {
		toSerialize["list_info"] = o.ListInfo
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateListResponse struct {
	value *TemplateListResponse
	isSet bool
}

func (v NullableTemplateListResponse) Get() *TemplateListResponse {
	return v.value
}

func (v *NullableTemplateListResponse) Set(val *TemplateListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateListResponse(val *TemplateListResponse) *NullableTemplateListResponse {
	return &NullableTemplateListResponse{value: val, isSet: true}
}

func (v NullableTemplateListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

