package elasticsearch

import "github.com/olivere/elastic/v7"

type Client interface {
	Session(string) *elastic.Client
	Names() []string
	Close()
}
