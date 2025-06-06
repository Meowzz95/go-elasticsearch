// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 8.16.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newQueryRulesTestFunc(t Transport) QueryRulesTest {
	return func(body io.Reader, ruleset_id string, o ...func(*QueryRulesTestRequest)) (*Response, error) {
		var r = QueryRulesTestRequest{Body: body, RulesetID: ruleset_id}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// QueryRulesTest tests a query ruleset to identify the rules that would match input criteria
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/test-query-ruleset.html.
type QueryRulesTest func(body io.Reader, ruleset_id string, o ...func(*QueryRulesTestRequest)) (*Response, error)

// QueryRulesTestRequest configures the Query Rules Test API request.
type QueryRulesTestRequest struct {
	Body io.Reader

	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r QueryRulesTestRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "query_rules.test")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_query_rules") + 1 + len(r.RulesetID) + 1 + len("_test"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_query_rules")
	path.WriteString("/")
	path.WriteString(r.RulesetID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "ruleset_id", r.RulesetID)
	}
	path.WriteString("/")
	path.WriteString("_test")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "query_rules.test")
		if reader := instrument.RecordRequestBody(ctx, "query_rules.test", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "query_rules.test")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f QueryRulesTest) WithContext(v context.Context) func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f QueryRulesTest) WithPretty() func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f QueryRulesTest) WithHuman() func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f QueryRulesTest) WithErrorTrace() func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f QueryRulesTest) WithFilterPath(v ...string) func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f QueryRulesTest) WithHeader(h map[string]string) func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f QueryRulesTest) WithOpaqueID(s string) func(*QueryRulesTestRequest) {
	return func(r *QueryRulesTestRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
