package sgb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type options struct {
	timeout time.Duration
	uri     string
}

type sgb struct {
	key string
	options
}

type option interface {
	apply(*options)
}

type optionFunc func(*options)

const (
	getMethod  = "GET"
	postMethod = "POST"
)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func New(key string, opts ...option) *sgb {
	const (
		_defaultTimeout = 200 * time.Second
		_defaultURI     = "https://api.silvergoldbull.com/v1/"
	)

	options := options{
		timeout: _defaultTimeout,
		uri:     _defaultURI,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &sgb{key, options}
}

func (s *sgb) request(reqMethod, entity string) ([]byte, error) {
	const _authField = "X-API-KEY"

	url := s.uri + entity
	client := &http.Client{Timeout: s.timeout}
	req, err := http.NewRequest(reqMethod, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(_authField, s.key)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}