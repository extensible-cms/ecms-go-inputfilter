package ecms_go_inputfilter

import (
	"errors"
	"fmt"
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
	Result         bool
	Messages       map[string][]string
	ValidResults   map[string]InputResult
	InvalidResults map[string]InputResult
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

	// Validate inputs
	for _, i := range inputF.Inputs {
		rslt, msgs, inputValueRslt := i.Validate(d[i.Name])

		if !rslt {
			vResult = false
			invalidResults[i.Name] = inputValueRslt
			messages[i.Name] = msgs
		}
		if rslt {
			validResults[i.Name] = inputValueRslt
		}
	}

	ir.Result = vResult
	ir.InvalidResults = invalidResults
	ir.ValidResults = validResults

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
