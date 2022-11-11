package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Функция рендеринга (отрисовки) главной страницы с использованием главного шаблона - "layout/main",
// пустой fiber.Map{} передаётся из-за требования функции c.Render к передаче 3-ёх параметров
func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layout/main")
}

// Функция рендеринга страницы создания животного с использованием главного шаблона - "layout/main",
// пустой fiber.Map{} передаётся из-за требования функции c.Render к передаче 3-ёх параметров
func createTechnicsGet(c *fiber.Ctx) error {
	return c.Render("create", fiber.Map{}, "layout/main")
}

// Функция создания новой техники и возвращение соответствующего статуса
func createTechnics(c *fiber.Ctx) error {
	// Объявление нового элемента Technics
	var item Technics

	// Запись в переменную item значений, полученных с клиента, возвращение статуса ошибки в случае, если она есть
	if err := c.BodyParser(&item); err != nil {
		return c.SendStatus(501)
	}

	// Генерация нового уникального id для техники
	item.Id = uuid.New().String()

	// Запись данных в JSON-файл, возвращение статуса ошибки в случае, если она есть
	if err := writeDataJSON(item); err != nil {
		return c.SendStatus(501)
	}

	// Возвращение статуса 200 об успешном выполнении запроса
	return c.SendStatus(200)
}

// Функция чтения списка техники
func readTechnics(c *fiber.Ctx) error {
	// Задание id = "all" по умолчанию
	id := "all"

	// Определение длины параметра id, переданного с клиента. Если id есть (длина больше 0), то обновление переменной id
	if len(c.Params("id")) != 0 {
		id = c.Params("id")
	}

	// Чтение данных из JSON-файла
	data := readDataJSON(id)

	// Рендеринг страницы с передачей в неё данных (data) о технике и использованием главного шаблона - "layout/main"
	return c.Render("read", fiber.Map{
		"Data": data,
	}, "layout/main")
}

// Функция обновления техники
func updateTechnics(c *fiber.Ctx) error {
	// Получение данных о технике по конкретному id
	item := readDataJSON(c.Params("id"))

	// Запись в переменную item[0] значений, полученных с клиента, возвращение статуса ошибки в случае, если она есть.
	// Передаётся именно item[0], т.к. из readDataJSON возвращается слайс
	if err := c.BodyParser(&item[0]); err != nil {
		return c.SendStatus(501)
	}

	// Обновление данных у определённой техники, возвращение статуса ошибки в случае, если она есть
	if err := updateDataJSON(item[0]); err != nil {
		return c.SendStatus(501)
	}

	// Возвращение статуса 200 об успешном выполнении запроса
	return c.SendStatus(200)
}

// Функция вывода данных о технике, которое нужно обновить
func updateTechnicsGet(c *fiber.Ctx) error {
	// Получение данных о технике по конкретному id
	data := readDataJSON(c.Params("id"))

	// Рендеринг страницы с передачей в неё данных (data) о технике и использованием главного шаблона - "layout/main"
	return c.Render("update", fiber.Map{
		"Data": data[0],
	}, "layout/main")
}

// Функция удаления техники
func deleteTechnics(c *fiber.Ctx) error {
	// Получение данных о всей технике
	data := readDataJSON("all")

	// Удаление из полученного слайса данных о той технике, которая была передана клиенту
	data = removeTechnicsFromSlice(data, c.Params("id"))

	// Преобразование слайса из структур Technics в набор байт, возвращение статуса ошибки в случае, если она есть
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return c.SendStatus(501)
	}

	// Запись данных в исходный файл data.json с правами доступа 0644, возвращение статуса ошибки в случае, если она есть
	if err := ioutil.WriteFile("./data.json", dataBytes, 0644); err != nil {
		return c.SendStatus(501)
	}

	// Возвращение статуса 200 об успешном выполнении запроса
	return c.SendStatus(200)
}
