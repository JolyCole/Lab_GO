package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

// Функция для обработки каждого пикселя (перевод в оттенки серого)
func filter(img draw.RGBA64Image, wg *sync.WaitGroup, y int) {
	defer wg.Done()
	width := img.Bounds().Max.X
	for x := 0; x < width; x++ {
		c := img.At(x, y).(color.RGBA64)
		gray := (c.R + c.G + c.B) / 3
		img.SetRGBA64(x, y, color.RGBA64{R: gray, G: gray, B: gray, A: c.A})
	}
}

// Главная функция
func main() {
	// Открытие изображения
	file, err := os.Open("image.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Декодирование изображения
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	// Преобразуем изображение в draw.RGBA64Image для редактирования
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Ошибка преобразования изображения")
		return
	}

	// Измеряем время обработки
	start := time.Now()

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	bounds := drawImg.Bounds()
	height := bounds.Max.Y

	// Обрабатываем каждую строку изображения параллельно
	for y := 0; y < height; y++ {
		wg.Add(1)
		go filter(drawImg, &wg, y)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Измеряем время выполнения
	fmt.Println("Время обработки:", time.Since(start))

	// Сохранение обработанного изображения
	outFile, err := os.Create("output_parallel.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outFile.Close()

	// Кодируем изображение в новый файл
	err = png.Encode(outFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
	}
}
