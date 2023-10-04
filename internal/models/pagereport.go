package models

import (
	"net/url"
)

type PageReport struct {
	Id                 int64
	URL                string
	ParsedURL          *url.URL
	RedirectURL        string
	Refresh            string
	StatusCode         int
	ContentType        string
	MediaType          string
	Lang               string
	Title              string
	Description        string
	Robots             string
	Noindex            bool
	Nofollow           bool
	Canonical          string
	H1                 string
	H2                 string
	Links              []Link
	ExternalLinks      []Link
	Words              int
	Hreflangs          []Hreflang
	Size               int
	Images             []Image
	Scripts            []string
	Styles             []string
	Iframes            []string
	Audios             []string
	Videos             []string
	ValidHeadings      bool
	BlockedByRobotstxt bool
	Crawled            bool
	InSitemap          bool
	InternalLinks      []InternalLink
	ValidLang          bool
}
