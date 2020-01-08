package validator

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
)

type TestValid struct {
	Validator
}

func NewTestValid() IValidator {
	return &TestValid{}
}

func (v TestValid) Log(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input

	v.ErrorHandle(valid)
}

func (v TestValid) Mysql(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input

	v.ErrorHandle(valid)
}

func (v TestValid) Redis(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input

	v.ErrorHandle(valid)
}

func (v TestValid) Alarm(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input

	v.ErrorHandle(valid)
}

func (v TestValid) Mail(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input

	v.ErrorHandle(valid)
}

func init() {
	Register(TestValidator, NewTestValid)
}
