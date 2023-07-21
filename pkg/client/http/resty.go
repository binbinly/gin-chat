package http

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

// docs: https://github.com/go-resty/resty

type restyClient struct {
	timeout time.Duration
}

// NewRestyClient 创建http client客户端
func NewRestyClient() Client {
	return &restyClient{
		timeout: defaultTimeout,
	}
}

// Get 发送get请求
func (r *restyClient) Get(ctx context.Context, url string, options ...ClientOption) ([]byte, error) {
	client := resty.New()
	r.setting(ctx, client, options...)
	//client.SetDebug(true)
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Post 发送form post请求
func (r *restyClient) Post(ctx context.Context, url string, data map[string]string, options ...ClientOption) ([]byte, error) {
	options = append(options, WithHTTPHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8"))
	client := resty.New()
	r.setting(ctx, client, options...)
	resp, err := client.R().SetFormData(data).Post(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// PostJSON 发送post raw json 请求
func (r *restyClient) PostJSON(ctx context.Context, url string, body []byte, options ...ClientOption) ([]byte, error) {
	options = append(options, WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))
	client := resty.New()
	r.setting(ctx, client, options...)
	resp, err := client.R().SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Upload 上传
func (r *restyClient) Upload(ctx context.Context, url string, form UploadForm, options ...ClientOption) (b []byte, err error) {
	return
}

func (r *restyClient) setting(ctx context.Context, client *resty.Client, options ...ClientOption) {
	settings := &httpSettings{timeout: r.timeout}

	if len(options) != 0 {
		settings.headers = make(map[string]string)

		for _, f := range options {
			f(settings)
		}
	}

	if r.timeout != 0 {
		client.SetTimeout(settings.timeout)
	}

	// headers
	if len(settings.headers) != 0 {
		client.SetHeaders(settings.headers)
	}

	// cookies
	if len(settings.cookies) != 0 {
		client.SetCookies(settings.cookies)
	}

	if settings.close {
		client.SetCloseConnection(true)
	}
}
