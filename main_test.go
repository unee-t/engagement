package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestRedirection(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&medium=email&url=https%3A%2F%2Fdev.case.unee-t.com%2Fcase%2F61914&user=21").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "https://dev.case.unee-t.com/case/61914", r.HeaderMap.Get("Location"))
			assert.Equal(t, http.StatusSeeOther, r.Code)
		})
}

func TestRedirectionTwo(t *testing.T) {
	r := gofight.New()
	r.GET("/").
		SetQuery(gofight.H{
			"id":     "foobar-321",
			"url":    "https://www.google.com/search?q=cute+fluffy+animals&source=lnms&tbm=isch",
			"medium": "email",
			"user":   "2",
		}).
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "https://www.google.com/search?q=cute+fluffy+animals&source=lnms&tbm=isch", r.HeaderMap.Get("Location"))
			assert.Equal(t, http.StatusSeeOther, r.Code)
		})
}

func TestBadURL(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&medium=email&url=bad&user=21").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestEmptyURL(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&medium=email&url=&user=21").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestMissingUser(t *testing.T) {
	r := gofight.New()
	r.GET("/?id=foobar-5311&medium=email&url=https://example.com").
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}
