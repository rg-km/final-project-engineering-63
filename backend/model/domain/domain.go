package domain

type UserDomain struct {
	Id       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsLogin  bool   `json:"is_login"`
}

type ProductDomain struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}
