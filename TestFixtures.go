package ecms_go_inputfilter

import (
	ecms_validator "github.com/extensible-cms/ecms-go-validator"
)

var ContactFormInputFilter InputFilter

func init() {
	nameInput := NewInput("name")
	nameInput.Required = true
	nameInput.RequiredMessage = "Name is required."
	nameInput.AddValidator(Validators[NameValidator])

	emailInput := NewInput("email")
	emailInput.Required = true
	emailInput.RequiredMessage = "Email is required."
	emailInput.AddValidator(Validators[EmailValidator])

	phoneInput := NewInput("phone")
	phoneInput.AddValidator(func() ecms_validator.Validator {
		lenOps := ecms_validator.NewLengthValidatorOptions()
		lenOps.Min = 10
		lenOps.Max = 10
		return ecms_validator.LengthValidator(lenOps)
	}())
	phoneInput.AddValidator(Validators[DigitValidator])

	subjInput := NewInput("subject")
	subjInput.AddValidator(func() ecms_validator.Validator {
		lenOps := ecms_validator.NewLengthValidatorOptions()
		lenOps.Min = 3
		lenOps.Max = 55
		return ecms_validator.LengthValidator(lenOps)
	}())
	subjInput.AddValidator(Validators[DescriptionValidator])

	msgInput := NewInput("message")
	msgInput.Required = true
	msgInput.RequiredMessage = "Message is required."
	msgInput.AddValidator(Validators[DescriptionValidator])

	ContactFormInputFilter = InputFilter{
		Inputs: map[string]*Input{
			nameInput.Name:  nameInput,
			emailInput.Name: emailInput,
			subjInput.Name:  subjInput,
			msgInput.Name:   msgInput,
		},
		BreakOnFailure: false, // validate all inputs
	}
}