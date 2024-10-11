package models

type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type LoginForm struct {
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password" binding:"required"`
	NotificationToken string `json:"token" binding:"omitempty"`
}

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Category struct {
	ID int    `json:"category_id" binding:"omitempty"`
	Tm string `json:"tm" binding:"required"`
	Ru string `json:"ru" binding:"required"`
}

type Film struct {
	Id            int    `json:"id"`
	SubCategoryId int    `json:"sub_category_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Path          string `json:"path"`
	ImagePath     string `json:"image_path"`
	CreatedAt     string `json:"created_at"`
}

type Music struct {
	Id            int    `json:"id"`
	SubCategoryId int    `json:"sub_category_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Path          string `json:"path"`
	ImagePath     string `json:"image_path"`
}

type Book struct {
	Id            int    `json:"id"`
	SubCategoryId int    `json:"sub_category_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Path          string `json:"path"`
	ImagePath     string `json:"image_path"`
}

type Response struct {
	Error  error       `json:"error"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Message struct {
	Message ResultMessage `json:"message"`
}
type ResultMessage struct {
	Tk string `json:"tk"`
	Ru string `json:"ru"`
	En string `json:"en"`
}

var InvalidInput = Message{
	Message: ResultMessage{
		Tk: "Nädogry maglumat",
		Ru: "Неверное тело запроса",
		En: "Invalid request body",
	},
}

var NotFound = Message{
	Message: ResultMessage{
		Tk: "Maglumat Tapylmady",
		Ru: "Ничего не найдено",
		En: "Nothing found",
	},
}

var Conflict = Message{
	Message: ResultMessage{
		Tk: "Eýýäm bar",
		Ru: "уже существует",
		En: "already exists",
	},
}

var InternalServerError = Message{
	Message: ResultMessage{
		Tk: "Serwer ýalňyşlygy",
		Ru: "Ошибка сервера",
		En: "Internal server error",
	},
}

type Translate struct {
	TM *string `json:"tm"`
	RU *string `json:"ru"`
}

type Categories struct {
	ID            int              `json:"id"`
	Name          Translate        `json:"name"`
	SubCategories []map[string]any `json:"sub_categories"`
}

type ElementData struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
	Language    string `json:"language" binding:"required"`
	Status      bool   `json:"status" binding:"required"`
}
