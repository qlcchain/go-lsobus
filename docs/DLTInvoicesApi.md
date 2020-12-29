# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DltInvoiceGenerateByBuyerPost**](DLTInvoicesApi.md#V1DltInvoiceGenerateByBuyerPost) | **Post** /v1/dlt/invoice/generate-by-buyer | 
[**V1DltInvoiceGenerateByOrderIdPost**](DLTInvoicesApi.md#V1DltInvoiceGenerateByOrderIdPost) | **Post** /v1/dlt/invoice/generate-by-order-id | 
[**V1DltInvoiceGenerateByProductIdPost**](DLTInvoicesApi.md#V1DltInvoiceGenerateByProductIdPost) | **Post** /v1/dlt/invoice/generate-by-product-id | 

# **V1DltInvoiceGenerateByBuyerPost**
> DltInvoiceGenerateInvoiceByBuyerRes V1DltInvoiceGenerateByBuyerPost(ctx, body)


Generate invoice by buyer's qlc address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltInvoiceGenerateInvoiceByBuyerReq**](DltInvoiceGenerateInvoiceByBuyerReq.md)|  | 

### Return type

[**DltInvoiceGenerateInvoiceByBuyerRes**](dltInvoiceGenerateInvoiceByBuyerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltInvoiceGenerateByOrderIdPost**
> DltInvoiceGenerateInvoiceByOrderIdRes V1DltInvoiceGenerateByOrderIdPost(ctx, body)


Generate invoice by order id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltInvoiceGenerateInvoiceByOrderIdReq**](DltInvoiceGenerateInvoiceByOrderIdReq.md)|  | 

### Return type

[**DltInvoiceGenerateInvoiceByOrderIdRes**](dltInvoiceGenerateInvoiceByOrderIdRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltInvoiceGenerateByProductIdPost**
> DltInvoiceGenerateInvoiceByProductIdRes V1DltInvoiceGenerateByProductIdPost(ctx, body)


Generate invoice by product id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltInvoiceGenerateInvoiceByProductIdReq**](DltInvoiceGenerateInvoiceByProductIdReq.md)|  | 

### Return type

[**DltInvoiceGenerateInvoiceByProductIdRes**](dltInvoiceGenerateInvoiceByProductIdRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

