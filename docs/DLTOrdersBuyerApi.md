# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DltOrderBuyerChangeOrderBlockPost**](DLTOrdersBuyerApi.md#V1DltOrderBuyerChangeOrderBlockPost) | **Post** /v1/dlt/order/buyer/change-order-block | 
[**V1DltOrderBuyerCreateOrderBlockPost**](DLTOrdersBuyerApi.md#V1DltOrderBuyerCreateOrderBlockPost) | **Post** /v1/dlt/order/buyer/create-order-block | 
[**V1DltOrderBuyerTerminateOrderBlockPost**](DLTOrdersBuyerApi.md#V1DltOrderBuyerTerminateOrderBlockPost) | **Post** /v1/dlt/order/buyer/terminate-order-block | 
[**V1DltOrderBuyerUpdateOrderInfoBlockPost**](DLTOrdersBuyerApi.md#V1DltOrderBuyerUpdateOrderInfoBlockPost) | **Post** /v1/dlt/order/buyer/update-order-info-block | 

# **V1DltOrderBuyerChangeOrderBlockPost**
> DltOrderBuyerChangeOrderBlockRes V1DltOrderBuyerChangeOrderBlockPost(ctx, body)


Generate a block to create a request for changing a connection or other product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderBuyerChangeOrderBlockReq**](DltOrderBuyerChangeOrderBlockReq.md)|  | 

### Return type

[**DltOrderBuyerChangeOrderBlockRes**](dltOrderBuyerChangeOrderBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderBuyerCreateOrderBlockPost**
> DltOrderBuyerCreateOrderBlockRes V1DltOrderBuyerCreateOrderBlockPost(ctx, body)


Generate a block to create a request for creating a connection or other product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderBuyerCreateOrderBlockReq**](DltOrderBuyerCreateOrderBlockReq.md)|  | 

### Return type

[**DltOrderBuyerCreateOrderBlockRes**](dltOrderBuyerCreateOrderBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderBuyerTerminateOrderBlockPost**
> DltOrderBuyerTerminateOrderBlockRes V1DltOrderBuyerTerminateOrderBlockPost(ctx, body)


Generate a block to create a request for terminating a connection or other product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderBuyerTerminateOrderBlockReq**](DltOrderBuyerTerminateOrderBlockReq.md)|  | 

### Return type

[**DltOrderBuyerTerminateOrderBlockRes**](dltOrderBuyerTerminateOrderBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderBuyerUpdateOrderInfoBlockPost**
> DltOrderBuyerUpdateOrderInfoBlockRes V1DltOrderBuyerUpdateOrderInfoBlockPost(ctx, body)


Generate a block to update order id and product id to chain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderBuyerUpdateOrderInfoBlockReq**](DltOrderBuyerUpdateOrderInfoBlockReq.md)|  | 

### Return type

[**DltOrderBuyerUpdateOrderInfoBlockRes**](dltOrderBuyerUpdateOrderInfoBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

