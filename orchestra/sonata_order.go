package orchestra

type sonataOrderImpl struct {
	sonataBaseImpl
}

func newSonataOrderImpl() *sonataOrderImpl {
	s := &sonataOrderImpl{}
	return s
}

func (s *sonataOrderImpl) Init() error {
	return s.sonataBaseImpl.Init()
}
