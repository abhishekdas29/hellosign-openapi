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

// AccountResponseQuotas Details concerning remaining monthly quotas.
type AccountResponseQuotas struct {
	// API signature requests remaining.
	ApiSignatureRequestsLeft NullableInt32 `json:"api_signature_requests_left,omitempty"`
	// Signature requests remaining.
	DocumentsLeft NullableInt32 `json:"documents_left,omitempty"`
	// Total API templates allowed.
	TemplatesTotal NullableInt32 `json:"templates_total,omitempty"`
	// API templates remaining.
	TemplatesLeft NullableInt32 `json:"templates_left,omitempty"`
	// SMS verifications  remaining.
	SmsVerificationsLeft NullableInt32 `json:"sms_verifications_left,omitempty"`
}

// NewAccountResponseQuotas instantiates a new AccountResponseQuotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponseQuotas() *AccountResponseQuotas {
	this := AccountResponseQuotas{}
	return &this
}

// NewAccountResponseQuotasWithDefaults instantiates a new AccountResponseQuotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponseQuotasWithDefaults() *AccountResponseQuotas {
	this := AccountResponseQuotas{}
	return &this
}

// GetApiSignatureRequestsLeft returns the ApiSignatureRequestsLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponseQuotas) GetApiSignatureRequestsLeft() int32 {
	if o == nil || o.ApiSignatureRequestsLeft.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ApiSignatureRequestsLeft.Get()
}

// GetApiSignatureRequestsLeftOk returns a tuple with the ApiSignatureRequestsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponseQuotas) GetApiSignatureRequestsLeftOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApiSignatureRequestsLeft.Get(), o.ApiSignatureRequestsLeft.IsSet()
}

// HasApiSignatureRequestsLeft returns a boolean if a field has been set.
func (o *AccountResponseQuotas) HasApiSignatureRequestsLeft() bool {
	if o != nil && o.ApiSignatureRequestsLeft.IsSet() {
		return true
	}

	return false
}

// SetApiSignatureRequestsLeft gets a reference to the given NullableInt32 and assigns it to the ApiSignatureRequestsLeft field.
func (o *AccountResponseQuotas) SetApiSignatureRequestsLeft(v int32) {
	o.ApiSignatureRequestsLeft.Set(&v)
}
// SetApiSignatureRequestsLeftNil sets the value for ApiSignatureRequestsLeft to be an explicit nil
func (o *AccountResponseQuotas) SetApiSignatureRequestsLeftNil() {
	o.ApiSignatureRequestsLeft.Set(nil)
}

// UnsetApiSignatureRequestsLeft ensures that no value is present for ApiSignatureRequestsLeft, not even an explicit nil
func (o *AccountResponseQuotas) UnsetApiSignatureRequestsLeft() {
	o.ApiSignatureRequestsLeft.Unset()
}

// GetDocumentsLeft returns the DocumentsLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponseQuotas) GetDocumentsLeft() int32 {
	if o == nil || o.DocumentsLeft.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DocumentsLeft.Get()
}

// GetDocumentsLeftOk returns a tuple with the DocumentsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponseQuotas) GetDocumentsLeftOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentsLeft.Get(), o.DocumentsLeft.IsSet()
}

// HasDocumentsLeft returns a boolean if a field has been set.
func (o *AccountResponseQuotas) HasDocumentsLeft() bool {
	if o != nil && o.DocumentsLeft.IsSet() {
		return true
	}

	return false
}

// SetDocumentsLeft gets a reference to the given NullableInt32 and assigns it to the DocumentsLeft field.
func (o *AccountResponseQuotas) SetDocumentsLeft(v int32) {
	o.DocumentsLeft.Set(&v)
}
// SetDocumentsLeftNil sets the value for DocumentsLeft to be an explicit nil
func (o *AccountResponseQuotas) SetDocumentsLeftNil() {
	o.DocumentsLeft.Set(nil)
}

// UnsetDocumentsLeft ensures that no value is present for DocumentsLeft, not even an explicit nil
func (o *AccountResponseQuotas) UnsetDocumentsLeft() {
	o.DocumentsLeft.Unset()
}

// GetTemplatesTotal returns the TemplatesTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponseQuotas) GetTemplatesTotal() int32 {
	if o == nil || o.TemplatesTotal.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TemplatesTotal.Get()
}

// GetTemplatesTotalOk returns a tuple with the TemplatesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponseQuotas) GetTemplatesTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplatesTotal.Get(), o.TemplatesTotal.IsSet()
}

// HasTemplatesTotal returns a boolean if a field has been set.
func (o *AccountResponseQuotas) HasTemplatesTotal() bool {
	if o != nil && o.TemplatesTotal.IsSet() {
		return true
	}

	return false
}

