// Code generated by goa v3.11.0, DO NOT EDIT.
//
// front HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/front/design -o
// services/front

package server

import (
	"context"
	"errors"
	"net/http"

	front "goa.design/clue/example/weather/services/front/gen/front"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeForecastResponse returns an encoder for responses returned by the
// front forecast endpoint.
func EncodeForecastResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*front.Forecast2)
		enc := encoder(ctx, w)
		body := NewForecastResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeForecastRequest returns a decoder for requests sent to the front
// forecast endpoint.
func DecodeForecastRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			ip  string
			err error

			params = mux.Vars(r)
		)
		ip = params["ip"]
		err = goa.MergeErrors(err, goa.ValidateFormat("ip", ip, goa.FormatIP))
		if err != nil {
			return nil, err
		}
		payload := ip

		return payload, nil
	}
}

// EncodeForecastError returns an encoder for errors returned by the forecast
// front endpoint.
func EncodeForecastError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_usa":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewForecastNotUsaResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalFrontLocationToLocationResponseBody builds a value of type
// *LocationResponseBody from a value of type *front.Location.
func marshalFrontLocationToLocationResponseBody(v *front.Location) *LocationResponseBody {
	res := &LocationResponseBody{
		Lat:   v.Lat,
		Long:  v.Long,
		City:  v.City,
		State: v.State,
	}

	return res
}

// marshalFrontPeriodToPeriodResponseBody builds a value of type
// *PeriodResponseBody from a value of type *front.Period.
func marshalFrontPeriodToPeriodResponseBody(v *front.Period) *PeriodResponseBody {
	res := &PeriodResponseBody{
		Name:            v.Name,
		StartTime:       v.StartTime,
		EndTime:         v.EndTime,
		Temperature:     v.Temperature,
		TemperatureUnit: v.TemperatureUnit,
		Summary:         v.Summary,
	}

	return res
}
