package ecms_go_inputfilter

import (
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
				ExpectEqual(t2, "`err` should not be `nil`:\n", err != nil, true)
			}
			if !tc.ExpectErr {
				ExpectEqual(t2, "`err` should be `nil`:\n", err, nil)
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
				ExpectEqual(t2, "`err` should not be `nil`:\n", err != nil, true)
			}
			if !tc.ExpectErr {
				ExpectEqual(t2, "`err` should be `nil`:\n", err, nil)
			}
		})
	}
}

func TestInputFilter_Validate(t *testing.T) {
	type TestCaseInputFilterValidate struct {
		Name        string
		InputFilter *InputFilter
		Data        map[string]interface{}
		Expected    *InputFilterResult
	}

	for _, tc := range []TestCaseInputFilterValidate{
		{
			"Only required fields (should pass)",
			&ContactFormInputFilter,
			map[string]interface{}{
				"name":    "Hello World",
				"email":   "abc@abc.com",
				"message": "Some description here.",
			},
			NewInputFilterResult(),
		},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			resultF := tc.InputFilter.Validate(tc.Data)
			ExpectEqual(t2, "Result:", resultF.Result, tc.Expected.Result)
			t2.Logf("%v", resultF.InvalidResults["name"])
		})
	}
}
