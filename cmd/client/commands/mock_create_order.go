package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/go-resty/resty/v2"
	qlcchain "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/random"
	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/qlcchain/qlc-go-sdk/pkg/util"

	u "github.com/qlcchain/go-lsobus/cmd/util"
)

const (
	POQSample = `{
	"projectId": "MyProject-03",
	"provideAlternative": false,
	"requestedResponseDate": "2021-01-12T08:19:10.528Z",
	"instantSyncQualification": true,
	"relatedParty": [
		{
			"id": "1",
			"role": [
				"Buyer"
			],
			"name": "Mr Buyer XXX",
			"emailAddress": "mrbuyer@foobar.com",
			"number": "85211112222",
			"numberExtension": ""
		}
	],
	"productOfferingQualificationItem": [
		{
			"id": "1",
			"product": {
				"id": "PRD0017HCVF44ED",
				"productSpecification": {
					"id": "UNISpec",
					"describing": {
						"physicalLayer": [
							"1000BASE-LX"
						],
						"@type": "UNISpec"
					}
				},
				"place": [
					{
						"role": "UNI_LOCATION",
						"id": "32431201",
						"geographicSite": {
							"floor": "1",
							"room": "1506"
						}
					}
				]
			},
			"productOffering": {
				"id": "UNI001-POQ-03"
			}
		},
		{
			"id": "2",
			"product": {
				"id": "PRD00FXUZY1GCVE",
				"productRelationship": [
					{
						"type": "RELIES_ON",
						"product": {
							"id": "QLINK01"
						}
					}
				],
				"productSpecification": {
					"id": "ELineSpec",
					"describing": {
						"@type": "ELineSpec",
						"sVlanId": "123",
						"ENNIIngressBWProfile": {
							"cir": {
								"amount": 10,
								"unit": "Mbps"
							}
						},
						"UNIIngressBWProfile": {
							"cir": {
								"amount": 10,
								"unit": "Mbps"
							}
						},
						"svlanId": "123"
					}
				}
			},
			"productOffering": {
				"id": "ELINE001-POQ-03"
			},
			"productOfferingQualificationItemRelationship": [
				{
					"type": "RELIES_ON",
					"id": "1"
				}
			]
		}
	]
}`

	QuoteSample = `{
	"projectId": "MyProject-03",
	"quoteLevel": "FIRM",
	"instantSyncQuoting": true,
	"description": "",
	"requestedQuoteCompletionDate": "2021-02-19T08:19:10.528Z",
	"expectedFulfillmentStartDate": "",
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
			"name": "Buyer",
			"emailAddress": "buyer@me.com",
			"number": "12123",
			"numberExtension": "1"
		}
	],
	"quoteItem": [
		{
			"id": "1",
			"action": "INSTALL",
			"productOffering": {
				"id": "UNI001-POQ-03"
			},
			"requestedQuoteItemTerm": {
				"name": "",
				"description": "",
				"duration": {
					"value": 12,
					"unit": "MONTH"
				}
			},
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
						"role": "UNI_LOCATION",
						"id": "32443101"
					}
				]
			},
			"qualification": {
				"id": "8e32872f156b4f25814720d17f2fc0c8",
				"href": "",
				"qualificationItem": "1"
			},
			"quoteItemRelationship": [
				{
					"type": "RELIES_ON",
					"id": ""
				}
			],
			"note": [
				{
					"date": "2021-01-12T12:08:22.936Z",
					"author": "Sample Note Author 1",
					"text": ""
				}
			],
			"relatedParty": [
				{
					"id": "1",
					"role": [
						"buyer"
					],
					"name": "Buyer",
					"emailAddress": "buyer@buyer.com",
					"number": "213123",
					"numberExtension": "1",
					"@referredType": ""
				}
			]
		},
		{
			"id": "2",
			"action": "INSTALL",
			"productOffering": {
				"id": "ELINE001-POQ-03"
			},
			"requestedQuoteItemTerm": {
				"name": "",
				"description": "",
				"duration": {
					"value": 12,
					"unit": "MONTH"
				}
			},
			"product": {
				"id": "PRD00FXUZY1GCVE",
				"href": "",
				"productSpecification": {
					"id": "ELineSpec",
					"describing": {
						"type": "ELineSpec",
						"sVlanId": "123",
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
			"qualification": {
				"id": "8e32872f156b4f25814720d17f2fc0c8",
				"href": "",
				"qualificationItem": "2"
			},
			"quoteItemRelationship": [
				{
					"type": "RELIES_ON",
					"id": "1"
				}
			],
			"relatedParty": [
				{
					"id": "",
					"role": [
						"buyer"
					],
					"name": "",
					"emailAddress": "",
					"number": "",
					"numberExtension": ""
				}
			]
		}
	]
}`
	OrderSample = `{
  "externalId": "d2e019e027f19e250daa9a5d4d68d40cf160d40928028b6771053f7b54e6450d",
  "buyerRequestDate": "2021-01-07T14:00:00.000Z",
  "requestedCompletionDate": "2021-01-07T15:00:00.000Z",
  "orderActivity": "INSTALL",
  "desiredResponse": "CONFIRMATION_AND_ENGINEERING_DESIGN",
  "orderVersion": "1",
  "projectId": "MyProject-02",
  "relatedParty": [
    {
      "id": "1",
      "role": [
        "Buyer"
      ],
      "name": "Buyer",
      "emailAddress": "buyer@me.com",
      "number": "12123",
      "numberExtension": "1"
    },
    {
      "id": "2",
      "role": [
        "Notification Contact"
      ],
      "name": "Buyer",
      "emailAddress": "buyer@me.com",
      "number": "12123",
      "numberExtension": "1"
    }
  ],
  "note": [],
  "orderItem": [
    {
      "id": "1",
      "action": "INSTALL",
      "productOffering": {
        "id": "294d59af1508440b8c2989856bd4eedd"
      },
      "relatedParty": [
        {
          "id": "1",
          "role": [
            "UNISpecContact"
          ],
          "name": "test contact",
          "emailAddress": "who@me.com",
          "number": "661763285401"
        }
      ],
      "product": {
        "productSpecification": {
          "id": "UNISpec"
        },
        "place": [
          {
            "id": "32443101",
            "role": "UNI_LOCATION"
          }
        ],
        "productRelationship": []
      },
      "orderItemRelationship": [],
      "qualification": {
        "id": "294d59af1508440b8c2989856bd4eedd",
        "qualificationItem": "1"
      },
      "quote": {
        "id": "81e08d2909f94d0fb60582c362384bca",
        "quoteItem": "1"
      }
    },
    {
      "id": "2",
      "action": "INSTALL",
      "productOffering": {
        "id": "294d59af1508440b8c2989856bd4eedd"
      },
      "relatedParty": [
        {
          "id": "1",
          "role": [
            "ELineSpecContact"
          ],
          "name": "Buyer",
          "emailAddress": "buyer@me.com",
          "number": "12123",
          "numberExtension": "1"
        }
      ],
      "product": {
        "productSpecification": {
          "id": "ELineSpec"
        },
        "place": [],
        "productRelationship": [
          {
            "type": "RELIES_ON",
            "product": {
              "id": "QLINK01",
              "productSpecification": {
                "id": "ENNI"
              }
            }
          }
        ]
      },
      "orderItemRelationship": [
        {
          "type": "RELIES_ON",
          "id": "1"
        }
      ],
      "qualification": {
        "id": "294d59af1508440b8c2989856bd4eedd",
        "qualificationItem": "2"
      },
      "quote": {
        "id": "81e08d2909f94d0fb60582c362384bca",
        "quoteItem": "2"
      }
    }
  ],
  "createDate": "2021-01-07T14:00:00.000Z"
}`

	SmartContractSample = `{
	"buyer": {
		"address": "qlc_17xd74qtbz5cu7teg9nx4m418qoc3pwq4a6kxixassed3af4ux3houii3ur7",
		"seed": "b41c5db45a2b9d2bce0a19445f74b768af2a546d8d9f5b59d6680067a1ed00be",
		"name": "Buyer"
	}, 
  "orderItem": [
    {
      "id": "PRODUCT-6eeddcba-e423-43e9-a0d4-71fb5ce427df",
      "productOffering": {
        "id": "294d59af1508440b8c2989856bd4eedd"
      },
      "quote": {
        "id": "7a1407e8f50b4f7c86636fc354400f48",
        "quoteItem": "1"
      },
			"detail": {
				"buyerProductId": "PRD0017HCVF44ED"
			}
    },
		{
      "id": "PRODUCT-85fbb62d-f24c-42e7-8632-56b7fffaf2cd",
      "productOffering": {
        "id": "294d59af1508440b8c2989856bd4eedd"
      },
      "quote": {
        "id": "7a1407e8f50b4f7c86636fc354400f48",
        "quoteItem": "2"
      },
			"detail": {
				"buyerProductId": "PRD00FXUZY1GCVE",
				"connectionName" : "MyProject-03",
				"paymentType" : "invoice",
				"billingType" : "DOD",
				"bandwidth" : "100",
				"unit": "Mbps",
				"billingUnit" : null,
				"dateStartUnix": "1610467115",
				"dateEndUnix": "1611467115",
				"serviceClass" : "bronze" 
			}
    }
  ]
}`
)

