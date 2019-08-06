# ecms-go-inputfilter
Input Filter package for validating user submitted inputs with validators (ecms-go-validator and such).

## Docs
See sources (actual docs coming later).

### Input.Validate(x interface{}) (bool, []string, InputResult) {}
Validates given `Input` element.
*Note*: Filtering and obscuring only occurs if validation success (see source).
