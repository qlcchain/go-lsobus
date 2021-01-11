# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DltLedgerBlockConfirmedStatusHashGet**](DLTLedgerApi.md#V1DltLedgerBlockConfirmedStatusHashGet) | **Get** /v1/dlt/ledger/block-confirmed-status/{hash} | 
[**V1DltLedgerProcessPost**](DLTLedgerApi.md#V1DltLedgerProcessPost) | **Post** /v1/dlt/ledger/process | 

# **V1DltLedgerBlockConfirmedStatusHashGet**
> DltLedgerBlockConfirmedStatusRes V1DltLedgerBlockConfirmedStatusHashGet(ctx, hash)


Return block confirmed status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Blocks hash | 

### Return type

[**DltLedgerBlockConfirmedStatusRes**](dltLedgerBlockConfirmedStatusRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DltLedgerProcessPost**
> DltLedgerProcessRes V1DltLedgerProcessPost(ctx, body)


Check block base info, update chain info for the block, and broadcast block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DltLedgerProcessReq**](DltLedgerProcessReq.md)|  | 

### Return type

[**DltLedgerProcessRes**](dltLedgerProcessRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