func addMockOrderCmdByShell(parentCmd *ishell.Cmd) {
	buyerSeed := u.Flag{
		Name:  "seed",
		Must:  true,
		Usage: "buyer's seed hex string",
		Value: "",
	}

	vendor := u.Flag{
		Name:  "vendor",
		Must:  false,
		Usage: "sonata vendor",
		Value: "HGC",
	}

	apiToken := u.Flag{
		Name:  "apiToken",
		Must:  false,
		Usage: "DoD backend API token",
		Value: "f93713d0-beac-4661-879c-3d479aaa7333",
	}

	clientToken := u.Flag{
		Name:  "clientToken",
		Must:  false,
		Usage: "DoD backend API token",
		Value: "9bdb61d9-8ecb-421f-9495-c79c066dd613",
	}

	args := []u.Flag{buyerSeed, vendor}
	cmd := &ishell.Cmd{
		Name:                "mock",
		Help:                "mock an order from buyer side",
		CompleterWithPrefix: u.OptsCompleter(args),
		Func: func(c *ishell.Context) {
			if u.HelpText(c, args) {
				return
			}
			err := u.CheckArgs(c, args)
			if err != nil {
				u.Warn(err)
				return
			}

			buyerSeedP := u.StringVar(c.Args, buyerSeed)
			vendorP := u.StringVar(c.Args, vendor)
			apiTokenP := u.StringVar(c.Args, apiToken)
			clientTokenP := u.StringVar(c.Args, clientToken)
			u.Info(fmt.Sprintf("buyer:%s, vendor: %s, apiToken: %s,clientToken: %s", buyerSeedP, vendorP,
				apiTokenP, clientTokenP))

			if err := mockOrderForBuyer(buyerSeedP, vendorP, apiTokenP, clientTokenP); err != nil {
				u.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func mockOrderForBuyer(seed, vendor, apiToken, clientToken string) error {
	//var account *pkg.Account
	//if bytes, err := hex.DecodeString(seed); err != nil {
	//	return err
	//} else {
	//	if s, err := pkg.BytesToSeed(bytes); err != nil {
	//		return err
	//	} else {
	//		var err error
	//		if account, err = s.Account(0); err != nil {
	//			return err
	//		}
	//	}
	//}

	client := resty.New().SetHeader("CLIENT-KEY", clientToken).SetHeader("API-KEY", apiToken).
		SetHeader("Content-Type", "application/json").EnableTrace().SetDebug(true)

	switch strings.ToUpper(vendor) {
	case "HGC":
		return mockHGCOrderForBuyer(client, seed)
	default:
		return fmt.Errorf("unsupport sonata vendor %s", vendor)
	}
}

type Inventory struct {
	ID                   string    `json:"id"`
	Status               string    `json:"status"`
	StartDate            time.Time `json:"startDate"`
	ProductSpecification struct {
		ID string `json:"id"`
	} `json:"productSpecification"`
}

type POQRequest struct {
	ProjectID                string    `json:"projectId"`
	ProvideAlternative       bool      `json:"provideAlternative"`
	RequestedResponseDate    time.Time `json:"requestedResponseDate"`
	InstantSyncQualification bool      `json:"instantSyncQualification"`
	RelatedParty             []struct {
		ID              string   `json:"id"`
		Role            []string `json:"role"`
		Name            string   `json:"name"`
		EmailAddress    string   `json:"emailAddress"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
	} `json:"relatedParty"`
	ProductOfferingQualificationItem []struct {
		ID      string `json:"id"`
		Product struct {
			ID                   string `json:"id"`
			ProductSpecification struct {
				ID         string `json:"id"`
				Describing struct {
					PhysicalLayer []string `json:"physicalLayer"`
					Type          string   `json:"@type"`
				} `json:"describing"`
			} `json:"productSpecification"`
			Place []struct {
				Role           string `json:"role"`
				ID             string `json:"id"`
				GeographicSite struct {
					Floor string `json:"floor"`
					Room  string `json:"room"`
				} `json:"geographicSite"`
			} `json:"place"`
		} `json:"product,omitempty"`
		ProductOffering struct {
			ID string `json:"id"`
		} `json:"productOffering"`
		ProductOfferingQualificationItemRelationship []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"productOfferingQualificationItemRelationship,omitempty"`
	} `json:"productOfferingQualificationItem"`
}

type POQResponse struct {
	ID                                   string    `json:"id"`
	InstantSyncQualification             bool      `json:"instantSyncQualification"`
	State                                string    `json:"state"`
	ProjectID                            string    `json:"projectId"`
	ProvideAlternative                   bool      `json:"provideAlternative"`
	RequestedResponseDate                time.Time `json:"requestedResponseDate"`
	ExpectedResponseDate                 string    `json:"expectedResponseDate"`
	EffectiveQualificationCompletionDate string    `json:"effectiveQualificationCompletionDate"`
	RelatedParty                         []struct {
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		Role            []string `json:"role"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
		EmailAddress    string   `json:"emailAddress"`
	} `json:"relatedParty"`
	ProductOfferingQualificationItem []struct {
		ID                       string `json:"id"`
		State                    string `json:"state"`
		ServiceabilityConfidence string `json:"serviceabilityConfidence"`
		ServiceConfidenceReason  string `json:"serviceConfidenceReason"`
		InstallationInterval     struct {
			Amount   int    `json:"amount"`
			TimeUnit string `json:"timeUnit"`
		} `json:"installationInterval"`
		GuaranteedUntilDate string `json:"guaranteedUntilDate"`
		Product             struct {
			ID                   string `json:"id"`
			ProductSpecification struct {
				ID         string `json:"id"`
				Describing struct {
					Type          string   `json:"@type"`
					PhysicalLayer []string `json:"physicalLayer"`
				} `json:"describing"`
			} `json:"productSpecification"`
			Place []struct {
				Role string `json:"role"`
				ID   string `json:"id"`
				Type string `json:"@type"`
			} `json:"place"`
		} `json:"product,omitempty"`
	} `json:"productOfferingQualificationItem"`
}

type QuoteRequest struct {
	ProjectID                    string    `json:"projectId"`
	QuoteLevel                   string    `json:"quoteLevel"`
	InstantSyncQuoting           bool      `json:"instantSyncQuoting"`
	Description                  string    `json:"description"`
	RequestedQuoteCompletionDate time.Time `json:"requestedQuoteCompletionDate"`
	ExpectedFulfillmentStartDate string    `json:"expectedFulfillmentStartDate"`
	Agreement                    []struct {
		ID   string `json:"id"`
		Href string `json:"href"`
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"agreement"`
	RelatedParty []struct {
		ID              string   `json:"id"`
		Role            []string `json:"role"`
		Name            string   `json:"name"`
		EmailAddress    string   `json:"emailAddress"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
	} `json:"relatedParty"`
	QuoteItem []struct {
		ID              string `json:"id"`
		Action          string `json:"action"`
		ProductOffering struct {
			ID string `json:"id"`
		} `json:"productOffering"`
		RequestedQuoteItemTerm struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    struct {
				Value int    `json:"value"`
				Unit  string `json:"unit"`
			} `json:"duration"`
		} `json:"requestedQuoteItemTerm"`
		Product struct {
			ID                   string `json:"id"`
			Href                 string `json:"href,omitempty"`
			ProductSpecification struct {
				ID         string `json:"id"`
				Describing struct {
					PhysicalLayer       []string `json:"physicalLayer"`
					InterconnectionType string   `json:"interconnectionType"`
				} `json:"describing"`
			} `json:"productSpecification"`
			ProductRelationship []struct {
				Type    string `json:"type"`
				Product struct {
					ID string `json:"id"`
				} `json:"product"`
			} `json:"productRelationship"`
			Place []interface{} `json:"place"`
		} `json:"product,omitempty"`
		Qualification struct {
			ID                string `json:"id"`
			Href              string `json:"href"`
			QualificationItem string `json:"qualificationItem"`
		} `json:"qualification"`
		QuoteItemRelationship []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"quoteItemRelationship"`
		Note []struct {
			Date   time.Time `json:"date"`
			Author string    `json:"author"`
			Text   string    `json:"text"`
		} `json:"note,omitempty"`
		RelatedParty []struct {
			ID              string   `json:"id"`
			Role            []string `json:"role"`
			Name            string   `json:"name"`
			EmailAddress    string   `json:"emailAddress"`
			Number          string   `json:"number"`
			NumberExtension string   `json:"numberExtension"`
			ReferredType    string   `json:"@referredType"`
		} `json:"relatedParty"`
	} `json:"quoteItem"`
}

type QuoteResponse struct {
	HTTPStatus                   string    `json:"httpStatus"`
	ID                           string    `json:"id"`
	Href                         string    `json:"href"`
	ProjectID                    string    `json:"projectId"`
	Description                  string    `json:"description"`
	State                        string    `json:"state"`
	QuoteDate                    time.Time `json:"quoteDate"`
	InstantSyncQuoting           bool      `json:"instantSyncQuoting"`
	QuoteLevel                   string    `json:"quoteLevel"`
	RequestedQuoteCompletionDate time.Time `json:"requestedQuoteCompletionDate"`
	ExpectedQuoteCompletionDate  time.Time `json:"expectedQuoteCompletionDate"`
	ExpectedFulfillmentStartDate string    `json:"expectedFulfillmentStartDate"`
	EffectiveQuoteCompletionDate time.Time `json:"effectiveQuoteCompletionDate"`
	ValidFor                     struct {
		StartDate time.Time `json:"startDate"`
		EndDate   time.Time `json:"endDate"`
	} `json:"validFor"`
	Note []struct {
		Date   time.Time `json:"date"`
		Author string    `json:"author"`
		Text   string    `json:"text"`
	} `json:"note"`
	Agreement []struct {
		ID   string `json:"id"`
		Href string `json:"href"`
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"agreement"`
	RelatedParty []struct {
		ID              string   `json:"id"`
		Role            []string `json:"role"`
		Name            string   `json:"name"`
		EmailAddress    string   `json:"emailAddress"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
	} `json:"relatedParty"`
	QuoteItem []struct {
		ID      string `json:"id"`
		State   string `json:"state"`
		Action  string `json:"action"`
		Product struct {
			ID                   string `json:"id"`
			ProductSpecification struct {
				ID         string `json:"id"`
				Describing struct {
					PhysicalLayer       []string `json:"physicalLayer"`
					InterconnectionType string   `json:"interconnectionType"`
				} `json:"describing"`
			} `json:"productSpecification"`
			ProductRelationship []struct {
				Type    string `json:"type"`
				Product struct {
					ID string `json:"id"`
				} `json:"product"`
			} `json:"productRelationship"`
			Place []struct {
				ID   string `json:"id"`
				Role string `json:"role"`
				Type string `json:"@type"`
			} `json:"place"`
		} `json:"product,omitempty"`
		QuoteItemPrice []struct {
			PriceType             string `json:"priceType"`
			RecurringChargePeriod string `json:"recurringChargePeriod,omitempty"`
			Name                  string `json:"name,omitempty"`
			Price                 struct {
				PreTaxAmount struct {
					Value int    `json:"value"`
					Unit  string `json:"unit"`
				} `json:"preTaxAmount"`
			} `json:"price"`
		} `json:"quoteItemPrice"`
		RequestedQuoteItemTerm struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    struct {
				Value int    `json:"value"`
				Unit  string `json:"unit"`
			} `json:"duration"`
		} `json:"requestedQuoteItemTerm"`
		QuoteItemTerm struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    struct {
				Value int    `json:"value"`
				Unit  string `json:"unit"`
			} `json:"duration"`
		} `json:"quoteItemTerm"`
		QuoteItemRelationship []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"quoteItemRelationship"`
		Note []struct {
			Date   time.Time `json:"date"`
			Author string    `json:"author"`
			Text   string    `json:"text"`
		} `json:"note,omitempty"`
		Qualification []struct {
			ID                string `json:"id"`
			Href              string `json:"href"`
			QualificationItem string `json:"qualificationItem"`
		} `json:"qualification"`
		RelatedParty []struct {
			ID              string   `json:"id"`
			Role            []string `json:"role"`
			Name            string   `json:"name"`
			EmailAddress    string   `json:"emailAddress"`
			Number          string   `json:"number"`
			NumberExtension string   `json:"numberExtension"`
			ReferredType    string   `json:"@referredType"`
		} `json:"relatedParty"`
	} `json:"quoteItem"`
}

type OrderRequest struct {
	ExternalID              string    `json:"externalId"`
	BuyerRequestDate        time.Time `json:"buyerRequestDate"`
	RequestedCompletionDate time.Time `json:"requestedCompletionDate"`
	OrderActivity           string    `json:"orderActivity"`
	DesiredResponse         string    `json:"desiredResponse"`
	OrderVersion            string    `json:"orderVersion"`
	ProjectID               string    `json:"projectId"`
	RelatedParty            []struct {
		ID              string   `json:"id"`
		Role            []string `json:"role"`
		Name            string   `json:"name"`
		EmailAddress    string   `json:"emailAddress"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
	} `json:"relatedParty"`
	Note      []interface{} `json:"note"`
	OrderItem []struct {
		ID              string `json:"id"`
		Action          string `json:"action"`
		ProductOffering struct {
			ID string `json:"id"`
		} `json:"productOffering"`
		RelatedParty []struct {
			ID           string   `json:"id"`
			Role         []string `json:"role"`
			Name         string   `json:"name"`
			EmailAddress string   `json:"emailAddress"`
			Number       string   `json:"number"`
		} `json:"relatedParty"`
		Product struct {
			ProductSpecification struct {
				ID string `json:"id"`
			} `json:"productSpecification"`
			Place []struct {
				ID   string `json:"id"`
				Role string `json:"role"`
			} `json:"place"`
			ProductRelationship []interface{} `json:"productRelationship"`
		} `json:"product"`
		OrderItemRelationship []interface{} `json:"orderItemRelationship"`
		Qualification         struct {
			ID                string `json:"id"`
			QualificationItem string `json:"qualificationItem"`
		} `json:"qualification"`
		Quote struct {
			ID        string `json:"id"`
			QuoteItem string `json:"quoteItem"`
		} `json:"quote"`
	} `json:"orderItem"`
	CreateDate time.Time `json:"createDate"`
}

type OrderResponse struct {
	ApimXUserID             string    `json:"apimXUserId"`
	ID                      string    `json:"id"`
	ExternalID              string    `json:"externalId"`
	State                   string    `json:"state"`
	BuyerRequestDate        time.Time `json:"buyerRequestDate"`
	RequestedCompletionDate time.Time `json:"requestedCompletionDate"`
	OrderActivity           string    `json:"orderActivity"`
	DesiredResponse         string    `json:"desiredResponse"`
	OrderVersion            string    `json:"orderVersion"`
	ProjectID               string    `json:"projectId"`
	RelatedParty            []struct {
		ID              string   `json:"id"`
		Role            []string `json:"role"`
		Name            string   `json:"name"`
		EmailAddress    string   `json:"emailAddress"`
		Number          string   `json:"number"`
		NumberExtension string   `json:"numberExtension"`
	} `json:"relatedParty"`
	Note      []interface{} `json:"note"`
	OrderItem []struct {
		ID              string `json:"id"`
		Action          string `json:"action"`
		ProductOffering struct {
			ID string `json:"id"`
		} `json:"productOffering"`
		RelatedParty []struct {
			ID           string   `json:"id"`
			Role         []string `json:"role"`
			Name         string   `json:"name"`
			EmailAddress string   `json:"emailAddress"`
			Number       string   `json:"number"`
		} `json:"relatedParty"`
		Product struct {
			ProductSpecification struct {
				ID string `json:"id"`
			} `json:"productSpecification"`
			Place []struct {
				ID           string `json:"id"`
				Role         string `json:"role"`
				AddressLine1 string `json:"addressLine1"`
			} `json:"place"`
			ProductRelationship []interface{} `json:"productRelationship"`
		} `json:"product"`
		OrderItemRelationship []interface{} `json:"orderItemRelationship"`
		Qualification         struct {
			ID                string `json:"id"`
			QualificationItem string `json:"qualificationItem"`
		} `json:"qualification"`
		Quote struct {
			ID        string `json:"id"`
			QuoteItem string `json:"quoteItem"`
		} `json:"quote"`
	} `json:"orderItem"`
	CreateDate string `json:"createDate"`
}

type User struct {
	Address string `json:"address"`
	Seed    string `json:"seed"`
	Name    string `json:"name"`
}

type SmartContractOrderItem struct {
	ID              string `json:"id"`
	ProductOffering struct {
		ID string `json:"id"`
	} `json:"productOffering"`
	Quote struct {
		ID        string `json:"id"`
		QuoteItem string `json:"quoteItem"`
	} `json:"quote"`
	Detail struct {
		BuyerProductID string      `json:"buyerProductId"`
		ConnectionName string      `json:"connectionName"`
		PaymentType    string      `json:"paymentType"`
		BillingType    string      `json:"billingType"`
		Bandwidth      string      `json:"bandwidth"`
		Unit           string      `json:"unit"`
		BillingUnit    interface{} `json:"billingUnit"`
		DateStartUnix  string      `json:"dateStartUnix"`
		DateEndUnix    string      `json:"dateEndUnix"`
		ServiceClass   string      `json:"serviceClass"`
	} `json:"detail,omitempty"`
}

type SmartContractRequest struct {
	Buyer     User                      `json:"buyer"`
	OrderItem []*SmartContractOrderItem `json:"orderItem"`
}

type SmartContractResponse struct {
	TxID string `json:"txId"`
}

type OrderInfo struct {
	Jsonrpc string                       `json:"jsonrpc"`
	ID      string                       `json:"id"`
	Result  *qlcchain.DoDSettleOrderInfo `json:"result"`
}

type OrderInfoItem struct {
	ItemID      string `json:"itemId"`
	OrderItemID string `json:"orderItemId"`
}

type UpdateOrderRequest struct {
	Buyer       string           `json:"buyer"`
	InternalID  string           `json:"internalId"`
	OrderID     string           `json:"orderId"`
	OrderItemID []*OrderInfoItem `json:"orderItemId"`
	Status      string           `json:"status"`
	FailReason  string           `json:"failReason"`
}

func mockHGCOrderForBuyer(client *resty.Client, seed string) error {
	u.Info("STEP1, query inventory")
	var inventory []*Inventory
	resp, err := client.R().SetResult(&inventory).Get(fmt.Sprintf("%s/v1/orders/product-inventory", endpointP))
	if err != nil {
		return err
	}
	//u.Info(u.ToIndentString(inventory))

	u.Info("STEP2, create POQ")

	poqReq := mockPOQRequest(inventory)
	resp, err = client.R().SetBody(poqReq).SetResult(&POQResponse{}).Post(fmt.Sprintf("%s/v1/poq", endpointP))
	if err != nil {
		return err
	}
	poq := resp.Result().(*POQResponse)
	//u.Info(u.ToIndentString(poq))

	u.Info("STEP3, create quote")
	quoteReq := strings.ReplaceAll(QuoteSample, "MyProject-03", poq.ProjectID)
	quoteReq = strings.ReplaceAll(quoteReq, "8e32872f156b4f25814720d17f2fc0c8", poq.ID)
	quoteReq = strings.ReplaceAll(quoteReq, "2021-02-19T08:19:10.528Z", time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05.000Z"))
	//u.Info(quoteReq)
	resp, err = client.R().SetBody(quoteReq).SetResult(&QuoteResponse{}).Post(fmt.Sprintf("%s/v1/quotes", endpointP))
	if err != nil {
		return err
	}
	quote := resp.Result().(*QuoteResponse)

	u.Info("STEP4, create smart contract to save order info")
	bytes, _ := hex.DecodeString(seed)
	s, _ := pkg.BytesToSeed(bytes)
	account, _ := s.Account(0)
	buyer := account.Address().String()
	smReq := mockSmartContract(quote, buyer, seed)
	resp, err = client.R().SetBody(smReq).SetResult(&SmartContractResponse{}).Post(fmt.Sprintf("%s/v1/orders/smart-contract", endpointP))
	if err != nil {
		return err
	}
	tx := resp.Result().(*SmartContractResponse)
	u.Info("smart contract tx: ", tx.TxID)

	u.Info("STEP5, place an order")
	orderReq := mockOrderRequest(quote, tx.TxID)
	resp, err = client.R().SetBody(orderReq).SetResult(&OrderResponse{}).Post(fmt.Sprintf("%s/v1/orders", endpointP))
	if err != nil {
		return err
	}
	order := resp.Result().(*OrderResponse)
	u.Info(fmt.Sprintf("order id: %s, externalid: %s", order.ID, order.ExternalID))

	u.Info("STEP6, waiting the seller to sign the contract")
	oi := &OrderInfo{}
	client.DisableTrace()
	client.Debug = false
	for {
		resp, err = client.R().SetBody(fmt.Sprintf(`{"internalId": "%s"}`, order.ExternalID)).SetResult(&OrderInfo{}).
			Post(fmt.Sprintf("%s/v1/dlt/order/info/by-internal-id", endpointP))

		if err != nil {
			return err
		}
		oi = resp.Result().(*OrderInfo)
		if oi.Result.ContractState == qlcchain.DoDSettleContractStateConfirmed {
			u.Info(fmt.Sprintf("order: %s status is %s", order.ExternalID, qlcchain.DoDSettleContractStateConfirmed))
			break
		}
		time.Sleep(time.Second)
	}

	client.EnableTrace()
	client.Debug = true

	param := &UpdateOrderRequest{
		Buyer:      buyer,
		InternalID: order.ExternalID,
		OrderID:    order.ID,
		Status:     "success",
		FailReason: "",
	}
	for _, connection := range oi.Result.Connections {
		param.OrderItemID = append(param.OrderItemID, &OrderInfoItem{
			ItemID: connection.ItemId,
			//FIXME: remove the fake data
			OrderItemID: order.OrderItem[1].ID,
		})
	}
	u.Info("STEP7, update order info for buyer")
	resp, err = client.R().SetBody(param).SetResult(&SmartContractResponse{}).
		Post(fmt.Sprintf("%s/v1/dlt/order/buyer/update-order-info-block", endpointP))

	if err != nil {
		return err
	}
	tx = resp.Result().(*SmartContractResponse)
	u.Info("update order info tx id: ", tx.TxID)
	return nil
}

func mockPOQRequest(inventory []*Inventory) *POQRequest {
	var poqRequest POQRequest
	json.Unmarshal([]byte(POQSample), &poqRequest)
	poqRequest.ProjectID = generateID("Prj", 3)
	poqRequest.RequestedResponseDate = time.Now().AddDate(0, 0, 1)
	poqRequest.ProductOfferingQualificationItem[0].Product.ProductSpecification.ID = generateID("UNI001-POQ", 3)
	poqRequest.ProductOfferingQualificationItem[1].Product.ProductSpecification.ID = generateID("ELINE001-POQ", 3)

	return &poqRequest
}

func mockOrderRequest(quote *QuoteResponse, id string) *OrderRequest {
	orderRequest := &OrderRequest{}
	json.Unmarshal([]byte(OrderSample), &orderRequest)
	orderRequest.ProjectID = quote.ProjectID
	orderRequest.ExternalID = id
	for i := 0; i < len(orderRequest.OrderItem); i++ {
		orderRequest.OrderItem[i].Quote.ID = quote.ID
		orderRequest.OrderItem[i].Qualification.ID = quote.QuoteItem[i].Qualification[0].ID
	}
	now := time.Now()
	orderRequest.CreateDate = now
	orderRequest.BuyerRequestDate = now
	orderRequest.RequestedCompletionDate = now.AddDate(0, 0, 2)

	return orderRequest
}

func mockSmartContract(quote *QuoteResponse, buyer, seed string) *SmartContractRequest {
	smRequest := &SmartContractRequest{}
	json.Unmarshal([]byte(SmartContractSample), &smRequest)
	smRequest.Buyer = User{Address: buyer, Seed: seed, Name: "LSOBus Bot"}

	for i := 0; i < len(quote.QuoteItem); i++ {
		smRequest.OrderItem[i].ID = generateID("PRD", 32)
		smRequest.OrderItem[i].ProductOffering.ID = quote.QuoteItem[i].Qualification[0].ID
		smRequest.OrderItem[i].Quote.ID = quote.ID
	}

	return smRequest
}

func generateID(prefix string, length int) string {
	return fmt.Sprintf("%s-%s", prefix, util.RandomFixedString(length))
}

func Hash() pkg.Hash {
	h := pkg.Hash{}
	_ = random.Bytes(h[:])
	return h
}
