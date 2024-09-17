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
