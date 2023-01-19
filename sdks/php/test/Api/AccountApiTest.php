<?php

declare(strict_types=1);

namespace HelloSign\Test\Api;

use GuzzleHttp;
use HelloSign\Api;
use HelloSign\ApiException;
use HelloSign\Configuration;
use HelloSign\Model;
use HelloSign\Test\HelloTestCase;
use HelloSign\Test\TestUtils;

class AccountApiTest extends HelloTestCase
{
    /** @var Api\AccountApi */
    protected $api;

    protected function setUp(): void
    {
        parent::setUp();

        $this->client = new GuzzleHttp\Client([
            'handler' => GuzzleHttp\HandlerStack::create($this->handler),
        ]);

        $this->api = new Api\AccountApi(
            Configuration::getDefaultConfiguration(),
            $this->client,
        );
    }

    public function testHttpCodeRange()
    {
        $requestClass = Model\AccountVerifyRequest::class;
        $requestData = TestUtils::getFixtureData($requestClass)['default'];

        $responseClass = Model\ErrorResponse::class;
        $responseData = TestUtils::getFixtureData($responseClass)['default'];

        $this->setExpectedResponse($responseData, mt_rand(400, 499));

        $obj = Model\AccountVerifyRequest::init($requestData);

        $error = null;

        try {
            $this->api->accountVerify($obj);
        } catch (ApiException $e) {
            $error = $e->getResponseObject();
        }

        $this->assertInstanceOf($responseClass, $error);
        $this->assertEquals($responseData, TestUtils::toArray($error));
    }

    public function testAccountCreate()
    {
        $requestClass = Model\AccountCreateRequest::class;
        $requestData = TestUtils::getFixtureData($requestClass)['default'];

        $responseClass = Model\AccountCreateResponse::class;
        $responseData = TestUtils::getFixtureData($responseClass)['default'];

        $this->setExpectedResponse($responseData);

        $obj = Model\AccountCreateRequest::init($requestData);

        $response = $this->api->accountCreate($obj);

        $this->assertInstanceOf($responseClass, $response);
        $this->assertEquals($responseData, TestUtils::toArray($response));
    }

    public function testAccountGet()
    {
        $responseClass = Model\AccountGetResponse::class;
        $responseData = TestUtils::getFixtureData($responseClass)['default'];

        $this->setExpectedResponse($responseData);

        $response = $this->api->accountGet();

        $this->assertInstanceOf($responseClass, $response);
        $this->assertEquals($responseData, TestUtils::toArray($response));
    }

    public function testAccountUpdate()
    {
        $requestClass = Model\AccountUpdateRequest::class;
        $requestData = TestUtils::getFixtureData($requestClass)['default'];

        $responseClass = Model\AccountGetResponse::class;
        $responseData = TestUtils::getFixtureData($responseClass)['default'];

        $this->setExpectedResponse($responseData);

        $obj = Model\AccountUpdateRequest::init($requestData);

        $response = $this->api->accountUpdate($obj);

        $this->assertInstanceOf($responseClass, $response);
        $this->assertEquals($responseData, TestUtils::toArray($response));
    }

    public function testAccountVerify()
    {
        $requestClass = Model\AccountVerifyRequest::class;
        $requestData = TestUtils::getFixtureData($requestClass)['default'];

        $responseClass = Model\AccountVerifyResponse::class;
        $responseData = TestUtils::getFixtureData($responseClass)['default'];

        $this->setExpectedResponse($responseData);

        $obj = Model\AccountVerifyRequest::init($requestData);

        $response = $this->api->accountVerify($obj);

        $this->assertInstanceOf($responseClass, $response);
        $this->assertEquals($responseData, TestUtils::toArray($response));
    }
}
