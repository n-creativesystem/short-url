// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"net/url"
	"time"
)

type MetadataType struct {
	Prev  string `json:"prev"`
	Self  string `json:"self"`
	Next  string `json:"next"`
	Count int    `json:"count"`
}

type OAuthApplication struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Domain string `json:"domain"`
}

type OAuthApplicationInput struct {
	Name string `json:"name"`
}

type OAuthApplicationType struct {
	Result   []*OAuthApplication `json:"result"`
	Metadata *MetadataType       `json:"_metadata"`
}

type URL struct {
	Key       string    `json:"key"`
	URL       url.URL   `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type URLType struct {
	Result []*URL `json:"result"`
}
