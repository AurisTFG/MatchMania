package validators

func Validate(anyStruct any) error {
	err := ValidateErrorHandler(Validator.Struct(anyStruct))

	if err == nil {
		return nil
	}

	return err
}
