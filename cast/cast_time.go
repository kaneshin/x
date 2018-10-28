package cast

import (
	"time"
)

// PtrTime casts an interface value to a pointer time.Time.
func PtrTime(v interface{}) *time.Time {
	switch v := v.(type) {
	case time.Time:
		return &v
	case *time.Time:
		return v
	}
	return nil
}
