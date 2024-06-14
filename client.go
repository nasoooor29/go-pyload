package gopyload

import (
	"net/http"
	"net/http/cookiejar"
)

type client struct {
	http.Client
	headers   headersMap
	ourCookie http.CookieJar
}

func NewClient() *client {
	jar, _ := cookiejar.New(nil)

	return &client{
		headers:   defaultPyloadHeaders,
		ourCookie: jar,
		Client: http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return nil
			},
		},
	}
}

func (c *client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addMapToHeaders(req, c.headers)
	return c.Client.Do(req)
}

func (c *client) Post(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addMapToHeaders(req, c.headers)
	return c.Client.Do(req)
}
