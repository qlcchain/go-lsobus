# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrdersIdGet**](OrdersApi.md#V1OrdersIdGet) | **Get** /v1/orders/{id} | 
[**V1OrdersOrderPost**](OrdersApi.md#V1OrdersOrderPost) | **Post** /v1/orders/order | 
[**V1OrdersPost**](OrdersApi.md#V1OrdersPost) | **Post** /v1/orders | 
[**V1OrdersProductInventoryGet**](OrdersApi.md#V1OrdersProductInventoryGet) | **Get** /v1/orders/product-inventory | 
[**V1OrdersProductInventoryIdGet**](OrdersApi.md#V1OrdersProductInventoryIdGet) | **Get** /v1/orders/product-inventory/{id} | 
[**V1OrdersSmartContractPost**](OrdersApi.md#V1OrdersSmartContractPost) | **Post** /v1/orders/smart-contract | 

# **V1OrdersIdGet**
> OrderGetRes V1OrdersIdGet(ctx, id)


Get order details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Order id to query with | 

### Return type

[**OrderGetRes**](orderGetRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrdersOrderPost**
> OrderGetRes V1OrdersOrderPost(ctx, body)


Create a DoD order with given SC tx id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderCreateOrderReq**](OrderCreateOrderReq.md)|  | 

### Return type

[**OrderGetRes**](orderGetRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrdersPost**
> OrderGetRes V1OrdersPost(ctx, body)


Create a DoD order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderCreateReq**](OrderCreateReq.md)|  | 

### Return type

[**OrderGetRes**](orderGetRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrdersProductInventoryGet**
> []ProductInventoryGetModel V1OrdersProductInventoryGet(ctx, )


Get product inventory

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductInventoryGetModel**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrdersProductInventoryIdGet**
> ProductInventoryGetModel V1OrdersProductInventoryIdGet(ctx, id)


Get product inventory by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id to query with | 

### Return type

[**ProductInventoryGetModel**](productInventoryGetModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1OrdersSmartContractPost**
> OrderSmartContractRes V1OrdersSmartContractPost(ctx, body)


Create new smart contract

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderCreateSmartContractReq**](OrderCreateSmartContractReq.md)|  | 

### Return type

[**OrderSmartContractRes**](orderSmartContractRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

