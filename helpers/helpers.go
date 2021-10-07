package helpers

import (
	"strconv"
)

type AnyType string

//Function to convert a float to a string
func Int64ToString(inputNum int64) string {
	return strconv.FormatInt(inputNum, 10)
}

//Function to convert a integer to a string
func IntToString(inputNum int) string {
	return strconv.Itoa(inputNum)
}

//Function to convert any type to string
func (c CustomType) String() string {
	return string(c)
}
