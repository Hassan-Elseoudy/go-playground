# go-playground

`go run` command compiles the code into a temporary directory, and then deletes it after the program finished.
`go build` creates executable file, that can be used to be run later.

## The Zero Value
Go has a concept of zero value. When a variable is declared, it is assigned a zero value. The zero value is:
- 0 for numeric types
- false for the boolean type
- "" for strings
- nil for interfaces, slices, channels, maps, pointers and functions

## Capacity and Length
- Capacity is the total number of elements that can be stored in the slice.
- Length is the number of elements that are currently stored in the slice.
