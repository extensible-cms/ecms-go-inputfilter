package ecms_go_inputfilter

type Filter = func(interface{}) interface{}

type Validatable interface {
	Validate (value interface{}) (bool, []string)
	ValidateIO (value interface{}) (bool, []string)
}