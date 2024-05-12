package parse

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

func text(n *html.Node) string {
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(n)
	return buf.String()
}

func comment(source []string) []string {
	// desc := make([]string, 0)
	// for _, v := range source {
	// 	text := strings.TrimSpace(v)
	// 	if text == "" {
	// 		continue
	// 	}
	// 	desc = append(desc, text)
	// }
	return source
}

type Source struct {
	FileName string
	Types    []Type
	Comment  []string
}

type catalog struct {
	title string
	desc  []string

	chapter []chapter
}

func (c *catalog) conv() Source {
	types := make([]Type, 0)
	for _, v := range c.chapter {
		conv := v.conv()
		types = append(types, conv)
	}
	name := strings.ReplaceAll(c.title, " ", "_")
	name = strings.ToLower(name) + ".go"
	return Source{
		FileName: name,
		Types:    types,
		Comment:  comment(c.desc),
	}
}
func (c *catalog) add(n *html.Node) {
	if n.Data == "h4" {
		fmt.Println("add", text(n))
		c.chapter = append(c.chapter, chapter{
			title: text(n),
			desc:  make([]string, 0),
			th:    make([]string, 0),
			td:    make([]string, 0),
		})
		return
	}
	num := len(c.chapter)
	if num > 0 {
		c.chapter[num-1].add(n)
		return
	}
	c.desc = append(c.desc, text(n))
}

type chapter struct {
	href  string
	title string
	desc  []string
	th    []string
	td    []string
}

func (c *chapter) add(n *html.Node) {
	if c.href == "" && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				c.href = attr.Val
				return
			}
		}
	}
	switch n.Data {
	case "p":
		c.desc = append(c.desc, text(n))
	case "th":
		c.th = append(c.th, text(n))
	case "td":
		c.td = append(c.td, text(n))
	case "li":
		c.desc = append(c.desc, "  - "+text(n))
	}
}

func (c *chapter) conv() Type {
	var rt reflect.Value
	fields := make([]*Field, 0)
	num := len(c.th)
	for i, v := range c.td {
		index := i % num
		if index == 0 {
			filed := &Field{}
			fields = append(fields, filed)
			rt = reflect.ValueOf(filed).Elem()
		}
		f := rt.FieldByName(c.th[index])
		if !f.IsValid() {
			continue
		}
		f.SetString(v)
	}
	result := Type{
		Name:    strings.ReplaceAll(c.title, " ", "_"),
		Href:    c.href,
		Fields:  fields,
		Comment: comment(c.desc),
	}
	for _, v := range fields {
		v.conv()
	}
	result.init()
	return result
}
