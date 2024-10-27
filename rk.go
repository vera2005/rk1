package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Output struct {
	Result string `json: result`
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("non expected method"))
		return
	}

	ourString := r.URL.Query().Get("src_string")
	if !r.URL.Query().Has("src_string") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no nedeing atribute"))
		return
	}
	if len(ourString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no data"))
		return
	}

	ourSringRune := []rune(ourString)
	ans := ""
	curCount := 0
	curElem := ourSringRune[0]
	//abbbbaaaa
	for _, elem := range ourSringRune {
		if elem == curElem {
			curCount += 1
		} else {
			ans += string(curElem)
			ans += strconv.Itoa(curCount)
			curCount = 1
			curElem = elem
		}
	}
	ans += string(curElem)
	ans += strconv.Itoa(curCount)
	var answer Output
	answer.Result = ans
	itog, _ := json.Marshal(answer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(itog)
}

func main() {
	http.HandleFunc("/encode", handler)
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("error...")
	}
}
