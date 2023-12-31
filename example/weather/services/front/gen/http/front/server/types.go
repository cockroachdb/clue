// Code generated by goa v3.11.0, DO NOT EDIT.
//
// front HTTP server types
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/front/design -o
// services/front

package server

import (
	front "goa.design/clue/example/weather/services/front/gen/front"
	goa "goa.design/goa/v3/pkg"
)

// ForecastResponseBody is the type of the "front" service "forecast" endpoint
// HTTP response body.
type ForecastResponseBody struct {
	// Forecast location
	Location *LocationResponseBody `form:"location" json:"location" xml:"location"`
	// Weather forecast periods
	Periods []*PeriodResponseBody `form:"periods" json:"periods" xml:"periods"`
}

// ForecastNotUsaResponseBody is the type of the "front" service "forecast"
// endpoint HTTP response body for the "not_usa" error.
type ForecastNotUsaResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// LocationResponseBody is used to define fields on response body types.
type LocationResponseBody struct {
	// Latitude
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// Longitude
	Long float64 `form:"long" json:"long" xml:"long"`
	// City
	City string `form:"city" json:"city" xml:"city"`
	// State
	State string `form:"state" json:"state" xml:"state"`
}

// PeriodResponseBody is used to define fields on response body types.
type PeriodResponseBody struct {
	// Period name
	Name string `form:"name" json:"name" xml:"name"`
	// Start time
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
	// End time
	EndTime string `form:"endTime" json:"endTime" xml:"endTime"`
	// Temperature
	Temperature int `form:"temperature" json:"temperature" xml:"temperature"`
	// Temperature unit
	TemperatureUnit string `form:"temperatureUnit" json:"temperatureUnit" xml:"temperatureUnit"`
	// Summary
	Summary string `form:"summary" json:"summary" xml:"summary"`
}

// NewForecastResponseBody builds the HTTP response body from the result of the
// "forecast" endpoint of the "front" service.
func NewForecastResponseBody(res *front.Forecast2) *ForecastResponseBody {
	body := &ForecastResponseBody{}
	if res.Location != nil {
		body.Location = marshalFrontLocationToLocationResponseBody(res.Location)
	}
	if res.Periods != nil {
		body.Periods = make([]*PeriodResponseBody, len(res.Periods))
		for i, val := range res.Periods {
			body.Periods[i] = marshalFrontPeriodToPeriodResponseBody(val)
		}
	}
	return body
}

// NewForecastNotUsaResponseBody builds the HTTP response body from the result
// of the "forecast" endpoint of the "front" service.
func NewForecastNotUsaResponseBody(res *goa.ServiceError) *ForecastNotUsaResponseBody {
	body := &ForecastNotUsaResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
