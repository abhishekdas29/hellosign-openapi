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

// ApiAppResponseWhiteLabelingOptions An object with options to customize the app's signer page
type ApiAppResponseWhiteLabelingOptions struct {
	HeaderBackgroundColor *string `json:"header_background_color,omitempty"`
	LegalVersion *string `json:"legal_version,omitempty"`
	LinkColor *string `json:"link_color,omitempty"`
	PageBackgroundColor *string `json:"page_background_color,omitempty"`
	PrimaryButtonColor *string `json:"primary_button_color,omitempty"`
	PrimaryButtonColorHover *string `json:"primary_button_color_hover,omitempty"`
	PrimaryButtonTextColor *string `json:"primary_button_text_color,omitempty"`
	PrimaryButtonTextColorHover *string `json:"primary_button_text_color_hover,omitempty"`
	SecondaryButtonColor *string `json:"secondary_button_color,omitempty"`
	SecondaryButtonColorHover *string `json:"secondary_button_color_hover,omitempty"`
	SecondaryButtonTextColor *string `json:"secondary_button_text_color,omitempty"`
	SecondaryButtonTextColorHover *string `json:"secondary_button_text_color_hover,omitempty"`
	TextColor1 *string `json:"text_color1,omitempty"`
	TextColor2 *string `json:"text_color2,omitempty"`
}

// NewApiAppResponseWhiteLabelingOptions instantiates a new ApiAppResponseWhiteLabelingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAppResponseWhiteLabelingOptions() *ApiAppResponseWhiteLabelingOptions {
	this := ApiAppResponseWhiteLabelingOptions{}
	return &this
}

// NewApiAppResponseWhiteLabelingOptionsWithDefaults instantiates a new ApiAppResponseWhiteLabelingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAppResponseWhiteLabelingOptionsWithDefaults() *ApiAppResponseWhiteLabelingOptions {
	this := ApiAppResponseWhiteLabelingOptions{}
	return &this
}

// GetHeaderBackgroundColor returns the HeaderBackgroundColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetHeaderBackgroundColor() string {
	if o == nil || o.HeaderBackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.HeaderBackgroundColor
}

// GetHeaderBackgroundColorOk returns a tuple with the HeaderBackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetHeaderBackgroundColorOk() (*string, bool) {
	if o == nil || o.HeaderBackgroundColor == nil {
		return nil, false
	}
	return o.HeaderBackgroundColor, true
}

// HasHeaderBackgroundColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasHeaderBackgroundColor() bool {
	if o != nil && o.HeaderBackgroundColor != nil {
		return true
	}

	return false
}

// SetHeaderBackgroundColor gets a reference to the given string and assigns it to the HeaderBackgroundColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetHeaderBackgroundColor(v string) {
	o.HeaderBackgroundColor = &v
}

// GetLegalVersion returns the LegalVersion field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetLegalVersion() string {
	if o == nil || o.LegalVersion == nil {
		var ret string
		return ret
	}
	return *o.LegalVersion
}

// GetLegalVersionOk returns a tuple with the LegalVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetLegalVersionOk() (*string, bool) {
	if o == nil || o.LegalVersion == nil {
		return nil, false
	}
	return o.LegalVersion, true
}

// HasLegalVersion returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasLegalVersion() bool {
	if o != nil && o.LegalVersion != nil {
		return true
	}

	return false
}

// SetLegalVersion gets a reference to the given string and assigns it to the LegalVersion field.
func (o *ApiAppResponseWhiteLabelingOptions) SetLegalVersion(v string) {
	o.LegalVersion = &v
}

// GetLinkColor returns the LinkColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetLinkColor() string {
	if o == nil || o.LinkColor == nil {
		var ret string
		return ret
	}
	return *o.LinkColor
}

// GetLinkColorOk returns a tuple with the LinkColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetLinkColorOk() (*string, bool) {
	if o == nil || o.LinkColor == nil {
		return nil, false
	}
	return o.LinkColor, true
}

// HasLinkColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasLinkColor() bool {
	if o != nil && o.LinkColor != nil {
		return true
	}

	return false
}

// SetLinkColor gets a reference to the given string and assigns it to the LinkColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetLinkColor(v string) {
	o.LinkColor = &v
}

