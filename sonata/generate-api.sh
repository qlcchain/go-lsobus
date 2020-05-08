#!/usr/bin/env bash

# Client
## Serviceability
### Address
swagger generate client -t ./address -f ./spec/MEF_api_geographicAddressManagement_3.0.0.json
### Site
swagger generate client -t ./site -f ./spec/MEF_api_geographicSiteManagement_3.0.0.json
### POQ
swagger generate client -t ./poq -f ./spec/MEF_api_productOfferingQualificationManagement_3.0.1.json

## Quote
swagger generate client -t ./quote -f ./spec/MEF_api_quoteManagement_2.0.0.json

## Product Order
swagger generate client -t ./order/ -f ./spec/MEF_api_productOrderManagement_3.0.1.json

## Product Inventory
swagger generate client -t ./inventory -f ./spec/MEF_api_productInventoryManagement_3.0.0.json

## Product Specification
swagger generate model -t ./common -f ./spec/MEF_api_productSpec.json

# Server
## POQ Notification
swagger generate server --exclude-main --skip-support --exclude-spec -t ./notify/poq -f ./spec/MEF_api_productOfferingQualificationNotification_3.0.0.json

## Quote Notification
swagger generate server --exclude-main --skip-support --exclude-spec -t ./notify/quote -f ./spec/MEF_api_quoteNotification_1.0.0.json

## Product Order Notification
swagger generate server --exclude-main --skip-support --exclude-spec -t ./notify/order/ -f ./spec/MEF_api_productOrderNotification_3.0.0.json
