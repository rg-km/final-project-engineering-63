package web

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}

type ProductResponse struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}
