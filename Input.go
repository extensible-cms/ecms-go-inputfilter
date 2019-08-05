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
	Obscurer     Filter
}

type InputValueResult struct {
	Value         interface{}
	RawValue      interface{}
	ObscuredValue interface{}
	FilteredValue interface{}
}

func NewInputValueResult(x interface{}) InputValueResult {
	return InputValueResult{x, x, x, x}
}

type InputInterface interface {
	AddValidator(fn ecmsValidator.Validator)
	AddValidators(validators []ecmsValidator.Validator)
	AddFilter(fn func(interface{}) interface{})
	AddFilters(filters []func(interface{}) interface{})
	Validate(x interface{}) (bool, []string, InputValueResult)
	Filter(x interface{}) interface{}
}

func (i *Input) Validate(x interface{}) (bool, []string, InputValueResult) {
	ivResult := InputValueResult{
		Value:         x,
		RawValue:      x,
		ObscuredValue: x,
		FilteredValue: x,
	}

	if i.Validators == nil || len(i.Validators) == 0 {
		return true, nil, ivResult
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

	ivResult.FilteredValue = i.Filter(x)

	if i.Obscurer != nil {
		ivResult.ObscuredValue = i.Obscurer(x)
	}

	return vResult,
		sliceof.SliceOfStringConcat(messageSlices),
		ivResult
}

func (i *Input) Filter(x interface{}) interface{} {
	if i.Filters == nil {
		return x
	}
	last := x
	for _, f := range i.Filters {
		last = f(last)
	}
	return last
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
