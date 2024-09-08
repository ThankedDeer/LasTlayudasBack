package app

import (
	"context"
	"errors"
	"fmt"
	db "github/thankeddeer/lastlayudas/db/sqlc"
	"github/thankeddeer/lastlayudas/internal/domain/dto"

	"golang.org/x/crypto/bcrypt"
)

type UserApp struct {
	store *db.Store
}

func NewUserApp(store *db.Store) UserApp {
	return UserApp{
		store: store,
	}
}

func (u *UserApp) CreateUser(data dto.CreateUserRequest) (*db.Users, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	data.Password = string(hashedPassword)

	user, err := u.store.CreateUser(context.Background(), db.CreateUserParams(data))
	if err != nil {
		return nil, err
	}

	userRole, err := u.store.CreateUserRole(context.Background(), db.CreateUserRoleParams{
		UserID: user.UserID,
		RoleID: 1,
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(userRole)

	return &user, nil
}

func (u *UserApp) GetUsers() ([]db.GetUsersWithRolesRow, error) {

	users, err := u.store.GetUsersWithRoles(context.Background())
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no se encontraron usuarios")
	}

	return users, nil
}

func (u *UserApp) UpdateUser(data db.UpdatUserParams) error {

	arg := db.UpdatUserParams{
		UserID:    data.UserID,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Password:  data.Password,
		Email:     data.Email,
	}

	err := u.store.UpdatUser(context.Background(), arg)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserApp) CreateTestimonial(data dto.CreateTestimonial) (*db.Testimonials, error) {

	arg := db.CreateTestimonialParams{
		Title:       data.Title,
		Testimonial: data.Testimonial,
		UserID:      5,
	}
	testimonial, err := u.store.CreateTestimonial(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return &testimonial, nil
}

func (u *UserApp) GetTestimonial() ([]db.Testimonials, error) {

	testimonial, err := u.store.GetTestimonials(context.Background())
	if err != nil {
		return nil, err
	}
	if len(testimonial) == 0 {
		return nil, errors.New("no existen testimonios")
	}

	return testimonial, nil
}
