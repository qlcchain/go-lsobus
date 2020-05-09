package contract

func (cs *ContractService) GetTerminateOrderBlock() {
	/* TODO: Genearate a block to terminate an order
		1. call dod_settlement_getTerminateOrderBlock to terminate an order,need order's id generated before
	    2. sign orderBlock and process it to the chain
		3. periodically check whether this order has been signed and confirmed through internal id
		4. if order has been signed and confirmed,call orchestra interface to order at the sonata service
		5. call orchestra interface to periodically check whether the resource of this order has been ready?
		6. if resource is ready,call dod_settlement_getResourceReadyBlock periodically check whether the resource of this order has been ready?
	*/
}

func (cs *ContractService) CheckTerminateOrderContractSignStatus(internalId string) bool {
	return true
}

func (cs *ContractService) CheckTerminateOrderResourceReady(externalId string) bool {
	return true
}
