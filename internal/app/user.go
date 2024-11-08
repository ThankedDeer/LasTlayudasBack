package app

import "github/thankeddeer/lastlayudas/internal/store/sqlc"

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"github/thankeddeer/lastlayudas/internal/domain/dto"
// 	"github/thankeddeer/lastlayudas/internal/store/sqlc"
//

// 	"golang.org/x/crypto/bcrypt"

// )

type UserApp struct {
	store *sqlc.Store
}

// func NewUserApp(store *sqlc.Store) UserApp {
// 	return UserApp{
// 		store: store,
// 	}
// }

// func (u *UserApp) CreateUser(data dto.CreateUserRequest) (*sqlc.Users, error) {

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	data.Password = string(hashedPassword)

// 	user, err := u.store.CreateUser(context.Background(), sqlc.CreateUserParams(data))
// 	if err != nil {
// 		return nil, err
// 	}

// 	userRole, err := u.store.CreateUserRole(context.Background(), sqlc.CreateUserRoleParams{
// 		UserID: user.UserID,
// 		RoleID: 1,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(userRole)

// 	return &user, nil
// }

// func (u *UserApp) GetUsers() ([]sqlc.GetUsersWithRolesRow, error) {

// 	users, err := u.store.GetUsersWithRoles(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(users) == 0 {
// 		return nil, errors.New("no se encontraron usuarios")
// 	}

// 	return users, nil
// }

// func (u *UserApp) UpdateUser(data sqlc.UpdatUserParams) error {

// 	arg := sqlc.UpdatUserParams{
// 		UserID:    data.UserID,
// 		Firstname: data.Firstname,
// 		Lastname:  data.Lastname,
// 		Password:  data.Password,
// 		Email:     data.Email,
// 	}

// 	err := u.store.UpdatUser(context.Background(), arg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *UserApp) CreateTestimonial(data dto.CreateTestimonial) (*sqlc.Testimonials, error) {

// 	arg := sqlc.CreateTestimonialParams{
// 		Title:       data.Title,
// 		Testimonial: data.Testimonial,
// 		UserID:      5,
// 	}
// 	testimonial, err := u.store.CreateTestimonial(context.Background(), arg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &testimonial, nil
// }

// func (u *UserApp) GetTestimonial() ([]sqlc.Testimonials, error) {

// 	testimonial, err := u.store.GetTestimonials(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(testimonial) == 0 {
// 		return nil, errors.New("no existen testimonios")
// 	}

// 	return testimonial, nil
// }
