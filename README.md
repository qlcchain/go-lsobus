# LSOBUS
Business Application for MEF LSO Framework.

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

Orchestra is responsible for managing the responses and update status to chain.

### Contract
Contract is responsible for make a series interactions to blockchain smart contract.

### RPC Server
RPC Server is responsible for providing APIs to Front Web UI.

## Process Flows

### UNI & E-Line Order
#### External
Uploading order to chain:

User -> Front Web UI -> LSOBUS -> QLC Chain.

Sending order to partner:

User -> Front Web UI -> LSOBUS -> Sonata Server.

#### Internal
Uploading order to chain:

RPC -> Contract -> QLC Chain.

Sending order to partner:

RPC -> Orchestra -> Sonata Client -> Sonata Server.
