# convert
A set of helper functions for data type convertion in Go programming language.

## Installation
```
go get github.com/vence722/convert
```
## Functions

- Int2Str() --- Parse any integer types into string

- Float2Str() --- Parse any float types into string

- Bool2Str() --- Parse bool type to string

- Str2Int64() --- Parse string to int64, will throw panic if the parsing fails

- Str2Int64NoPanic() --- Parse string to int64, will return 0 instead of throw panic if the parsing fails

- Str2Int32() --- Parse string to int32, will throw panic if the parsing fails

- Str2Int32NoPanic() --- Parse string to int32, will return 0 instead of throw panic if the parsing fails
