package orchestra

type sonataInvImpl struct {
	sonataBaseImpl
}

func newSonataInvImpl() *sonataInvImpl {
	s := &sonataInvImpl{}
	return s
}

func (s *sonataInvImpl) Init() error {
	return s.sonataBaseImpl.Init()
}
