package xconverts

import (
	"database/sql"
	"fmt"
	"time"
)

func FormatAny(i interface{}) (string, error) {
	if b, ok := IsInt(i); ok {
		return FormatInt64(b), nil
	}
	if b, ok := IsUint(i); ok {
		return FormatUint64(b), nil
	}
	switch b := i.(type) {
	case float32:
		return formatFloat32(b), nil
	case float64:
		return formatFloat64(b), nil
	case string:
		return b, nil
	case *string:
		return *b, nil
	case []byte:
		return string(b), nil
	case *[]byte:
		return string(*b), nil
	case time.Time:
		return formatTime(b), nil
	case *time.Time:
		return formatTime(*b), nil
	case sql.RawBytes:
		return string(b), nil
	case *sql.RawBytes:
		return string(*b), nil
	case TextMarshaler:
		text, err := b.MarshalText()
		return string(text), err
	case BinaryMarshaler:
		text, err := b.MarshalBinary()
		return string(text), err
	}
	return fmt.Sprintf("%v", i), nil
}

func FormatAnyToBytes(i interface{}) ([]byte, error) {
	if b, ok := IsInt(i); ok {
		return FormatInt64ToBytes(b), nil
	}
	if b, ok := IsUint(i); ok {
		return FormatUint64ToBytes(b), nil
	}
	switch b := i.(type) {
	case float32:
		return formatFloat32ToBytes(b), nil
	case float64:
		return formatFloat64ToBytes(b), nil
	case string:
		return []byte(b), nil
	case *string:
		return []byte(*b), nil
	case []byte:
		return b, nil
	case *[]byte:
		return *b, nil
	case sql.RawBytes:
		return []byte(b), nil
	case *sql.RawBytes:
		return []byte(*b), nil
	case TextMarshaler:
		text, err := b.MarshalText()
		return text, err
	case BinaryMarshaler:
		text, err := b.MarshalBinary()
		return text, err
	}
	return []byte(fmt.Sprintf("%v", i)), nil
}
