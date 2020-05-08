# Virtual-LSOBUS
Virtual Business Application for MEF LSO Framework.

## Modules
### Sonata
MEF Sonata API client and models.

### Orchestra
Orchestra is responsible for making a series of MEF calls to the partner:
- to determine that the address provided is valid and provide an interface to handle invalid addresses
- to get a site at the customer address, or allow selection of a new site
- to determine if the requested service is feasible
- to get a quote for the requested service
- to create ProductOrder request
- to check for ProductOrder updates
ODO is responsible for managing the responses and providing email notifications to appropriate OperatorA users.

### Contract
Contract is responsible for make a series interactions to blockchain smart contract.

### RPC Server
RPC Server is responsible for providing APIs to Front Web UI.

## Process Flows

### E-Line Order
#### External
User -> Front Web UI -> QLC Chain -> Virtual-LSOBUS.
#### Internal
QLC Chain -> Contract -> Orchestra -> Sonata -> Partner.
