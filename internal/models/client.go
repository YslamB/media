package models

type CreateUser struct {
	Fullname string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Fullname string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Password string `json:"password" binding:"omitempty"`
}

type BuySubscription struct {
	SubscriptionID int    `json:"subscription_id" binding:"gt=0"`
	CardType       string `json:"card_type" binding:"oneof=senagat halkbank"`
	Expiry         string `json:"expiry" binding:"alphanum,min=6,max=6"`
	Pan            string `json:"pan" binding:"alphanum,min=16,max=16"`
	Cvc            string `json:"cvc" binding:"alphanum,min=3,max=3"`
	Language       string `json:"language" binding:"oneof=ru tk en"`
}

type ConfirmBuySubscription struct {
	SubscriptionID int    `json:"subscription_id" binding:"gt=0"`
	MDOrder        string `json:"md_order" binding:"required"`
	OTP            string `json:"otp" binding:"alphanum,min=5,max=5"`
}
