package xconverts

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func ConvertToInt8(i interface{}) (int8, error) {
	b, err := ConvertToInt64(i)
	return int8(b), err
}

func ConvertToInt16(i interface{}) (int16, error) {
	b, err := ConvertToInt64(i)
	return int16(b), err
}

func ConvertToInt(i interface{}) (int, error) {
	b, err := ConvertToInt64(i)
	return int(b), err
}

func ConvertToInt32(i interface{}) (int32, error) {
	b, err := ConvertToInt64(i)
	return int32(b), err
}

func ConvertToInt64(i interface{}) (int64, error) {
	if b, ok := IsInt(i); ok {
		return b, nil
	}
	if b, ok := IsUint(i); ok {
		return int64(b), nil
	}
	switch b := i.(type) {
	case string:
		return ParseInt64(b)
	case *string:
		return ParseInt64(*b)
	case []byte:
		return ParseInt64(string(b))
	case *[]byte:
		return ParseInt64(string(*b))
	case bool:
		if b {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrConvertFail
	}
}

func ConvertToUint8(i interface{}) (uint8, error) {
	b, err := ConvertToUint64(i)
	return uint8(b), err
}

func ConvertToUint16(i interface{}) (uint16, error) {
	b, err := ConvertToUint64(i)
	return uint16(b), err
}

func ConvertToUint(i interface{}) (uint, error) {
	b, err := ConvertToUint64(i)
	return uint(b), err
}

func ConvertToUint32(i interface{}) (uint32, error) {
	b, err := ConvertToUint64(i)
	return uint32(b), err
}

func ConvertToUint64(i interface{}) (uint64, error) {
	if b, ok := IsInt(i); ok {
		return uint64(b), nil
	}
	if b, ok := IsUint(i); ok {
		return b, nil
	}
	switch b := i.(type) {
	case string:
		return ParseUint64(b)
	case *string:
		return ParseUint64(*b)
	case []byte:
		return ParseUint64(string(b))
	case *[]byte:
		return ParseUint64(string(*b))
	case bool:
		if b {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrConvertFail
	}
}

func ConvertToFloat32(i interface{}) (float32, error) {
	if b, ok := IsFloat32(i); ok {
		return b, nil
	}
	if b, ok := IsFloat64(i); ok {
		return float32(b), nil
	}
	if b, ok := IsInt(i); ok {
		return float32(b), nil
	}
	if b, ok := IsUint(i); ok {
		return float32(b), nil
	}
	switch b := i.(type) {
	case string:
		return ParseFloat32(b)
	case *string:
		return ParseFloat32(*b)
	case []byte:
		return ParseFloat32(string(b))
	case *[]byte:
		return ParseFloat32(string(*b))
	case bool:
		if b {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrConvertFail
	}
}

func ConvertToFloat64(i interface{}) (float64, error) {
	if b, ok := IsFloat32(i); ok {
		return float64(b), nil
	}
	if b, ok := IsFloat64(i); ok {
		return b, nil
	}
	if b, ok := IsInt(i); ok {
		return float64(b), nil
	}
	if b, ok := IsUint(i); ok {
		return float64(b), nil
	}
	switch b := i.(type) {
	case string:
		return ParseFloat64(b)
	case *string:
		return ParseFloat64(*b)
	case []byte:
		return ParseFloat64(string(b))
	case *[]byte:
		return ParseFloat64(string(*b))
	case bool:
		if b {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrConvertFail
	}
}

func ConvertToStr(i interface{}) (string, error) {
	switch b := i.(type) {
	case int8, int16, int, int32, int64:
		return fmt.Sprintf("%d", b), nil
	case uint8, uint16, uint, uint32, uint64:
		return fmt.Sprintf("%d", b), nil
	case string:
		return b, nil
	case *string:
		return *b, nil
	case []byte:
		return string(b), nil
	case *[]byte:
		return string(*b), nil
	case sql.RawBytes:
		return string(b), nil
	case *sql.RawBytes:
		return string(*b), nil
	case bool:
		if b {
			return "true", nil
		}
		return "false", nil
	default:
		return "", ErrConvertFail
	}
}

func ConvertToBytes(i interface{}) ([]byte, error) {
	switch b := i.(type) {
	case string:
		return []byte(b), nil
	case *string:
		return []byte(*b), nil
	case []byte:
		return b, nil
	case *[]byte:
		return *b, nil
	case sql.RawBytes:
		return b, nil
	case *sql.RawBytes:
		return *b, nil
	default:
		return nil, ErrConvertFail
	}
}

func ConvertToTime(i interface{}, layout string) (time.Time, error) {
	switch b := i.(type) {
	case string:
		return ParseTime(b, layout)
	case *string:
		return ParseTime(*b, layout)
	case []byte:
		return ParseBytesToTime(b, layout)
	case *[]byte:
		return ParseBytesToTime(*b, layout)
	case sql.RawBytes:
		return ParseBytesToTime(b, layout)
	case *sql.RawBytes:
		return ParseBytesToTime(*b, layout)
	default:
		return time.Time{}, ErrConvertFail
	}
}

func ConvertToBytesSlice(i interface{}) ([][]byte, error) {
	var s [][]byte
	switch b := i.(type) {
	case [][]byte:
		return b, nil
	case *[][]byte:
		return *b, nil
	case []string:
		s = make([][]byte, 0, len(b))
		for i := 0; i < len(b); i++ {
			s = append(s, []byte(b[i]))
		}
		return s, nil
	case *[]string:
		return ConvertToBytesSlice(*b)
	case []sql.RawBytes:
		s = make([][]byte, 0, len(b))
		for i := 0; i < len(b); i++ {
			s = append(s, b[i])
		}
		return s, nil
	case *sql.RawBytes:
		return ConvertToBytesSlice(*b)
	default:
		return nil, ErrConvertFail
	}
}

func ConvertToStrs(i interface{}) ([]string, error) {
	var s []string
	switch b := i.(type) {
	case []string:
		return b, nil
	case *[]string:
		return *b, nil
	case [][]byte:
		s = make([]string, 0, len(b))
		for i := 0; i < len(b); i++ {
			s = append(s, string(b[i]))
		}
		return s, nil
	case *[][]byte:
		return ConvertToStrs(*b)
	case []sql.RawBytes:
		s = make([]string, 0, len(b))
		for i := 0; i < len(b); i++ {
			s = append(s, string(b[i]))
		}
		return s, nil
	case *sql.RawBytes:
		return ConvertToStrs(*b)
	default:
		return nil, ErrConvertFail
	}
}

func ConvertToInt8s(i interface{}) ([]int8, error) {
	s, err := ConvertToInt64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]int8, 0, len(s))
	for _, v := range s {
		t = append(t, int8(v))
	}
	return t, nil
}

func ConvertToInt16s(i interface{}) ([]int16, error) {
	s, err := ConvertToInt64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]int16, 0, len(s))
	for _, v := range s {
		t = append(t, int16(v))
	}
	return t, nil
}

