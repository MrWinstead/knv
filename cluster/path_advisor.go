package cluster

import "path"

var _ PathAdvisor = &etcdPathAdvisor{}

type PathAdvisor interface {
	ExpandPath(parts ...string) string
}

type etcdPathAdvisor struct {
	prefix string
}

func NewEtcdPathAdvisor(prefix string) PathAdvisor {
	created := &etcdPathAdvisor{
		prefix: prefix,
	}
	return created
}

func (a *etcdPathAdvisor) ExpandPath(parts ...string) string {
	fullParts := []string{a.prefix}
	fullParts = append(fullParts, parts...)
	expanded := path.Join(fullParts...)
	if '/' != expanded[0] {
		expanded = "/" + expanded
	}
	return expanded
}
