# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1QualificationProductMetronetPost**](QualificationsApi.md#V1QualificationProductMetronetPost) | **Post** /v1/qualification/product-metronet | 
[**V1QualificationValidateAddressPost**](QualificationsApi.md#V1QualificationValidateAddressPost) | **Post** /v1/qualification/validate-address | 

# **V1QualificationProductMetronetPost**
> QualificationProductMetronetRes V1QualificationProductMetronetPost(ctx, body)


Product metronet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**QualificationProductMetronetReq**](QualificationProductMetronetReq.md)|  | 

### Return type

[**QualificationProductMetronetRes**](qualificationProductMetronetRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1QualificationValidateAddressPost**
> QualificationValidateAddressRes V1QualificationValidateAddressPost(ctx, optional)


Validate geo address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QualificationsApiV1QualificationValidateAddressPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualificationsApiV1QualificationValidateAddressPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QualificationValidateAddressReq**](QualificationValidateAddressReq.md)|  | 

### Return type

[**QualificationValidateAddressRes**](qualificationValidateAddressRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

