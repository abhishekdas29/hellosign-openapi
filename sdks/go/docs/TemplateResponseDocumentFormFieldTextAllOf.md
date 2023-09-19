# TemplateResponseDocumentFormFieldTextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this form field. See [field types](/api/reference/constants/#field-types).  * Text Field uses &#x60;TemplateResponseDocumentFormFieldText&#x60; * Dropdown Field uses &#x60;TemplateResponseDocumentFormFieldDropdown&#x60; * Hyperlink Field uses &#x60;TemplateResponseDocumentFormFieldHyperlink&#x60; * Checkbox Field uses &#x60;TemplateResponseDocumentFormFieldCheckbox&#x60; * Radio Field uses &#x60;TemplateResponseDocumentFormFieldRadio&#x60; * Signature Field uses &#x60;TemplateResponseDocumentFormFieldSignature&#x60; * Date Signed Field uses &#x60;TemplateResponseDocumentFormFieldDateSigned&#x60; * Initials Field uses &#x60;TemplateResponseDocumentFormFieldInitials&#x60; | [optional] [default to "text"]
**AvgTextLength** | Pointer to [**TemplateResponseFieldAvgTextLength**](TemplateResponseFieldAvgTextLength.md) |  | [optional] 
**IsMultiline** | Pointer to **bool** | Whether this form field is multiline text. | [optional] 
**OriginalFontSize** | Pointer to **int32** | Original font size used in this form field&#39;s text. | [optional] 
**FontFamily** | Pointer to **string** | Font family used in this form field&#39;s text. | [optional] 
**ValidationType** | Pointer to **NullableString** | Each text field may contain a &#x60;validation_type&#x60; parameter. Check out the list of [validation types](https://faq.hellosign.com/hc/en-us/articles/217115577) to learn more about the possible values. | [optional] 

## Methods

### NewTemplateResponseDocumentFormFieldTextAllOf

`func NewTemplateResponseDocumentFormFieldTextAllOf() *TemplateResponseDocumentFormFieldTextAllOf`

NewTemplateResponseDocumentFormFieldTextAllOf instantiates a new TemplateResponseDocumentFormFieldTextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseDocumentFormFieldTextAllOfWithDefaults

`func NewTemplateResponseDocumentFormFieldTextAllOfWithDefaults() *TemplateResponseDocumentFormFieldTextAllOf`

NewTemplateResponseDocumentFormFieldTextAllOfWithDefaults instantiates a new TemplateResponseDocumentFormFieldTextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAvgTextLength

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetAvgTextLength() TemplateResponseFieldAvgTextLength`

GetAvgTextLength returns the AvgTextLength field if non-nil, zero value otherwise.

### GetAvgTextLengthOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetAvgTextLengthOk() (*TemplateResponseFieldAvgTextLength, bool)`

GetAvgTextLengthOk returns a tuple with the AvgTextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTextLength

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetAvgTextLength(v TemplateResponseFieldAvgTextLength)`

SetAvgTextLength sets AvgTextLength field to given value.

### HasAvgTextLength

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasAvgTextLength() bool`

HasAvgTextLength returns a boolean if a field has been set.

### GetIsMultiline

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetIsMultiline() bool`

GetIsMultiline returns the IsMultiline field if non-nil, zero value otherwise.

### GetIsMultilineOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetIsMultilineOk() (*bool, bool)`

GetIsMultilineOk returns a tuple with the IsMultiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiline

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetIsMultiline(v bool)`

SetIsMultiline sets IsMultiline field to given value.

### HasIsMultiline

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasIsMultiline() bool`

HasIsMultiline returns a boolean if a field has been set.

### GetOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetOriginalFontSize() int32`

GetOriginalFontSize returns the OriginalFontSize field if non-nil, zero value otherwise.

### GetOriginalFontSizeOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetOriginalFontSizeOk() (*int32, bool)`

GetOriginalFontSizeOk returns a tuple with the OriginalFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetOriginalFontSize(v int32)`

SetOriginalFontSize sets OriginalFontSize field to given value.

### HasOriginalFontSize

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasOriginalFontSize() bool`

HasOriginalFontSize returns a boolean if a field has been set.

### GetFontFamily

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### GetValidationType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetValidationType() string`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *TemplateResponseDocumentFormFieldTextAllOf) GetValidationTypeOk() (*string, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetValidationType(v string)`

SetValidationType sets ValidationType field to given value.

### HasValidationType

`func (o *TemplateResponseDocumentFormFieldTextAllOf) HasValidationType() bool`

HasValidationType returns a boolean if a field has been set.

### SetValidationTypeNil

`func (o *TemplateResponseDocumentFormFieldTextAllOf) SetValidationTypeNil(b bool)`

 SetValidationTypeNil sets the value for ValidationType to be an explicit nil

### UnsetValidationType
`func (o *TemplateResponseDocumentFormFieldTextAllOf) UnsetValidationType()`

UnsetValidationType ensures that no value is present for ValidationType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


