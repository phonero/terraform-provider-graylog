package dashboard

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(ctx context.Context, id string) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/dashboards/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Gets(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/dashboards",
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if data == nil {
		return nil, nil, errors.New("request body is nil")
	}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/dashboards",
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, id string, data map[string]interface{},
) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	if data == nil {
		return nil, errors.New("request body is nil")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "PUT",
		Path:        "/dashboards/" + id,
		RequestBody: data,
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context, id string) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/dashboards/" + id,
	})
	return resp, err
}
