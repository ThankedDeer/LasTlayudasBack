package dto

type CreateUserRequest struct {
	RoleID    int32  `json:"roleid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"isactive"`
}

type UpdateUserRequest struct {
	UserID    int32  `json:"userid"`
	RoleID    int32  `json:"roleid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"isactive"`
}
