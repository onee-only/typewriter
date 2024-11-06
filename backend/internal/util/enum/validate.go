package enum

type Validator interface {
	Valid() error
}

func Validate(v any) error {
	if validator, ok := v.(Validator); ok {
		//nolint: wrapcheck
		return validator.Valid()
	}

	return nil
}
