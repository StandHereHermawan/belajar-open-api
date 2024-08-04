# {{classname}}

All URIs are relative to *https://{environment}.ariefbelajarteknologi.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TodolistGet**](TodolistApi.md#TodolistGet) | **Get** /todolist | Get All Todolist
[**TodolistPost**](TodolistApi.md#TodolistPost) | **Post** /todolist | Create new todolist
[**TodolistTodolistIdDelete**](TodolistApi.md#TodolistTodolistIdDelete) | **Delete** /todolist/{todolistId} | Delele existing todolist
[**TodolistTodolistIdPut**](TodolistApi.md#TodolistTodolistIdPut) | **Put** /todolist/{todolistId} | Update existing todolist

# **TodolistGet**
> []Todolist TodolistGet(ctx, optional)
Get All Todolist

Get all active Todolist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TodolistApiTodolistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TodolistApiTodolistGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDone** | **optional.Bool**| Include done todolist in the result | [default to false]
 **name** | **optional.String**| Search todolist by name | 

### Return type

[**[]Todolist**](array.md)

### Authorization

[TodolistAuth](../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistPost**
> Todolist TodolistPost(ctx, body)
Create new todolist

Create new active todolist to stored in database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateTodolist**](CreateOrUpdateTodolist.md)|  | 

### Return type

[**Todolist**](Todolist.md)

### Authorization

[TodolistAuth](../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistTodolistIdDelete**
> InlineResponse200 TodolistTodolistIdDelete(ctx, todolistId)
Delele existing todolist

Delele existing todolist in database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todolistId** | **string**| Todolist Id for update todo in database | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[TodolistAuth](../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistTodolistIdPut**
> Todolist TodolistTodolistIdPut(ctx, body, todolistId)
Update existing todolist

Update existing todolist in database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateTodolist**](CreateOrUpdateTodolist.md)|  | 
  **todolistId** | **string**| Todolist Id for update todo in database | 

### Return type

[**Todolist**](Todolist.md)

### Authorization

[TodolistAuth](../README.md#TodolistAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

