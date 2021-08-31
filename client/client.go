package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/VitoChueng/vito_infra/logger"
)

const contentTypeJson = "application/json"

type Option struct {
	Token     string
	UrlPrefix string
	Timeout   time.Duration
}

type Client struct {
	Option
}

func NewClient(clientOption Option) *Client {
	return &Client{
		clientOption,
	}
}

func (c *Client) SetTimeOut(t time.Duration) {
	c.Option.Timeout = t
}

func (c *Client) tracingClient() *http.Client {
	httpClient := http.DefaultClient
	httpClient.Timeout = c.Timeout
	if httpClient.Timeout == 0 {
		httpClient.Timeout = 120 * time.Second
	}
	return httpClient
}

// Post http post method request api
func (c *Client) Post(route string, queryParams map[string]string, postParams interface{}) (string, error) {
	postJSON, err := json.Marshal(postParams)
	if err != nil {
		return "", errors.Wrapf(err, "序列化POST参数失败; postParams:%v; route:%s", postParams, route)
	}
	return c.post(contentTypeJson, route, queryParams, string(postJSON))
}

func (c *Client) generateURL(route string, params map[string]string) string {
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}
	return c.UrlPrefix + route + vals.Encode()
}

func (c *Client) post(contentType string, route string, queryParmas map[string]string, postParams string) (string, error) {
	logger.TransLogger.Sugar().Infof("request %s postParams [%s] and queryParams [%v]", route, postParams, queryParmas)
	reqUrl := c.generateURL(route, queryParmas)
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(postParams))
	if err != nil {
		return "", errors.WithStack(err)
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := c.tracingClient().Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "req error with url:[%s],params:[%v]", reqUrl, postParams)
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "ioutil.ReadAll url:%s params: %s", reqUrl, postParams)
	}
	return string(body), nil
}
