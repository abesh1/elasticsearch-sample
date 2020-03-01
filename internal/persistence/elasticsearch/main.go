package elasticsearch

import (
	"time"

	"github.com/olivere/elastic/v7"

	"log"
)

func NewClient(cfg Config) (Client, error) {
	c := &client{
		systems: make(map[string]*elastic.Client),
	}
	for k, v := range cfg {
		sess, err := newSession(v)
		if err != nil {
			return nil, err
		}
		c.systems[k] = sess
	}
	return c, nil
}

func MustNewClient(cfg Config) Client {
	c, err := NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func newSession(sys *systemConf) (*elastic.Client, error) {
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(sys.URL),
		elastic.SetGzip(false),
		elastic.SetHealthcheck(true),
		elastic.SetHealthcheckInterval(60 * time.Second),
		elastic.SetHealthcheckTimeout(1 * time.Second),
		elastic.SetHealthcheckTimeoutStartup(5 * time.Second),
		elastic.SetSniff(false),
		elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewConstantBackoff(2 * time.Second))),
	}

	esc, err := elastic.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return esc, nil
}
