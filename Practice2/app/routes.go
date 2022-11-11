package app

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/", index) // Главная страница
	
	app.Get("/technics/create", createTechnicsGet) // Страница создания новой техники
	app.Post("/technics/create", createTechnics) // Добавление новоой техники

	app.Get("/technics/:id?", readTechnics) // Страница вывода данных о множестве техники, либо одной с сохранением шаблона

	app.Get("/technics/update/:id", updateTechnicsGet) // Страница вывода данных для обновления данных о технике
	app.Post("/technics/update/:id", updateTechnics) // Обновление данных о технике

	app.Get("/technics/delete/:id", deleteTechnics) // Страница, которая вызывается для удаления данных о технике
}