// GetPageBackgroundColor returns the PageBackgroundColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetPageBackgroundColor() string {
	if o == nil || o.PageBackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.PageBackgroundColor
}

// GetPageBackgroundColorOk returns a tuple with the PageBackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetPageBackgroundColorOk() (*string, bool) {
	if o == nil || o.PageBackgroundColor == nil {
		return nil, false
	}
	return o.PageBackgroundColor, true
}

// HasPageBackgroundColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasPageBackgroundColor() bool {
	if o != nil && o.PageBackgroundColor != nil {
		return true
	}

	return false
}

// SetPageBackgroundColor gets a reference to the given string and assigns it to the PageBackgroundColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetPageBackgroundColor(v string) {
	o.PageBackgroundColor = &v
}

// GetPrimaryButtonColor returns the PrimaryButtonColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonColor() string {
	if o == nil || o.PrimaryButtonColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonColor
}

// GetPrimaryButtonColorOk returns a tuple with the PrimaryButtonColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonColor == nil {
		return nil, false
	}
	return o.PrimaryButtonColor, true
}

// HasPrimaryButtonColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasPrimaryButtonColor() bool {
	if o != nil && o.PrimaryButtonColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonColor gets a reference to the given string and assigns it to the PrimaryButtonColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetPrimaryButtonColor(v string) {
	o.PrimaryButtonColor = &v
}

// GetPrimaryButtonColorHover returns the PrimaryButtonColorHover field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonColorHover() string {
	if o == nil || o.PrimaryButtonColorHover == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonColorHover
}

// GetPrimaryButtonColorHoverOk returns a tuple with the PrimaryButtonColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonColorHoverOk() (*string, bool) {
	if o == nil || o.PrimaryButtonColorHover == nil {
		return nil, false
	}
	return o.PrimaryButtonColorHover, true
}

// HasPrimaryButtonColorHover returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasPrimaryButtonColorHover() bool {
	if o != nil && o.PrimaryButtonColorHover != nil {
		return true
	}

	return false
}

// SetPrimaryButtonColorHover gets a reference to the given string and assigns it to the PrimaryButtonColorHover field.
func (o *ApiAppResponseWhiteLabelingOptions) SetPrimaryButtonColorHover(v string) {
	o.PrimaryButtonColorHover = &v
}

// GetPrimaryButtonTextColor returns the PrimaryButtonTextColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonTextColor() string {
	if o == nil || o.PrimaryButtonTextColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonTextColor
}

// GetPrimaryButtonTextColorOk returns a tuple with the PrimaryButtonTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonTextColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonTextColor == nil {
		return nil, false
	}
	return o.PrimaryButtonTextColor, true
}

// HasPrimaryButtonTextColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasPrimaryButtonTextColor() bool {
	if o != nil && o.PrimaryButtonTextColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonTextColor gets a reference to the given string and assigns it to the PrimaryButtonTextColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetPrimaryButtonTextColor(v string) {
	o.PrimaryButtonTextColor = &v
}

// GetPrimaryButtonTextColorHover returns the PrimaryButtonTextColorHover field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonTextColorHover() string {
	if o == nil || o.PrimaryButtonTextColorHover == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonTextColorHover
}

// GetPrimaryButtonTextColorHoverOk returns a tuple with the PrimaryButtonTextColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetPrimaryButtonTextColorHoverOk() (*string, bool) {
	if o == nil || o.PrimaryButtonTextColorHover == nil {
		return nil, false
	}
	return o.PrimaryButtonTextColorHover, true
}

// HasPrimaryButtonTextColorHover returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasPrimaryButtonTextColorHover() bool {
	if o != nil && o.PrimaryButtonTextColorHover != nil {
		return true
	}

	return false
}

// SetPrimaryButtonTextColorHover gets a reference to the given string and assigns it to the PrimaryButtonTextColorHover field.
func (o *ApiAppResponseWhiteLabelingOptions) SetPrimaryButtonTextColorHover(v string) {
	o.PrimaryButtonTextColorHover = &v
}

