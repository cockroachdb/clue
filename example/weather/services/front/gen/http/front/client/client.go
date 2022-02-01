// Code generated by goa v3.5.4, DO NOT EDIT.
//
// front client HTTP transport
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/front/design -o
// services/front

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the front service endpoint HTTP clients.
type Client struct {
	// Forecast Doer is the HTTP client used to make requests to the forecast
	// endpoint.
	ForecastDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the front service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ForecastDoer:        doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Forecast returns an endpoint that makes HTTP requests to the front service
// forecast server.
func (c *Client) Forecast() goa.Endpoint {
	var (
		decodeResponse = DecodeForecastResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildForecastRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ForecastDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("front", "forecast", err)
		}
		return decodeResponse(resp)
	}
}