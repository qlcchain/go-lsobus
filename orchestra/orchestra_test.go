package orchestra

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/qlcchain/go-lsobus/config"
)

type mockOrchestraData struct {
	cm *config.CfgManager
}

func setupTestCase(t *testing.T) (*mockOrchestraData, func(t *testing.T)) {
	t.Log("setup test case")

	dir := filepath.Join(config.TestDataDir(), uuid.New().String())
	cm := config.NewCfgManager(dir)
	_, err := cm.Load()
	if err != nil {
		t.Fatal(err)
	}

	md := &mockOrchestraData{}
	md.cm = cm

	return md, func(t *testing.T) {
		t.Log("teardown test case")
		err := os.RemoveAll(dir)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func setupOrchestraConfig(o *Orchestra) {
	o.cfg.Partners = nil
	p1 := &config.PartnerCfg{
		Name:      "PCCW",
		ID:        "PCCW",
		SonataUrl: "http://127.0.0.1:7777",
		Username:  "test",
		Password:  "test",
	}
	o.cfg.Partners = append(o.cfg.Partners, p1)
}

func setupOrderParams() *OrderParams {
	createParams := &OrderParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}

	uniItem := &UNIItemParams{
		SiteID:    "site111",
		PortSpeed: 1000,
	}
	uniItem.ProdOfferID = "offer111"
	uniItem.QuoteID = "quote111"
	uniItem.Name = "port111"
	uniItem.BillingParams = &BillingParams{
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Add(24 * time.Hour).Unix(),
		Currency:  "USD",
		Price:     12.34,
	}
	createParams.UNIItems = append(createParams.UNIItems, uniItem)

	lineItem := &ELineItemParams{
		Bandwidth:     100,
		CosName:       "GOLD",
		SrcPortID:     "port111",
		DstPortID:     "port222",
		SrcLocationID: "loc111",
		DstLocationID: "loc222",
	}
	lineItem.ProdOfferID = "offer111"
	lineItem.QuoteID = "quote111"
	lineItem.Name = "line111"
	lineItem.BillingParams = &BillingParams{
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Add(24 * time.Hour).Unix(),
		Currency:  "USD",
		Price:     56.78,
	}
	createParams.ELineItems = append(createParams.ELineItems, lineItem)

	return createParams
}

func TestOrchestra_Site(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecSiteFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "SITE-TEST-1",
	}
	err = o.ExecSiteGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_Offer(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecOfferFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "OFFER-TEST-1",
	}
	err = o.ExecOfferGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_POQ(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	createParams := setupOrderParams()

	err = o.ExecPOQCreate(createParams)
	if err != nil {
		t.Fatal(err)
	}

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecPOQFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "POQ-TEST-1",
	}
	err = o.ExecPOQGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_Quote(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	createParams := setupOrderParams()

	err = o.ExecQuoteCreate(createParams)
	if err != nil {
		t.Fatal(err)
	}

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecQuoteFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "QUOTE-TEST-1",
	}
	err = o.ExecQuoteGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_Order(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	createParams := setupOrderParams()

	err = o.ExecOrderCreate(createParams)
	if err != nil {
		t.Fatal(err)
	}

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecOrderFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "ORDER-TEST-1",
	}
	err = o.ExecOrderGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_Inventory(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	findParams := &FindParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},
	}
	err = o.ExecInventoryFind(findParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := &GetParams{
		Buyer:  &PartnerParams{ID: "CBC", Name: "CBC"},
		Seller: &PartnerParams{ID: "PCCW", Name: "PCCW"},

		ID: "PRODUCT-TEST-1",
	}
	err = o.ExecInventoryGet(getParams)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrchestra_Login(t *testing.T) {
	md, tearDown := setupTestCase(t)
	defer tearDown(t)

	o := NewOrchestra(md.cm.ConfigFile)

	setupOrchestraConfig(o)

	err := o.Init()
	if err != nil {
		t.Fatal(err)
	}
	o.SetFakeMode(true)

	p := o.GetPartnerImpl("PCCW")
	p.ClearApiToken()

	if p.GetApiToken() == "" {
		t.Fatal("token empty")
	}

	p.RenewApiToken()
	if p.GetApiToken() == "" {
		t.Fatal("token empty")
	}
}
