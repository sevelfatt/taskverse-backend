package middlewares

import "net/http"

func Test(next http.HandlerFunc)http.HandlerFunc{
	goNext := true
	if goNext == true {
		return next
	}
	return nil
}