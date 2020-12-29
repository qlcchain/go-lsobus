# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1QuotesIdGet**](QuotesApi.md#V1QuotesIdGet) | **Get** /v1/quotes/{id} | 
[**V1QuotesPost**](QuotesApi.md#V1QuotesPost) | **Post** /v1/quotes | 

# **V1QuotesIdGet**
> QuoteRes V1QuotesIdGet(ctx, id)


Get quote by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id to query with | 

### Return type

[**QuoteRes**](quoteRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1QuotesPost**
> QuoteRes V1QuotesPost(ctx, body)


Creates a quote

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QuoteCreateReq**](QuoteCreateReq.md)|  | 

### Return type

[**QuoteRes**](quoteRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

