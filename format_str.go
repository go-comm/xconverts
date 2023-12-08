package xconverts

import (
	"strconv"
	"time"
)

func FormatInt8(n int8) string {
	return strconv.FormatInt(int64(n), 10)
}

func FormatUint8(n uint8) string {
	return strconv.FormatUint(uint64(n), 10)
}

func FormatInt16(n int16) string {
	return strconv.FormatInt(int64(n), 10)
}

func FormatUint16(n uint16) string {
	return strconv.FormatUint(uint64(n), 10)
}

func FormatInt(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

func FormatUint(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

func FormatInt32(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func FormatUint32(n uint32) string {
	return strconv.FormatUint(uint64(n), 10)
}

func FormatInt64(n int64) string {
	return strconv.FormatInt(n, 10)
}

func FormatUint64(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func FormatFloat32(n float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(n), fmt, prec, 32)
}

func formatFloat32(n float32) string {
	return FormatFloat32(n, 'f', 10)
}

func FormatFloat64(n float64, fmt byte, prec int) string {
	return strconv.FormatFloat(n, fmt, prec, 64)
}

func formatFloat64(n float64) string {
	return FormatFloat64(n, 'f', 10)
}

func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

func formatTime(t time.Time) string {
	return FormatTime(t, time.RFC3339)
}
