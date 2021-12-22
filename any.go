package xconverts

func Int8(x interface{}, err error) (int8, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToInt8(x)
}

func Int16(x interface{}, err error) (int16, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToInt16(x)
}

func Int(x interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToInt(x)
}

func Int32(x interface{}, err error) (int32, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToInt32(x)
}

func Int64(x interface{}, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToInt64(x)
}

func Uint8(x interface{}, err error) (uint8, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToUint8(x)
}

func Uint16(x interface{}, err error) (uint16, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToUint16(x)
}

func Uint(x interface{}, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToUint(x)
}

func Uint32(x interface{}, err error) (uint32, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToUint32(x)
}

func Uint64(x interface{}, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToUint64(x)
}

func Float32(x interface{}, err error) (float32, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToFloat32(x)
}

func Float64(x interface{}, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return ConvertToFloat64(x)
}

func Str(x interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return ConvertToStr(x)
}

func Bytes(x interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToBytes(x)
}

func Strs(x interface{}, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToStrs(x)
}

func BytesSlice(x interface{}, err error) ([][]byte, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToBytesSlice(x)
}

func Int8s(x interface{}, err error) ([]int8, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToInt8s(x)
}

func Int16s(x interface{}, err error) ([]int16, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToInt16s(x)
}

func Ints(x interface{}, err error) ([]int, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToInts(x)
}

func Int32s(x interface{}, err error) ([]int32, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToInt32s(x)
}

func Int64s(x interface{}, err error) ([]int64, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToInt64s(x)
}

func Uint8s(x interface{}, err error) ([]uint8, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToUint8s(x)
}

func Uint16s(x interface{}, err error) ([]uint16, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToUint16s(x)
}

func Uints(x interface{}, err error) ([]uint, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToUints(x)
}

func Uint32s(x interface{}, err error) ([]uint32, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToUint32s(x)
}

func Uint64s(x interface{}, err error) ([]uint64, error) {
	if err != nil {
		return nil, err
	}
	return ConvertToUint64s(x)
}

func MustInt8(x interface{}, def int8) int8 {
	v, err := ConvertToInt8(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt16(x interface{}, def int16) int16 {
	v, err := ConvertToInt16(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt(x interface{}, def int) int {
	v, err := ConvertToInt(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt32(x interface{}, def int32) int32 {
	v, err := ConvertToInt32(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt64(x interface{}, def int64) int64 {
	v, err := ConvertToInt64(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint8(x interface{}, def uint8) uint8 {
	v, err := ConvertToUint8(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint16(x interface{}, def uint16) uint16 {
	v, err := ConvertToUint16(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint(x interface{}, def uint) uint {
	v, err := ConvertToUint(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint32(x interface{}, def uint32) uint32 {
	v, err := ConvertToUint32(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint64(x interface{}, def uint64) uint64 {
	v, err := ConvertToUint64(x)
	if err != nil {
		return def
	}
	return v
}

func MustFloat32(x interface{}, def float32) float32 {
	v, err := ConvertToFloat32(x)
	if err != nil {
		return def
	}
	return v
}

func MustFloat64(x interface{}, def float64) float64 {
	v, err := ConvertToFloat64(x)
	if err != nil {
		return def
	}
	return v
}

func MustStrs(x interface{}, def []string) []string {
	v, err := ConvertToStrs(x)
	if err != nil {
		return def
	}
	return v
}

func MustBytesSlice(x interface{}, def [][]byte) [][]byte {
	v, err := ConvertToBytesSlice(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt8s(x interface{}, def []int8) []int8 {
	v, err := ConvertToInt8s(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt16s(x interface{}, def []int16) []int16 {
	v, err := ConvertToInt16s(x)
	if err != nil {
		return def
	}
	return v
}

func MustInts(x interface{}, def []int) []int {
	v, err := ConvertToInts(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt32s(x interface{}, def []int32) []int32 {
	v, err := ConvertToInt32s(x)
	if err != nil {
		return def
	}
	return v
}

func MustInt64s(x interface{}, def []int64) []int64 {
	v, err := ConvertToInt64s(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint8s(x interface{}, def []uint8) []uint8 {
	v, err := ConvertToUint8s(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint16s(x interface{}, def []uint16) []uint16 {
	v, err := ConvertToUint16s(x)
	if err != nil {
		return def
	}
	return v
}

func MustUints(x interface{}, def []uint) []uint {
	v, err := ConvertToUints(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint32s(x interface{}, def []uint32) []uint32 {
	v, err := ConvertToUint32s(x)
	if err != nil {
		return def
	}
	return v
}

func MustUint64s(x interface{}, def []uint64) []uint64 {
	v, err := ConvertToUint64s(x)
	if err != nil {
		return def
	}
	return v
}
