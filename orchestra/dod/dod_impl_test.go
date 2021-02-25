package dod

import (
	"encoding/json"
	"testing"

	sw "github.com/qlcchain/go-lsobus/generated/dod"
)

func Test_quoteItem2Quote(t *testing.T) {
	data := `{"apimXUserId":"587055","httpStatus":"OK","id":"7a1407e8f50b4f7c86636fc354400f48","href":"https://test-api.hgc.com.hk/quote/v1/quotes/7a1407e8f50b4f7c86636fc354400f48","projectId":"MyProject-03","description":"","state":"READY","quoteDate":"2021-01-16T12:43:02.006Z","instantSyncQuoting":true,"quoteLevel":"FIRM","requestedQuoteCompletionDate":"2021-01-19T08:19:10.528Z","expectedQuoteCompletionDate":"2021-01-17T12:43:02.006Z","expectedFulfillmentStartDate":"","effectiveQuoteCompletionDate":"2021-01-16T12:43:02.006Z","validFor":{"startDate":"2021-01-16T12:43:02.006Z","endDate":"2021-02-16T12:43:02.006Z"},"note":[{"date":"2021-01-16T12:43:02.120Z","author":"HGC","text":"Interconnection Type = PTP, Interface Type = FE"}],"agreement":[{"id":"1","href":"","name":"Buyer","path":"/"}],"relatedParty":[{"id":"1","role":["buyer"],"name":"Buyer","emailAddress":"buyer@me.com","number":"12123","numberExtension":"1"}],"quoteItem":[{"id":"1","state":"READY","action":"INSTALL","product":{"id":"PRD0017HCVF44ED","productSpecification":{"id":"UNISpec","describing":{"physicalLayer":["1000BASE-LX"],"interconnectionType":"GatewayInterconnect"}},"productRelationship":[{"type":"UNISpec","product":{"id":"QLINK01"}}],"place":[{"id":"32443101","role":"UNI_LOCATION","@type":"FieldedAddress"}]},"quoteItemPrice":[{"priceType":"RECURRING","recurringChargePeriod":"MONTH","name":"Quote Item Price","price":{"preTaxAmount":{"value":0,"unit":"USD"}}},{"priceType":"NON_RECURRING","price":{"preTaxAmount":{"value":0,"unit":"USD"}}}],"requestedQuoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemRelationship":[{"type":"RELIES_ON","id":""}],"note":[{"date":"2021-01-12T12:08:22.936Z","author":"Sample Note Author 1","text":""}],"qualification":[{"id":"8e32872f156b4f25814720d17f2fc0c8","href":"","qualificationItem":"1"}],"relatedParty":[{"id":"1","role":["buyer"],"name":"Buyer","emailAddress":"buyer@buyer.com","number":"213123","numberExtension":"1","@referredType":""}]},{"id":"2","state":"READY","action":"INSTALL","product":{"id":"PRD00FXUZY1GCVE","href":"","productSpecification":{"id":"ELineSpec","describing":{"sVlanId":"123","ENNIIngressBWProfile":{"cir":{"amount":100,"unit":"Mbps"}},"UNIIngressBWProfile":{"cir":{"amount":100,"unit":"Mbps"}}}},"productRelationship":[{"type":"RELIES_ON","product":{"id":"QLINK01"}}],"place":[]},"quoteItemPrice":[{"priceType":"RECURRING","recurringChargePeriod":"MONTH","name":"Quote Item Price for METRONET_METROPORT. Interconnection Type = PTP, Interface Type = FE","price":{"preTaxAmount":{"value":396,"unit":"USD"}}},{"priceType":"NON_RECURRING","price":{"preTaxAmount":{"value":0,"unit":"USD"}}}],"requestedQuoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemTerm":{"name":"","description":"","duration":{"value":12,"unit":"MONTH"}},"quoteItemRelationship":[{"type":"RELIES_ON","id":"1"}],"qualification":[{"id":"8e32872f156b4f25814720d17f2fc0c8","href":"","qualificationItem":"2"}],"relatedParty":[{"id":"","role":["buyer"],"name":"","emailAddress":"","number":"","numberExtension":""}]}]}`
	var quote sw.QuoteRes
	if err := json.Unmarshal([]byte(data), &quote); err != nil {
		t.Fatal(err)
	} else {
		t.Log(quote.ID)
	}
}

