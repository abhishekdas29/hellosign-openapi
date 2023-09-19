# SubFormFieldsPerDocumentCheckboxAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A yes/no checkbox. Use the &#x60;SubFormFieldsPerDocumentCheckbox&#x60; class. | [default to "checkbox"]
**IsChecked** | **bool** | &#x60;true&#x60; for checking the checkbox field by default, otherwise &#x60;false&#x60;. | 
**Group** | Pointer to **string** | String referencing group defined in &#x60;form_field_groups&#x60; parameter. | [optional] 

## Methods

### NewSubFormFieldsPerDocumentCheckboxAllOf

`func NewSubFormFieldsPerDocumentCheckboxAllOf(type_ string, isChecked bool, ) *SubFormFieldsPerDocumentCheckboxAllOf`

NewSubFormFieldsPerDocumentCheckboxAllOf instantiates a new SubFormFieldsPerDocumentCheckboxAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubFormFieldsPerDocumentCheckboxAllOfWithDefaults

`func NewSubFormFieldsPerDocumentCheckboxAllOfWithDefaults() *SubFormFieldsPerDocumentCheckboxAllOf`

NewSubFormFieldsPerDocumentCheckboxAllOfWithDefaults instantiates a new SubFormFieldsPerDocumentCheckboxAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetIsChecked

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetIsChecked() bool`

GetIsChecked returns the IsChecked field if non-nil, zero value otherwise.

### GetIsCheckedOk

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetIsCheckedOk() (*bool, bool)`

GetIsCheckedOk returns a tuple with the IsChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChecked

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) SetIsChecked(v bool)`

SetIsChecked sets IsChecked field to given value.


### GetGroup

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SubFormFieldsPerDocumentCheckboxAllOf) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


