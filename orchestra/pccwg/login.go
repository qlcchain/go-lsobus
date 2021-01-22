package pccwg

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-lsobus/common/rest"

	"github.com/qlcchain/go-lsobus/api"
)

func (p *PCCWGImpl) TryUpdateApiToken() {
	reqParams := &api.LoginParams{}
	reqParams.Username = p.user
	reqParams.Password = p.password

	err := p.ExecAuthLogin(reqParams)
	if err == nil {
		p.logger.Infof("update api token, got new token %s", reqParams.RspLogin.Data)
		p.apiToken = reqParams.RspLogin.Data
	} else {
		p.logger.Errorf("ExecAuthLogin err %s", err)
	}
}

func (p *PCCWGImpl) SetApiToken(token string) {
	p.apiToken = token
}

func (p *PCCWGImpl) GetAPIToken() string {
	if p.apiToken == "" {
		p.TryUpdateApiToken()
	}

	return p.apiToken
}

func (p *PCCWGImpl) RenewAPIToken() string {
	p.apiToken = ""
	p.TryUpdateApiToken()
	return p.apiToken
}

func (p *PCCWGImpl) ClearAPIToken() {
	p.apiToken = ""
}

func (p *PCCWGImpl) ExecAuthLogin(params *api.LoginParams) error {
	var err error

	req := rest.Request{Method: rest.Post}
	req.BaseURL = p.GetSellerConfig().BackEndURL + "/api/login"
	req.Body, err = json.Marshal(params)
	if err != nil {
		return err
	}

	p.logger.Debugf("send login, url %s, username %s", req.BaseURL, params.Username)

	rsp, err := rest.Send(req)
	if p.GetSellerConfig().IsFake {
		rsp = &rest.Response{}
		rsp.StatusCode = 200
		rsp.Body = "{\"data\": \"12345678\"}"
	} else if err != nil {
		return err
	}
	if rsp.StatusCode != 200 {
		return fmt.Errorf("receive response, StatusCode(%d) is not OK", rsp.StatusCode)
	}

	resp := &api.LoginResponse{}
	bdIO := bytes.NewBuffer([]byte(rsp.Body))
	cs := runtime.JSONConsumer()
	err = cs.Consume(bdIO, resp)
	if err != nil {
		return err
	}

	params.RspLogin = resp
	return nil
}
