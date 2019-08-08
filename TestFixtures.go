package ecms_go_inputfilter

import (
	"fmt"
	ecms_validator "github.com/extensible-cms/ecms-go-validator"
	"regexp"
)

var (
	ContactFormInputFilter InputFilter
	NameInput *Input
	EmailInput *Input
	SubjInput *Input
	MessageInput *Input
)

func init() {
	nameValidatorOps := ecms_validator.NewRegexValidatorOptions()
	nameValidatorOps.Pattern = regexp.MustCompile("^[a-zA-Z][a-zA-Z\\s'\"]{4,54}$")
	NameValidator := ecms_validator.RegexValidator(nameValidatorOps)

	nametestRslt, nametestMsgs := NameValidator("abcd")
	fmt.Printf("nametestMsgs: %v; nametest: %v", nametestMsgs, nametestRslt)

	NameInput = NewInput("name")
	NameInput.Required = true
	NameInput.RequiredMessage = "Name is required."
	NameInput.AddValidator(NameValidator)

	EmailInput = NewInput("email")
	EmailInput.Required = true
	EmailInput.RequiredMessage = "Email is required."

	fakeEmailValidatorOps := ecms_validator.NewRegexValidatorOptions()
	fakeEmailValidatorOps.Pattern = regexp.MustCompile("^[^@]{1,55}@[^@]{1,55}$")
	fakeEmailValidator := ecms_validator.RegexValidator(fakeEmailValidatorOps)

	EmailInput.AddValidator(fakeEmailValidator)

	DescrLenValidatorOps := ecms_validator.NewLengthValidatorOptions()
	DescrLenValidatorOps.Min = 1
	DescrLenValidatorOps.Max = 2048
	DescrLenValidator := ecms_validator.LengthValidator(DescrLenValidatorOps)

	SubjInput = NewInput("subject")
	SubjInput.AddValidator(func() ecms_validator.Validator {
		lenOps := ecms_validator.NewLengthValidatorOptions()
		lenOps.Min = 3
		lenOps.Max = 55
		return ecms_validator.LengthValidator(lenOps)
	}())

	MessageInput = NewInput("message")
	MessageInput.Required = true
	MessageInput.RequiredMessage = "Message is required."
	MessageInput.AddValidator(DescrLenValidator)

	ContactFormInputFilter = InputFilter{
		Inputs: map[string]*Input{
			NameInput.Name:    NameInput,
			EmailInput.Name:   EmailInput,
			SubjInput.Name:    SubjInput,
			MessageInput.Name: MessageInput,
		},
		BreakOnFailure: false, // validate all inputs
	}
}
