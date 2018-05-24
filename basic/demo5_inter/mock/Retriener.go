package mock

type Retriener struct {
	Context string
}

func (r Retriener) String() string {
	panic("implement me")
}

func (r Retriener) Get(url string) string {
	return r.Context
}
