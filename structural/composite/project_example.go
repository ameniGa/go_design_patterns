package composite

import "strings"

// TodoList is the component
// will be implemented by the leaf and the composite
type TodoList interface {
	GetHtml() string
}

// Todo is the leaf
type Todo struct {
	Text string
}

func (l *Todo) GetHtml() string {
	return l.Text
}

// Project is the composite
type Project struct {
	Title string
	Todos []TodoList
}

func (p *Project) GetHtml() string {
	html := strings.Builder{}
	html.WriteString("<h1>")
	html.WriteString(p.Title)
	html.WriteString("</h1>")
	html.WriteString("<ul>")
	for _, todo := range p.Todos {
		html.WriteString("<li>")
		html.WriteString(todo.GetHtml())
		html.WriteString("</li>")
	}
	html.WriteString("</ul>")
	return html.String()
}
