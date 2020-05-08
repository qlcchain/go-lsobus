package orchestra

type sonataSiteImpl struct {
	sonataBaseImpl
}

func newSonataSiteImpl() *sonataSiteImpl {
	s := &sonataSiteImpl{}
	return s
}

func (s *sonataSiteImpl) Init() error {
	return s.sonataBaseImpl.Init()
}
