package convert

import(
	"fmt"
	"strconv"
	"errors"
	"unsafe"
	"reflect"
)

// Convert integers to string, return empty string if the convertion fails
func Int2Str(val interface{}) string{
	switch val.(type){
		// Integers
		case int:
			return fmt.Sprintf("%d", val)
		case int8:
			return fmt.Sprintf("%d", val)
		case int16:
			return fmt.Sprintf("%d", val)
		case int32:
			return fmt.Sprintf("%d", val)
		case int64:
			return fmt.Sprintf("%d", val)
		case uint:
			return fmt.Sprintf("%d", val)
		case uint8:
			return fmt.Sprintf("%d", val)
		case uint16:
			return fmt.Sprintf("%d", val)
		case uint32:
			return fmt.Sprintf("%d", val)
		case uint64:
			return fmt.Sprintf("%d", val)
		// Type is not integers, return empty string
		default: 
			return ""
	}
	
}

// Convert floats to string, parameter length specifies the number of digits behind the float point, return empty string if the convertion fails
func Float2Str(val interface{}, length int) string {
	switch val.(type){
		// Floats
		case float32:
			return fmt.Sprintf("%." + Int2Str(length) + "f", val)
		case float64:
			return fmt.Sprintf("%." + Int2Str(length) + "f", val)
			// Type is not floats, return empty string
		default:
			return ""
	}
}

// Convert bool to string
func Bool2Str(val bool) string {
	if val {
		return "true"
	}
	return "false"
}

func internalStr2Int(val string, bitSize int, noPanic bool) int64 {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil{
		if noPanic {
			return 0
		}else {
			panic(err)
		}
	}
	return res
}

// Convert string to int64, panic when the string is not of a valid int64 format.
func Str2Int64(val string) int64 {
	return int64(internalStr2Int(val, 64, false))
}

// Convert string to int64, return 0 if the convertion fails
func Str2Int64NoPanic(val string) int64 {
	return int64(internalStr2Int(val, 64, true))
}

// Convert string to int32, panic when the string is not of a valid int32 format.
func Str2Int32(val string) int32 {
	return int32(internalStr2Int(val, 32, false))
}

// Convert string to int32, return 0 if the convertion fails
func Str2Int32NoPanic(val string) int32 {
	return int32(internalStr2Int(val, 32, true))
}

// Convert string to int16, panic when the string is not of a valid int16 format.
func Str2Int16(val string) int16 {
	return int16(internalStr2Int(val, 16, false))
}

// Convert string to int16, return 0 if the convertion fails
func Str2Int16NoPanic(val string) int16 {
	return int16(internalStr2Int(val, 16, true))
}

// Convert string to int8, panic when the string is not of a valid int8 format.
func Str2Int8(val string) int8 {
	return int8(internalStr2Int(val, 8, false))
}

// Convert string to int8, return 0 if the convertion fails
func Str2Int8NoPanic(val string) int8 {
	return int8(internalStr2Int(val, 8, true))
}

// Convert string to int, panic when the string is not of a valid int format.
func Str2Int(val string) int {
	return int(internalStr2Int(val, 0, false))
}

// Convert string to int, return 0 if the convertion fails
func Str2IntNoPanic(val string) int {
	return int(internalStr2Int(val, 0, true))
}

func internalStr2Uint(val string, bitSize int, noPanic bool) uint64 {
	res, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil{
		if noPanic {
			return 0
		}else {
			panic(err)
		}
	}
	return res
}

// Convert string to uint64, panic when the string is not of a valid uint64 format.
func Str2Uint64(val string) uint64 {
	return uint64(internalStr2Uint(val, 64, false))
}

// Convert string to uint64, return 0 if the convertion fails
func Str2Uint64NoPanic(val string) uint64 {
	return uint64(internalStr2Uint(val, 64, true))
}

// Convert string to uint32, panic when the string is not of a valid uint32 format.
func Str2Uint32(val string) uint32 {
	return uint32(internalStr2Uint(val, 32, false))
}

// Convert string to uint32, return 0 if the convertion fails
func Str2Uint32NoPanic(val string) uint32 {
	return uint32(internalStr2Uint(val, 32, true))
}

// Convert string to uint16, panic when the string is not of a valid uint16 format.
func Str2Uint16(val string) uint16 {
	return uint16(internalStr2Uint(val, 16, false))
}

// Convert string to uint16, return 0 if the convertion fails
func Str2Uint16NoPanic(val string) uint16 {
	return uint16(internalStr2Uint(val, 16, true))
}

// Convert string to uint8, panic when the string is not of a valid uint8 format.
func Str2Uint8(val string) uint8 {
	return uint8(internalStr2Uint(val, 8, false))
}

// Convert string to uint8, return 0 if the convertion fails
func Str2Uint8NoPanic(val string) uint8 {
	return uint8(internalStr2Uint(val, 8, true))
}

// Convert string to uint, panic when the string is not of a valid uint format.
func Str2Uint(val string) uint {
	return uint(internalStr2Uint(val, 0, false))
}

// Convert string to uint, return 0 if the convertion fails
func Str2UintNoPanic(val string) uint {
	return uint(internalStr2Uint(val, 0, true))
}

func internalStr2Float(val string, bitSize int, noPanic bool) float64 {
	res, err := strconv.ParseFloat(val, bitSize)
	if err != nil{
		if noPanic {
			return 0
		}else {
			panic(err)
		}
	}
	return res
}

// Convert string to float64, panic when the string is not of a valid float64 format.
func Str2Float64(val string) float64 {
	return float64(internalStr2Float(val, 64, false))
}

// Convert string to float32, panic when the string is not of a valid float32 format.
func Str2Float32(val string) float32 {
	return float32(internalStr2Float(val, 32, false))
}

// Convert string to bool, panic when the string is not of a valid bool format
func Str2Bool(val string) bool {
	if val == "true" || val == "True" || val == "TRUE" {
		return true
	}
	if val == "false" || val == "False" || val == "FALSE" {
		return false
	}
	panic(errors.New("could not resolve bool value:" + val))
}

// Convert bool to int, never panics
func Bool2Int(val bool) int {
	if val == true {
		return 1
	}
	if val == false {
		return 0
	}
}


// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
func Str2ByteSliceNonCopy(val string) []byte {
	pslc := (*reflect.SliceHeader)(unsafe.Pointer(&val))
	pslc.Cap = pslc.Len
	return *(*[]byte)(unsafe.Pointer(pslc))
}

// Zero-copy convert from byte slice to a string
func BytesSlice2StrNonCopy(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}
