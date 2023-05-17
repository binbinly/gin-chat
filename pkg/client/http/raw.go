package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// raw 使用原生包封装的 http client
// see: https://github.com/shenghui0779/gochat/blob/master/wx/http.go

// rawClient
type rawClient struct {
	client  *http.Client
	timeout time.Duration
}

// NewRawClient 实例化 http client
func NewRawClient(tlsCfg ...*tls.Config) Client {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if len(tlsCfg) != 0 {
		t.TLSClientConfig = tlsCfg[0]
	}

	return &rawClient{
		client: &http.Client{
			Transport: t,
		},
		timeout: defaultTimeout,
	}
}

// Get http get request
func (r *rawClient) Get(ctx context.Context, url string, options ...ClientOption) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return r.do(ctx, req, options...)
}

// Post http post request
func (r *rawClient) Post(ctx context.Context, url string, data map[string]string, options ...ClientOption) ([]byte, error) {
	options = append(options, WithHTTPHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8"))
	formData := r.buildForms(data).Encode()
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	return r.do(ctx, req, options...)
}

// PostJSON http json post request
func (r *rawClient) PostJSON(ctx context.Context, url string, body []byte, options ...ClientOption) ([]byte, error) {
	options = append(options, WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	return r.do(ctx, req, options...)
}

// Upload 文件上传
func (r *rawClient) Upload(ctx context.Context, url string, form UploadForm, options ...ClientOption) ([]byte, error) {
	media, err := form.Buffer()

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4kb
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile(form.FieldName(), form.FileName())

	if err != nil {
		return nil, err
	}

	if _, err = fw.Write(media); err != nil {
		return nil, err
	}

	// add extra fields
	if extraFields := form.ExtraFields(); len(extraFields) != 0 {
		for k, v := range extraFields {
			if err = w.WriteField(k, v); err != nil {
				return nil, err
			}
		}
	}

	options = append(options, WithHTTPHeader("Content-Type", w.FormDataContentType()))

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, buf)

	if err != nil {
		return nil, err
	}

	return r.do(ctx, req, options...)
}

func (r *rawClient) do(ctx context.Context, req *http.Request, options ...ClientOption) ([]byte, error) {
	settings := &httpSettings{timeout: r.timeout}

	if len(options) != 0 {
		settings.headers = make(map[string]string)

		for _, f := range options {
			f(settings)
		}
	}

	// headers
	if len(settings.headers) != 0 {
		for k, v := range settings.headers {
			req.Header.Set(k, v)
		}
	}

	// cookies
	if len(settings.cookies) != 0 {
		for _, v := range settings.cookies {
			req.AddCookie(v)
		}
	}

	if settings.close {
		req.Close = true
	}

	// timeout
	ctx, cancel := context.WithTimeout(ctx, settings.timeout)
	defer cancel()

	resp, err := r.client.Do(req.WithContext(ctx))

	if err != nil {
		// If the context has been canceled, the context's error is probably more useful.
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)

		return nil, fmt.Errorf("error http code: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// buildForms build post Form data
func (r *rawClient) buildForms(data map[string]string) (Forms url.Values) {
	Forms = url.Values{}
	for key, value := range data {
		Forms.Add(key, value)
	}
	return Forms
}
