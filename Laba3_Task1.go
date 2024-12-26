package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", greetingHandler)
	fmt.Println("Сервер запущен на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем значения query-параметров
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// Проверяем, что параметры не пустые
	if name == "" || age == "" {
		http.Error(w, "Параметры 'name' и 'age' обязательны", http.StatusBadRequest)
		return
	}

	// Формируем ответ
	response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)

	// Отправляем ответ клиенту
	fmt.Fprintln(w, response)
}
