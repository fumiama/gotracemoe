package trmoe

import "testing"

func TestSearch(t *testing.T) {
	m := NewMoe("https://api.trace.moe", "https://media.trace.moe/", "")
	r, err := m.Search("https://trace.moe/img/flipped-good.jpg", true, true)
	t.Log(r)
	t.Fatal(err)
}