// GetSecondaryButtonColor returns the SecondaryButtonColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonColor() string {
	if o == nil || o.SecondaryButtonColor == nil {
		var ret string
		return ret
	}
	return *o.SecondaryButtonColor
}

// GetSecondaryButtonColorOk returns a tuple with the SecondaryButtonColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonColorOk() (*string, bool) {
	if o == nil || o.SecondaryButtonColor == nil {
		return nil, false
	}
	return o.SecondaryButtonColor, true
}

// HasSecondaryButtonColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasSecondaryButtonColor() bool {
	if o != nil && o.SecondaryButtonColor != nil {
		return true
	}

	return false
}

// SetSecondaryButtonColor gets a reference to the given string and assigns it to the SecondaryButtonColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetSecondaryButtonColor(v string) {
	o.SecondaryButtonColor = &v
}

// GetSecondaryButtonColorHover returns the SecondaryButtonColorHover field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonColorHover() string {
	if o == nil || o.SecondaryButtonColorHover == nil {
		var ret string
		return ret
	}
	return *o.SecondaryButtonColorHover
}

// GetSecondaryButtonColorHoverOk returns a tuple with the SecondaryButtonColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonColorHoverOk() (*string, bool) {
	if o == nil || o.SecondaryButtonColorHover == nil {
		return nil, false
	}
	return o.SecondaryButtonColorHover, true
}

// HasSecondaryButtonColorHover returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasSecondaryButtonColorHover() bool {
	if o != nil && o.SecondaryButtonColorHover != nil {
		return true
	}

	return false
}

// SetSecondaryButtonColorHover gets a reference to the given string and assigns it to the SecondaryButtonColorHover field.
func (o *ApiAppResponseWhiteLabelingOptions) SetSecondaryButtonColorHover(v string) {
	o.SecondaryButtonColorHover = &v
}

// GetSecondaryButtonTextColor returns the SecondaryButtonTextColor field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonTextColor() string {
	if o == nil || o.SecondaryButtonTextColor == nil {
		var ret string
		return ret
	}
	return *o.SecondaryButtonTextColor
}

// GetSecondaryButtonTextColorOk returns a tuple with the SecondaryButtonTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonTextColorOk() (*string, bool) {
	if o == nil || o.SecondaryButtonTextColor == nil {
		return nil, false
	}
	return o.SecondaryButtonTextColor, true
}

// HasSecondaryButtonTextColor returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasSecondaryButtonTextColor() bool {
	if o != nil && o.SecondaryButtonTextColor != nil {
		return true
	}

	return false
}

// SetSecondaryButtonTextColor gets a reference to the given string and assigns it to the SecondaryButtonTextColor field.
func (o *ApiAppResponseWhiteLabelingOptions) SetSecondaryButtonTextColor(v string) {
	o.SecondaryButtonTextColor = &v
}

// GetSecondaryButtonTextColorHover returns the SecondaryButtonTextColorHover field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonTextColorHover() string {
	if o == nil || o.SecondaryButtonTextColorHover == nil {
		var ret string
		return ret
	}
	return *o.SecondaryButtonTextColorHover
}

// GetSecondaryButtonTextColorHoverOk returns a tuple with the SecondaryButtonTextColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetSecondaryButtonTextColorHoverOk() (*string, bool) {
	if o == nil || o.SecondaryButtonTextColorHover == nil {
		return nil, false
	}
	return o.SecondaryButtonTextColorHover, true
}

// HasSecondaryButtonTextColorHover returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasSecondaryButtonTextColorHover() bool {
	if o != nil && o.SecondaryButtonTextColorHover != nil {
		return true
	}

	return false
}

// SetSecondaryButtonTextColorHover gets a reference to the given string and assigns it to the SecondaryButtonTextColorHover field.
func (o *ApiAppResponseWhiteLabelingOptions) SetSecondaryButtonTextColorHover(v string) {
	o.SecondaryButtonTextColorHover = &v
}

// GetTextColor1 returns the TextColor1 field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetTextColor1() string {
	if o == nil || o.TextColor1 == nil {
		var ret string
		return ret
	}
	return *o.TextColor1
}

// GetTextColor1Ok returns a tuple with the TextColor1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetTextColor1Ok() (*string, bool) {
	if o == nil || o.TextColor1 == nil {
		return nil, false
	}
	return o.TextColor1, true
}

