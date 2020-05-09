package contract

func (cs *ContractService) GetCreateOrderBlock() {
	/* TODO: Generate a block to create an order
		1. call dod_settlement_getCreateOrderBlock to creat an order,it will return an internal id
	    2. sign orderBlock and process it to the chain
		3. periodically check whether this order has been signed and confirmed through internal id
		4. if order has been signed and confirmed,call orchestra interface to order at the sonata service,
	       will return an external order id
		5. call dod_settlement_getUpdateOrderInfoBlock to update real orderId to qlc chain
		6. call orchestra interface to periodically check whether the resource of this order has been ready?
		7. if resource is ready,call dod_settlement_getResourceReadyBlock periodically check whether the resource of this order has been ready?
	*/
}

func (cs *ContractService) CheckCreateOrderContractSignStatus(internalId string) bool {
	return true
}

func (cs *ContractService) CheckCreateOrderResourceReady(externalId string) bool {
	return true
}
