// Copyright © 2019 Martin Tournoij – This file is part of GoatCounter and
// published under the terms of a slightly modified EUPL v1.2 license, which can
// be found in the LICENSE file or at https://license.goatcounter.com

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebsiteTpl(t *testing.T) {
	tests := []struct {
		path, want string
	}{
		{"/", "doesn’t track users with"},
		{"/help", "I don’t see my pageviews?"},
		{"/privacy", "Screen size"},
		{"/gdpr", "consult a lawyer"},
		{"/terms", "The “services” are any software, application, product, or service"},
		{"/contact", "Public Telegram Group"},
		{"/contribute", "One-time donation"},
		{"/code", "Setting the endpoint in JavaScript"},
		{"/why", "Footnotes"},
		{"/data", "CSV format with a header"},
		{"/api", "Backend integration"},
		{"/design", "Firefox on iOS is just displayed as Safari"},
		{"/status", "uptime"},
		{"/signup", `<label for="email">Email address</label>`},
		{"/user/forgot", "Forgot domain"},
		// {"/api.html", "GoatCounter API documentation"},
		// {"/api.json", `"description": "API for GoatCounter"`},
	}

	for _, tt := range tests {
		runTest(t, handlerTest{
			name:     tt.path,
			path:     tt.path,
			router:   newWebsite,
			wantCode: 200,
			wantBody: tt.want,
		}, nil)
	}
}

func TestWebsiteSignup(t *testing.T) {
	tests := []handlerTest{
		{
			name:         "basic",
			method:       "POST",
			router:       newWebsite,
			path:         "/signup",
			body:         signupArgs{Code: "xxx", Email: "m@example.com", TuringTest: "9", Password: "coconuts"},
			wantCode:     303,
			wantFormCode: 303,
		},

		{
			name:         "no-code",
			method:       "POST",
			router:       newWebsite,
			path:         "/signup",
			body:         signupArgs{Email: "m@example.com", TuringTest: "9", Password: "coconuts"},
			wantCode:     200,
			wantBody:     "", // TODO: should return JSON
			wantFormCode: 200,
			wantFormBody: "Error: must be set, must be longer than 2 characters",
		},
	}

	for _, tt := range tests {
		runTest(t, tt, func(t *testing.T, rr *httptest.ResponseRecorder, r *http.Request) {
			// TODO: test state
		})
	}
}
