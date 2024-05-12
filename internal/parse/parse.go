package parse

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

type Field struct {
	Field       string
	Type        string
	Description string

	Parameter string
	Required  string

	Tag string
}

var tymap = map[string]string{
	"Integer": "int",
	"String":  "string",
	"Boolean": "bool",
	"True":    "bool",
	"Float":   "float64",

	"[]Integer": "[]int",
	"[]String":  "[]string",

	"Integer or String":   "string",
	"InputFile or String": "string",

	"InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply": "InlineKeyboardMarkup",
	"[]InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo":       "InputMediaAudio",
}

func (f *Field) conv() {
	if f.Parameter != "" {
		f.Field = f.Parameter
	}
	f.Tag = f.Field
	fs := make([]string, 0)
	for _, v := range strings.Split(f.Field, "_") {
		fs = append(fs, strings.Title(v))
	}
	f.Field = strings.Join(fs, "")
	if strings.HasPrefix(f.Type, "Array of ") {
		f.Type = strings.Replace(f.Type, "Array of ", "[]", 1)
	}
	if strings.HasPrefix(f.Type, "[]Array of ") {
		f.Type = strings.Replace(f.Type, "[]Array of ", "[]", 1)
	}
	if val, ok := tymap[f.Type]; ok {
		f.Type = val
	}
	if f.Field == "ReplyToMessage" || f.Field == "GiveawayMessage" {
		f.Type = "*" + f.Type
	}
	if f.Required != "" {
		f.Description = "Required: " + f.Required + "\n    //\n    // " + f.Description
	}
}

type Type struct {
	Name    string
	Href    string
	Fields  []*Field
	Comment []string

	Method     bool
	MethodName string
}

func (t *Type) init() {
	m, _ := regexp.Match("^[a-z]", []byte(t.Name))
	if m {
		t.Method = true
		t.MethodName = t.Name
		t.Name = strings.Title(t.Name)
	}
}

func parse(body io.Reader) {
	doc, err := html.Parse(body)
	if err != nil {
		return
	}
	num := 0
	cs := make([]*catalog, 0)
	var deal func(*html.Node)
	deal = func(n *html.Node) {
		if n.Data == "h3" {
			cs = append(cs, &catalog{title: text(n), desc: make([]string, 0), chapter: make([]chapter, 0)})
			fmt.Println("-----------------", text(n), "-----------------")
			num++
		} else if num > 0 {
			val := cs[num-1]
			val.add(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			deal(c)
		}
	}
	deal(doc)
	t, _ := template.ParseFiles("golang.tpl")
	for i, c := range cs {
		if i < 4 {
			continue
		}
		d := c.conv()
		f, _ := os.Create("../../" + d.FileName)
		t.Execute(f, d)
	}
}
