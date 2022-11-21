/**
 * The MIT License (MIT)
 *
 * Copyright (C) 2022 hellosign.com
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import { RequestFile, AttributeTypeMap } from "./models";
import { TemplateResponseFieldAvgTextLength } from "./templateResponseFieldAvgTextLength";

export class TemplateResponseDocumentFormField {
  /**
   * A unique id for the form field.
   */
  "apiId"?: string;
  /**
   * The name of the form field.
   */
  "name"?: string;
  /**
   * The type of this form field. See [field types](/api/reference/constants/#field-types).
   */
  "type"?: TemplateResponseDocumentFormField.TypeEnum;
  /**
   * The signer of the Form Field.
   */
  "signer"?: string;
  /**
   * The horizontal offset in pixels for this form field.
   */
  "x"?: number;
  /**
   * The vertical offset in pixels for this form field.
   */
  "y"?: number;
  /**
   * The width in pixels of this form field.
   */
  "width"?: number;
  /**
   * The height in pixels of this form field.
   */
  "height"?: number;
  /**
   * Boolean showing whether or not this field is required.
   */
  "required"?: boolean;
  /**
   * The name of the group this field is in. If this field is not a group, this defaults to `null`.
   */
  "group"?: string | null;
  "avgTextLength"?: TemplateResponseFieldAvgTextLength;
  /**
   * Whether this form field is multiline text.
   */
  "isMultiline"?: boolean | null;
  /**
   * Original font size used in this form field\'s text.
   */
  "originalFontSize"?: number | null;
  /**
   * Font family used in this form field\'s text.
   */
  "fontFamily"?: string | null;

  static discriminator: string | undefined = undefined;

  static attributeTypeMap: AttributeTypeMap = [
    {
      name: "apiId",
      baseName: "api_id",
      type: "string",
    },
    {
      name: "name",
      baseName: "name",
      type: "string",
    },
    {
      name: "type",
      baseName: "type",
      type: "TemplateResponseDocumentFormField.TypeEnum",
    },
    {
      name: "signer",
      baseName: "signer",
      type: "string",
    },
    {
      name: "x",
      baseName: "x",
      type: "number",
    },
    {
      name: "y",
      baseName: "y",
      type: "number",
    },
    {
      name: "width",
      baseName: "width",
      type: "number",
    },
    {
      name: "height",
      baseName: "height",
      type: "number",
    },
    {
      name: "required",
      baseName: "required",
      type: "boolean",
    },
    {
      name: "group",
      baseName: "group",
      type: "string",
    },
    {
      name: "avgTextLength",
      baseName: "avg_text_length",
      type: "TemplateResponseFieldAvgTextLength",
    },
    {
      name: "isMultiline",
      baseName: "isMultiline",
      type: "boolean",
    },
    {
      name: "originalFontSize",
      baseName: "originalFontSize",
      type: "number",
    },
    {
      name: "fontFamily",
      baseName: "fontFamily",
      type: "string",
    },
  ];

  static getAttributeTypeMap(): AttributeTypeMap {
    return TemplateResponseDocumentFormField.attributeTypeMap;
  }
}

export namespace TemplateResponseDocumentFormField {
  export enum TypeEnum {
    Checkbox = "checkbox",
    CheckboxMerge = "checkbox-merge",
    DateSigned = "date_signed",
    Dropdown = "dropdown",
    Hyperlink = "hyperlink",
    Initials = "initials",
    Signature = "signature",
    Radio = "radio",
    Text = "text",
    TextMerge = "text-merge",
  }
}