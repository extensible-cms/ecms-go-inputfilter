package ecms_go_inputfilter

import (
	"fmt"
	"github.com/extensible-cms/ecms-go-inputfilter/test"
	ecms_validator "github.com/extensible-cms/ecms-go-validator"
	"strconv"
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
		{Name: "`Input{}` (passing case)",
			Input:                 &Input{},
			IncomingValue:         "",
			ExpectedValue:         "",
			ExpectedRawValue:      "",
			ExpectedFilteredValue: "",
			ExpectedObscuredValue: "",
			ExpectedResult:        true,
		},
		{Name: "`Input{Required}` (failing case)",
			Input: func() *Input {
				i := &Input{}
				i.Required = true
				return i
			}(),
			IncomingValue:         nil,
			ExpectedValue:         nil,
			ExpectedRawValue:      nil,
			ExpectedFilteredValue: nil,
			ExpectedObscuredValue: nil,
			ExpectedMessageLen:    1,
			ExpectedResult:        false,
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
		{Name: "`Input{Validators(1),Filters(1)}` (validator failing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return 99
				})
				return i
			}(),
			IncomingValue:         0,
			ExpectedValue:         0,
			ExpectedRawValue:      0,
			ExpectedFilteredValue: 0,
			ExpectedObscuredValue: 0,
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),Filters(1)}` (validator passing)",
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
		{Name: "`Input{Validators(1),Filters(1),Obscurer}` (validator(s) passing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return "00000" + x.(string)
				})
				i.Obscurer = func(x interface{}) interface{} {
					return ecms_validator.ObscurateLeft(5, x.(string))
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
		{Name: "`Input{Validators(1),Filters(1),Obscurer}` (with validator failing)",
			Input: func() *Input {
				i := &Input{}
				i.Validators = append(i.Validators, test.Validators[test.Last4Social])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return "00000" + x.(string)
				})
				i.Obscurer = func(x interface{}) interface{} {
					return ecms_validator.ObscurateLeft(5, x.(string))
				}
				return i
			}(),
			IncomingValue:         "321",
			ExpectedValue:         "321",
			ExpectedRawValue:      "321",
			ExpectedFilteredValue: "321",
			ExpectedObscuredValue: "321",
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),Filters(1),Obscurer,Required}` (all passing)",
			Input: func() *Input {
				i := &Input{}
				i.Required = true
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return 99
				})
				i.Obscurer = func(x interface{}) interface{} {
					return "*" + strconv.Itoa(x.(int))[1:]
				}
				return i
			}(),
			IncomingValue:         1,
			ExpectedValue:         99,
			ExpectedRawValue:      1,
			ExpectedFilteredValue: 99,
			ExpectedObscuredValue: "*9",
			ExpectedResult:        true,
			ExpectedMessageLen:    0,
		},
		{Name: "`Input{Validators(1),Filters(1),Obscurer,Required}` (all passing)",
			Input: func() *Input {
				i := &Input{}
				i.Required = true
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return "00000" + x.(string)
				})
				i.Obscurer = func(x interface{}) interface{} {
					return ecms_validator.ObscurateLeft(5, x.(string))
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
		{Name: "`Input{Validators(1),Filters(1),Obscurer,Required}` (validators failing)",
			Input: func() *Input {
				i := &Input{}
				i.Required = true
				i.Validators = append(i.Validators, test.Validators[test.Last4Social])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return 99
				})
				i.Obscurer = func(x interface{}) interface{} {
					return "*" + strconv.Itoa(x.(int))[1:]
				}
				return i
			}(),
			IncomingValue:         "1",
			ExpectedValue:         "1",
			ExpectedRawValue:      "1",
			ExpectedFilteredValue: "1",
			ExpectedObscuredValue: "1",
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),Filters(1),Obscurer,Required}` (validators failing)",
			Input: func() *Input {
				i := &Input{}
				i.Required = true
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Filters = append(i.Filters, func(x interface{}) interface{} {
					return "00000" + x.(string)
				})
				i.Obscurer = func(x interface{}) interface{} {
					return ecms_validator.ObscurateLeft(5, x.(string))
				}
				return i
			}(),
			IncomingValue:         "",
			ExpectedValue:         "",
			ExpectedRawValue:      "",
			ExpectedFilteredValue: "",
			ExpectedObscuredValue: "",
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),BreakOnFailure}` (validators failing)",
			Input: func() *Input {
				i := &Input{}
				i.BreakOnFailure = true
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				i.Validators = append(i.Validators, test.Validators[test.Last4Social])
				return i
			}(),
			IncomingValue:         nil,
			ExpectedValue:         nil,
			ExpectedRawValue:      nil,
			ExpectedFilteredValue: nil,
			ExpectedObscuredValue: nil,
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
		},
		{Name: "`Input{Validators(1),BreakOnFailure}` (validators passing)",
			Input: func() *Input {
				i := &Input{}
				i.BreakOnFailure = true
				i.Validators = append(i.Validators, test.Validators[test.Last4Social])
				i.Validators = append(i.Validators, test.Validators[test.NotEmptyValidator])
				return i
			}(),
			IncomingValue:         "",
			ExpectedValue:         "",
			ExpectedRawValue:      "",
			ExpectedFilteredValue: "",
			ExpectedObscuredValue: "",
			ExpectedResult:        false,
			ExpectedMessageLen:    1,
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

func TestInput_AddFilter(t *testing.T) {
	type TestCaseInputAddFilter struct {
		Name               string
		Input              *Input
		Filters            []Filter
		ExpectedFiltersLen int
	}

	identityFilter := func(x interface{}) interface{} {
		return x
	}

	for _, tc := range func() []TestCaseInputAddFilter {
		out := make([]TestCaseInputAddFilter, 0)
		rangeStr := "aeiou"
		for i, _ := range rangeStr {
			input := &Input{}
			filters := make([]Filter, 0)
			for j := 0; j < i+1; j += 1 {
				filters = append(filters, identityFilter)
			}
			out = append(out, TestCaseInputAddFilter{
				fmt.Sprintf("Input.AddFilter(%v)", i+1),
				input,
				filters,
				i + 1,
			})
		}
		return out
	}() {
		t.Run(tc.Name, func(t2 *testing.T) {
			for _, f := range tc.Filters {
				tc.Input.AddFilter(f)
			}
			test.ExpectEqual(t2, fmt.Sprintf("len(Input.Filters) === %v:", tc.ExpectedFiltersLen),
				len(tc.Input.Filters), tc.ExpectedFiltersLen)
		})
	}

	t.Run("Should not add 'nil' values", func(t2 *testing.T) {
		i:= Input{}
		i.AddFilter(nil)
		test.ExpectEqual(t2, "len(Input.Filters) === 0;", len(i.Filters), 0)
	})
}

func TestInput_AddFilters(t *testing.T) {
	type TestCaseInputAddFilter struct {
		Name               string
		Input              *Input
		Filters            []Filter
		ExpectedFiltersLen int
	}

	identityFilter := func(x interface{}) interface{} {
		return x
	}

	for _, tc := range func() []TestCaseInputAddFilter {
		out := make([]TestCaseInputAddFilter, 0)
		for i, _ := range "aeiou" {
			input := &Input{}
			filters := make([]Filter, 0)
			for j := 0; j < i+1; j += 1 {
				filters = append(filters, identityFilter)
			}
			out = append(out, TestCaseInputAddFilter{
				fmt.Sprintf("Input.AddFilter(%v)", i+1),
				input,
				filters,
				i + 1,
			})
		}
		return out
	}() {
		t.Run(tc.Name, func(t2 *testing.T) {
			tc.Input.AddFilters(tc.Filters)
			test.ExpectEqual(t2, fmt.Sprintf("len(Input.Filters) === %v:", tc.ExpectedFiltersLen),
				len(tc.Input.Filters), tc.ExpectedFiltersLen)
		})
	}

	t.Run("Should not add 'nil' values", func(t2 *testing.T) {
		i:= Input{}
		i.AddFilters(nil)
		test.ExpectEqual(t2, "len(Input.Filters) === 0;", len(i.Filters), 0)
	})
}