func Test_quoteItem2Quote_DCC_Port(t *testing.T) {
	dccPort := `{
    "httpStatus": "ACCEPTED",
    "id": "f887b213-482a-4bed-974f-5a303b221b11",
    "href": "",
    "projectId": "test",
    "description": "",
    "state": "READY",
    "quoteDate": "2021-01-07T12:15:13.201Z",
    "instantSyncQuoting": true,
    "quoteLevel": "FIRM",
    "requestedQuoteCompletionDate": "2021-02-24 01:03:01 AM +01:00",
    "expectedQuoteCompletionDate": "2021-02-25T00:03:01.000Z",
    "expectedFulfillmentStartDate": "",
    "effectiveQuoteCompletionDate": "2021-02-25T00:03:01.000Z",
    "validFor": {
        "startDate": "2021-02-24 01:03:01 AM +01:00",
        "endDate": "2021-03-24T00:03:01.000Z"
    },
    "agreement": [
        {
            "id": "1",
            "href": "",
            "name": "Buyer",
            "path": "/"
        }
    ],
    "relatedParty": [
        {
            "id": "1",
            "role": [
                "buyer"
            ],
            "name": "HGC",
            "emailAddress": "info@hgc.com.hk",
            "number": "111111111",
            "numberExtension": ""
        }
    ],
    "quoteItem": [
        {
            "id": "1",
            "state": "READY",
            "action": "INSTALL",
            "product": {
                "id": 30584,
                "productSpecification": {
                    "id": "UNISpec",
                    "describing": {
                        "physicalLayer": [
                            "1000BASE-LX"
                        ],
                        "interconnectionType": "GatewayInterconnect"
                    }
                },
                "productRelationship": [
                    {
                        "type": "UNISpec",
                        "product": {
                            "id": "QLINK01"
                        }
                    }
                ],
                "place": [
                    {
                        "role": "UNI_LOCATION",
                        "id": "76455"
                    }
                ]
            },
            "quoteItemPrice": [
                {
                    "priceType": "RECURRING",
                    "recurringChargePeriod": "MONTH",
                    "name": "Quote Item Price",
                    "price": {
                        "preTaxAmount": {
                            "value": 100,
                            "unit": "USD"
                        }
                    }
                },
                {
                    "priceType": "NON_RECURRING",
                    "price": {
                        "preTaxAmount": {
                            "value": 0,
                            "unit": "USD"
                        }
                    }
                }
            ],
            "requestedQuoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 12,
                    "unit": "MONTH"
                }
            },
            "quoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 12,
                    "unit": "MONTH"
                }
            },
            "quoteItemRelationship": [
                {
                    "type": "RELIES_ON",
                    "id": ""
                }
            ],
            "qualification": {
                "id": "ff30cebd-086e-4fe2-b12c-d4b5b52d7dd4",
                "href": "",
                "qualificationItem": "1"
            },
            "relatedParty": [
                {
                    "id": "1",
                    "role": [
                        "buyer"
                    ],
                    "name": "HGC",
                    "emailAddress": "info@hgc.com.hk",
                    "number": "111111111",
                    "numberExtension": "",
                    "@referredType": ""
                }
            ],
            "productOffering": {
                "id": "test"
            },
            "note": [
                {
                    "date": "2021-02-24 01:03:01 AM +01:00",
                    "author": "HGC",
                    "text": ""
                }
            ]
        }
    ]
}`
	q := &sw.QuoteRes{}
	if err := json.Unmarshal([]byte(dccPort), q); err != nil {
		t.Fatal(err)
	} else {
		t.Log(q.QuoteItem)
	}
}
func Test_quoteItem2Quote_DCC_Conn(t *testing.T) {
	dccConn := `{
    "httpStatus": "ACCEPTED",
    "id": "c1aa7738-4005-43c4-9722-6e027c870ed5",
    "href": "",
    "projectId": "New Connection Project",
    "description": "",
    "state": "READY",
    "quoteDate": "2021-01-07T12:15:13.201Z",
    "instantSyncQuoting": true,
    "quoteLevel": "FIRM",
    "requestedQuoteCompletionDate": "2021-02-10T08:00:00.000Z",
    "expectedQuoteCompletionDate": "2021-02-11T08:00:00.000Z",
    "expectedFulfillmentStartDate": "",
    "effectiveQuoteCompletionDate": "2021-02-11T08:00:00.000Z",
    "validFor": {
        "startDate": "2021-02-10T08:00:00.000Z",
        "endDate": "2021-03-10T08:00:00.000Z"
    },
    "agreement": [
        {
            "id": "1",
            "href": "",
            "name": "Buyer",
            "path": "/"
        }
    ],
    "relatedParty": [
        {
            "id": "1",
            "role": [
                "buyer"
            ],
            "name": "HGC",
            "emailAddress": "info@hgc.com.hk",
            "number": "44411111444",
            "numberExtension": "1"
        }
    ],
    "quoteItem": [
        {
            "id": "1",
            "state": "READY",
            "action": "INSTALL",
            "product": {
                "id": "10000",
                "href": "",
                "productSpecification": {
                    "id": "ELineSpec",
                    "describing": {
                        "type": "ELineSpec",
                        "sVlanId": "10",
                        "ENNIIngressBWProfile": {
                            "cir": {
                                "amount": 100,
                                "unit": "Mbps"
                            }
                        },
                        "UNIIngressBWProfile": {
                            "cir": {
                                "amount": 100,
                                "unit": "Mbps"
                            }
                        },
                        "ports": {
                            "sdnSrcPort": 30906,
                            "sdnDestPort": 31025,
                            "sdnAVlan": 10,
                            "sdnZVlan": 10
                        }
                    }
                },
                "productRelationship": [
                    {
                        "type": "RELIES_ON",
                        "product": {
                            "id": "any"
                        }
                    }
                ],
                "place": []
            },
            "quoteItemPrice": [
                {
                    "priceType": "RECURRING",
                    "recurringChargePeriod": "MONTH",
                    "name": "Quote Item Price for METRONET_METROPORT. Interconnection Type = PTP, Interface Type = FE",
                    "price": {
                        "preTaxAmount": {
                            "value": 17515,
                            "unit": "USD"
                        }
                    }
                },
                {
                    "priceType": "NON_RECURRING",
                    "price": {
                        "preTaxAmount": {
                            "value": 0,
                            "unit": "USD"
                        }
                    }
                }
            ],
            "requestedQuoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 12,
                    "unit": "MONTH"
                }
            },
            "quoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 12,
                    "unit": "MONTH"
                }
            },
            "quoteItemRelationship": [
                {
                    "type": "RELIES_ON",
                    "id": "1"
                }
            ],
            "qualification": {
                "id": "17a52c94-4831-4e45-bfff-007dc7516787",
                "href": "",
                "qualificationItem": "1"
            },
            "relatedParty": [
                {
                    "id": "",
                    "role": [
                        "buyer"
                    ],
                    "name": "HGC",
                    "emailAddress": "info@hgc.com.hk",
                    "number": "44411111444",
                    "numberExtension": ""
                }
            ],
            "productOffering": {
                "id": "New Connection POQ ID 02"
            }
        }
    ]
}`
	q := &sw.QuoteRes{}
	if err := json.Unmarshal([]byte(dccConn), q); err != nil {
		t.Fatal(err)
	} else {
		t.Log(q.QuoteItem)
	}
}

