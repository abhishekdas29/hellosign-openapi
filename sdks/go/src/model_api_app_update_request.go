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
	"os"
)

// ApiAppUpdateRequest struct for ApiAppUpdateRequest
type ApiAppUpdateRequest struct {
	// The URL at which the API App should receive event callbacks.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// An image file to use as a custom logo in embedded contexts. (Only applies to some API plans)
	CustomLogoFile **os.File `json:"custom_logo_file,omitempty"`
	// The domain names the ApiApp will be associated with.
	Domains *[]string `json:"domains,omitempty"`
	// The name you want to assign to the ApiApp.
	Name *string `json:"name,omitempty"`
	Oauth *SubOAuth `json:"oauth,omitempty"`
	Options *SubOptions `json:"options,omitempty"`
	WhiteLabelingOptions *SubWhiteLabelingOptions `json:"white_labeling_options,omitempty"`
}

// NewApiAppUpdateRequest instantiates a new ApiAppUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAppUpdateRequest() *ApiAppUpdateRequest {
	this := ApiAppUpdateRequest{}
	return &this
}

// NewApiAppUpdateRequestWithDefaults instantiates a new ApiAppUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAppUpdateRequestWithDefaults() *ApiAppUpdateRequest {
	this := ApiAppUpdateRequest{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *ApiAppUpdateRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetCustomLogoFile returns the CustomLogoFile field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetCustomLogoFile() *os.File {
	if o == nil || o.CustomLogoFile == nil {
		var ret *os.File
		return ret
	}
	return *o.CustomLogoFile
}

// GetCustomLogoFileOk returns a tuple with the CustomLogoFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetCustomLogoFileOk() (**os.File, bool) {
	if o == nil || o.CustomLogoFile == nil {
		return nil, false
	}
	return o.CustomLogoFile, true
}

// HasCustomLogoFile returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasCustomLogoFile() bool {
	if o != nil && o.CustomLogoFile != nil {
		return true
	}

	return false
}

// SetCustomLogoFile gets a reference to the given *os.File and assigns it to the CustomLogoFile field.
func (o *ApiAppUpdateRequest) SetCustomLogoFile(v *os.File) {
	o.CustomLogoFile = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetDomainsOk() (*[]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *ApiAppUpdateRequest) SetDomains(v []string) {
	o.Domains = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAppUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetOauth() SubOAuth {
	if o == nil || o.Oauth == nil {
		var ret SubOAuth
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetOauthOk() (*SubOAuth, bool) {
	if o == nil || o.Oauth == nil {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasOauth() bool {
	if o != nil && o.Oauth != nil {
		return true
	}

	return false
}

// SetOauth gets a reference to the given SubOAuth and assigns it to the Oauth field.
func (o *ApiAppUpdateRequest) SetOauth(v SubOAuth) {
	o.Oauth = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetOptions() SubOptions {
	if o == nil || o.Options == nil {
		var ret SubOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetOptionsOk() (*SubOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SubOptions and assigns it to the Options field.
func (o *ApiAppUpdateRequest) SetOptions(v SubOptions) {
	o.Options = &v
}

// GetWhiteLabelingOptions returns the WhiteLabelingOptions field value if set, zero value otherwise.
func (o *ApiAppUpdateRequest) GetWhiteLabelingOptions() SubWhiteLabelingOptions {
	if o == nil || o.WhiteLabelingOptions == nil {
		var ret SubWhiteLabelingOptions
		return ret
	}
	return *o.WhiteLabelingOptions
}

// GetWhiteLabelingOptionsOk returns a tuple with the WhiteLabelingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppUpdateRequest) GetWhiteLabelingOptionsOk() (*SubWhiteLabelingOptions, bool) {
	if o == nil || o.WhiteLabelingOptions == nil {
		return nil, false
	}
	return o.WhiteLabelingOptions, true
}

// HasWhiteLabelingOptions returns a boolean if a field has been set.
func (o *ApiAppUpdateRequest) HasWhiteLabelingOptions() bool {
	if o != nil && o.WhiteLabelingOptions != nil {
		return true
	}

	return false
}

// SetWhiteLabelingOptions gets a reference to the given SubWhiteLabelingOptions and assigns it to the WhiteLabelingOptions field.
func (o *ApiAppUpdateRequest) SetWhiteLabelingOptions(v SubWhiteLabelingOptions) {
	o.WhiteLabelingOptions = &v
}

func (o ApiAppUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.CustomLogoFile != nil {
		toSerialize["custom_logo_file"] = o.CustomLogoFile
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Oauth != nil {
		toSerialize["oauth"] = o.Oauth
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.WhiteLabelingOptions != nil {
		toSerialize["white_labeling_options"] = o.WhiteLabelingOptions
	}
	return json.Marshal(toSerialize)
}

type NullableApiAppUpdateRequest struct {
	value *ApiAppUpdateRequest
	isSet bool
}

func (v NullableApiAppUpdateRequest) Get() *ApiAppUpdateRequest {
	return v.value
}

func (v *NullableApiAppUpdateRequest) Set(val *ApiAppUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAppUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAppUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAppUpdateRequest(val *ApiAppUpdateRequest) *NullableApiAppUpdateRequest {
	return &NullableApiAppUpdateRequest{value: val, isSet: true}
}

func (v NullableApiAppUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAppUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

