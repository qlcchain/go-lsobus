package contract

func (cs *ContractService) GetChangeOrderBlock() {
	/* TODO: Generate a block to change order's service parameters
		1. call dod_settlement_getChangeOrderBlock  to change an order,need order's id generated before it will return an internal id
	    2. sign orderBlock and process it to the chain
		3. periodically check whether this order has been signed and confirmed through internal id
		4. if order has been signed and confirmed,call orchestra interface to order at the sonata service,
	       will return an real order id
		5. call dod_settlement_getUpdateOrderInfoBlock to update real orderId to qlc chain
		6. call orchestra interface to periodically check whether the resource of this order has been ready?
		7. if resource is ready,call dod_settlement_getResourceReadyBlock periodically check whether the resource of this order has been ready?
	*/
}

func (cs *ContractService) CheckChangeOrderContractSignStatus(internalId string) bool {
	return true
}

func (cs *ContractService) CheckChangeOrderResourceReady(externalId string) bool {
	return true
}
