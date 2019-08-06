package ecms_go_inputfilter

import (
	"errors"
	"fmt"
)

type InputFilter struct {
	Inputs         map[string]*Input
	BreakOnFailure bool
}

type InputFilterResult struct {
	Result              bool
	Messages            map[string][]string
	validInputs         map[string]*Input
	invalidInputs       map[string]*Input
	validInputsAsList   []*Input
	invalidInputsAsList []*Input
}

type InputFilterInterface interface {
	Validate(data map[string]interface{})
	AddInput(i *Input)
	AddInputs(i []*Input)
}

func NewInputFilter () *InputFilter {
	return &InputFilter{
		Inputs: make(map[string]*Input),
		BreakOnFailure: false,
	}
}

func (inputF *InputFilter) Validate(d map[string]interface{}) InputFilterResult {
	ir := InputFilterResult{}
	// @todo write body here
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
