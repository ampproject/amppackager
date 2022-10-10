package v2

func mapValueOrNil(src map[string]string, key string) *string {
	if x, found := src[key]; found {
		return &x
	}

	return nil
}
