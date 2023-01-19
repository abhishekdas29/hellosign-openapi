=begin
#HelloSign API

#HelloSign v3 API

The version of the OpenAPI document: 3.0.0
Contact: apisupport@hellosign.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.3.0

=end

require 'spec_helper'
require 'json_spec'
require_relative '../test_utils'

describe HelloSign::AccountApi do
  context 'AccountApiTest' do
    api = HelloSign::AccountApi.new

    it 'testHttpCodeRange' do
      request_class = 'AccountCreateRequest'
      request_data = get_fixture_data(request_class)[:default]

      response_class = 'ErrorResponse'
      response_data = get_fixture_data(response_class)[:default]

      code = rand(400..499)

      set_expected_response(code, JSON.dump(response_data))
      expected = HelloSign::ErrorResponse.init(response_data)
      obj = HelloSign::AccountCreateRequest.init(request_data)

      begin
        result = api.account_create(obj)
        fail_with("Should have thrown error: #{result}")
      rescue HelloSign::ApiError => e
        expect(e.response_body.class.to_s).to eq("HelloSign::#{response_class}")
        expect(e.response_body.to_json).to be_json_eql(JSON.dump(expected))
      end
    end

    it 'testAccountCreate' do
      request_class = 'AccountCreateRequest'
      request_data = get_fixture_data(request_class)[:default]

      response_class = 'AccountCreateResponse'
      response_data = get_fixture_data(response_class)[:default]

      set_expected_response(200, JSON.dump(response_data))
      expected = HelloSign::AccountCreateResponse.init(response_data)
      obj = HelloSign::AccountCreateRequest.init(request_data)

      result = api.account_create(obj)

      expect(result.class.to_s).to eq("HelloSign::#{response_class}")
      expect(result.to_json).to be_json_eql(JSON.dump(expected))
    end

    it 'testAccountGet' do
      response_class = 'AccountGetResponse'
      response_data = get_fixture_data(response_class)[:default]

      set_expected_response(200, JSON.dump(response_data))
      expected = HelloSign::AccountGetResponse.init(response_data)

      result = api.account_get({ email_address: "jack@example.com" })

      expect(result.class.to_s).to eq("HelloSign::#{response_class}")
      expect(result.to_json).to be_json_eql(JSON.dump(expected))
    end

    it 'testAccountUpdate' do
      request_class = 'AccountUpdateRequest'
      request_data = get_fixture_data(request_class)[:default]

      response_class = 'AccountGetResponse'
      response_data = get_fixture_data(response_class)[:default]

      set_expected_response(200, JSON.dump(response_data))
      expected = HelloSign::AccountGetResponse.init(response_data)
      obj = HelloSign::AccountUpdateRequest.init(request_data)

      result = api.account_update(obj)

      expect(result.class.to_s).to eq("HelloSign::#{response_class}")
      expect(result.to_json).to be_json_eql(JSON.dump(expected))
    end

    it 'testAccountVerify' do
      request_class = 'AccountVerifyRequest'
      request_data = get_fixture_data(request_class)[:default]

      response_class = 'AccountVerifyResponse'
      response_data = get_fixture_data(response_class)[:default]

      set_expected_response(200, JSON.dump(response_data))
      expected = HelloSign::AccountVerifyResponse.init(response_data)
      obj = HelloSign::AccountVerifyRequest.init(request_data)

      result = api.account_verify(obj)

      expect(result.class.to_s).to eq("HelloSign::#{response_class}")
      expect(result.to_json).to be_json_eql(JSON.dump(expected))
    end
  end
end