func ConvertToInts(i interface{}) ([]int, error) {
	s, err := ConvertToInt64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]int, 0, len(s))
	for _, v := range s {
		t = append(t, int(v))
	}
	return t, nil
}

func ConvertToInt32s(i interface{}) ([]int32, error) {
	s, err := ConvertToInt64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]int32, 0, len(s))
	for _, v := range s {
		t = append(t, int32(v))
	}
	return t, nil
}

func ConvertToInt64s(i interface{}) ([]int64, error) {
	var s []int64
	switch b := i.(type) {
	case []int64:
		return b, nil
	case *[]int64:
		return *b, nil
	case [][]byte:
		s = make([]int64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseInt(string(b[i]), 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[][]byte:
		return ConvertToInt64s(*b)
	case []string:
		s = make([]int64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseInt(b[i], 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[]string:
		return ConvertToInt64s(*b)
	case []sql.RawBytes:
		s = make([]int64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseInt(string(b[i]), 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[]sql.RawBytes:
		return ConvertToInt64s(*b)
	default:
		return nil, ErrConvertFail
	}
}

func ConvertToUint8s(i interface{}) ([]uint8, error) {
	s, err := ConvertToUint64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]uint8, 0, len(s))
	for _, v := range s {
		t = append(t, uint8(v))
	}
	return t, nil
}

func ConvertToUint16s(i interface{}) ([]uint16, error) {
	s, err := ConvertToUint64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]uint16, 0, len(s))
	for _, v := range s {
		t = append(t, uint16(v))
	}
	return t, nil
}

func ConvertToUints(i interface{}) ([]uint, error) {
	s, err := ConvertToUint64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]uint, 0, len(s))
	for _, v := range s {
		t = append(t, uint(v))
	}
	return t, nil
}

func ConvertToUint32s(i interface{}) ([]uint32, error) {
	s, err := ConvertToUint64s(i)
	if err != nil {
		return nil, err
	}
	t := make([]uint32, 0, len(s))
	for _, v := range s {
		t = append(t, uint32(v))
	}
	return t, nil
}

func ConvertToUint64s(i interface{}) ([]uint64, error) {
	var s []uint64
	switch b := i.(type) {
	case []uint64:
		return b, nil
	case *[]uint64:
		return *b, nil
	case [][]byte:
		s = make([]uint64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseUint(string(b[i]), 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[][]byte:
		return ConvertToUint64s(*b)
	case []string:
		s = make([]uint64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseUint(b[i], 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[]string:
		return ConvertToUint64s(*b)
	case []sql.RawBytes:
		s = make([]uint64, 0, len(b))
		for i := 0; i < len(b); i++ {
			v, err := strconv.ParseUint(string(b[i]), 10, 0)
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	case *[]sql.RawBytes:
		return ConvertToUint64s(*b)
	default:
		return nil, ErrConvertFail
	}
}
