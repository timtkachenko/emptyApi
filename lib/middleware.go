package lib

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"log"
)

var debugFlag = true

func Middleware(handler http.Handler) http.Handler {
	return RequestLogging(handler)
}

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if debugFlag {
			buf := GetBuffer()
			defer PutBuffer(buf)

			_, err := buf.ReadFrom(r.Body)
			r.Body.Close()
			if err != nil {
				log.Println(err)
			} else {
				val := buf.Bytes()
				log.Println("Received request [", r.URL.Path, "] with body: ", string(val))
				r.Body = ioutil.NopCloser(bytes.NewReader(val))
			}
		}
		next.ServeHTTP(w, r)
	})
}
