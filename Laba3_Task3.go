package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestData struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", countHandler)
	fmt.Println("Сервер запущен на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только метод POST поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка при чтении тела запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Парсинг JSON
	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Ошибка при парсинге JSON", http.StatusBadRequest)
		return
	}

	// Подсчёт вхождений символов
	charCount := make(map[string]int)
	for _, char := range requestData.Text {
		charCount[string(char)]++
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправка результата в формате JSON
	jsonResponse, err := json.Marshal(charCount)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON ответа", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
