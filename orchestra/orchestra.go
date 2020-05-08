package orchestra

type Orchestra struct {
	sonataPOQImpl   *sonataPOQImpl
	sonataQuoteImpl *sonataQuoteImpl
	sonataOrderImpl *sonataOrderImpl
}

func NewOrchestra() *Orchestra {
	o := &Orchestra{}
	o.sonataPOQImpl = newSonataPOQImpl()
	o.sonataQuoteImpl = newSonataQuoteImpl()
	o.sonataOrderImpl = newSonataOrderImpl()
	return o
}

func (o *Orchestra) Init() error {
	err := o.sonataPOQImpl.Init()
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

	return nil
}
