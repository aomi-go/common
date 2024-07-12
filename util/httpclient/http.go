package httpclient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/go-resty/resty/v2"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

type WebClientOptions struct {
	BaseUrl  string         `json:"base_url" yaml:"base_url"`
	Timeout  *time.Duration `json:"timeout" yaml:"timeout"`
	ProxyUrl string         `json:"proxy_url" yaml:"proxy_url"`
	TLS      *TLS           `json:"tls" yaml:"tls"`
}

type TLS struct {
	InsecureSkipVerify bool `json:"insecure_skip_verify" yaml:"insecure_skip_verify" desc:"Skip SSL verification"`
}

func NewWebClient(options *WebClientOptions) *WebClient {
	client := resty.New()
	if "" != options.BaseUrl {
		client.SetBaseURL(options.BaseUrl)
	}
	if nil != options.Timeout {
		client.SetTimeout(*options.Timeout)
	}
	if "" != options.ProxyUrl {
		client.SetProxy(options.ProxyUrl)
	}
	if nil != options.TLS {
		cfg := tls.Config{}
		cfg.InsecureSkipVerify = options.TLS.InsecureSkipVerify
		client.SetTLSClientConfig(&cfg)
	}

	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.IsError() {
			return &ApiError{
				Code:         EXCEPTION,
				Msg:          "HTTP请求处理失败",
				HttpResponse: response,
			}
		}

		return nil
	})

	return &WebClient{
		options:  options,
		provider: client,
	}
}

type WebClient struct {
	options  *WebClientOptions
	provider *resty.Client
}

func (c *WebClient) Post(ctx context.Context, url string, headers map[string]string, body any) (*resty.Response, error) {
	return c.provider.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(body).
		Post(url)
}

func (c *WebClient) Execute(
	ctx context.Context,
	method string,
	url string,
	headers map[string]string,
	queryParams url.Values,
	formParams url.Values,
	body any,
) (*resty.Response, error) {
	return c.provider.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetQueryParamsFromValues(queryParams).
		SetFormDataFromValues(formParams).
		SetBody(body).
		Execute(method, url)
}

func (c *WebClient) GetProvider() *resty.Client {
	return c.provider
}

func (c *WebClient) Decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}
