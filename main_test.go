package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestRedirection(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&url=https%3A%2F%2Fdev.case.unee-t.com%2Fcase%2F61914&email=foo@example.com").
		SetDebug(true).
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "https://dev.case.unee-t.com/case/61914", r.HeaderMap.Get("Location"))
			assert.Equal(t, http.StatusSeeOther, r.Code)
		})
}

func TestRedirectionNotUneeT(t *testing.T) {
	r := gofight.New()
	r.GET("/").
		SetQuery(gofight.H{
			"id":    "foobar-321",
			"url":   "https://unee-t.example.com",
			"email": "foobar@example.com",
		}).
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "https://unee-t.example.com is not a valid unee-t.com URL", strings.TrimSpace(r.Body.String()))
		})
}

func TestBadURL(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&url=bad&user=21&email=foo@example.com").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "bad is not a valid URL", strings.TrimSpace(r.Body.String()))
		})
}

func TestEmptyURL(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&medium=email&url=&user=21").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "url searchParam is empty", strings.TrimSpace(r.Body.String()))
		})
}

func TestMissingEmail(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&url=https://example.unee-t.com").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "Missing email parameter", strings.TrimSpace(r.Body.String()))
		})
}
