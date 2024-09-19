package models

type Film struct {
	Id            int    `json:"id"`
	SubCategoryId int    `json:"sub_category_id"`
	Language      string `json:"language"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Path          string `json:"path"`
	ImagePath     string `json:"image_path"`
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
