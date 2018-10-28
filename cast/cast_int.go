package cast

import (
	"reflect"
	"strconv"
	"time"
)

// PtrInt casts an interface value to a pointer int.
func PtrInt(v interface{}) *int {
	if ref := reflect.ValueOf(v); ref.Kind() != reflect.Ptr {
		val := Int(v)
		return &val
	}

	switch v := v.(type) {
	case *int:
		return v

	case *int64:
		vv := int(*v)
		return &vv

	}

	return nil
}

// PtrInt64 casts an interface value to a pointer int.
func PtrInt64(v interface{}) *int64 {
	if ref := reflect.ValueOf(v); ref.Kind() != reflect.Ptr {
		val := Int64(v)
		return &val
	}

	switch v := v.(type) {
	case *int:
		vv := int64(*v)
		return &vv

	case *int64:
		return v

	}

	return nil
}

// Int casts an interface value to a int.
func Int(v interface{}) int {
	switch v := v.(type) {
	case int:
		return v
	case int8:
		return (int)(v)
	case int16:
		return (int)(v)
	case int32:
		return (int)(v)
	case int64:
		return (int)(v)
	case uint8:
		return (int)(v)
	case uint16:
		return (int)(v)
	case uint32:
		return (int)(v)
	case uint64:
		return (int)(v)
	case string:
		n, _ := strconv.Atoi(v)
		return n
	case *string:
		if v != nil {
			n, _ := strconv.Atoi(*v)
			return n
		}
	}
	return 0
}

// Int64 casts an interface value to a int64.
func Int64(v interface{}) int64 {
	switch v := v.(type) {
	case int:
		return (int64)(v)
	case int8:
		return (int64)(v)
	case int16:
		return (int64)(v)
	case int32:
		return (int64)(v)
	case int64:
		return v
	case uint8:
		return (int64)(v)
	case uint16:
		return (int64)(v)
	case uint32:
		return (int64)(v)
	case uint64:
		return (int64)(v)
	case string:
		n, _ := strconv.ParseInt(v, 10, 64)
		return n
	case *string:
		if v != nil {
			n, _ := strconv.ParseInt(*v, 10, 64)
			return n
		}
	case time.Time:
		if v.IsZero() {
			return 0
		}
		return Int64(v.Format("20060102150405"))
	}
	return 0
}

// Uint64 casts an interface value to a uint64.
func Uint64(v interface{}) uint64 {
	switch v := v.(type) {
	case int:
		return (uint64)(v)
	case int8:
		return (uint64)(v)
	case int16:
		return (uint64)(v)
	case int32:
		return (uint64)(v)
	case int64:
		return (uint64)(v)
	case uint8:
		return (uint64)(v)
	case uint16:
		return (uint64)(v)
	case uint32:
		return (uint64)(v)
	case uint64:
		return v
	case string:
		n, _ := strconv.ParseUint(v, 10, 64)
		return n
	case *string:
		if v != nil {
			n, _ := strconv.ParseUint(*v, 10, 64)
			return n
		}
	case time.Time:
		if v.IsZero() {
			return 0
		}
		return Uint64(v.Format("20060102150405"))
	}
	return 0
}
