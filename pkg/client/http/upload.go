package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// see: https://github.com/shenghui0779/gochat/blob/master/wx/http.go

// UploadForm is the interface for http upload
type UploadForm interface {
	// FieldName returns field name for upload
	FieldName() string

	// FileName returns filename for upload
	FileName() string

	// ExtraFields returns extra fields for upload
	ExtraFields() map[string]string

	// Buffer returns the buffer of media
	Buffer() ([]byte, error)
}

type httpUpload struct {
	fieldName   string
	filename    string
	resourceURL string
	extraFields map[string]string
}

func (u *httpUpload) FieldName() string {
	return u.fieldName
}

func (u *httpUpload) FileName() string {
	return u.filename
}

func (u *httpUpload) ExtraFields() map[string]string {
	return u.extraFields
}

func (u *httpUpload) Buffer() ([]byte, error) {
	if len(u.resourceURL) != 0 {
		resp, err := http.Get(u.resourceURL)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error http code: %d", resp.StatusCode)
		}

		return io.ReadAll(resp.Body)
	}

	path, err := filepath.Abs(u.filename)

	if err != nil {
		return nil, err
	}

	return os.ReadFile(path)
}

// UploadOption configures how we set up the http upload from.
type UploadOption func(u *httpUpload)

// WithResourceURL specifies http upload by resource url.
func WithResourceURL(url string) UploadOption {
	return func(u *httpUpload) {
		u.resourceURL = url
	}
}

// WithExtraField specifies the extra field to http upload from.
func WithExtraField(key, value string) UploadOption {
	return func(u *httpUpload) {
		u.extraFields[key] = value
	}
}

// NewUploadForm returns new upload form
func NewUploadForm(fieldName, filename string, options ...UploadOption) UploadForm {
	form := &httpUpload{
		fieldName: fieldName,
		filename:  filename,
	}

	if len(options) != 0 {
		form.extraFields = make(map[string]string)

		for _, f := range options {
			f(form)
		}
	}

	return form
}
