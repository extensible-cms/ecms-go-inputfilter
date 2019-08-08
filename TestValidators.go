package ecms_go_inputfilter

import (
	ecms_validator "github.com/extensible-cms/ecms-go-validator"
	"regexp"
)

var (
	Validators map[int]ecms_validator.Validator
)


const (
	IdValidatorKey = iota
	SlugValidatorKey
	AliasValidatorKey
	NameValidatorKey
	DescriptionValidatorKey
	NotEmptyValidatorKey
	Last4SocialKey
	EmailValidatorKey
	DigitValidatorKey
)

func init() {
	IdValidatorOps := ecms_validator.NewIntRangeValidatorOptions()
	IdValidatorOps.Min = 1
	IdValidatorOps.Max = 20
	IdValidator := ecms_validator.IntRangeValidator(IdValidatorOps)

	DescrLenValidatorOps := ecms_validator.NewIntRangeValidatorOptions()
	DescrLenValidatorOps.Min = 1
	DescrLenValidatorOps.Max = 2048
	DescrLenValidator := ecms_validator.IntRangeValidator(DescrLenValidatorOps)

	slugValidatorOps := ecms_validator.NewRegexValidatorOptions()
	slugValidatorOps.Pattern = regexp.MustCompile("^[a-z][a-z_\\-\\d]{1,54}$i")
	SlugValidator :=  ecms_validator.RegexValidator(slugValidatorOps)

	nameValidatorOps := ecms_validator.NewRegexValidatorOptions()
	nameValidatorOps.Pattern = regexp.MustCompile("^[a-z][a-z\\s'\"]{5,54}$i")
	NameValidator :=  ecms_validator.RegexValidator(nameValidatorOps)

	NotEmptyValidator :=  ecms_validator.NotEmptyValidator1()

	last4SocialValidatorOps :=  ecms_validator.NewRegexValidatorOptions()
	last4SocialValidatorOps.Pattern = regexp.MustCompile("^\\d{4}$")
	Last4SocialValidator :=  ecms_validator.RegexValidator(last4SocialValidatorOps)

	fakeEmailValidatorOps := ecms_validator.NewRegexValidatorOptions()
	fakeEmailValidatorOps.Pattern = regexp.MustCompile("^[^@]{1,55}@[^@]{1,55}$")
	FakeEmailValidator :=  ecms_validator.RegexValidator(fakeEmailValidatorOps)

	DigitValidator :=  ecms_validator.DigitValidator1()

	Validators = map[int]ecms_validator.Validator{
		IdValidatorKey:          IdValidator,
		SlugValidatorKey:        SlugValidator,
		AliasValidatorKey:       SlugValidator,
		NameValidatorKey:        NameValidator,     // add name validator
		DescriptionValidatorKey: DescrLenValidator, // add description/content validator
		NotEmptyValidatorKey:    NotEmptyValidator,
		Last4SocialKey:          Last4SocialValidator,
		EmailValidatorKey:       FakeEmailValidator, // over-simplified version of email validation (not for production!!!)
		DigitValidatorKey:       DigitValidator,
	}
}
