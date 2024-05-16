// Code generated by goa v3.16.1, DO NOT EDIT.
//
// Forecaster HTTP client types
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/forecaster/design -o
// services/forecaster

package client

import (
	goa "goa.design/goa/v3/pkg"
	forecaster "goa.design/pulse/examples/weather/services/forecaster/gen/forecaster"
)

// ForecastResponseBody is the type of the "Forecaster" service "forecast"
// endpoint HTTP response body.
type ForecastResponseBody struct {
	// Forecast location
	Location *LocationResponseBody `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Weather forecast periods
	Periods []*PeriodResponseBody `form:"periods,omitempty" json:"periods,omitempty" xml:"periods,omitempty"`
}

// ForecastTimeoutResponseBody is the type of the "Forecaster" service
// "forecast" endpoint HTTP response body for the "timeout" error.
type ForecastTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LocationResponseBody is used to define fields on response body types.
type LocationResponseBody struct {
	// Latitude
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// Longitude
	Long *float64 `form:"long,omitempty" json:"long,omitempty" xml:"long,omitempty"`
	// City
	City *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	// State
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// PeriodResponseBody is used to define fields on response body types.
type PeriodResponseBody struct {
	// Period name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Start time
	StartTime *string `form:"startTime,omitempty" json:"startTime,omitempty" xml:"startTime,omitempty"`
	// End time
	EndTime *string `form:"endTime,omitempty" json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Temperature
	Temperature *int `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// Temperature unit
	TemperatureUnit *string `form:"temperatureUnit,omitempty" json:"temperatureUnit,omitempty" xml:"temperatureUnit,omitempty"`
	// Summary
	Summary *string `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
}

// NewForecast2OK builds a "Forecaster" service "forecast" endpoint result from
// a HTTP "OK" response.
func NewForecast2OK(body *ForecastResponseBody) *forecaster.Forecast2 {
	v := &forecaster.Forecast2{}
	v.Location = unmarshalLocationResponseBodyToForecasterLocation(body.Location)
	v.Periods = make([]*forecaster.Period, len(body.Periods))
	for i, val := range body.Periods {
		v.Periods[i] = unmarshalPeriodResponseBodyToForecasterPeriod(val)
	}

	return v
}

// NewForecastTimeout builds a Forecaster service forecast endpoint timeout
// error.
func NewForecastTimeout(body *ForecastTimeoutResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateForecastResponseBody runs the validations defined on
// ForecastResponseBody
func ValidateForecastResponseBody(body *ForecastResponseBody) (err error) {
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.Periods == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("periods", "body"))
	}
	if body.Location != nil {
		if err2 := ValidateLocationResponseBody(body.Location); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.Periods {
		if e != nil {
			if err2 := ValidatePeriodResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateForecastTimeoutResponseBody runs the validations defined on
// forecast_timeout_response_body
func ValidateForecastTimeoutResponseBody(body *ForecastTimeoutResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLocationResponseBody runs the validations defined on
// LocationResponseBody
func ValidateLocationResponseBody(body *LocationResponseBody) (err error) {
	if body.Lat == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lat", "body"))
	}
	if body.Long == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("long", "body"))
	}
	if body.City == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("city", "body"))
	}
	if body.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("state", "body"))
	}
	return
}

// ValidatePeriodResponseBody runs the validations defined on PeriodResponseBody
func ValidatePeriodResponseBody(body *PeriodResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.StartTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("startTime", "body"))
	}
	if body.EndTime == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("endTime", "body"))
	}
	if body.Temperature == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temperature", "body"))
	}
	if body.TemperatureUnit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temperatureUnit", "body"))
	}
	if body.Summary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("summary", "body"))
	}
	if body.StartTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.startTime", *body.StartTime, goa.FormatDateTime))
	}
	if body.EndTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.endTime", *body.EndTime, goa.FormatDateTime))
	}
	return
}
