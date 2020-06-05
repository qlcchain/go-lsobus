package offer

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-lsobus/common/rest"
	"github.com/qlcchain/go-lsobus/log"
)

type APIProductOfferingManagement struct {
	BaseURL string
	Client  *rest.Client
	logger  *zap.SugaredLogger
}

func NewAPIProductOfferingManagement(baseUrl string) *APIProductOfferingManagement {
	a := &APIProductOfferingManagement{
		BaseURL: baseUrl,
		Client: &rest.Client{
			HTTPClient: &http.Client{},
		},
	}
	a.logger = log.NewLogger("APIProductOffering")
	return a
}

type ProductOfferingFindParams struct {
	Provider *string `json:"provider,omitempty"`
	Type     *string `json:"type,omitempty"`
	Deleted  *bool   `json:"deleted,omitempty"`
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
}

type ProductOfferingGetParams struct {
	ProductOfferingID string
}

func (a *APIProductOfferingManagement) ProductOfferingFind(params *ProductOfferingFindParams) (*FindResponse, error) {
	req := rest.Request{
		Method:      rest.Get,
		BaseURL:     a.BaseURL,
		Headers:     nil,
		QueryParams: nil,
	}

	if params != nil {
		req.QueryParams = make(map[string]string)
		if params.Provider != nil {
			req.QueryParams["provider"] = *params.Provider
		}
		if params.Type != nil {
			req.QueryParams["type"] = *params.Type
		}
		if params.Deleted != nil {
			req.QueryParams["deleted"] = strconv.FormatBool(*params.Deleted)
		}
		if params.Page != nil {
			req.QueryParams["page"] = strconv.Itoa(*params.Page)
		}
		if params.PageSize != nil {
			req.QueryParams["page_size"] = strconv.Itoa(*params.PageSize)
		}
	}

	rsp, err := a.Client.Send(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("receive response, StatusCode(%d) is not OK", rsp.StatusCode)
	}

	resp := &FindResponse{}

	// response header X-Result-Count
	xrcHdrs := rsp.Headers["X-Result-Count"]
	if len(xrcHdrs) > 0 {
		resp.XResultCount = xrcHdrs[0]
	}

	// response header X-Total_Count
	xtcHdrs := rsp.Headers["X-Total_Count"]
	if len(xtcHdrs) > 0 {
		resp.XTotalCount = xtcHdrs[0]
	}

	// response payload
	bdIO := bytes.NewBuffer([]byte(rsp.Body))
	cs := runtime.JSONConsumer()
	err = cs.Consume(bdIO, &resp.Payload)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *APIProductOfferingManagement) ProductOfferingGet(params *ProductOfferingGetParams) (*GetResponse, error) {
	if params == nil || params.ProductOfferingID == "" {
		return nil, errors.New("invalid get params")
	}

	req := rest.Request{
		Method:      rest.Get,
		BaseURL:     fmt.Sprintf("%s/%s", a.BaseURL, params.ProductOfferingID),
		Headers:     nil,
		QueryParams: nil,
	}

	rsp, err := a.Client.Send(req)
	if err != nil {
		return nil, err
	}

	resp := &GetResponse{}
	bdIO := bytes.NewBuffer([]byte(rsp.Body))
	cs := runtime.JSONConsumer()
	err = cs.Consume(bdIO, &(resp.Payload))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
