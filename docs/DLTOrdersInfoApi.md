# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DltOrderInfoByInternalIdPost**](DLTOrdersInfoApi.md#V1DltOrderInfoByInternalIdPost) | **Post** /v1/dlt/order/info/by-internal-id | 
[**V1DltOrderInfoBySellerAndOrderIdPost**](DLTOrdersInfoApi.md#V1DltOrderInfoBySellerAndOrderIdPost) | **Post** /v1/dlt/order/info/by-seller-and-order-id | 
[**V1DltOrderInfoInternalIdByOrderIdPost**](DLTOrdersInfoApi.md#V1DltOrderInfoInternalIdByOrderIdPost) | **Post** /v1/dlt/order/info/internal-id-by-order-id | 
[**V1DltOrderInfoOrderCountByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderCountByAddressAndSellerPost) | **Post** /v1/dlt/order/info/order-count-by-address-and-seller | 
[**V1DltOrderInfoOrderCountByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderCountByAddressPost) | **Post** /v1/dlt/order/info/order-count-by-address | 
[**V1DltOrderInfoOrderIdListByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderIdListByAddressAndSellerPost) | **Post** /v1/dlt/order/info/order-id-list-by-address-and-seller | 
[**V1DltOrderInfoOrderIdListByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderIdListByAddressPost) | **Post** /v1/dlt/order/info/order-id-list-by-address | 
[**V1DltOrderInfoOrderInfoByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderInfoByAddressAndSellerPost) | **Post** /v1/dlt/order/info/order-info-by-address-and-seller | 
[**V1DltOrderInfoOrderInfoByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoOrderInfoByAddressPost) | **Post** /v1/dlt/order/info/order-info-by-address | 
[**V1DltOrderInfoPlacingOrderPost**](DLTOrdersInfoApi.md#V1DltOrderInfoPlacingOrderPost) | **Post** /v1/dlt/order/info/placing-order | 
[**V1DltOrderInfoProductCountByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductCountByAddressAndSellerPost) | **Post** /v1/dlt/order/info/product-count-by-address-and-seller | 
[**V1DltOrderInfoProductCountByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductCountByAddressPost) | **Post** /v1/dlt/order/info/product-count-by-address | 
[**V1DltOrderInfoProductIdListByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductIdListByAddressAndSellerPost) | **Post** /v1/dlt/order/info/product-id-list-by-address-and-seller | 
[**V1DltOrderInfoProductIdListByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductIdListByAddressPost) | **Post** /v1/dlt/order/info/product-id-list-by-address | 
[**V1DltOrderInfoProductInfoByAddressAndSellerPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductInfoByAddressAndSellerPost) | **Post** /v1/dlt/order/info/product-info-by-address-and-seller | 
[**V1DltOrderInfoProductInfoByAddressPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductInfoByAddressPost) | **Post** /v1/dlt/order/info/product-info-by-address | 
[**V1DltOrderInfoProductInfoBySellerAndProductIdPost**](DLTOrdersInfoApi.md#V1DltOrderInfoProductInfoBySellerAndProductIdPost) | **Post** /v1/dlt/order/info/product-info-by-seller-and-product-id | 

# **V1DltOrderInfoByInternalIdPost**
> DltOrderInfoOrderDetailRes V1DltOrderInfoByInternalIdPost(ctx, body)


Get order info by internal id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoByInternalIdReq**](DltOrderInfoByInternalIdReq.md)|  | 

### Return type

[**DltOrderInfoOrderDetailRes**](dltOrderInfoOrderDetailRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoBySellerAndOrderIdPost**
> DltOrderInfoOrderDetailRes V1DltOrderInfoBySellerAndOrderIdPost(ctx, body)


Get order info by seller address and order id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoBySellerAndOrderIdReq**](DltOrderInfoBySellerAndOrderIdReq.md)|  | 

### Return type

[**DltOrderInfoOrderDetailRes**](dltOrderInfoOrderDetailRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoInternalIdByOrderIdPost**
> DltOrderInfoInternalIdByOrderIdRes V1DltOrderInfoInternalIdByOrderIdPost(ctx, body)


Get internal id by seller address and order id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoInternalIdByOrderIdReq**](DltOrderInfoInternalIdByOrderIdReq.md)|  | 

### Return type

[**DltOrderInfoInternalIdByOrderIdRes**](dltOrderInfoInternalIdByOrderIdRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderCountByAddressAndSellerPost**
> DltOrderInfoOrderCountByAddressAndSellerRes V1DltOrderInfoOrderCountByAddressAndSellerPost(ctx, body)


Get order count by buyer's address and seller's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderCountByAddressAndSellerReq**](DltOrderInfoOrderCountByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoOrderCountByAddressAndSellerRes**](dltOrderInfoOrderCountByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderCountByAddressPost**
> DltOrderInfoOrderCountByAddressRes V1DltOrderInfoOrderCountByAddressPost(ctx, body)


Get order count by buyer's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderCountByAddressReq**](DltOrderInfoOrderCountByAddressReq.md)|  | 

### Return type

[**DltOrderInfoOrderCountByAddressRes**](dltOrderInfoOrderCountByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderIdListByAddressAndSellerPost**
> DltOrderInfoOrderIdListByAddressAndSellerRes V1DltOrderInfoOrderIdListByAddressAndSellerPost(ctx, body)


Get buyer's all order ids with specified seller

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderIdListByAddressAndSellerReq**](DltOrderInfoOrderIdListByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoOrderIdListByAddressAndSellerRes**](dltOrderInfoOrderIdListByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderIdListByAddressPost**
> DltOrderInfoOrderIdListByAddressRes V1DltOrderInfoOrderIdListByAddressPost(ctx, body)


Get buyer's all order ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderIdListByAddressReq**](DltOrderInfoOrderIdListByAddressReq.md)|  | 

### Return type

[**DltOrderInfoOrderIdListByAddressRes**](dltOrderInfoOrderIdListByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderInfoByAddressAndSellerPost**
> DltOrderInfoOrderInfoByAddressAndSellerRes V1DltOrderInfoOrderInfoByAddressAndSellerPost(ctx, body)


Get order info by buyer's address and seller's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderInfoByAddressAndSellerReq**](DltOrderInfoOrderInfoByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoOrderInfoByAddressAndSellerRes**](dltOrderInfoOrderInfoByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoOrderInfoByAddressPost**
> DltOrderInfoOrderInfoByAddressRes V1DltOrderInfoOrderInfoByAddressPost(ctx, body)


Get order info by buyer's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoOrderInfoByAddressReq**](DltOrderInfoOrderInfoByAddressReq.md)|  | 

### Return type

[**DltOrderInfoOrderInfoByAddressRes**](dltOrderInfoOrderInfoByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoPlacingOrderPost**
> DltOrderInfoPlacingOrderRes V1DltOrderInfoPlacingOrderPost(ctx, body)


Get all placing orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoPlacingOrderReq**](DltOrderInfoPlacingOrderReq.md)|  | 

### Return type

[**DltOrderInfoPlacingOrderRes**](dltOrderInfoPlacingOrderRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductCountByAddressAndSellerPost**
> DltOrderInfoProductCountByAddressAndSellerRes V1DltOrderInfoProductCountByAddressAndSellerPost(ctx, body)


Get product count by buyer's address and seller's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductCountByAddressAndSellerReq**](DltOrderInfoProductCountByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoProductCountByAddressAndSellerRes**](dltOrderInfoProductCountByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductCountByAddressPost**
> DltOrderInfoProductCountByAddressRes V1DltOrderInfoProductCountByAddressPost(ctx, body)


Get product count by buyer's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductCountByAddressReq**](DltOrderInfoProductCountByAddressReq.md)|  | 

### Return type

[**DltOrderInfoProductCountByAddressRes**](dltOrderInfoProductCountByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductIdListByAddressAndSellerPost**
> DltOrderInfoProductIdListByAddressAndSellerRes V1DltOrderInfoProductIdListByAddressAndSellerPost(ctx, body)


Get buyer's all product ids with specified seller

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductIdListByAddressAndSellerReq**](DltOrderInfoProductIdListByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoProductIdListByAddressAndSellerRes**](dltOrderInfoProductIdListByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductIdListByAddressPost**
> DltOrderInfoProductIdListByAddressRes V1DltOrderInfoProductIdListByAddressPost(ctx, body)


Get buyer's all product ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductIdListByAddressReq**](DltOrderInfoProductIdListByAddressReq.md)|  | 

### Return type

[**DltOrderInfoProductIdListByAddressRes**](dltOrderInfoProductIdListByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductInfoByAddressAndSellerPost**
> DltOrderInfoProductInfoByAddressAndSellerRes V1DltOrderInfoProductInfoByAddressAndSellerPost(ctx, body)


Get product info by buyer's address and seller's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductInfoByAddressAndSellerReq**](DltOrderInfoProductInfoByAddressAndSellerReq.md)|  | 

### Return type

[**DltOrderInfoProductInfoByAddressAndSellerRes**](dltOrderInfoProductInfoByAddressAndSellerRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductInfoByAddressPost**
> DltOrderInfoProductInfoByAddressRes V1DltOrderInfoProductInfoByAddressPost(ctx, body)


Get product info by buyer's address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductInfoByAddressReq**](DltOrderInfoProductInfoByAddressReq.md)|  | 

### Return type

[**DltOrderInfoProductInfoByAddressRes**](dltOrderInfoProductInfoByAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltOrderInfoProductInfoBySellerAndProductIdPost**
> DltOrderInfoProductInfoBySellerAndProductIdRes V1DltOrderInfoProductInfoBySellerAndProductIdPost(ctx, body)


Get product info by seller address and product id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltOrderInfoProductInfoBySellerAndProductIdReq**](DltOrderInfoProductInfoBySellerAndProductIdReq.md)|  | 

### Return type

[**DltOrderInfoProductInfoBySellerAndProductIdRes**](dltOrderInfoProductInfoBySellerAndProductIdRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

