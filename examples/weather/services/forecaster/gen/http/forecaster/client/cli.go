// Code generated by goa v3.16.1, DO NOT EDIT.
//
// Forecaster HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/forecaster/design -o
// services/forecaster

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	forecaster "goa.design/pulse/examples/weather/services/forecaster/gen/forecaster"
)

// BuildForecastPayload builds the payload for the Forecaster forecast endpoint
// from CLI flags.
func BuildForecastPayload(forecasterForecastState string, forecasterForecastCity string) (*forecaster.ForecastPayload, error) {
	var err error
	var state string
	{
		state = forecasterForecastState
		if utf8.RuneCountInString(state) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, true))
		}
		if utf8.RuneCountInString(state) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var city string
	{
		city = forecasterForecastCity
	}
	v := &forecaster.ForecastPayload{}
	v.State = state
	v.City = city

	return v, nil
}
