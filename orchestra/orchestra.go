package orchestra

import ologger "github.com/go-openapi/runtime/logger"

type Orchestra struct {
	sonataSiteImpl  *sonataSiteImpl
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
	sonataInvImpl   *sonataInvImpl
}

func NewOrchestra() *Orchestra {
	o := &Orchestra{}
	o.sonataSiteImpl = newSonataSiteImpl()
	o.sonataPOQImpl = newSonataPOQImpl()
	o.sonataQuoteImpl = newSonataQuoteImpl()
	o.sonataOrderImpl = newSonataOrderImpl()
	o.sonataInvImpl = newSonataInvImpl()
	return o
}

func (o *Orchestra) Init() error {
	err := o.sonataSiteImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataPOQImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataQuoteImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataOrderImpl.Init()
	if err != nil {
		return err
	}

	err = o.sonataInvImpl.Init()
	if err != nil {
		return err
	}

	ologger.DebugEnabled()

	return nil
}

func (o *Orchestra) ExecPOQCreate(params *OrderParams) error {
	return o.sonataPOQImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecPOQFind(params *FindParams) error {
	return o.sonataPOQImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecQuoteCreate(params *OrderParams) error {
	return o.sonataQuoteImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecQuoteFind(params *FindParams) error {
	return o.sonataQuoteImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecOrderCreate(params *OrderParams) error {
	return o.sonataOrderImpl.SendCreateRequest(params)
}

func (o *Orchestra) ExecOrderFind(params *FindParams) error {
	return o.sonataOrderImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecInventoryFind(params *FindParams) error {
	return o.sonataInvImpl.SendFindRequest(params)
}

func (o *Orchestra) ExecSiteFind(params *FindParams) error {
	return o.sonataSiteImpl.SendFindRequest(params)
}
