package ecms_go_inputfilter

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

func (inputF *InputFilter) Validate(d map[string]interface{}) InputFilterResult {
	ir := InputFilterResult{}
	// @todo write body here
	return ir
}

func (inputF *InputFilter) AddInput(i *Input) {
	if i == nil {
		return
	}
	inputF.Inputs[i.Name] = i
}

func (inputF *InputFilter) AddInputs(inputs []*Input) {
	for _, i := range inputs {
		inputF.AddInput(i)
	}
}
