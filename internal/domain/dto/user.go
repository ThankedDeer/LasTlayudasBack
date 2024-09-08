package dto

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type CreateTestimonial struct {
	Title       string `json:"title"`
	Testimonial string `json:"testimonial"`
	UserID      int32  `json:"user_id"`
}
