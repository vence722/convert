# convert
A set of helper functions for data type conversion in Go programming language.

## Installation
```
go get github.com/vence722/convert
```
## Functions

- Int2Str() --- Parse any integer types into string

- Float2Str() --- Parse any float types into string

- Bool2Str() --- Parse bool type to string

- Str2IntXX() --- Parse string to intXX, will throw panic if the parsing fails (IntXX can be int, int8, int16, int32, int64)

- Str2IntXXNoPanic() --- Same as above, but will return 0 instead of throw panic if the parsing fails

- Str2UintXX() --- Parse string to int32, will throw panic if the parsing fails (UintXX can be uint, uint8, uint16, uint32, uint64)

- Str2Int32NoPanic() --- Same as above, but will return 0 instead of throw panic if the parsing fails

- Str2FloatXX() --- Parse string to floatXX, will throw panic if the parsing fails (FloatXX can be float32, float64)

- Str2Bool() --- Parse string to bool, will throw panic if the parsing fails

- Str2ByteSliceNonCopy() --- Parse string to []byte without copy (the native []byte(str) will copy the bytes internally)
