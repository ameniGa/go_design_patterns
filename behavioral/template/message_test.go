package template_test

import (
	"github.com/ameniGa/go_design_patterns/behavioral/template"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateImpl_ExecuteAlgorithm(t *testing.T) {
	templImpl := template.TemplateImpl{}
	myImp := mySecondStepImp{msg: "testing"}
	s := templImpl.ExecuteAlgorithm(&myImp)
	assert.Equal(t, "hello testing template", s)
}

type mySecondStepImp struct {
	msg string
}

func (msi *mySecondStepImp) Message() string {
	return msi.msg
}
