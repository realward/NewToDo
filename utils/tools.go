package utils

import (
	"crypto/md5"
	"regexp"
	"fmt"
	"io"
)

//加密  MD5
func Md5(str string) string {
	//STR
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func FilterEmail(str string)(ok bool){

	Pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(Pattern)
	ok = reg.MatchString(str)
	return
}


func FilterPhonenumber(str string)(ok bool){

	Pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(Pattern)
	ok = reg.MatchString(str)
	return
}