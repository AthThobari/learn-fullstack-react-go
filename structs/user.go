package structs

// This structure is used to display user data as an API response
type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `jsonn"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

// This structure is used to receive data during the user creation process.
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

// This structure is used to receive data during the user update process.
type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password;omitempty"`
}

// This structure is used when the user perfirms the login process
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
