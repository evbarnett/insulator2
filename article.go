package main

type Article struct {
	Title        string            `json:"title"`
	Url          string            `json:"url"`
	TemplateHtml string            `json:"template"`
	Markdown     string            `json:"md"`
	UnixTime     uint64            `json:"unix_time"`
	Values       map[string]string `json:"values"`
}

func (article *Article) templateString() string {
	return "NIL" //TODO
}
