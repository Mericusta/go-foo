package httpfoo

import (
	"fmt"
	"net/http"
)

func RequestExample(index int) {
	r, e := http.Get("http://127.0.0.1:1099/loginNotice")
	if e != nil {
		fmt.Printf("index %v occurs error: %v\n", index, e)
		return
	}
	r.Body.Close()
}
