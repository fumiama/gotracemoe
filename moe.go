package trmoe

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	BASE = "https://api.trace.moe"
)

type Moe struct {
	token string
}

func NewMoe(token string) *Moe {
	return &Moe{token}
}

func (m *Moe) Me() ([]byte, error) {
	url := BASE + "/me"
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
	u := BASE + "/search"
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
			body := new(bytes.Buffer)
			mw := multipart.NewWriter(body)
			w, err1 := mw.CreateFormFile("image", path)
			if err1 != nil {
				return nil, err1
			}
			_, err1 = io.Copy(w, bytes.NewReader(d))
			if err1 != nil {
				return nil, err1
			}
			err1 = mw.Close()
			if err1 != nil {
				return nil, err1
			}
			req, err1 := http.NewRequest("POST", u, body)
			if err1 != nil {
				return nil, err1
			}
			req.Header.Set("Content-Type", mw.FormDataContentType())
			resp, err = http.DefaultClient.Do(req)
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
