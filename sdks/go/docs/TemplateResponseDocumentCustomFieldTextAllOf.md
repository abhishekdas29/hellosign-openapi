# TemplateResponseDocumentCustomFieldTextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this Custom Field. Only &#x60;text&#x60; and &#x60;checkbox&#x60; are currently supported.  * Text uses &#x60;TemplateResponseDocumentCustomFieldText&#x60; * Checkbox uses &#x60;TemplateResponseDocumentCustomFieldCheckbox&#x60; | [optional] [default to "text"]
**AvgTextLength** | Pointer to [**TemplateResponseFieldAvgTextLength**](TemplateResponseFieldAvgTextLength.md) |  | [optional] 
**IsMultiline** | Pointer to **bool** | Whether this form field is multiline text. | [optional] 
**OriginalFontSize** | Pointer to **int32** | Original font size used in this form field&#39;s text. | [optional] 
**FontFamily** | Pointer to **string** | Font family used in this form field&#39;s text. | [optional] 

## Methods

### NewTemplateResponseDocumentCustomFieldTextAllOf

`func NewTemplateResponseDocumentCustomFieldTextAllOf() *TemplateResponseDocumentCustomFieldTextAllOf`

NewTemplateResponseDocumentCustomFieldTextAllOf instantiates a new TemplateResponseDocumentCustomFieldTextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseDocumentCustomFieldTextAllOfWithDefaults

`func NewTemplateResponseDocumentCustomFieldTextAllOfWithDefaults() *TemplateResponseDocumentCustomFieldTextAllOf`

NewTemplateResponseDocumentCustomFieldTextAllOfWithDefaults instantiates a new TemplateResponseDocumentCustomFieldTextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAvgTextLength

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetAvgTextLength() TemplateResponseFieldAvgTextLength`

GetAvgTextLength returns the AvgTextLength field if non-nil, zero value otherwise.

### GetAvgTextLengthOk

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetAvgTextLengthOk() (*TemplateResponseFieldAvgTextLength, bool)`

GetAvgTextLengthOk returns a tuple with the AvgTextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTextLength

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) SetAvgTextLength(v TemplateResponseFieldAvgTextLength)`

SetAvgTextLength sets AvgTextLength field to given value.

### HasAvgTextLength

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) HasAvgTextLength() bool`

HasAvgTextLength returns a boolean if a field has been set.

### GetIsMultiline

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetIsMultiline() bool`

GetIsMultiline returns the IsMultiline field if non-nil, zero value otherwise.

### GetIsMultilineOk

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetIsMultilineOk() (*bool, bool)`

GetIsMultilineOk returns a tuple with the IsMultiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiline

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) SetIsMultiline(v bool)`

SetIsMultiline sets IsMultiline field to given value.

### HasIsMultiline

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) HasIsMultiline() bool`

HasIsMultiline returns a boolean if a field has been set.

### GetOriginalFontSize

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetOriginalFontSize() int32`

GetOriginalFontSize returns the OriginalFontSize field if non-nil, zero value otherwise.

### GetOriginalFontSizeOk

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetOriginalFontSizeOk() (*int32, bool)`

GetOriginalFontSizeOk returns a tuple with the OriginalFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFontSize

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) SetOriginalFontSize(v int32)`

SetOriginalFontSize sets OriginalFontSize field to given value.

### HasOriginalFontSize

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) HasOriginalFontSize() bool`

HasOriginalFontSize returns a boolean if a field has been set.

### GetFontFamily

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *TemplateResponseDocumentCustomFieldTextAllOf) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


