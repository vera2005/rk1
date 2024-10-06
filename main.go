package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	FirstNumber  *int    `json:"first_number"`
	SecondNumber *int    `json:"second_number"`
	Operator     *string `json:"operator"`
}

type Output struct {
	Result float64 `json:"result"`
}

// Обработчик HTTP-запроса
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}

	var input Input

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	if input.FirstNumber == nil {
		w.WriteHeader(400)
		w.Write([]byte("first_number is missing"))
		return
	}
	if input.SecondNumber == nil {
		w.WriteHeader(400)
		w.Write([]byte("second_number is missing"))
		return
	}
	if input.Operator == nil {
		w.WriteHeader(400)
		w.Write([]byte("operator is missing"))
		return
	}

	var output Output

	switch *input.Operator {
	case "+":
		output.Result = float64(*input.FirstNumber) + float64(*input.SecondNumber)
	case "-":
		output.Result = float64(*input.FirstNumber) - float64(*input.SecondNumber)
	case "*":
		output.Result = float64(*input.FirstNumber) * float64(*input.SecondNumber)
	case "/":
		if *input.SecondNumber == 0 {
			w.WriteHeader(400)
			w.Write([]byte("division by zero is not allowed"))
			return
		}
		output.Result = float64(*input.FirstNumber) / float64(*input.SecondNumber)
	default:
		w.WriteHeader(400)
		w.Write([]byte("unknown operator"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	respBytes, _ := json.Marshal(output)
	w.Write(respBytes)
}

func main() {
	// Регистрируем обработчик для пути "/calculate"
	http.HandleFunc("/calculate", CalculateHandler)

	// Запускаем веб-сервер на порту 8081
	fmt.Println("starting server...")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
