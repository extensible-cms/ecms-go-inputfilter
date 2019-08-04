package ecms_go_inputfilter

type InputFilter struct {
	Inputs *map[string]*Input
	BreakOnFailure bool
}

type InputFilterResult struct {
	Result bool
	Messages map[string][]string
	validInputs map[string]*Input
	invalidInputs map[string]*Input
	validInputsAsList []*Input
	invalidInputsAsList []*Input
}
