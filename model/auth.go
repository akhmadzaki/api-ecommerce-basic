package model

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email" example:"test@mail.com"`
	Password  string `json:"password" binding:"required,min=8" example:"password123"`
	FirstName string `json:"first_name" binding:"required" example:"FirstName"`
	LastName  string `json:"last_name" binding:"required" example:"LastName"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@mail.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type RegisterResponse struct {
	Message string `json:"message" example:"User created"`
	UserId  int64  `json:"user_id" example:"1"`
}

type LoginResponse struct {
	Message      string `json:"message" example:"Login successful"`
	UserId       int64  `json:"user_id" example:"1"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoiY3VzdG9tZXIiLCJleHAiOjE3OTIzOTk2NzMsImlhdCI6MTc5MjM5ODc3M30.ivu-dP4TLb4x1Oyt5MefFsWwrKZYwGN2GT_vJEqCAXY"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ0b2tlbl9pZCI6IjA0MjBiM2Q5LTgzOWYtNDcyMC04Njk5LTljOWZhNmU0MjY5YyIsImV4cCI6MTc5MzAwMzU3MywiaWF0IjoxNzkyMzk4NzczfQ.3ImJmyp91EfoRXcvygvB-J1BFzUxfkRze3wyoQUJFqA"`
}
