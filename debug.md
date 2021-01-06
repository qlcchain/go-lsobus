
## API spec erro
```
➜ docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate -i /local/docs/openapi.yaml
Validating spec (/local/docs/openapi.yaml)
Errors:
	- attribute components.schemas.200SuccessResponse.items is missing
	- attribute components.schemas.500ErrorResponse.items is missing
	- attribute components.schemas.400ErrorResponse.items is missing
Warnings:
	- Unused model: orderItemRelationshipModel

[error] Spec has 3 errors.

```

## generate code 
```
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/docs/openapi.yaml -g go -o /local/out/go --skip-validate-spec
```

## Missing QLCChain API

- https://docs.qlcchain.online/api/rpc/ledger.html#ledger-process
- https://docs.qlcchain.online/api/rpc/ledger.html#ledger-blockconfirmedstatus
- pov_getPovStatus
request:
```json
{
  "jsonrpc": "2.0",
  "id":3,
  "method":"pov_getPovStatus"
}
```
response:
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "povEnabled": true,
    "syncState": 2,
    "syncStateStr": "SyncDone"
  }
}
```
