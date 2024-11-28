package app

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

var ErrUserNotFound = errors.New("usuario no encontrado")

type UserApp struct {
	store *sqlc.Store
}

func NewUserApp(store *sqlc.Store) UserApp {
	return UserApp{
		store: store,
	}
}

// CreateUser crea un usuario con contraseña hasheada y retorna el usuario creado.
func (u *UserApp) CreateUser(data dto.CreateUserRequest) (*sqlc.User, error) {
	// Generar el hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Crear el usuario en la base de datos
	user, err := u.store.CreateUser(context.Background(), sqlc.CreateUserParams{
		RoleID:    data.RoleID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(hashedPassword),
		Column6:   data.IsActive,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers obtiene todos los usuarios ordenados por apellido y nombre.
func (u *UserApp) GetAllUsers() ([]sqlc.GetAllUsersRow, error) {
	users, err := u.store.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no se encontraron usuarios")
	}

	return users, nil
}

// GetUserByID obtiene un usuario por su ID.
func (u *UserApp) GetUserByID(userID int32) (*sqlc.User, error) {
	user, err := u.store.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail obtiene un usuario por su correo electrónico.
func (u *UserApp) GetUserByEmail(email string) (*sqlc.User, error) {
	user, err := u.store.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser actualiza la información de un usuario, incluyendo el hash de la nueva contraseña si se proporciona.
func (u *UserApp) UpdateUser(data dto.UpdateUserRequest) error {
	// Hashear la contraseña si se proporciona una nueva.
	var hashedPassword string
	if data.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		hashedPassword = string(hash)
	} else {
		// Si la contraseña no cambia, la dejamos vacía y la manejaremos como opcional en la consulta SQL.
		hashedPassword = data.Password
	}

	// Construir los parámetros para la actualización
	params := sqlc.UpdateUserParams{
		UserID:    data.UserID,
		RoleID:    data.RoleID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashedPassword, // Puede ser vacío si no se actualiza
		IsActive:  sql.NullBool{Bool: data.IsActive, Valid: true},
	}

	// Actualizar el usuario en la base de datos
	err := u.store.UpdateUser(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser elimina un usuario por su ID.
func (u *UserApp) DeleteUser(userID int32) error {
	err := u.store.DeleteUser(context.Background(), userID)
	if err != nil {
		return err
	}
	return nil
}
