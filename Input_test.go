package ecms_go_inputfilter

import (
	"github.com/extensible-cms/ecms-go-inputfilter/test"
	"testing"
)

func TestInput_Validate(t *testing.T) {
	type TestCaseInputValidate struct {
		Name                  string
		Input                 *Input
		IncomingValue         interface{}
		ExpectedValue         interface{}
		ExpectedRawValue      interface{}
		ExpectedFilteredValue interface{}
		ExpectedObscuredValue interface{}
		ExpectedResult        bool
		ExpectedMessageLen    int
	}

	for _, tc := range []TestCaseInputValidate{
		{Name: "`Input{}` (passing)",
			Input:                 &Input{},
			IncomingValue:         "",
			ExpectedValue:         "",
			ExpectedRawValue:      "",
			ExpectedFilteredValue: "",
			ExpectedObscuredValue: "",
			ExpectedResult:        true,
		},
		{Name: "`Input{Validators(1)}` (validator passing)",
			Input: func() *Input {
				i := &Input{}
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
		{Name: "`Input{Validators(2)}` (validators failing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.IdValidator])
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				return i
			}(),
			IncomingValue:         0,
			ExpectedValue:         0,
			ExpectedRawValue:      0,
			ExpectedFilteredValue: 0,
			ExpectedObscuredValue: 0,
			ExpectedResult:        false,
			ExpectedMessageLen:    2, // both validators should fail
		},
		{Name: "`Input{Validators(1),Filters(1)}` (validator failing, filter passing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return 99
				})
				return i
			}(),
			IncomingValue:         0,
			ExpectedValue:         99,
			ExpectedRawValue:      0,
			ExpectedFilteredValue: 99,
			ExpectedObscuredValue: 99,
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),Filters(1)}` (validator passing, filter passing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return 99
				})
				return i
			}(),
			IncomingValue:         1,
			ExpectedValue:         99,
			ExpectedRawValue:      1,
			ExpectedFilteredValue: 99,
			ExpectedObscuredValue: 99,
			ExpectedResult:        true,
			ExpectedMessageLen:    0,
		},
		{Name: "`Input{Validators(1),Filters(1)}` (validator passing, filter passing, obscurer passing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return "00000" + x.(string)
				})
				i.Obscurer = func(x interface{}) interface{} {
					return "*****" + x.(string)[5:]
				}
				return i
			}(),
			IncomingValue:         "4321",
			ExpectedValue:         "000004321",
			ExpectedRawValue:      "4321",
			ExpectedFilteredValue: "000004321",
			ExpectedObscuredValue: "*****4321",
			ExpectedResult:        true,
			ExpectedMessageLen:    0,
		},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			result, messages, inputResult := tc.Input.Validate(tc.IncomingValue)
			test.ExpectEqual(t2, "Result:", result, tc.ExpectedResult)
			test.ExpectEqual(t2, "len(Messages):", len(messages), tc.ExpectedMessageLen)
			test.ExpectEqual(t2, "Value:", inputResult.Value, tc.ExpectedValue)
			test.ExpectEqual(t2, "RawValue:", inputResult.RawValue, tc.ExpectedRawValue)
			test.ExpectEqual(t2, "FilteredValue:", inputResult.FilteredValue, tc.ExpectedFilteredValue)
			test.ExpectEqual(t2, "ObscuredValue:", inputResult.ObscuredValue, tc.ExpectedObscuredValue)
		})
	}
}
