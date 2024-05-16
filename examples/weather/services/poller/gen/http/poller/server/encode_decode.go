// Code generated by goa v3.16.1, DO NOT EDIT.
//
// Poller HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/poller/design -o
// services/poller

package server

import (
	"context"
	"errors"
	"net/http"
	"unicode/utf8"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	poller "goa.design/pulse/examples/weather/services/poller/gen/poller"
)

// EncodeAddLocationResponse returns an encoder for responses returned by the
// Poller add_location endpoint.
func EncodeAddLocationResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeAddLocationRequest returns a decoder for requests sent to the Poller
// add_location endpoint.
func DecodeAddLocationRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			city  string
			state string
			err   error
		)
		qp := r.URL.Query()
		city = qp.Get("city")
		if city == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("city", "query string"))
		}
		state = qp.Get("state")
		if state == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("state", "query string"))
		}
		if utf8.RuneCountInString(state) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, true))
		}
		if utf8.RuneCountInString(state) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddLocationCityAndState(city, state)

		return payload, nil
	}
}

// EncodeAddLocationError returns an encoder for errors returned by the
// add_location Poller endpoint.
func EncodeAddLocationError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "location_exists":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddLocationLocationExistsResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeSubscribeRequest returns a decoder for requests sent to the Poller
// subscribe endpoint.
func DecodeSubscribeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			city  string
			state string
			err   error
		)
		qp := r.URL.Query()
		city = qp.Get("city")
		if city == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("city", "query string"))
		}
		state = qp.Get("state")
		if state == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("state", "query string"))
		}
		if utf8.RuneCountInString(state) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, true))
		}
		if utf8.RuneCountInString(state) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("state", state, utf8.RuneCountInString(state), 2, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSubscribeCityAndState(city, state)

		return payload, nil
	}
}

// EncodeAddWorkerResponse returns an encoder for responses returned by the
// Poller add_worker endpoint.
func EncodeAddWorkerResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*poller.Worker)
		enc := encoder(ctx, w)
		body := NewAddWorkerResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeRemoveWorkerResponse returns an encoder for responses returned by the
// Poller remove_worker endpoint.
func EncodeRemoveWorkerResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// EncodeRemoveWorkerError returns an encoder for errors returned by the
// remove_worker Poller endpoint.
func EncodeRemoveWorkerError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "too_few":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRemoveWorkerTooFewResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusExpectationFailed)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeStatusResponse returns an encoder for responses returned by the Poller
// status endpoint.
func EncodeStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*poller.PollerStatus)
		enc := encoder(ctx, w)
		body := NewStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalPollerLocationToLocationResponseBody builds a value of type
// *LocationResponseBody from a value of type *poller.Location.
func marshalPollerLocationToLocationResponseBody(v *poller.Location) *LocationResponseBody {
	res := &LocationResponseBody{
		Lat:   v.Lat,
		Long:  v.Long,
		City:  v.City,
		State: v.State,
	}

	return res
}

// marshalPollerPeriodToPeriodResponseBody builds a value of type
// *PeriodResponseBody from a value of type *poller.Period.
func marshalPollerPeriodToPeriodResponseBody(v *poller.Period) *PeriodResponseBody {
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

// marshalPollerJobToJobResponseBody builds a value of type *JobResponseBody
// from a value of type *poller.Job.
func marshalPollerJobToJobResponseBody(v *poller.Job) *JobResponseBody {
	res := &JobResponseBody{
		Key:       v.Key,
		Payload:   v.Payload,
		CreatedAt: v.CreatedAt,
	}

	return res
}

// marshalPollerWorkerToWorkerResponseBody builds a value of type
// *WorkerResponseBody from a value of type *poller.Worker.
func marshalPollerWorkerToWorkerResponseBody(v *poller.Worker) *WorkerResponseBody {
	res := &WorkerResponseBody{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
	}
	if v.Jobs != nil {
		res.Jobs = make([]*JobResponseBody, len(v.Jobs))
		for i, val := range v.Jobs {
			res.Jobs[i] = marshalPollerJobToJobResponseBody(val)
		}
	} else {
		res.Jobs = []*JobResponseBody{}
	}

	return res
}
