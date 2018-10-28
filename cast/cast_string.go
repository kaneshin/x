package cast

// PtrString casts an interface value to a pointer string.
func PtrString(v interface{}) *string {
	switch v := v.(type) {
	case string:
		return &v
	case *string:
		return v
	}
	return nil
}
