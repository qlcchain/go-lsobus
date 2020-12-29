# DltInvoiceGenerateInvoiceByOrderIdReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QlcAddressSeller** | **string** |  | [default to null]
**OrderId** | **string** |  | [default to null]
**StartTime** | **float64** |  | [default to null]
**EndTime** | **float64** |  | [default to null]
**InFlight** | **bool** | Order allowed (true &#x3D; in-flight order will be included) | [optional] [default to null]
**Split** | **bool** | Order allowed (true &#x3D; cacl completed duration, false &#x3D; calc whole duration or exclueded, depends on if the startTime of the order was in the invoice duration) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

