
## API spec error
```
➜ docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli validate -i /local/docs/openapi.json
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
docker run --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3:latest generate -i /local/docs/openapi.json -l go -o /local/generated/dod
```
## DEBUG

### swagger doc

- missing

```js
router.post('/v1/rpc/qlc/dod-settlement/update-order-info-block', dodSettlement.updateOrderInfoBlock);
router.post('/v1/rpc/qlc/dod-settlement/update-product-info-block', dodSettlement.updateProductInfoBlock);
router.post('/v1/rpc/qlc/dod-settlement/update-order-info-reward-block', dodSettlement.updateOrderInfoRewardBlock);
router.post('/v1/rpc/qlc/dod-settlement/change-order-block', dodSettlement.changeOrderBlock);
router.post('/v1/rpc/qlc/dod-settlement/terminate-order-block', dodSettlement.terminateOrderBlock);
router.post('/v1/rpc/qlc/dod-settlement/create-order-reward-block', dodSettlement.createOrderRewardBlock);
router.post('/v1/rpc/qlc/dod-settlement/change-order-reward-block', dodSettlement.changeOrderRewardBlock);
router.post('/v1/rpc/qlc/dod-settlement/terminate-order-reward-block', dodSettlement.terminateOrderRewardBlock);
```

- miss match

  `v1/dlt/order/seller/pending-request` etc.

### QLC Node

chain node have to enable PoV, otherwise the DoD contract won't work

### just for test

- add an endpoint to quickly create an order from buyer side