// SetTemplatesTotal gets a reference to the given NullableInt32 and assigns it to the TemplatesTotal field.
func (o *AccountResponseQuotas) SetTemplatesTotal(v int32) {
	o.TemplatesTotal.Set(&v)
}
// SetTemplatesTotalNil sets the value for TemplatesTotal to be an explicit nil
func (o *AccountResponseQuotas) SetTemplatesTotalNil() {
	o.TemplatesTotal.Set(nil)
}

// UnsetTemplatesTotal ensures that no value is present for TemplatesTotal, not even an explicit nil
func (o *AccountResponseQuotas) UnsetTemplatesTotal() {
	o.TemplatesTotal.Unset()
}

// GetTemplatesLeft returns the TemplatesLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponseQuotas) GetTemplatesLeft() int32 {
	if o == nil || o.TemplatesLeft.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TemplatesLeft.Get()
}

// GetTemplatesLeftOk returns a tuple with the TemplatesLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponseQuotas) GetTemplatesLeftOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplatesLeft.Get(), o.TemplatesLeft.IsSet()
}

// HasTemplatesLeft returns a boolean if a field has been set.
func (o *AccountResponseQuotas) HasTemplatesLeft() bool {
	if o != nil && o.TemplatesLeft.IsSet() {
		return true
	}

	return false
}

// SetTemplatesLeft gets a reference to the given NullableInt32 and assigns it to the TemplatesLeft field.
func (o *AccountResponseQuotas) SetTemplatesLeft(v int32) {
	o.TemplatesLeft.Set(&v)
}
// SetTemplatesLeftNil sets the value for TemplatesLeft to be an explicit nil
func (o *AccountResponseQuotas) SetTemplatesLeftNil() {
	o.TemplatesLeft.Set(nil)
}

// UnsetTemplatesLeft ensures that no value is present for TemplatesLeft, not even an explicit nil
func (o *AccountResponseQuotas) UnsetTemplatesLeft() {
	o.TemplatesLeft.Unset()
}

// GetSmsVerificationsLeft returns the SmsVerificationsLeft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponseQuotas) GetSmsVerificationsLeft() int32 {
	if o == nil || o.SmsVerificationsLeft.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SmsVerificationsLeft.Get()
}

// GetSmsVerificationsLeftOk returns a tuple with the SmsVerificationsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponseQuotas) GetSmsVerificationsLeftOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SmsVerificationsLeft.Get(), o.SmsVerificationsLeft.IsSet()
}

// HasSmsVerificationsLeft returns a boolean if a field has been set.
func (o *AccountResponseQuotas) HasSmsVerificationsLeft() bool {
	if o != nil && o.SmsVerificationsLeft.IsSet() {
		return true
	}

	return false
}

// SetSmsVerificationsLeft gets a reference to the given NullableInt32 and assigns it to the SmsVerificationsLeft field.
func (o *AccountResponseQuotas) SetSmsVerificationsLeft(v int32) {
	o.SmsVerificationsLeft.Set(&v)
}
// SetSmsVerificationsLeftNil sets the value for SmsVerificationsLeft to be an explicit nil
func (o *AccountResponseQuotas) SetSmsVerificationsLeftNil() {
	o.SmsVerificationsLeft.Set(nil)
}

// UnsetSmsVerificationsLeft ensures that no value is present for SmsVerificationsLeft, not even an explicit nil
func (o *AccountResponseQuotas) UnsetSmsVerificationsLeft() {
	o.SmsVerificationsLeft.Unset()
}

func (o AccountResponseQuotas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiSignatureRequestsLeft.IsSet() {
		toSerialize["api_signature_requests_left"] = o.ApiSignatureRequestsLeft.Get()
	}
	if o.DocumentsLeft.IsSet() {
		toSerialize["documents_left"] = o.DocumentsLeft.Get()
	}
	if o.TemplatesTotal.IsSet() {
		toSerialize["templates_total"] = o.TemplatesTotal.Get()
	}
	if o.TemplatesLeft.IsSet() {
		toSerialize["templates_left"] = o.TemplatesLeft.Get()
	}
	if o.SmsVerificationsLeft.IsSet() {
		toSerialize["sms_verifications_left"] = o.SmsVerificationsLeft.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountResponseQuotas struct {
	value *AccountResponseQuotas
	isSet bool
}

func (v NullableAccountResponseQuotas) Get() *AccountResponseQuotas {
	return v.value
}

func (v *NullableAccountResponseQuotas) Set(val *AccountResponseQuotas) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponseQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponseQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponseQuotas(val *AccountResponseQuotas) *NullableAccountResponseQuotas {
	return &NullableAccountResponseQuotas{value: val, isSet: true}
}

func (v NullableAccountResponseQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponseQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


