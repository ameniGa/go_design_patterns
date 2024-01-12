package template

import "fmt"

type MessageRetriever interface {
	Message() string
}

type UserImpl struct {
}

func (ui *UserImpl) Message() string {
	return ""
}

type Template interface {
	first() string
	third() string
	// ExecuteAlgorithm is the template method
	ExecuteAlgorithm(retriever MessageRetriever) string
}

type TemplateImpl struct {
}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(retriever MessageRetriever) string {
	return fmt.Sprintf("%s %s %s", t.first(), retriever.Message(), t.third())
}
