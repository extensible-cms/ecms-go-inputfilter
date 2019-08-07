package ecms_go_inputfilter

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
	NotEmptyValidator
	Last4Social
	EmailValidator
	DigitValidator
)

func init() {
	idValidatorOps := ecms_validator.NewIntRangeValidatorOptions()
	idValidatorOps.Min = 1
	idValidatorOps.Max = 20

	descrLenValidatorOps := ecms_validator.NewIntRangeValidatorOptions()
	descrLenValidatorOps.Min = 1
	descrLenValidatorOps.Max = 2048
	descrLenValidator := ecms_validator.IntRangeValidator(descrLenValidatorOps)

	slugValidatorOps := ecms_validator.NewRegexValidatorOptions()
	slugValidatorOps.Pattern = regexp.MustCompile("^[a-z][a-z_\\-\\d]{1,54}$")
	slugValidator := ecms_validator.RegexValidator(slugValidatorOps)
	nameValidatorOps := ecms_validator.NewRegexValidatorOptions()
	nameValidatorOps.Pattern = regexp.MustCompile("^[a-z][a-z_\\s'\"]{1,54}$")
	nameValidator := ecms_validator.RegexValidator(nameValidatorOps)

	notEmptyValidator := ecms_validator.NotEmptyValidator1()

	last4SocialValidatorOps := ecms_validator.NewRegexValidatorOptions()
	last4SocialValidatorOps.Pattern = regexp.MustCompile("^\\d{4}$")
	last4SocialValidator := ecms_validator.RegexValidator(last4SocialValidatorOps)

	fakeEmailValidatorOps := ecms_validator.NewRegexValidatorOptions()
	fakeEmailValidatorOps.Pattern = regexp.MustCompile("^[^\\@]{1,89}\\@[^\\@]{1,89}$")
	fakeEmailValidator := ecms_validator.RegexValidator(fakeEmailValidatorOps)

	digitValidator := ecms_validator.DigitValidator1()

	Validators = map[int]ecms_validator.Validator{
		IdValidator:          ecms_validator.IntRangeValidator(idValidatorOps),
		SlugValidator:        slugValidator,
		AliasValidator:       slugValidator,
		NameValidator:        nameValidator, // add name validator
		DescriptionValidator: descrLenValidator, // add description/content validator
		NotEmptyValidator:    notEmptyValidator,
		Last4Social:          last4SocialValidator,
		EmailValidator:       fakeEmailValidator, // over-simplified version of email validation (not for production!!!)
		DigitValidator:       digitValidator,
	}
}
