package orchestra

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-lsobus/common/rest"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`

	RspLogin *LoginResponse
}

type LoginResponse struct {
	Data string `json:"data"`
}

func (p *PartnerImpl) TryUpdateApiToken() {
	if p.cfg.Username == "" || p.cfg.Password == "" {
		p.logger.Infof("partner %s username is empty", p.cfg.Name)
		return
	}

	reqParams := &LoginParams{}
	reqParams.Username = p.cfg.Username
	reqParams.Password = p.cfg.Password

	err := p.ExecAuthLogin(reqParams)
	if err == nil {
		p.logger.Infof("partner %s update api token, got new token %s", p.cfg.Name, reqParams.RspLogin.Data)
		p.apiToken = reqParams.RspLogin.Data
	} else {
		p.logger.Errorf("partner %s ExecAuthLogin err %s", p.cfg.Name, err)
	}
}

func (p *PartnerImpl) SetApiToken(token string) {
	p.apiToken = token
}

func (p *PartnerImpl) GetApiToken() string {
	if p.apiToken == "" {
		p.TryUpdateApiToken()
	}

	return p.apiToken
}

func (p *PartnerImpl) RenewApiToken() string {
	p.apiToken = ""
	p.TryUpdateApiToken()
	return p.apiToken
}

func (p *PartnerImpl) ClearApiToken() {
	p.apiToken = ""
}

func (p *PartnerImpl) ExecAuthLogin(params *LoginParams) error {
	var err error

	req := rest.Request{Method: rest.Post}
	req.BaseURL = p.GetSonataUrl() + "/api/login"
	req.Body, err = json.Marshal(params)
	if err != nil {
		return err
	}

	p.logger.Debugf("send login, url %s, username %s", req.BaseURL, params.Username)

	rsp, err := rest.Send(req)
	if p.GetFakeMode() {
		rsp = &rest.Response{}
		rsp.StatusCode = 200
		rsp.Body = "{\"data\": \"12345678\"}"
	} else if err != nil {
		return err
	}
	if rsp.StatusCode != 200 {
		return fmt.Errorf("receive response, StatusCode(%d) is not OK", rsp.StatusCode)
	}

	resp := &LoginResponse{}
	bdIO := bytes.NewBuffer([]byte(rsp.Body))
	cs := runtime.JSONConsumer()
	err = cs.Consume(bdIO, resp)
	if err != nil {
		return err
	}

	params.RspLogin = resp
	return nil
}
