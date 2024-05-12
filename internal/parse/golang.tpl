package tgbot
{{range .Types}}
// https://core.telegram.org/bots/api{{.Href}}{{if .Comment}}{{range .Comment}}
//
// {{.}}{{end}}
type {{.Name}} struct {
    {{range .Fields}}{{if .Field}}// {{.Description}}
    {{.Field}} {{.Type}} `json:"{{.Tag}},omitempty"`
    {{end}}{{end}}
}{{else}}{{end}}
{{if .Method}}
func (m *{{.Name}}) Method() string {
	return "{{.MethodName}}"
}{{end}}
{{end}}