package services

import (
	"crypto/rand"
	"fmt"
	"strings"
)

//Func return random string in upper-case
func RandString(str string) string {
	b := make([]byte, 3)
	rand.Read(b)
	id := fmt.Sprintf(str+"%x", b)
	return strings.ToUpper(id)
}
