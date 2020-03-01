package elasticsearch

import (
	"github.com/olivere/elastic/v7"
)

type client struct {
	systems map[string]*elastic.Client
}

func (c client) Session(name string) *elastic.Client {
	return c.systems[name]
}

func (c client) Names() []string {
	l := make([]string, 0, len(c.systems))
	for n := range c.systems {
		l = append(l, n)
	}
	return l
}

func (c client) Close() {
	for _, v := range c.systems {
		v.Stop()
	}
}
