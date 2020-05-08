package orchestra

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

	return nil
}
