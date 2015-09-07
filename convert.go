package convert

import(
	"fmt"
	"strconv"
)

// Convert integers to string
func Int2String(val interface{}) string{
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

// Convert floats to string, parameter length specifies the number of digits behind the float point
func Float2String(val interface{}, length int) string {
	switch val.(type){
		// Floats
		case float32:
			return fmt.Sprintf("%." + Int2String(length) + "f", val)
		case float64:
			return fmt.Sprintf("%." + Int2String(length) + "f", val)
			// Type is not floats, return empty string
		default:
			return ""
	}
}

func internalString2Int(val string, bitSize int, nonPanic bool) int64 {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil{
		if nonPanic {
			return 0
		}else {
			panic(err)
		}
	}
	return int64(res)
}

// Convert string to int64, panic when the string is not of a valid int64 format.
func String2Int64(val string) int64 {
	return int64(internalString2Int(val, 64, false))
}

// Convert string to int64, return 0 if the convertion fails
func String2Int64NonPanic(val string) int64 {
	return int64(internalString2Int(val, 64, true))
}

// Convert string to int32, panic when the string is not of a valid int32 format.
func String2Int32(val string) int32 {
	return int32(internalString2Int(val, 32, false))
}

// Convert string to int32, return 0 if the convertion fails
func String2Int32NonPanic(val string) int32 {
	return int32(internalString2Int(val, 32, true))
}

// Convert string to int16, panic when the string is not of a valid int16 format.
func String2Int16(val string) int16 {
	return int16(internalString2Int(val, 16, false))
}

// Convert string to int16, return 0 if the convertion fails
func String2Int16NonPanic(val string) int16 {
	return int16(internalString2Int(val, 16, true))
}