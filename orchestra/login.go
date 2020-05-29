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

func (o *Orchestra) updateApiToken() {
	reqParams := &LoginParams{Username: "yuangui.wen@qlink.mobi", Password: "Wyg12345"}
	err := o.ExecAuthLogin(reqParams)
	if err == nil {
		o.logger.Infof("update api token, got new token %s", reqParams.RspLogin.Data)
		o.apiToken = reqParams.RspLogin.Data
	} else {
		o.logger.Errorf("ExecAuthLogin err %s", err)
	}
}

func (o *Orchestra) GetApiToken() string {
	if o.apiToken == "" {
		o.updateApiToken()
	}

	return o.apiToken
}

func (o *Orchestra) ExecAuthLogin(params *LoginParams) error {
	var err error

	req := rest.Request{Method: rest.Post}
	req.BaseURL = o.GetSonataUrl("") + "/api/login"
	req.Body, err = json.Marshal(params)
	if err != nil {
		return err
	}

	o.logger.Debugf("send login, url %s, username %s", req.BaseURL, params.Username)

	rsp, err := rest.Send(req)
	if err != nil {
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
