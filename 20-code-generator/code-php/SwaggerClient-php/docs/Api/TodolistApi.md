# Swagger\Client\TodolistApi

All URIs are relative to *https://{environment}.ariefbelajarteknologi.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**todolistGet**](TodolistApi.md#todolistget) | **GET** /todolist | Get All Todolist
[**todolistPost**](TodolistApi.md#todolistpost) | **POST** /todolist | Create new todolist
[**todolistTodolistIdDelete**](TodolistApi.md#todolisttodolistiddelete) | **DELETE** /todolist/{todolistId} | Delele existing todolist
[**todolistTodolistIdPut**](TodolistApi.md#todolisttodolistidput) | **PUT** /todolist/{todolistId} | Update existing todolist

# **todolistGet**
> \Swagger\Client\Model\ArrayTodolist todolistGet($include_done, $name)

Get All Todolist

Get all active Todolist

### Example
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
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include_done** | **bool**| Include done todolist in the result | [optional] [default to false]
 **name** | **string**| Search todolist by name | [optional]

### Return type

[**\Swagger\Client\Model\ArrayTodolist**](../Model/ArrayTodolist.md)

### Authorization

[TodolistAuth](../../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **todolistPost**
> \Swagger\Client\Model\Todolist todolistPost($body)

Create new todolist

Create new active todolist to stored in database

### Example
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
$body = new \Swagger\Client\Model\CreateOrUpdateTodolist(); // \Swagger\Client\Model\CreateOrUpdateTodolist | 

try {
    $result = $apiInstance->todolistPost($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistPost: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\Swagger\Client\Model\CreateOrUpdateTodolist**](../Model/CreateOrUpdateTodolist.md)|  |

### Return type

[**\Swagger\Client\Model\Todolist**](../Model/Todolist.md)

### Authorization

[TodolistAuth](../../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **todolistTodolistIdDelete**
> \Swagger\Client\Model\InlineResponse200 todolistTodolistIdDelete($todolist_id)

Delele existing todolist

Delele existing todolist in database

### Example
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
$todolist_id = "todolist_id_example"; // string | Todolist Id for update todo in database

try {
    $result = $apiInstance->todolistTodolistIdDelete($todolist_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TodolistApi->todolistTodolistIdDelete: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todolist_id** | **string**| Todolist Id for update todo in database |

### Return type

[**\Swagger\Client\Model\InlineResponse200**](../Model/InlineResponse200.md)

### Authorization

[TodolistAuth](../../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **todolistTodolistIdPut**
> \Swagger\Client\Model\Todolist todolistTodolistIdPut($body, $todolist_id)

Update existing todolist

Update existing todolist in database

### Example
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

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\Swagger\Client\Model\CreateOrUpdateTodolist**](../Model/CreateOrUpdateTodolist.md)|  |
 **todolist_id** | **string**| Todolist Id for update todo in database |

### Return type

[**\Swagger\Client\Model\Todolist**](../Model/Todolist.md)

### Authorization

[TodolistAuth](../../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

