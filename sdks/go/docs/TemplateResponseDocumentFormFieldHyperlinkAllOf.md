# TemplateResponseDocumentFormFieldHyperlinkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this form field. See [field types](/api/reference/constants/#field-types).  * Text Field uses &#x60;TemplateResponseDocumentFormFieldText&#x60; * Dropdown Field uses &#x60;TemplateResponseDocumentFormFieldDropdown&#x60; * Hyperlink Field uses &#x60;TemplateResponseDocumentFormFieldHyperlink&#x60; * Checkbox Field uses &#x60;TemplateResponseDocumentFormFieldCheckbox&#x60; * Radio Field uses &#x60;TemplateResponseDocumentFormFieldRadio&#x60; * Signature Field uses &#x60;TemplateResponseDocumentFormFieldSignature&#x60; * Date Signed Field uses &#x60;TemplateResponseDocumentFormFieldDateSigned&#x60; * Initials Field uses &#x60;TemplateResponseDocumentFormFieldInitials&#x60; | [optional] [default to "hyperlink"]
**AvgTextLength** | Pointer to [**TemplateResponseFieldAvgTextLength**](TemplateResponseFieldAvgTextLength.md) |  | [optional] 
**IsMultiline** | Pointer to **bool** | Whether this form field is multiline text. | [optional] 
**OriginalFontSize** | Pointer to **int32** | Original font size used in this form field&#39;s text. | [optional] 
**FontFamily** | Pointer to **string** | Font family used in this form field&#39;s text. | [optional] 

## Methods

### NewTemplateResponseDocumentFormFieldHyperlinkAllOf

`func NewTemplateResponseDocumentFormFieldHyperlinkAllOf() *TemplateResponseDocumentFormFieldHyperlinkAllOf`

NewTemplateResponseDocumentFormFieldHyperlinkAllOf instantiates a new TemplateResponseDocumentFormFieldHyperlinkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseDocumentFormFieldHyperlinkAllOfWithDefaults

`func NewTemplateResponseDocumentFormFieldHyperlinkAllOfWithDefaults() *TemplateResponseDocumentFormFieldHyperlinkAllOf`

NewTemplateResponseDocumentFormFieldHyperlinkAllOfWithDefaults instantiates a new TemplateResponseDocumentFormFieldHyperlinkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAvgTextLength

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetAvgTextLength() TemplateResponseFieldAvgTextLength`

GetAvgTextLength returns the AvgTextLength field if non-nil, zero value otherwise.

### GetAvgTextLengthOk

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetAvgTextLengthOk() (*TemplateResponseFieldAvgTextLength, bool)`

GetAvgTextLengthOk returns a tuple with the AvgTextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTextLength

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) SetAvgTextLength(v TemplateResponseFieldAvgTextLength)`

SetAvgTextLength sets AvgTextLength field to given value.

### HasAvgTextLength

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) HasAvgTextLength() bool`

HasAvgTextLength returns a boolean if a field has been set.

### GetIsMultiline

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetIsMultiline() bool`

GetIsMultiline returns the IsMultiline field if non-nil, zero value otherwise.

### GetIsMultilineOk

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetIsMultilineOk() (*bool, bool)`

GetIsMultilineOk returns a tuple with the IsMultiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiline

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) SetIsMultiline(v bool)`

SetIsMultiline sets IsMultiline field to given value.

### HasIsMultiline

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) HasIsMultiline() bool`

HasIsMultiline returns a boolean if a field has been set.

### GetOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetOriginalFontSize() int32`

GetOriginalFontSize returns the OriginalFontSize field if non-nil, zero value otherwise.

### GetOriginalFontSizeOk

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetOriginalFontSizeOk() (*int32, bool)`

GetOriginalFontSizeOk returns a tuple with the OriginalFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) SetOriginalFontSize(v int32)`

SetOriginalFontSize sets OriginalFontSize field to given value.

### HasOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) HasOriginalFontSize() bool`

HasOriginalFontSize returns a boolean if a field has been set.

### GetFontFamily

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *TemplateResponseDocumentFormFieldHyperlinkAllOf) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


