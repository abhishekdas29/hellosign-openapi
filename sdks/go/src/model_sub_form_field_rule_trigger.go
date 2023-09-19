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

// SubFormFieldRuleTrigger struct for SubFormFieldRuleTrigger
type SubFormFieldRuleTrigger struct {
	// Must reference the `api_id` of an existing field defined within `form_fields_per_document`. Trigger and action fields and groups must belong to the same signer.
	Id string `json:"id"`
	// Different field types allow different `operator` values: - Field type of **text**:   - **is**: exact match   - **not**: not exact match   - **match**: regular expression, without /. Example:     - OK `[a-zA-Z0-9]`     - Not OK `/[a-zA-Z0-9]/` - Field type of **dropdown**:   - **is**: exact match, single value   - **not**: not exact match, single value   - **any**: exact match, array of values.   - **none**: not exact match, array of values. - Field type of **checkbox**:   - **is**: exact match, single value   - **not**: not exact match, single value - Field type of **radio**:   - **is**: exact match, single value   - **not**: not exact match, single value
	Operator string `json:"operator"`
	// **value** or **values** is required, but not both.  The value to match against **operator**.  - When **operator** is one of the following, **value** must be `String`:   - `is`   - `not`   - `match`  Otherwise, - **checkbox**: When **type** of trigger is **checkbox**, **value** must be `0` or `1` - **radio**: When **type** of trigger is **radio**, **value** must be `1`
	Value *string `json:"value,omitempty"`
	// **values** or **value** is required, but not both.  The values to match against **operator** when it is one of the following:  - `any` - `none`
	Values *[]string `json:"values,omitempty"`
}

// NewSubFormFieldRuleTrigger instantiates a new SubFormFieldRuleTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldRuleTrigger(id string, operator string) *SubFormFieldRuleTrigger {
	this := SubFormFieldRuleTrigger{}
	this.Id = id
	this.Operator = operator
	return &this
}

// NewSubFormFieldRuleTriggerWithDefaults instantiates a new SubFormFieldRuleTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldRuleTriggerWithDefaults() *SubFormFieldRuleTrigger {
	this := SubFormFieldRuleTrigger{}
	return &this
}

// GetId returns the Id field value
func (o *SubFormFieldRuleTrigger) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleTrigger) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubFormFieldRuleTrigger) SetId(v string) {
	o.Id = v
}

// GetOperator returns the Operator field value
func (o *SubFormFieldRuleTrigger) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleTrigger) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SubFormFieldRuleTrigger) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SubFormFieldRuleTrigger) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleTrigger) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SubFormFieldRuleTrigger) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SubFormFieldRuleTrigger) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SubFormFieldRuleTrigger) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleTrigger) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SubFormFieldRuleTrigger) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SubFormFieldRuleTrigger) SetValues(v []string) {
	o.Values = &v
}

func (o SubFormFieldRuleTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableSubFormFieldRuleTrigger struct {
	value *SubFormFieldRuleTrigger
	isSet bool
}

func (v NullableSubFormFieldRuleTrigger) Get() *SubFormFieldRuleTrigger {
	return v.value
}

func (v *NullableSubFormFieldRuleTrigger) Set(val *SubFormFieldRuleTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldRuleTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldRuleTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldRuleTrigger(val *SubFormFieldRuleTrigger) *NullableSubFormFieldRuleTrigger {
	return &NullableSubFormFieldRuleTrigger{value: val, isSet: true}
}

func (v NullableSubFormFieldRuleTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldRuleTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

