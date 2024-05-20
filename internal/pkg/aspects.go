package pkg

type AspectList struct {
	List []*Aspect `yaml:"list"`
}

func (al AspectList) Each(handler func(fragment *AspectFragment, err error)) {
	for _, aspect := range al.List {
		fragment, err := aspect.FetchData()
		handler(fragment, err)
	}
}
