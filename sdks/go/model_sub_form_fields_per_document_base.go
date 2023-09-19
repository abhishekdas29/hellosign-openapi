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

// SubFormFieldsPerDocumentBase The fields that should appear on the document, expressed as an array of objects. (For more details you can read about it here: [Using Form Fields per Document](/docs/openapi/form-fields-per-document).)  **NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.  * Text Field use `SubFormFieldsPerDocumentText` * Dropdown Field use `SubFormFieldsPerDocumentDropdown` * Hyperlink Field use `SubFormFieldsPerDocumentHyperlink` * Checkbox Field use `SubFormFieldsPerDocumentCheckbox` * Radio Field use `SubFormFieldsPerDocumentRadio` * Signature Field use `SubFormFieldsPerDocumentSignature` * Date Signed Field use `SubFormFieldsPerDocumentDateSigned` * Initials Field use `SubFormFieldsPerDocumentInitials` * Text Merge Field use `SubFormFieldsPerDocumentTextMerge` * Checkbox Merge Field use `SubFormFieldsPerDocumentCheckboxMerge`
type SubFormFieldsPerDocumentBase struct {
	// Represents the integer index of the `file` or `file_url` document the field should be attached to.
	DocumentIndex int32 `json:"document_index"`
	// An identifier for the field that is unique across all documents in the request.
	ApiId string `json:"api_id"`
	// Size of the field in pixels.
	Height int32 `json:"height"`
	// Whether this field is required.
	Required bool `json:"required"`
	// Signer index identified by the offset in the signers parameter (0-based indexing), indicating which signer should fill out the field.  **NOTE**: To set the value of the field as the preparer you must set this to `me_now`  **NOTE**: If type is `text-merge` or `checkbox-merge`, you must set this to sender in order to use pre-filled data.
	Signer string `json:"signer"`
	Type string `json:"type"`
	// Size of the field in pixels.
	Width int32 `json:"width"`
	// Location coordinates of the field in pixels.
	X int32 `json:"x"`
	// Location coordinates of the field in pixels.
	Y int32 `json:"y"`
	// Display name for the field.
	Name *string `json:"name,omitempty"`
	// Page in the document where the field should be placed (requires documents be PDF files).  - When the page number parameter is supplied, the API will use the new coordinate system. - Check out the differences between both [coordinate systems](https://faq.hellosign.com/hc/en-us/articles/217115577) and how to use them.
	Page NullableInt32 `json:"page,omitempty"`
}

// NewSubFormFieldsPerDocumentBase instantiates a new SubFormFieldsPerDocumentBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentBase(documentIndex int32, apiId string, height int32, required bool, signer string, type_ string, width int32, x int32, y int32) *SubFormFieldsPerDocumentBase {
	this := SubFormFieldsPerDocumentBase{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewSubFormFieldsPerDocumentBaseWithDefaults instantiates a new SubFormFieldsPerDocumentBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentBaseWithDefaults() *SubFormFieldsPerDocumentBase {
	this := SubFormFieldsPerDocumentBase{}
	return &this
}

// GetDocumentIndex returns the DocumentIndex field value
func (o *SubFormFieldsPerDocumentBase) GetDocumentIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DocumentIndex
}

// GetDocumentIndexOk returns a tuple with the DocumentIndex field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetDocumentIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentIndex, true
}

// SetDocumentIndex sets field value
func (o *SubFormFieldsPerDocumentBase) SetDocumentIndex(v int32) {
	o.DocumentIndex = v
}

// GetApiId returns the ApiId field value
func (o *SubFormFieldsPerDocumentBase) GetApiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetApiIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiId, true
}

// SetApiId sets field value
func (o *SubFormFieldsPerDocumentBase) SetApiId(v string) {
	o.ApiId = v
}

// GetHeight returns the Height field value
func (o *SubFormFieldsPerDocumentBase) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *SubFormFieldsPerDocumentBase) SetHeight(v int32) {
	o.Height = v
}

// GetRequired returns the Required field value
func (o *SubFormFieldsPerDocumentBase) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *SubFormFieldsPerDocumentBase) SetRequired(v bool) {
	o.Required = v
}

// GetSigner returns the Signer field value
func (o *SubFormFieldsPerDocumentBase) GetSigner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signer
}

// GetSignerOk returns a tuple with the Signer field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetSignerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Signer, true
}

// SetSigner sets field value
func (o *SubFormFieldsPerDocumentBase) SetSigner(v string) {
	o.Signer = v
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentBase) SetType(v string) {
	o.Type = v
}

// GetWidth returns the Width field value
func (o *SubFormFieldsPerDocumentBase) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetWidthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *SubFormFieldsPerDocumentBase) SetWidth(v int32) {
	o.Width = v
}

// GetX returns the X field value
func (o *SubFormFieldsPerDocumentBase) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetXOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *SubFormFieldsPerDocumentBase) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *SubFormFieldsPerDocumentBase) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetYOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *SubFormFieldsPerDocumentBase) SetY(v int32) {
	o.Y = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubFormFieldsPerDocumentBase) SetName(v string) {
	o.Name = &v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubFormFieldsPerDocumentBase) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubFormFieldsPerDocumentBase) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentBase) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *SubFormFieldsPerDocumentBase) SetPage(v int32) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *SubFormFieldsPerDocumentBase) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *SubFormFieldsPerDocumentBase) UnsetPage() {
	o.Page.Unset()
}

func (o SubFormFieldsPerDocumentBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["document_index"] = o.DocumentIndex
	}
	if true {
		toSerialize["api_id"] = o.ApiId
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["signer"] = o.Signer
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSubFormFieldsPerDocumentBase struct {
	value *SubFormFieldsPerDocumentBase
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentBase) Get() *SubFormFieldsPerDocumentBase {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentBase) Set(val *SubFormFieldsPerDocumentBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentBase(val *SubFormFieldsPerDocumentBase) *NullableSubFormFieldsPerDocumentBase {
	return &NullableSubFormFieldsPerDocumentBase{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


