package app

// Структура, описывающая поля данных каждой техники
type Technics struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Age   int    `json:"age"`
	Price int   `json:"price"`
}