package ecms_go_inputfilter

import (
	"github.com/extensible-cms/ecms-go-inputfilter/sliceof"
	ecmsValidator "github.com/extensible-cms/ecms-go-validator"
)

type Input struct {
	Name           string
	Required       bool
	Filters        []Filter
	Validators     []ecmsValidator.Validator
	BreakOnFailure bool
	Obscurer       Filter
}

type InputResult struct {
	Name          string
	Result        bool
	messages      []string
	Value         interface{}
	RawValue      interface{}
	ObscuredValue interface{}
	FilteredValue interface{}
}

func NewInputResult(name string, x interface{}) InputResult {
	return InputResult{
		Name:          name,
		Result:        false,
		messages:      nil,
		Value:         x,
		RawValue:      x,
		ObscuredValue: x,
		FilteredValue: x,
	}
}

type InputInterface interface {
	AddValidator(fn ecmsValidator.Validator)
	AddValidators(validators []ecmsValidator.Validator)
	AddFilter(fn func(interface{}) interface{})
	AddFilters(filters []func(interface{}) interface{})
	Validate(x interface{}) (bool, []string, InputResult)
}

func RunValidators(i *Input, x interface{}) (bool, []string) {
	hasValidators := i.Validators != nil && len(i.Validators) > 0

	if !i.Required && !hasValidators {
		return true, nil
	}

	if i.Required && !hasValidators {
		// @todo add noempty validator
		return false, []string{"\"" + i.Name + "\" is required.  Value received `nil`."}
	}

	vResult := true
	messageSlices := make([][]string, 0)
	for _, v := range i.Validators {
		result, messages := v(x)
		if !result {
			messageSlices = append(messageSlices, messages)
			vResult = false
		}
		if i.BreakOnFailure {
			break
		}
	}
	return vResult, sliceof.SliceOfStringConcat(messageSlices)
}

func RunFilters(i *Input, x interface{}) interface{} {
	if i.Filters == nil {
		return x
	}
	last := x
	for _, f := range i.Filters {
		last = f(last)
	}
	return last
}

func (i *Input) Validate(x interface{}) (bool, []string, InputResult) {
	iResult := NewInputResult(i.Name, x)
	vResult, messages := RunValidators(i, x)
	iResult.Value = RunFilters(i, x)
	iResult.FilteredValue = iResult.Value
	iResult.ObscuredValue = iResult.FilteredValue

	if i.Obscurer != nil {
		iResult.ObscuredValue = i.Obscurer(x)
	}

	return vResult,
		messages,
		iResult
}

func (i *Input) AddValidator(fn ecmsValidator.Validator) {
	i.Validators = append(i.Validators, fn)
}

func (i *Input) AddValidators(vs []ecmsValidator.Validator) {
	for _, v := range vs {
		i.Validators = append(i.Validators, v)
	}
}

func (i *Input) AddFilter(fn Filter) {
	i.Filters = append(i.Filters, fn)
}

func (i *Input) AddFilters(fs []Filter) {
	for _, f := range fs {
		i.Filters = append(i.Filters, f)
	}
}
