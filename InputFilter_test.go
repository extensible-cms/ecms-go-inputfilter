package ecms_go_inputfilter

import (
	"github.com/extensible-cms/ecms-go-inputfilter/test"
	"testing"
)

func TestInputFilter_AddInput(t *testing.T) {
	type TestCaseInputFilterAddInput struct {
		Name      string
		InputF    *InputFilter
		Input     *Input
		ExpectErr bool
	}

	persistedInputF := NewInputFilter()

	for _, tc := range []TestCaseInputFilterAddInput{
		{"Empty `Input{}`", persistedInputF, &Input{}, true},
		{"New `Input{Name}`", persistedInputF, NewInput("foo"), false},
		{"Already existing `Input{Name}`", persistedInputF, NewInput("foo"), true},
		{"New `Input{Name}`", persistedInputF, NewInput("bar"), false},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			err := tc.InputF.AddInput(tc.Input)
			if tc.ExpectErr {
				test.ExpectEqual(t2, "`err` should not be `nil`:\n", err != nil, true)
			}
			if !tc.ExpectErr {
				test.ExpectEqual(t2, "`err` should be `nil`:\n", err, nil)
			}
		})
	}
}

func TestInputFilter_AddInputs(t *testing.T) {
	type TestCaseInputFilterAddInputs struct {
		Name      string
		InputF    *InputFilter
		Inputs    []*Input
		ExpectErr bool
	}

	persistedInputF := NewInputFilter()
	emptyInput := &Input{}
	fooInput := NewInput("foo")
	barInput := NewInput("bar")
	bar2Input := NewInput("bar2")
	bar3Input := NewInput("bar3")

	for _, tc := range []TestCaseInputFilterAddInputs{
		{"With empty inputs: `[]*Input{Input{}}`", persistedInputF, []*Input{emptyInput, fooInput}, true},
		{"With new inputs: `[]*Input{onlyNewInputs...}`", persistedInputF, []*Input{fooInput, barInput}, false},
		{"Already existing inputs: `[]*Input{existingInputs...}`", persistedInputF, []*Input{fooInput, barInput}, true},
		{"With new inputs: `[]*Input{newInputs...}`", persistedInputF, []*Input{bar2Input, bar3Input}, false},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			err := tc.InputF.AddInputs(tc.Inputs)
			if tc.ExpectErr {
				test.ExpectEqual(t2, "`err` should not be `nil`:\n", err != nil, true)
			}
			if !tc.ExpectErr {
				test.ExpectEqual(t2, "`err` should be `nil`:\n", err, nil)
			}
		})
	}
}

/*func TestInputFilter_Validate(t *testing.T) {
	nameInput := NewInput("name")
	nameInput.Required = true
	nameInput.RequiredMessage = "Name is required."
	nameInput.AddValidator(test.Validators[test.NameValidator])

	emailInput := NewInput("email")
	emailInput.Required = true
	emailInput.RequiredMessage = "Email is required."
	emailInput.AddValidator(test.Validators[test.EmailValidator])

	phoneInput := NewInput("phone")
	phoneInput.AddValidator(func() ecms_validator.Validator {
		lenOps := ecms_validator.NewLengthValidatorOptions()
		lenOps.Min = 10
		lenOps.Max = 10
		return ecms_validator.LengthValidator(lenOps)
	}())
	phoneInput.AddValidator(test.Validators[test.DigitValidator])

	subjInput := NewInput("subject")
	subjInput.AddValidator(func() ecms_validator.Validator {
		lenOps := ecms_validator.NewLengthValidatorOptions()
		lenOps.Min = 3
		lenOps.Max = 55
		return ecms_validator.LengthValidator(lenOps)
	}())
	subjInput.AddValidator(test.Validators[test.DescriptionValidator])

	msgInput := NewInput("message")
	msgInput.Required = true
	msgInput.RequiredMessage = "Message is required."
	msgInput.AddValidator(test.Validators[test.DescriptionValidator])

	contactFormInputF := InputFilter{
		Inputs: map[string]*Input{
			nameInput.Name:  nameInput,
			emailInput.Name: emailInput,
			subjInput.Name:  subjInput,
			msgInput.Name:   msgInput,
		},
		BreakOnFailure: false, // validate all inputs
	}
}*/
