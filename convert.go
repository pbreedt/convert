package convert

import (
	"fmt"
	"strconv"
)

// 2 String
func SliceByte2String(b []byte) string {
	return string(b)
}

func Float642String(f float64, format string) string {
	return fmt.Sprintf(format, f)
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}

// 2 []byte
func String2SliceByte(str string) []byte {
	return []byte(str)
}

// 2 float64
func Int2Float64(i int) float64 {
	return float64(i)
}

func String2Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// 2 int
func Float642Int(f float64) int {
	return int(f)
}

func String2Int(s string) (int, error) {
	return strconv.Atoi(s)
}
