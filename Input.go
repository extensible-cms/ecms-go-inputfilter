package ecms_go_inputfilter

import ecmsValidator "github.com/extensible-cms/ecms-go-validator"

type Input struct {
	Name string
	Required bool
	Filters []Filter
	Validators []ecmsValidator.Validator
	BreakOnFailure bool
}

