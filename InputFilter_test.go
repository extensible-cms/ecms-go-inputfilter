package ecms_go_inputfilter

import (
	"fmt"
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
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = true
				// 'subject' input is skipped due to being `not required` (see 'TestFixtures') and `nil`
				o.ValidResults = map[string]InputResult{
					"name":    NewInputResult("name", "Hello World"),
					"message": NewInputResult("name", "Some description here."),
					"email":   NewInputResult("name", "abc@abc.com"),
				}
				return o
			}(),
		},
		{
			"No required fields (should fail)",
			&ContactFormInputFilter,
			map[string]interface{}{
				"subject": "Hello World",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = false
				o.InvalidResults = map[string]InputResult{
					"name": func() InputResult {
						o := NewInputResult("name", nil)
						o.Result = false
						return o
					}(),
					"message": func() InputResult {
						o := NewInputResult("message", nil)
						o.Result = false
						return o
					}(),
					"email": func() InputResult {
						o := NewInputResult("email", nil)
						o.Result = false
						return o
					}(),
				}
				o.ValidResults = map[string]InputResult{
					"subject": NewInputResult("subject", nil),
				}
				return o
			}(),
		},
		{
			"All fields valid (should pass)",
			&ContactFormInputFilter,
			map[string]interface{}{
				"name":    "Masambula",
				"email":   "masambula@aol.com",
				"subject": "Hello World!",
				"message": "Greetings from the hither world!",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = true
				o.ValidResults = map[string]InputResult{
					"name":    NewInputResult("name", "Hello World"),
					"email":   NewInputResult("email", "masambula@aol.com"),
					"subject": NewInputResult("subject", "Hello World!"),
					"message": NewInputResult("message", "Greetings from the hither world!"),
				}
				return o
			}(),
		},
		{
			"All fields invalid (should fail)",
			&ContactFormInputFilter,
			map[string]interface{}{
				"name":    "999",
				"email":   "999",
				"subject": "",
				"message": "",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = false
				o.InvalidResults = map[string]InputResult{
					"name": func() InputResult {
						o := NewInputResult("name", "999")
						o.Result = false
						return o
					}(),
					"email": func() InputResult {
						o := NewInputResult("email", "999")
						o.Result = false
						return o
					}(),
					//"subject": ... // subject is `.Required` so should get skipped
					"message": func() InputResult {
						o := NewInputResult("message", "")
						o.Result = false
						return o
					}(),
				}
				return o
			}(),
		},
		{
			"All fields valid (`len(InputFilter{BreakOnFailure}).InvalidResults === 0`)",
			func() *InputFilter {
				o := NewInputFilter()
				o.Inputs = ContactFormInputFilter.Inputs
				o.BreakOnFailure = true
				return o
			}(),
			map[string]interface{}{
				"name":    "Masambula",
				"email":   "masambula@aol.com",
				"subject": "Hello World!",
				"message": "Greetings from the hither world!",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = true
				o.ValidResults = map[string]InputResult{
					"name":    NewInputResult("name", "Hello World"),
					"email":   NewInputResult("email", "masambula@aol.com"),
					"subject": NewInputResult("subject", "Hello World!"),
					"message": NewInputResult("message", "Greetings from the hither world!"),
				}
				return o
			}(),
		},
		{
			"All fields invalid (`len(InputFilter{BreakOnFailure}).InvalidResults === 1`)",
			func() *InputFilter {
				o := NewInputFilter()
				o.Inputs = ContactFormInputFilter.Inputs
				o.BreakOnFailure = true
				return o
			}(),
			map[string]interface{}{
				"name":    "999",
				"email":   "999",
				"subject": "",
				"message": "",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = false
				o.InvalidResults = map[string]InputResult{
					"name": func() InputResult {
						o := NewInputResult("name", "999")
						o.Result = false
						return o
					}(),
					// `BreakOnFailure` flag of InputFilter will force a return on first failing input
					//  (hence why only expecting one invalid input result)
				}
				return o
			}(),
		},
		{
			"Only required fields (`len(InputFilter{BreakOnFailure}.InvalidResults) === 0`)",
			func() *InputFilter {
				o := NewInputFilter()
				o.BreakOnFailure = true
				o.Inputs["name"] = ContactFormInputFilter.Inputs["name"]
				o.Inputs["email"] = ContactFormInputFilter.Inputs["email"]
				o.Inputs["subject"] = ContactFormInputFilter.Inputs["subject"]
				o.Inputs["message"] = ContactFormInputFilter.Inputs["message"]
				return o
			}(),
			map[string]interface{}{
				"name":    "Hello World",
				"email":   "abc@abc.com",
				"message": "Some description here.",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = true
				// 'subject' input is skipped due to being `not required` (see 'TestFixtures') and `nil`
				o.ValidResults = map[string]InputResult{
					"name":    NewInputResult("name", "Hello World"),
					"message": NewInputResult("name", "Some description here."),
					"email":   NewInputResult("name", "abc@abc.com"),
				}
				return o
			}(),
		},
		{
			"No required fields #2 (`len(InputFilter{BreakOnFailure}.InvalidResults) === 1`)",
			func() *InputFilter {
				o := NewInputFilter()
				o.BreakOnFailure = true
				o.Inputs["name"] = ContactFormInputFilter.Inputs["name"]
				o.Inputs["email"] = ContactFormInputFilter.Inputs["email"]
				o.Inputs["subject"] = ContactFormInputFilter.Inputs["subject"]
				o.Inputs["message"] = ContactFormInputFilter.Inputs["message"]
				return o
			}(),
			map[string]interface{}{
				"subject": "Hello World",
			},
			func() *InputFilterResult {
				o := NewInputFilterResult()
				o.Result = false
				o.Messages = map[string][]string{"name": []string{""}}
				o.InvalidResults = map[string]InputResult{
					"name": func() InputResult {
						o := NewInputResult("name", nil)
						o.Result = false
						return o
					}(),
					// `BreakOnFailure` flag of InputFilter will force a return on first failing input
					//  (hence why only expecting one invalid input result)
				}
				// `BreakOnFailure` causes exit on first `Input` so we never get to validate the one valid
				// input here (hence why code below is commented out).
				//o.ValidResults = map[string]InputResult{
				//	"subject": NewInputResult("subject", nil),
				//}
				return o
			}(),
		},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			resultF := tc.InputFilter.Validate(tc.Data)
			invalidResultsLen := len(resultF.InvalidResults)
			expectedInvalidResultsLen := len(tc.Expected.InvalidResults)
			messagesLen := len(resultF.Messages)
			ExpectEqual(t2, "Result == Expected Result:", resultF.Result, tc.Expected.Result)
			ExpectEqual(t2, "len(InvalidResults) == len(Expected...)", invalidResultsLen, expectedInvalidResultsLen)
			ExpectEqual(t2, "len(ValidResults) == len(Expected...)", len(resultF.ValidResults), len(tc.Expected.ValidResults))
			ExpectEqual(t2, "len(Messages) == len(InvalidResults)", invalidResultsLen, messagesLen)
			t2.Run("Inspect invalid results", func(t3 *testing.T) {
				for k, ir := range resultF.InvalidResults {
					n := fmt.Sprintf(
						"InputResult{\"%v\"}.Result === Expected.InvalidResults[\"%v\"].Result",
						k, k,
					)
					t3.Run(n, func(t4 *testing.T) {
						ExpectEqual(t4, n, ir.Result, tc.Expected.InvalidResults[k].Result)
					})
				}
			})
			t2.Run("Inspect valid results", func(t3 *testing.T) {
				for k, ir := range resultF.ValidResults {
					resultCheckName := fmt.Sprintf(
						"InputResult{\"%v\"}.Result === Expected.ValidResults[\"%v\"].Result",
						k, k,
					)
					rawValueCheckName := fmt.Sprintf(
						"InputResult{\"%v\"}.Result === Expected.ValidResults[\"%v\"].Result",
						k, k,
					)
					t3.Run(resultCheckName, func(t4 *testing.T) {
						ExpectEqual(t4, resultCheckName, ir.Result, tc.Expected.ValidResults[k].Result)
					})
					t3.Run(rawValueCheckName, func(t4 *testing.T) {
						ExpectEqual(t4, rawValueCheckName, ir.RawValue, tc.Data[k])
					})
				}
			})
			//t2.Logf("Invalid input results: %v", resultF.InvalidResults)
			//t2.Logf("Valid input results: %v", resultF.ValidResults)
		})
	}
}