func Test_quoteItem2Quote_HGC(t *testing.T) {
	hgc := `{
    "httpStatus": "ACCEPTED",
    "id": "76b033afd43b429b80a3afac27bee41e",
    "href": "https://test-api.hgc.com.hk/quote/v1/quotes/76b033afd43b429b80a3afac27bee41e",
    "projectId": "New Project Name",
    "description": "",
    "state": "READY",
    "quoteDate": "2021-02-24T13:04:30.508Z",
    "instantSyncQuoting": true,
    "quoteLevel": "FIRM",
    "requestedQuoteCompletionDate": "2021-02-25T08:00:00.000Z",
    "expectedQuoteCompletionDate": "2021-02-25T13:04:30.508Z",
    "expectedFulfillmentStartDate": "",
    "effectiveQuoteCompletionDate": "2021-02-24T13:04:30.508Z",
    "validFor": {
        "startDate": "2021-02-24T13:04:30.508Z",
        "endDate": "2021-03-24T13:04:30.508Z"
    },
    "note": [
        {
            "date": "2021-02-24T13:04:30.681Z",
            "author": "HGC",
            "text": "Interconnection Type = PTP, Interface Type = FE"
        }
    ],
    "agreement": [
        {
            "id": "1",
            "href": "",
            "name": "Buyer",
            "path": "/"
        }
    ],
    "relatedParty": [
        {
            "id": "1",
            "role": [
                "buyer"
            ],
            "name": "DCConnect",
            "emailAddress": "info@dcconnected.com",
            "number": "88811111888",
            "numberExtension": "1"
        }
    ],
    "quoteItem": [
        {
            "id": "1",
            "state": "READY",
            "action": "INSTALL",
            "product": {
                "id": "PRD0017HCVF44ED",
                "productSpecification": {
                    "id": "UNISpec",
                    "describing": {
                        "physicalLayer": [
                            "1000BASE-LX"
                        ],
                        "interconnectionType": "GatewayInterconnect"
                    }
                },
                "productRelationship": [
                    {
                        "type": "UNISpec",
                        "product": {
                            "id": "QLINK01"
                        }
                    }
                ],
                "place": [
                    {
                        "id": "28402501",
                        "role": "UNI_LOCATION",
                        "@type": "FieldedAddress"
                    }
                ]
            },
            "quoteItemPrice": [
                {
                    "priceType": "RECURRING",
                    "recurringChargePeriod": "MONTH",
                    "name": "Quote Item Price",
                    "price": {
                        "preTaxAmount": {
                            "value": 0,
                            "unit": "USD"
                        }
                    }
                },
                {
                    "priceType": "NON_RECURRING",
                    "price": {
                        "preTaxAmount": {
                            "value": 0,
                            "unit": "USD"
                        }
                    }
                }
            ],
            "requestedQuoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 1,
                    "unit": "MONTH"
                }
            },
            "quoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 1,
                    "unit": "MONTH"
                }
            },
            "quoteItemRelationship": [
                {
                    "type": "RELIES_ON",
                    "id": ""
                }
            ],
            "note": [
                {
                    "date": "2021-02-10T08:00:00.000Z",
                    "author": "Sample Note Author 1",
                    "text": ""
                }
            ],
            "qualification": [
                {
                    "id": "b4dfc1eb498c415fb9f9296cce5c10b9",
                    "href": "",
                    "qualificationItem": "1"
                }
            ],
            "relatedParty": [
                {
                    "id": "1",
                    "role": [
                        "buyer"
                    ],
                    "name": "DCConnect",
                    "emailAddress": "info@dcconnected.com",
                    "number": "88811111888",
                    "numberExtension": "1",
                    "@referredType": ""
                }
            ]
        },
        {
            "id": "2",
            "state": "READY",
            "action": "INSTALL",
            "product": {
                "id": "PRD00FXUZY1GCVE",
                "href": "",
                "productSpecification": {
                    "id": "ELineSpec",
                    "describing": {
                        "sVlanId": "10",
                        "ENNIIngressBWProfile": {
                            "cir": {
                                "amount": 100,
                                "unit": "Mbps"
                            }
                        },
                        "UNIIngressBWProfile": {
                            "cir": {
                                "amount": 100,
                                "unit": "Mbps"
                            }
                        }
                    }
                },
                "productRelationship": [
                    {
                        "type": "RELIES_ON",
                        "product": {
                            "id": "QLINK01"
                        }
                    }
                ],
                "place": []
            },
            "quoteItemPrice": [
                {
                    "priceType": "RECURRING",
                    "recurringChargePeriod": "MONTH",
                    "name": "Quote Item Price for METRONET_METROPORT. Interconnection Type = PTP, Interface Type = FE",
                    "price": {
                        "preTaxAmount": {
                            "value": 396,
                            "unit": "USD"
                        }
                    }
                },
                {
                    "priceType": "NON_RECURRING",
                    "price": {
                        "preTaxAmount": {
                            "value": 0,
                            "unit": "USD"
                        }
                    }
                }
            ],
            "requestedQuoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 1,
                    "unit": "MONTH"
                }
            },
            "quoteItemTerm": {
                "name": "",
                "description": "",
                "duration": {
                    "value": 1,
                    "unit": "MONTH"
                }
            },
            "quoteItemRelationship": [
                {
                    "type": "RELIES_ON",
                    "id": "1"
                }
            ],
            "qualification": [
                {
                    "id": "b4dfc1eb498c415fb9f9296cce5c10b9",
                    "href": "",
                    "qualificationItem": "2"
                }
            ],
            "relatedParty": [
                {
                    "id": "",
                    "role": [
                        "buyer"
                    ],
                    "name": "DCConnect",
                    "emailAddress": "info@dcconnected.com",
                    "number": "88811111888",
                    "numberExtension": ""
                }
            ]
        }
    ]
}`
	q := &sw.QuoteRes{}
	if err := json.Unmarshal([]byte(hgc), q); err != nil {
		t.Fatal(err)
	} else {
		t.Log(q.QuoteItem)
	}
}
