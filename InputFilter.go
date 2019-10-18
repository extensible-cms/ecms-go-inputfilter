package ecms_go_inputfilter

import (
	"errors"
	"fmt"
	"github.com/extensible-cms/ecms-go-validator/is"
)

type InputFilter struct {
	Inputs         map[string]*Input
	BreakOnFailure bool
}

func NewInputFilter() *InputFilter {
	return &InputFilter{
		Inputs:         make(map[string]*Input),
		BreakOnFailure: false,
	}
}

type InputFilterResult struct {
	Result         bool                   `json:"result"`
	Messages       map[string][]string    `json:"messages"`
	ValidResults   map[string]InputResult `json:"validResults"`
	InvalidResults map[string]InputResult `json:"invalidResults"`
}

func NewInputFilterResult() *InputFilterResult {
	return &InputFilterResult{
		true,
		nil,
		nil,
		nil,
	}
}

type InputFilterInterface interface {
	Validate(data map[string]interface{})
	AddInput(i *Input)
	AddInputs(i []*Input)
}

func (inputF *InputFilter) Validate(d map[string]interface{}) InputFilterResult {
	ir := *NewInputFilterResult()
	inputFInputsLen := len(inputF.Inputs)

	if len(d) == 0 && inputFInputsLen == 0 {
		ir.Result = true
		return ir
	}

	messages := make(map[string][]string)
	validResults := make(map[string]InputResult)
	invalidResults := make(map[string]InputResult)
	vResult := true

	ir.InvalidResults = invalidResults
	ir.ValidResults = validResults

	// Validate inputs
	for _, i := range inputF.Inputs {
		if !i.Required && is.Empty(d[i.Name]) {
			continue
		}

		inputValueRslt := i.Validate(d[i.Name])

		if !inputValueRslt.Result {
			vResult = false
			invalidResults[i.Name] = inputValueRslt
			messages[i.Name] = inputValueRslt.Messages

			if inputF.BreakOnFailure {
				ir.Result = vResult
				return ir
			}
		}
		if inputValueRslt.Result {
			validResults[i.Name] = inputValueRslt
		}
	}

	ir.Result = vResult
	return ir
}

func (inputF *InputFilter) AddInput(i *Input) error {
	if i == nil {
		return errors.New("empty `Input`'s not allowed")
	}
	if len(i.Name) == 0 {
		return errors.New(
			fmt.Sprintf(
				"`Input` objects require name values.  "+
					"Recieved `Input` %v", i,
			),
		)
	}
	if inputF.Inputs[i.Name] != nil {
		return errors.New(
			fmt.Sprintf(
				"`Input` with name \"%s\" already exists in input filter",
				i.Name,
			),
		)
	}
	inputF.Inputs[i.Name] = i
	return nil
}

func (inputF *InputFilter) AddInputs(inputs []*Input) error {
	for _, i := range inputs {
		if err := inputF.AddInput(i); err != nil {
			return err
		}
	}
	return nil
}
