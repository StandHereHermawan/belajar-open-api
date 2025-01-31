# SwaggerClient-php
OpenAPI for Todolist RESTful API

This PHP package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 1
- Build package: io.swagger.codegen.v3.generators.php.PhpClientCodegen
For more information, please visit [https://web.facebook.com/ARIEF.KARDITYA](https://web.facebook.com/ARIEF.KARDITYA)

## Requirements

PHP 5.5 and later

## Installation & Usage
### Composer

To install the bindings via [Composer](http://getcomposer.org/), add the following to `composer.json`:

```
{
  "repositories": [
    {
      "type": "git",
      "url": "https://github.com/git_user_id/git_repo_id.git"
    }
  ],
  "require": {
    "git_user_id/git_repo_id": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
    require_once('/path/to/SwaggerClient-php/vendor/autoload.php');
```

## Tests

To run the unit tests:

```
composer install
./vendor/bin/phpunit
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: TodolistAuth
$config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('X-API-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-API-KEY', 'Bearer');

$apiInstance = new Swagger\Client\Api\TodolistApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$include_done = false; // bool | Include done todolist in the result
$name = "name_example"; // string | Search todolist by name

try {
    $result = $apiInstance->todolistGet($include_done, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistGet: ', $e->getMessage(), PHP_EOL;
}

// Configure API key authorization: TodolistAuth
$config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('X-API-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-API-KEY', 'Bearer');

$apiInstance = new Swagger\Client\Api\TodolistApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$body = new \Swagger\Client\Model\CreateOrUpdateTodolist(); // \Swagger\Client\Model\CreateOrUpdateTodolist | 

try {
    $result = $apiInstance->todolistPost($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistPost: ', $e->getMessage(), PHP_EOL;
}

// Configure API key authorization: TodolistAuth
$config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('X-API-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-API-KEY', 'Bearer');

$apiInstance = new Swagger\Client\Api\TodolistApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$todolist_id = "todolist_id_example"; // string | Todolist Id for update todo in database

try {
    $result = $apiInstance->todolistTodolistIdDelete($todolist_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistTodolistIdDelete: ', $e->getMessage(), PHP_EOL;
}

// Configure API key authorization: TodolistAuth
$config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('X-API-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-API-KEY', 'Bearer');

$apiInstance = new Swagger\Client\Api\TodolistApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$body = new \Swagger\Client\Model\CreateOrUpdateTodolist(); // \Swagger\Client\Model\CreateOrUpdateTodolist | 
$todolist_id = "todolist_id_example"; // string | Todolist Id for update todo in database

try {
    $result = $apiInstance->todolistTodolistIdPut($body, $todolist_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistTodolistIdPut: ', $e->getMessage(), PHP_EOL;
}
?>
```

## Documentation for API Endpoints

All URIs are relative to *https://{environment}.ariefbelajarteknologi.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*TodolistApi* | [**todolistGet**](docs/Api/TodolistApi.md#todolistget) | **GET** /todolist | Get All Todolist
*TodolistApi* | [**todolistPost**](docs/Api/TodolistApi.md#todolistpost) | **POST** /todolist | Create new todolist
*TodolistApi* | [**todolistTodolistIdDelete**](docs/Api/TodolistApi.md#todolisttodolistiddelete) | **DELETE** /todolist/{todolistId} | Delele existing todolist
*TodolistApi* | [**todolistTodolistIdPut**](docs/Api/TodolistApi.md#todolisttodolistidput) | **PUT** /todolist/{todolistId} | Update existing todolist

## Documentation For Models

 - [ArrayTodolist](docs/Model/ArrayTodolist.md)
 - [CreateOrUpdateTodolist](docs/Model/CreateOrUpdateTodolist.md)
 - [InlineResponse200](docs/Model/InlineResponse200.md)
 - [Todolist](docs/Model/Todolist.md)

## Documentation For Authorization


## TodolistAuth

- **Type**: API key
- **API key parameter name**: X-API-KEY
- **Location**: HTTP header


## Author

ariefkardityahrmwn@gmail.com

