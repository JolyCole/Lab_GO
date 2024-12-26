package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/mul", mulHandler)
	http.HandleFunc("/div", divHandler)

	fmt.Println("Сервер запущен на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseParams(r *http.Request) (float64, float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if aStr == "" || bStr == "" {
		return 0, 0, fmt.Errorf("Параметры 'a' и 'b' обязательны")
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Не удалось преобразовать 'a' в число")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Не удалось преобразовать 'b' в число")
	}

	return a, b, nil
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a + b
	fmt.Fprintf(w, "Результат сложения: %v", result)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a - b
	fmt.Fprintf(w, "Результат вычитания: %v", result)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a * b
	fmt.Fprintf(w, "Результат умножения: %v", result)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "Деление на ноль невозможно", http.StatusBadRequest)
		return
	}
	result := a / b
	fmt.Fprintf(w, "Результат деления: %v", result)
}
