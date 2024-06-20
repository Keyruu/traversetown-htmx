package utils

import "encoding/xml"

// <urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd http://www.google.com/schemas/sitemap-image/1.1 http://www.google.com/schemas/sitemap-image/1.1/sitemap-image.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
type Urlset struct {
	XMLName           xml.Name `xml:"urlset"`
	XmlnsXsi          string   `xml:"xmlns:xsi,attr"`
	XmlnsVideo        string   `xml:"xmlns:video,attr"`
	XmlnsXhtml        string   `xml:"xmlns:xhtml,attr"`
	XmlnsImage        string   `xml:"xmlns:image,attr"`
	XmlnsNews         string   `xml:"xmlns:news,attr"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	Urls              []Url    `xml:"url"`
}

type Url struct {
	Loc        string `xml:"loc"`
	Lastmod    string `xml:"lastmod,omitempty"`
	Changefreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

func NewUrlset() *Urlset {
	sitemap := &Urlset{
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsVideo:        "http://www.google.com/schemas/sitemap-video/1.1",
		XmlnsXhtml:        "http://www.w3.org/1999/xhtml",
		XmlnsImage:        "http://www.google.com/schemas/sitemap-image/1.1",
		XmlnsNews:         "http://www.google.com/schemas/sitemap-news/0.9",
		XsiSchemaLocation: "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd http://www.google.com/schemas/sitemap-image/1.1 http://www.google.com/schemas/sitemap-image/1.1/sitemap-image.xsd",
		Xmlns:             "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	return sitemap
}

func (s *Urlset) AddUrl(url Url) {
	s.Urls = append(s.Urls, url)
}
