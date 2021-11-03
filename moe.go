package trmoe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Moe struct {
	base  string
	media string
	token string
}

func NewMoe(baseurl string, mediaurl string, token string) *Moe {
	return &Moe{baseurl, mediaurl, token}
}

func (m *Moe) Me() ([]byte, error) {
	url := m.base + "/me"
	if m.token != "" {
		url += "?key=" + m.token
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Moe) Search(path string, cutBlackBorders bool, includeAnilistInfo bool) (*Result, error) {
	u := m.base + "/search"
	isurl := strings.HasPrefix(path, "http")
	var resp *http.Response
	var err error
	if m.token != "" {
		u += "?key=" + m.token
	}
	if cutBlackBorders {
		if m.token == "" {
			u += "?"
		} else {
			u += "&"
		}
		u += "cutBorders="
	}
	if includeAnilistInfo {
		if m.token == "" && !cutBlackBorders {
			u += "?"
		} else {
			u += "&"
		}
		u += "anilistInfo="
	}
	if isurl {
		if m.token == "" && !cutBlackBorders && !includeAnilistInfo {
			u += "?"
		} else {
			u += "&"
		}
		u += "url=" + path
		resp, err = http.Get(u)
	} else {
		d, err1 := os.ReadFile(path)
		if err1 != nil {
			vals := make(url.Values)
			vals["image"] = append(vals["image"], path)
			resp, err = http.PostForm(u, vals)
		} else {
			resp, err = http.Post(u, "multipart/form-data", bytes.NewReader(d))
		}
	}
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	r := new(Result)
	json.Unmarshal(data, r)
	return r, nil
}
