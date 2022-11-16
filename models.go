package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
	URL       string `json:"url"`
}

func (t *Todo) SetURL(r *http.Request) {
	scheme := "https://"
	if r.TLS == nil {
		scheme = "http://"
	}

	t.URL = fmt.Sprintf("%s%s%s/%d", scheme, r.Host, r.RequestURI, t.ID)
}
