package ecms_go_inputfilter

import (
	"github.com/extensible-cms/ecms-go-inputfilter/test"
	"testing"
)

func TestInput_Validate(t *testing.T) {
	type TestCaseInputValidate struct {
		Name                  string
		Input                 Input
		IncomingValue         interface{}
		ExpectedValue         interface{}
		ExpectedRawValue      interface{}
		ExpectedFilteredValue interface{}
		ExpectedObscuredValue interface{}
		ExpectedResult        bool
		ExpectedMessageLen    int
	}

	for _, tc := range []TestCaseInputValidate{
		{Name: "Input{}",
			IncomingValue:         "",
			ExpectedValue:         "",
			ExpectedRawValue:      "",
			ExpectedFilteredValue: "",
			ExpectedObscuredValue: "",
			ExpectedResult:        true,
		},
		{Name: "Input{Validators}",
			Input: func() Input {
				i := Input{}
				i.Validators = append(i.Validators, test.Validators[test.IdValidator])
				return i
			}(),
			IncomingValue:         20,
			ExpectedValue:         20,
			ExpectedRawValue:      20,
			ExpectedFilteredValue: 20,
			ExpectedObscuredValue: 20,
			ExpectedResult:        true,
		},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			result, messages, inputResult := tc.Input.Validate(tc.IncomingValue)
			test.ExpectEqual(t2, result, tc.ExpectedResult)
			test.ExpectEqual(t2, len(messages), tc.ExpectedMessageLen)
			test.ExpectEqual(t2, inputResult.Value, tc.ExpectedValue)
			test.ExpectEqual(t2, inputResult.RawValue, tc.ExpectedRawValue)
			test.ExpectEqual(t2, inputResult.FilteredValue, tc.ExpectedFilteredValue)
			test.ExpectEqual(t2, inputResult.ObscuredValue, tc.ExpectedObscuredValue)
		})
	}
}
