# {{classname}}

All URIs are relative to *https://api.openfigi.com/{basePath}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MappingPost**](DefaultApi.md#MappingPost) | **Post** /mapping | 
[**MappingValuesKeyGet**](DefaultApi.md#MappingValuesKeyGet) | **Get** /mapping/values/{key} | 

# **MappingPost**
> []MappingJobResult MappingPost(ctx, optional)


Allows mapping from third-party identifiers to FIGIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiMappingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiMappingPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []MappingJob**](MappingJob.md)| A list of third-party identifiers and extra filters. | 

### Return type

[**[]MappingJobResult**](array.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MappingValuesKeyGet**
> InlineResponse200 MappingValuesKeyGet(ctx, key)


Get values for enum-like fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Key of MappingJob for which to get possible values. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