// HasTextColor1 returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasTextColor1() bool {
	if o != nil && o.TextColor1 != nil {
		return true
	}

	return false
}

// SetTextColor1 gets a reference to the given string and assigns it to the TextColor1 field.
func (o *ApiAppResponseWhiteLabelingOptions) SetTextColor1(v string) {
	o.TextColor1 = &v
}

// GetTextColor2 returns the TextColor2 field value if set, zero value otherwise.
func (o *ApiAppResponseWhiteLabelingOptions) GetTextColor2() string {
	if o == nil || o.TextColor2 == nil {
		var ret string
		return ret
	}
	return *o.TextColor2
}

// GetTextColor2Ok returns a tuple with the TextColor2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseWhiteLabelingOptions) GetTextColor2Ok() (*string, bool) {
	if o == nil || o.TextColor2 == nil {
		return nil, false
	}
	return o.TextColor2, true
}

// HasTextColor2 returns a boolean if a field has been set.
func (o *ApiAppResponseWhiteLabelingOptions) HasTextColor2() bool {
	if o != nil && o.TextColor2 != nil {
		return true
	}

	return false
}

// SetTextColor2 gets a reference to the given string and assigns it to the TextColor2 field.
func (o *ApiAppResponseWhiteLabelingOptions) SetTextColor2(v string) {
	o.TextColor2 = &v
}

func (o ApiAppResponseWhiteLabelingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HeaderBackgroundColor != nil {
		toSerialize["header_background_color"] = o.HeaderBackgroundColor
	}
	if o.LegalVersion != nil {
		toSerialize["legal_version"] = o.LegalVersion
	}
	if o.LinkColor != nil {
		toSerialize["link_color"] = o.LinkColor
	}
	if o.PageBackgroundColor != nil {
		toSerialize["page_background_color"] = o.PageBackgroundColor
	}
	if o.PrimaryButtonColor != nil {
		toSerialize["primary_button_color"] = o.PrimaryButtonColor
	}
	if o.PrimaryButtonColorHover != nil {
		toSerialize["primary_button_color_hover"] = o.PrimaryButtonColorHover
	}
	if o.PrimaryButtonTextColor != nil {
		toSerialize["primary_button_text_color"] = o.PrimaryButtonTextColor
	}
	if o.PrimaryButtonTextColorHover != nil {
		toSerialize["primary_button_text_color_hover"] = o.PrimaryButtonTextColorHover
	}
	if o.SecondaryButtonColor != nil {
		toSerialize["secondary_button_color"] = o.SecondaryButtonColor
	}
	if o.SecondaryButtonColorHover != nil {
		toSerialize["secondary_button_color_hover"] = o.SecondaryButtonColorHover
	}
	if o.SecondaryButtonTextColor != nil {
		toSerialize["secondary_button_text_color"] = o.SecondaryButtonTextColor
	}
	if o.SecondaryButtonTextColorHover != nil {
		toSerialize["secondary_button_text_color_hover"] = o.SecondaryButtonTextColorHover
	}
	if o.TextColor1 != nil {
		toSerialize["text_color1"] = o.TextColor1
	}
	if o.TextColor2 != nil {
		toSerialize["text_color2"] = o.TextColor2
	}
	return json.Marshal(toSerialize)
}

type NullableApiAppResponseWhiteLabelingOptions struct {
	value *ApiAppResponseWhiteLabelingOptions
	isSet bool
}

func (v NullableApiAppResponseWhiteLabelingOptions) Get() *ApiAppResponseWhiteLabelingOptions {
	return v.value
}

func (v *NullableApiAppResponseWhiteLabelingOptions) Set(val *ApiAppResponseWhiteLabelingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAppResponseWhiteLabelingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAppResponseWhiteLabelingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAppResponseWhiteLabelingOptions(val *ApiAppResponseWhiteLabelingOptions) *NullableApiAppResponseWhiteLabelingOptions {
	return &NullableApiAppResponseWhiteLabelingOptions{value: val, isSet: true}
}

func (v NullableApiAppResponseWhiteLabelingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAppResponseWhiteLabelingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

