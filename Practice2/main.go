package main

import (
	"log"

	a "Site/app" // Добавление псевдонима "a" для пакета "app" из модуля сайта (Site)

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Добавление возможности работы с HTML-шаблонами
	engine := html.New("./views", ".html")

	// Запуск нового веб-сервера с заданной конфигурацией о включении шаблонизатора
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Добавление возможности отображения статических файлов (в рассматриваемом примере - CSS).
	// Пример ссылки: http://localhost:3000/index.css
	app.Static("/", "./public")

	// Работа со всеми ссылками приложения
	a.Routes(app)

	// Запуск сервера на порту :3000
	log.Fatal(app.Listen(":3000"))
}
