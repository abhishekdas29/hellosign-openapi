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

// ListInfoResponse Contains pagination information about the data returned.
type ListInfoResponse struct {
	// Total number of pages available.
	NumPages *int32 `json:"num_pages,omitempty"`
	// Total number of objects available.
	NumResults NullableInt32 `json:"num_results,omitempty"`
	// Number of the page being returned.
	Page *int32 `json:"page,omitempty"`
	// Objects returned per page.
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewListInfoResponse instantiates a new ListInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInfoResponse() *ListInfoResponse {
	this := ListInfoResponse{}
	return &this
}

// NewListInfoResponseWithDefaults instantiates a new ListInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInfoResponseWithDefaults() *ListInfoResponse {
	this := ListInfoResponse{}
	return &this
}

// GetNumPages returns the NumPages field value if set, zero value otherwise.
func (o *ListInfoResponse) GetNumPages() int32 {
	if o == nil || o.NumPages == nil {
		var ret int32
		return ret
	}
	return *o.NumPages
}

// GetNumPagesOk returns a tuple with the NumPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInfoResponse) GetNumPagesOk() (*int32, bool) {
	if o == nil || o.NumPages == nil {
		return nil, false
	}
	return o.NumPages, true
}

// HasNumPages returns a boolean if a field has been set.
func (o *ListInfoResponse) HasNumPages() bool {
	if o != nil && o.NumPages != nil {
		return true
	}

	return false
}

// SetNumPages gets a reference to the given int32 and assigns it to the NumPages field.
func (o *ListInfoResponse) SetNumPages(v int32) {
	o.NumPages = &v
}

// GetNumResults returns the NumResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListInfoResponse) GetNumResults() int32 {
	if o == nil || o.NumResults.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumResults.Get()
}

// GetNumResultsOk returns a tuple with the NumResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListInfoResponse) GetNumResultsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumResults.Get(), o.NumResults.IsSet()
}

// HasNumResults returns a boolean if a field has been set.
func (o *ListInfoResponse) HasNumResults() bool {
	if o != nil && o.NumResults.IsSet() {
		return true
	}

	return false
}

// SetNumResults gets a reference to the given NullableInt32 and assigns it to the NumResults field.
func (o *ListInfoResponse) SetNumResults(v int32) {
	o.NumResults.Set(&v)
}
// SetNumResultsNil sets the value for NumResults to be an explicit nil
func (o *ListInfoResponse) SetNumResultsNil() {
	o.NumResults.Set(nil)
}

// UnsetNumResults ensures that no value is present for NumResults, not even an explicit nil
func (o *ListInfoResponse) UnsetNumResults() {
	o.NumResults.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListInfoResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInfoResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListInfoResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ListInfoResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListInfoResponse) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInfoResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListInfoResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListInfoResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o ListInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumPages != nil {
		toSerialize["num_pages"] = o.NumPages
	}
	if o.NumResults.IsSet() {
		toSerialize["num_results"] = o.NumResults.Get()
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableListInfoResponse struct {
	value *ListInfoResponse
	isSet bool
}

func (v NullableListInfoResponse) Get() *ListInfoResponse {
	return v.value
}

func (v *NullableListInfoResponse) Set(val *ListInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInfoResponse(val *ListInfoResponse) *NullableListInfoResponse {
	return &NullableListInfoResponse{value: val, isSet: true}
}

func (v NullableListInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


