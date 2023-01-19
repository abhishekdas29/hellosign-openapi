=begin
#Dropbox Sign API

#Dropbox Sign v3 API

The version of the OpenAPI document: 3.0.0
Contact: apisupport@hellosign.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.3.0

=end

require 'date'
require 'time'

module Dropbox
end

module Dropbox::Sign
  class SignatureRequestResponseCustomFieldTypeEnum
    TEXT = "text".freeze
    CHECKBOX = "checkbox".freeze

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      constantValues = SignatureRequestResponseCustomFieldTypeEnum.constants.select { |c| SignatureRequestResponseCustomFieldTypeEnum::const_get(c) == value }
      raise "Invalid ENUM value #{value} for class #SignatureRequestResponseCustomFieldTypeEnum" if constantValues.empty?
      value
    end
  end
end