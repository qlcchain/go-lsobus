package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/go-resty/resty/v2"
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
	"requestedQuoteCompletionDate": "2021-01-19T08:19:10.528Z",
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
				"id": "{% response 'body', 'req_1d14aaa96cc64aa8b444b3e0ddb1046d', 'b64::JC5kYXRhLmlk::46b', 'never', 60 %}",
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
		Name:  "token",
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

			u.Info(fmt.Sprintf("buyer:%s, vendor: %s, api: %s", buyerSeedP, vendorP, apiTokenP))

			if err := mockOrderForBuyer(buyerSeedP, vendorP, apiTokenP); err != nil {
				u.Warn(err)
				return
			}
		},
	}
	parentCmd.AddCmd(cmd)
}

func mockOrderForBuyer(seed, vendor, token string) error {
	var account *pkg.Account
	if bytes, err := hex.DecodeString(seed); err != nil {
		return err
	} else {
		if s, err := pkg.BytesToSeed(bytes); err != nil {
			return err
		} else {
			var err error
			if account, err = s.Account(0); err != nil {
				return err
			}
		}
	}

	client := resty.New().SetHeader("CLIENT-KEY", token).
		SetHeader("Content-Type", "application/json").EnableTrace().SetDebug(true)

	switch strings.ToUpper(vendor) {
	case "HGC":
		return mockHGCOrderForBuyer(client, account)
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

type InventoryWrapper struct {
	Error string       `json:"error"`
	Data  []*Inventory `json:"data"`
	Code  string       `json:"code"`
	Meta  string       `json:"meta"`
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

type POQ struct {
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

type POQResponse struct {
	Error interface{} `json:"error"`
	Data  *POQ        `json:"data"`
	Code  interface{} `json:"code"`
	Meta  interface{} `json:"meta"`
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

type Quote struct {
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

type QuoteResponse struct {
	Error interface{} `json:"error"`
	Data  *Quote      `json:"data"`
	Code  interface{} `json:"code"`
	Meta  interface{} `json:"meta"`
}

func mockHGCOrderForBuyer(client *resty.Client, account *pkg.Account) error {
	u.Info("STEP1, query inventory")
	resp, err := client.R().SetResult(&InventoryWrapper{}).Get(fmt.Sprintf("%s/v1/orders/product-inventory", endpointP))
	if err != nil {
		return err
	}
	inventory := resp.Result().(*InventoryWrapper).Data
	//u.Info(u.ToIndentString(inventory))

	u.Info("STEP2, create POQ")

	poqReq := mockPOQRequest(inventory)
	resp, err = client.R().SetBody(poqReq).SetResult(&POQResponse{}).Post(fmt.Sprintf("%s/v1/poq", endpointP))
	if err != nil {
		return err
	}
	poq := resp.Result().(*POQResponse).Data
	//u.Info(u.ToIndentString(poq))

	u.Info("STEP3, create quote")
	quoteReq := strings.ReplaceAll(QuoteSample, "MyProject-03", poq.ProjectID)
	//quoteReq = strings.ReplaceAll(quoteReq, "${QUALIFICATION_ID}", poq.ID)
	//u.Info(quoteReq)
	resp, err = client.R().SetBody(quoteReq).SetResult(&QuoteResponse{}).Post(fmt.Sprintf("%s/v1/quotes", endpointP))
	if err != nil {
		return err
	}
	quote := resp.Result().(*QuoteResponse).Data
	u.Info(quote)

	u.Info("STEP4, place an order")

	u.Info("STEP5, create smart contract to save order info")

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

func generateID(prefix string, length int) string {
	return fmt.Sprintf("%s-%s", prefix, util.RandomFixedString(length))
}
