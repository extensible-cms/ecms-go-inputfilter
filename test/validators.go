package test

import (
	ecms_validator "github.com/extensible-cms/ecms-go-validator"
	"regexp"
)

var Validators map[int]ecms_validator.Validator

const (
	IdValidator = iota
	SlugValidator
	AliasValidator
	NameValidator
	DescriptionValidator
)

func init() {
	idValidatorOps := ecms_validator.NewIntRangeValidatorOptions()
	idValidatorOps.Min = 1
	idValidatorOps.Max = 20

	slugValidatorOps := ecms_validator.NewRegexValidatorOptions()
	slugValidatorOps.Pattern = regexp.MustCompile("^[a-z][a-z_\\-\\d]{1,54}$")
	slugValidator := ecms_validator.RegexValidator(slugValidatorOps)

	Validators = map[int]ecms_validator.Validator{
		IdValidator:    ecms_validator.IntRangeValidator(idValidatorOps),
		SlugValidator:  slugValidator,
		AliasValidator: slugValidator,
		NameValidator:  slugValidator, // add name validator
		DescriptionValidator: slugValidator, // add description/content validator
	}
}
