# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DltOrderSellerChangeOrderRewardBlockPost**](DLTOrdersSellerApi.md#V1DltOrderSellerChangeOrderRewardBlockPost) | **Post** /v1/dlt/order/seller/change-order-reward-block | 
[**V1DltOrderSellerCreateOrderRewardBlockPost**](DLTOrdersSellerApi.md#V1DltOrderSellerCreateOrderRewardBlockPost) | **Post** /v1/dlt/order/seller/create-order-reward-block | 
[**V1DltOrderSellerPendingRequestPost**](DLTOrdersSellerApi.md#V1DltOrderSellerPendingRequestPost) | **Post** /v1/dlt/order/seller/pending-request | 
[**V1DltOrderSellerPendingResourceCheckPost**](DLTOrdersSellerApi.md#V1DltOrderSellerPendingResourceCheckPost) | **Post** /v1/dlt/order/seller/pending-resource-check | 
[**V1DltOrderSellerTerminateOrderRewardBlockPost**](DLTOrdersSellerApi.md#V1DltOrderSellerTerminateOrderRewardBlockPost) | **Post** /v1/dlt/order/seller/terminate-order-reward-block | 
[**V1DltOrderSellerUpdateOrderInfoRewardBlockPost**](DLTOrdersSellerApi.md#V1DltOrderSellerUpdateOrderInfoRewardBlockPost) | **Post** /v1/dlt/order/seller/update-order-info-reward-block | 
[**V1DltOrderSellerUpdateProductInfoBlockPost**](DLTOrdersSellerApi.md#V1DltOrderSellerUpdateProductInfoBlockPost) | **Post** /v1/dlt/order/seller/update-product-info-block | 

# **V1DltOrderSellerChangeOrderRewardBlockPost**
> DltOrderSellerChangeOrderRewardBlockRes V1DltOrderSellerChangeOrderRewardBlockPost(ctx, body)


Generate a block to confirm or reject a changing request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerChangeOrderRewardBlockReq**](DltOrderSellerChangeOrderRewardBlockReq.md)|  | 

### Return type

[**DltOrderSellerChangeOrderRewardBlockRes**](dltOrderSellerChangeOrderRewardBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerCreateOrderRewardBlockPost**
> DltOrderSellerCreateOrderRewardBlockRes V1DltOrderSellerCreateOrderRewardBlockPost(ctx, body)


Generate a block to confirm or reject a creating request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerCreateOrderRewardBlockReq**](DltOrderSellerCreateOrderRewardBlockReq.md)|  | 

### Return type

[**DltOrderSellerCreateOrderRewardBlockRes**](dltOrderSellerCreateOrderRewardBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerPendingRequestPost**
> DltOrderSellerPendingRequestRes V1DltOrderSellerPendingRequestPost(ctx, body)


Get all pending requests for seller

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerPendingRequestReq**](DltOrderSellerPendingRequestReq.md)|  | 

### Return type

[**DltOrderSellerPendingRequestRes**](dltOrderSellerPendingRequestRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerPendingResourceCheckPost**
> DltOrderSellerPendingResourceCheckRes V1DltOrderSellerPendingResourceCheckPost(ctx, body)


Get all pending orders for seller to check. Seller need to check every product's status in each order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerPendingResourceCheckReq**](DltOrderSellerPendingResourceCheckReq.md)|  | 

### Return type

[**DltOrderSellerPendingResourceCheckRes**](dltOrderSellerPendingResourceCheckRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerTerminateOrderRewardBlockPost**
> DltOrderSellerTerminateOrderRewardBlockRes V1DltOrderSellerTerminateOrderRewardBlockPost(ctx, body)


Generate a block to confirm or reject a terminating request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerTerminateOrderRewardBlockReq**](DltOrderSellerTerminateOrderRewardBlockReq.md)|  | 

### Return type

[**DltOrderSellerTerminateOrderRewardBlockRes**](dltOrderSellerTerminateOrderRewardBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerUpdateOrderInfoRewardBlockPost**
> DltOrderSellerUpdateOrderInfoRewardBlockRes V1DltOrderSellerUpdateOrderInfoRewardBlockPost(ctx, body)


Generate a block to update order state to complete when all products in this order can be used normally

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerUpdateOrderInfoRewardBlockReq**](DltOrderSellerUpdateOrderInfoRewardBlockReq.md)|  | 

### Return type

[**DltOrderSellerUpdateOrderInfoRewardBlockRes**](dltOrderSellerUpdateOrderInfoRewardBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderSellerUpdateProductInfoBlockPost**
> DltOrderSellerUpdateProductInfoBlockRes V1DltOrderSellerUpdateProductInfoBlockPost(ctx, body)


Generate a block to update product id and product status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderSellerUpdateProductInfoBlockReq**](DltOrderSellerUpdateProductInfoBlockReq.md)|  | 

### Return type

[**DltOrderSellerUpdateProductInfoBlockRes**](dltOrderSellerUpdateProductInfoBlockRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

