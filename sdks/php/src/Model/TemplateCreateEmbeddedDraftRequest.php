<?php
/**
 * TemplateCreateEmbeddedDraftRequest
 *
 * PHP version 7.3
 *
 * @category Class
 * @author   OpenAPI Generator team
 * @see     https://openapi-generator.tech
 */

/**
 * Dropbox Sign API
 *
 * Dropbox Sign v3 API
 *
 * The version of the OpenAPI document: 3.0.0
 * Contact: apisupport@hellosign.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.3.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace HelloSign\Model;

use ArrayAccess;
use HelloSign\ObjectSerializer;
use InvalidArgumentException;
use JsonSerializable;
use SplFileObject;

/**
 * TemplateCreateEmbeddedDraftRequest Class Doc Comment
 *
 * @category Class
 * @author   OpenAPI Generator team
 * @see     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class TemplateCreateEmbeddedDraftRequest implements ModelInterface, ArrayAccess, JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
     * The original name of the model.
     *
     * @var string
     */
    protected static $openAPIModelName = 'TemplateCreateEmbeddedDraftRequest';

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $openAPITypes = [
        'client_id' => 'string',
        'files' => '\SplFileObject[]',
        'file_urls' => 'string[]',
        'allow_ccs' => 'bool',
        'allow_reassign' => 'bool',
        'attachments' => '\HelloSign\Model\SubAttachment[]',
        'cc_roles' => 'string[]',
        'editor_options' => '\HelloSign\Model\SubEditorOptions',
        'field_options' => '\HelloSign\Model\SubFieldOptions',
        'force_signer_roles' => 'bool',
        'force_subject_message' => 'bool',
        'form_field_groups' => '\HelloSign\Model\SubFormFieldGroup[]',
        'form_field_rules' => '\HelloSign\Model\SubFormFieldRule[]',
        'form_fields_per_document' => '\HelloSign\Model\SubFormFieldsPerDocumentBase[]',
        'merge_fields' => '\HelloSign\Model\SubMergeField[]',
        'message' => 'string',
        'metadata' => 'array<string,mixed>',
        'show_preview' => 'bool',
        'show_progress_stepper' => 'bool',
        'signer_roles' => '\HelloSign\Model\SubTemplateRole[]',
        'skip_me_now' => 'bool',
        'subject' => 'string',
        'test_mode' => 'bool',
        'title' => 'string',
        'use_preexisting_fields' => 'bool',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     * @phpstan-var array<string, string|null>
     * @psalm-var array<string, string|null>
     */
    protected static $openAPIFormats = [
        'client_id' => null,
        'files' => 'binary',
        'file_urls' => null,
        'allow_ccs' => null,
        'allow_reassign' => null,
        'attachments' => null,
        'cc_roles' => null,
        'editor_options' => null,
        'field_options' => null,
        'force_signer_roles' => null,
        'force_subject_message' => null,
        'form_field_groups' => null,
        'form_field_rules' => null,
        'form_fields_per_document' => null,
        'merge_fields' => null,
        'message' => null,
        'metadata' => null,
        'show_preview' => null,
        'show_progress_stepper' => null,
        'signer_roles' => null,
        'skip_me_now' => null,
        'subject' => null,
        'test_mode' => null,
        'title' => null,
        'use_preexisting_fields' => null,
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'client_id' => 'client_id',
        'files' => 'files',
        'file_urls' => 'file_urls',
        'allow_ccs' => 'allow_ccs',
        'allow_reassign' => 'allow_reassign',
        'attachments' => 'attachments',
        'cc_roles' => 'cc_roles',
        'editor_options' => 'editor_options',
        'field_options' => 'field_options',
        'force_signer_roles' => 'force_signer_roles',
        'force_subject_message' => 'force_subject_message',
        'form_field_groups' => 'form_field_groups',
        'form_field_rules' => 'form_field_rules',
        'form_fields_per_document' => 'form_fields_per_document',
        'merge_fields' => 'merge_fields',
        'message' => 'message',
        'metadata' => 'metadata',
        'show_preview' => 'show_preview',
        'show_progress_stepper' => 'show_progress_stepper',
        'signer_roles' => 'signer_roles',
        'skip_me_now' => 'skip_me_now',
        'subject' => 'subject',
        'test_mode' => 'test_mode',
        'title' => 'title',
        'use_preexisting_fields' => 'use_preexisting_fields',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'client_id' => 'setClientId',
        'files' => 'setFiles',
        'file_urls' => 'setFileUrls',
        'allow_ccs' => 'setAllowCcs',
        'allow_reassign' => 'setAllowReassign',
        'attachments' => 'setAttachments',
        'cc_roles' => 'setCcRoles',
        'editor_options' => 'setEditorOptions',
        'field_options' => 'setFieldOptions',
        'force_signer_roles' => 'setForceSignerRoles',
        'force_subject_message' => 'setForceSubjectMessage',
        'form_field_groups' => 'setFormFieldGroups',
        'form_field_rules' => 'setFormFieldRules',
        'form_fields_per_document' => 'setFormFieldsPerDocument',
        'merge_fields' => 'setMergeFields',
        'message' => 'setMessage',
        'metadata' => 'setMetadata',
        'show_preview' => 'setShowPreview',
        'show_progress_stepper' => 'setShowProgressStepper',
        'signer_roles' => 'setSignerRoles',
        'skip_me_now' => 'setSkipMeNow',
        'subject' => 'setSubject',
        'test_mode' => 'setTestMode',
        'title' => 'setTitle',
        'use_preexisting_fields' => 'setUsePreexistingFields',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'client_id' => 'getClientId',
        'files' => 'getFiles',
        'file_urls' => 'getFileUrls',
        'allow_ccs' => 'getAllowCcs',
        'allow_reassign' => 'getAllowReassign',
        'attachments' => 'getAttachments',
        'cc_roles' => 'getCcRoles',
        'editor_options' => 'getEditorOptions',
        'field_options' => 'getFieldOptions',
        'force_signer_roles' => 'getForceSignerRoles',
        'force_subject_message' => 'getForceSubjectMessage',
        'form_field_groups' => 'getFormFieldGroups',
        'form_field_rules' => 'getFormFieldRules',
        'form_fields_per_document' => 'getFormFieldsPerDocument',
        'merge_fields' => 'getMergeFields',
        'message' => 'getMessage',
        'metadata' => 'getMetadata',
        'show_preview' => 'getShowPreview',
        'show_progress_stepper' => 'getShowProgressStepper',
        'signer_roles' => 'getSignerRoles',
        'skip_me_now' => 'getSkipMeNow',
        'subject' => 'getSubject',
        'test_mode' => 'getTestMode',
        'title' => 'getTitle',
        'use_preexisting_fields' => 'getUsePreexistingFields',
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    /**
     * Associative array for storing property values
     *
     * @var array
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param array|null $data Associated array of property values
     *                         initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['client_id'] = $data['client_id'] ?? null;
        $this->container['files'] = $data['files'] ?? null;
        $this->container['file_urls'] = $data['file_urls'] ?? null;
        $this->container['allow_ccs'] = $data['allow_ccs'] ?? true;
        $this->container['allow_reassign'] = $data['allow_reassign'] ?? false;
        $this->container['attachments'] = $data['attachments'] ?? null;
        $this->container['cc_roles'] = $data['cc_roles'] ?? null;
        $this->container['editor_options'] = $data['editor_options'] ?? null;
        $this->container['field_options'] = $data['field_options'] ?? null;
        $this->container['force_signer_roles'] = $data['force_signer_roles'] ?? false;
        $this->container['force_subject_message'] = $data['force_subject_message'] ?? false;
        $this->container['form_field_groups'] = $data['form_field_groups'] ?? null;
        $this->container['form_field_rules'] = $data['form_field_rules'] ?? null;
        $this->container['form_fields_per_document'] = $data['form_fields_per_document'] ?? null;
        $this->container['merge_fields'] = $data['merge_fields'] ?? null;
        $this->container['message'] = $data['message'] ?? null;
        $this->container['metadata'] = $data['metadata'] ?? null;
        $this->container['show_preview'] = $data['show_preview'] ?? false;
        $this->container['show_progress_stepper'] = $data['show_progress_stepper'] ?? true;
        $this->container['signer_roles'] = $data['signer_roles'] ?? null;
        $this->container['skip_me_now'] = $data['skip_me_now'] ?? false;
        $this->container['subject'] = $data['subject'] ?? null;
        $this->container['test_mode'] = $data['test_mode'] ?? false;
        $this->container['title'] = $data['title'] ?? null;
        $this->container['use_preexisting_fields'] = $data['use_preexisting_fields'] ?? false;
    }

    /** @deprecated use ::init() */
    public static function fromArray(array $data): TemplateCreateEmbeddedDraftRequest
    {
        return self::init($data);
    }

    /** Attempt to instantiate and hydrate a new instance of this class */
    public static function init(array $data): TemplateCreateEmbeddedDraftRequest
    {
        /** @var TemplateCreateEmbeddedDraftRequest $obj */
        $obj = ObjectSerializer::deserialize(
            $data,
            TemplateCreateEmbeddedDraftRequest::class,
        );

        return $obj;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['client_id'] === null) {
            $invalidProperties[] = "'client_id' can't be null";
        }
        if (!is_null($this->container['message']) && (mb_strlen($this->container['message']) > 5000)) {
            $invalidProperties[] = "invalid value for 'message', the character length must be smaller than or equal to 5000.";
        }

        if (!is_null($this->container['subject']) && (mb_strlen($this->container['subject']) > 200)) {
            $invalidProperties[] = "invalid value for 'subject', the character length must be smaller than or equal to 200.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets client_id
     *
     * @return string
     */
    public function getClientId()
    {
        return $this->container['client_id'];
    }

    /**
     * Sets client_id
     *
     * @param string $client_id Client id of the app you're using to create this draft. Used to apply the branding and callback url defined for the app.
     *
     * @return self
     */
    public function setClientId(string $client_id)
    {
        $this->container['client_id'] = $client_id;

        return $this;
    }

    /**
     * Gets files
     *
     * @return SplFileObject[]|null
     */
    public function getFiles()
    {
        return $this->container['files'];
    }

    /**
     * Sets files
     *
     * @param SplFileObject[]|null $files Use `files[]` to indicate the uploaded file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both.
     *
     * @return self
     */
    public function setFiles(?array $files)
    {
        $this->container['files'] = $files;

        return $this;
    }

    /**
     * Gets file_urls
     *
     * @return string[]|null
     */
    public function getFileUrls()
    {
        return $this->container['file_urls'];
    }

    /**
     * Sets file_urls
     *
     * @param string[]|null $file_urls Use `file_urls[]` to have Dropbox Sign download the file(s) to send for signature.  This endpoint requires either **files** or **file_urls[]**, but not both.
     *
     * @return self
     */
    public function setFileUrls(?array $file_urls)
    {
        $this->container['file_urls'] = $file_urls;

        return $this;
    }

    /**
     * Gets allow_ccs
     *
     * @return bool|null
     */
    public function getAllowCcs()
    {
        return $this->container['allow_ccs'];
    }

    /**
     * Sets allow_ccs
     *
     * @param bool|null $allow_ccs this allows the requester to specify whether the user is allowed to provide email addresses to CC when creating a template
     *
     * @return self
     */
    public function setAllowCcs(?bool $allow_ccs)
    {
        $this->container['allow_ccs'] = $allow_ccs;

        return $this;
    }

    /**
     * Gets allow_reassign
     *
     * @return bool|null
     */
    public function getAllowReassign()
    {
        return $this->container['allow_reassign'];
    }

    /**
     * Sets allow_reassign
     *
     * @param bool|null $allow_reassign Allows signers to reassign their signature requests to other signers if set to `true`. Defaults to `false`.  **Note**: Only available for Premium plan and higher.
     *
     * @return self
     */
    public function setAllowReassign(?bool $allow_reassign)
    {
        $this->container['allow_reassign'] = $allow_reassign;

        return $this;
    }

    /**
     * Gets attachments
     *
     * @return SubAttachment[]|null
     */
    public function getAttachments()
    {
        return $this->container['attachments'];
    }

    /**
     * Sets attachments
     *
     * @param SubAttachment[]|null $attachments A list describing the attachments
     *
     * @return self
     */
    public function setAttachments(?array $attachments)
    {
        $this->container['attachments'] = $attachments;

        return $this;
    }

    /**
     * Gets cc_roles
     *
     * @return string[]|null
     */
    public function getCcRoles()
    {
        return $this->container['cc_roles'];
    }

    /**
     * Sets cc_roles
     *
     * @param string[]|null $cc_roles The CC roles that must be assigned when using the template to send a signature request
     *
     * @return self
     */
    public function setCcRoles(?array $cc_roles)
    {
        $this->container['cc_roles'] = $cc_roles;

        return $this;
    }

    /**
     * Gets editor_options
     *
     * @return SubEditorOptions|null
     */
    public function getEditorOptions()
    {
        return $this->container['editor_options'];
    }

    /**
     * Sets editor_options
     *
     * @param SubEditorOptions|null $editor_options editor_options
     *
     * @return self
     */
    public function setEditorOptions(?SubEditorOptions $editor_options)
    {
        $this->container['editor_options'] = $editor_options;

        return $this;
    }

    /**
     * Gets field_options
     *
     * @return SubFieldOptions|null
     */
    public function getFieldOptions()
    {
        return $this->container['field_options'];
    }

    /**
     * Sets field_options
     *
     * @param SubFieldOptions|null $field_options field_options
     *
     * @return self
     */
    public function setFieldOptions(?SubFieldOptions $field_options)
    {
        $this->container['field_options'] = $field_options;

        return $this;
    }

    /**
     * Gets force_signer_roles
     *
     * @return bool|null
     */
    public function getForceSignerRoles()
    {
        return $this->container['force_signer_roles'];
    }

    /**
     * Sets force_signer_roles
     *
     * @param bool|null $force_signer_roles provide users the ability to review/edit the template signer roles
     *
     * @return self
     */
    public function setForceSignerRoles(?bool $force_signer_roles)
    {
        $this->container['force_signer_roles'] = $force_signer_roles;

        return $this;
    }

    /**
     * Gets force_subject_message
     *
     * @return bool|null
     */
    public function getForceSubjectMessage()
    {
        return $this->container['force_subject_message'];
    }

    /**
     * Sets force_subject_message
     *
     * @param bool|null $force_subject_message provide users the ability to review/edit the template subject and message
     *
     * @return self
     */
    public function setForceSubjectMessage(?bool $force_subject_message)
    {
        $this->container['force_subject_message'] = $force_subject_message;

        return $this;
    }

    /**
     * Gets form_field_groups
     *
     * @return SubFormFieldGroup[]|null
     */
    public function getFormFieldGroups()
    {
        return $this->container['form_field_groups'];
    }

    /**
     * Sets form_field_groups
     *
     * @param SubFormFieldGroup[]|null $form_field_groups Group information for fields defined in `form_fields_per_document`. String-indexed JSON array with `group_label` and `requirement` keys. `form_fields_per_document` must contain fields referencing a group defined in `form_field_groups`.
     *
     * @return self
     */
    public function setFormFieldGroups(?array $form_field_groups)
    {
        $this->container['form_field_groups'] = $form_field_groups;

        return $this;
    }

    /**
     * Gets form_field_rules
     *
     * @return SubFormFieldRule[]|null
     */
    public function getFormFieldRules()
    {
        return $this->container['form_field_rules'];
    }

    /**
     * Sets form_field_rules
     *
     * @param SubFormFieldRule[]|null $form_field_rules conditional Logic rules for fields defined in `form_fields_per_document`
     *
     * @return self
     */
    public function setFormFieldRules(?array $form_field_rules)
    {
        $this->container['form_field_rules'] = $form_field_rules;

        return $this;
    }

    /**
     * Gets form_fields_per_document
     *
     * @return SubFormFieldsPerDocumentBase[]|null
     */
    public function getFormFieldsPerDocument()
    {
        return $this->container['form_fields_per_document'];
    }

    /**
     * Sets form_fields_per_document
     *
     * @param SubFormFieldsPerDocumentBase[]|null $form_fields_per_document The fields that should appear on the document, expressed as an array of objects. (We're currently fixing a bug where this property only accepts a two-dimensional array. You can read about it here: <a href=\"/docs/openapi/form-fields-per-document\" target=\"_blank\">Using Form Fields per Document</a>.)  **NOTE**: Fields like **text**, **dropdown**, **checkbox**, **radio**, and **hyperlink** have additional required and optional parameters. Check out the list of [additional parameters](/api/reference/constants/#form-fields-per-document) for these field types.  * Text Field use `SubFormFieldsPerDocumentText` * Dropdown Field use `SubFormFieldsPerDocumentDropdown` * Hyperlink Field use `SubFormFieldsPerDocumentHyperlink` * Checkbox Field use `SubFormFieldsPerDocumentCheckbox` * Radio Field use `SubFormFieldsPerDocumentRadio` * Signature Field use `SubFormFieldsPerDocumentSignature` * Date Signed Field use `SubFormFieldsPerDocumentDateSigned` * Initials Field use `SubFormFieldsPerDocumentInitials` * Text Merge Field use `SubFormFieldsPerDocumentTextMerge` * Checkbox Merge Field use `SubFormFieldsPerDocumentCheckboxMerge`
     *
     * @return self
     */
    public function setFormFieldsPerDocument(?array $form_fields_per_document)
    {
        $this->container['form_fields_per_document'] = $form_fields_per_document;

        return $this;
    }

    /**
     * Gets merge_fields
     *
     * @return SubMergeField[]|null
     */
    public function getMergeFields()
    {
        return $this->container['merge_fields'];
    }

    /**
     * Sets merge_fields
     *
     * @param SubMergeField[]|null $merge_fields Add merge fields to the template. Merge fields are placed by the user creating the template and used to pre-fill data by passing values into signature requests with the `custom_fields` parameter. If the signature request using that template *does not* pass a value into a merge field, then an empty field remains in the document.
     *
     * @return self
     */
    public function setMergeFields(?array $merge_fields)
    {
        $this->container['merge_fields'] = $merge_fields;

        return $this;
    }

    /**
     * Gets message
     *
     * @return string|null
     */
    public function getMessage()
    {
        return $this->container['message'];
    }

    /**
     * Sets message
     *
     * @param string|null $message the default template email message
     *
     * @return self
     */
    public function setMessage(?string $message)
    {
        if (!is_null($message) && (mb_strlen($message) > 5000)) {
            throw new InvalidArgumentException('invalid length for $message when calling TemplateCreateEmbeddedDraftRequest., must be smaller than or equal to 5000.');
        }

        $this->container['message'] = $message;

        return $this;
    }

    /**
     * Gets metadata
     *
     * @return array<string,mixed>|null
     */
    public function getMetadata()
    {
        return $this->container['metadata'];
    }

    /**
     * Sets metadata
     *
     * @param array<string,mixed>|null $metadata Key-value data that should be attached to the signature request. This metadata is included in all API responses and events involving the signature request. For example, use the metadata field to store a signer's order number for look up when receiving events for the signature request.  Each request can include up to 10 metadata keys (or 50 nested metadata keys), with key names up to 40 characters long and values up to 1000 characters long.
     *
     * @return self
     */
    public function setMetadata(?array $metadata)
    {
        $this->container['metadata'] = $metadata;

        return $this;
    }

    /**
     * Gets show_preview
     *
     * @return bool|null
     */
    public function getShowPreview()
    {
        return $this->container['show_preview'];
    }

    /**
     * Sets show_preview
     *
     * @param bool|null $show_preview This allows the requester to enable the editor/preview experience.  - `show_preview=true`: Allows requesters to enable the editor/preview experience. - `show_preview=false`: Allows requesters to disable the editor/preview experience.
     *
     * @return self
     */
    public function setShowPreview(?bool $show_preview)
    {
        $this->container['show_preview'] = $show_preview;

        return $this;
    }

    /**
     * Gets show_progress_stepper
     *
     * @return bool|null
     */
    public function getShowProgressStepper()
    {
        return $this->container['show_progress_stepper'];
    }

    /**
     * Sets show_progress_stepper
     *
     * @param bool|null $show_progress_stepper when only one step remains in the signature request process and this parameter is set to `false` then the progress stepper will be hidden
     *
     * @return self
     */
    public function setShowProgressStepper(?bool $show_progress_stepper)
    {
        $this->container['show_progress_stepper'] = $show_progress_stepper;

        return $this;
    }

    /**
     * Gets signer_roles
     *
     * @return SubTemplateRole[]|null
     */
    public function getSignerRoles()
    {
        return $this->container['signer_roles'];
    }

    /**
     * Sets signer_roles
     *
     * @param SubTemplateRole[]|null $signer_roles an array of the designated signer roles that must be specified when sending a SignatureRequest using this Template
     *
     * @return self
     */
    public function setSignerRoles(?array $signer_roles)
    {
        $this->container['signer_roles'] = $signer_roles;

        return $this;
    }

    /**
     * Gets skip_me_now
     *
     * @return bool|null
     */
    public function getSkipMeNow()
    {
        return $this->container['skip_me_now'];
    }

    /**
     * Sets skip_me_now
     *
     * @param bool|null $skip_me_now Disables the \"Me (Now)\" option for the person preparing the document. Does not work with type `send_document`. Defaults to `false`.
     *
     * @return self
     */
    public function setSkipMeNow(?bool $skip_me_now)
    {
        $this->container['skip_me_now'] = $skip_me_now;

        return $this;
    }

    /**
     * Gets subject
     *
     * @return string|null
     */
    public function getSubject()
    {
        return $this->container['subject'];
    }

    /**
     * Sets subject
     *
     * @param string|null $subject the template title (alias)
     *
     * @return self
     */
    public function setSubject(?string $subject)
    {
        if (!is_null($subject) && (mb_strlen($subject) > 200)) {
            throw new InvalidArgumentException('invalid length for $subject when calling TemplateCreateEmbeddedDraftRequest., must be smaller than or equal to 200.');
        }

        $this->container['subject'] = $subject;

        return $this;
    }

    /**
     * Gets test_mode
     *
     * @return bool|null
     */
    public function getTestMode()
    {
        return $this->container['test_mode'];
    }

    /**
     * Sets test_mode
     *
     * @param bool|null $test_mode Whether this is a test, the signature request created from this draft will not be legally binding if set to `true`. Defaults to `false`.
     *
     * @return self
     */
    public function setTestMode(?bool $test_mode)
    {
        $this->container['test_mode'] = $test_mode;

        return $this;
    }

    /**
     * Gets title
     *
     * @return string|null
     */
    public function getTitle()
    {
        return $this->container['title'];
    }

    /**
     * Sets title
     *
     * @param string|null $title the title you want to assign to the SignatureRequest
     *
     * @return self
     */
    public function setTitle(?string $title)
    {
        $this->container['title'] = $title;

        return $this;
    }

    /**
     * Gets use_preexisting_fields
     *
     * @return bool|null
     */
    public function getUsePreexistingFields()
    {
        return $this->container['use_preexisting_fields'];
    }

    /**
     * Sets use_preexisting_fields
     *
     * @param bool|null $use_preexisting_fields enable the detection of predefined PDF fields by setting the `use_preexisting_fields` to `true` (defaults to disabled, or `false`)
     *
     * @return self
     */
    public function setUsePreexistingFields(?bool $use_preexisting_fields)
    {
        $this->container['use_preexisting_fields'] = $use_preexisting_fields;

        return $this;
    }

    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param mixed $offset Offset
     *
     * @return bool
     */
    #[\ReturnTypeWillChange]
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param mixed $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param mixed $offset Offset
     * @param mixed $value Value to be set
     *
     * @return void
     */
    #[\ReturnTypeWillChange]
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param mixed $offset Offset
     *
     * @return void
     */
    #[\ReturnTypeWillChange]
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @see https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return scalar|object|array|null returns data which can be serialized by json_encode(), which is a value
     *                                  of any type other than a resource
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
        return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_UNESCAPED_SLASHES
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}
