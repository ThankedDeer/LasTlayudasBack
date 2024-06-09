package main

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"errors"
// 	db "github/thankeddeer/lastlayudas/db/sqlc"
// 	"log"
// 	"net/http"
// 	"unicode"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	_ "github.com/lib/pq"
// )

// type usernameRequest struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }
// type response struct {
// 	Status   int         `json:"status"`
// 	Message  string      `json:"message,omitempty"`
// 	Error    string      `json:"error,omitempty"`
// 	Data     interface{} `json:"data,omitempty"`
// 	ErrorLog error       `json:"errorLog"`
// }

// func main() {
// 	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/las_tlayudas?sslmode=disable")
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	queries := db.New(conn)

// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	userGroup := e.Group("/api")
// 	userGroup.Use(middleware.Logger())
// 	userGroup.Use(middleware.CORS())
// 	userGroup.POST("/user", func(c echo.Context) error {
// 		return CreateUser(c, queries)
// 	})
// 	userGroup.POST("/login", func(c echo.Context) error {
// 		return Login(c, queries)
// 	})
// 	userGroup.PUT("/password", func(c echo.Context) error {
// 		return updatePaswoord(c)
// 	})

// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func CreateUser(c echo.Context, queries *db.Queries) error {
// 	var jsonData usernameRequest

// 	decoder := json.NewDecoder(c.Request().Body)

// 	if err := decoder.Decode(&jsonData); err != nil {
// 		return c.JSON(http.StatusBadRequest, response{
// 			Status:   http.StatusBadRequest,
// 			Message:  "Invalid request data",
// 			ErrorLog: err,
// 		})
// 	}

// 	ctx := context.Background()

// 	existUser, err := queries.GetExistUser(ctx, jsonData.Username)

// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	if existUser != nil && existUser.UserID != 0 {
// 		return c.JSON(http.StatusBadRequest, response{
// 			Status:   http.StatusBadRequest,
// 			Message:  "Este usuario ya existe",
// 			ErrorLog: err,
// 		})
// 	}

// 	err = validatePassword(jsonData.Password)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response{
// 			Status: http.StatusBadRequest,
// 			Error:  err.Error(),
// 		})
// 	}

// 	user, err := queries.CreateUser(ctx, db.CreateUserParams{
// 		Username: jsonData.Username,
// 		Password: jsonData.Password,
// 		Email:    jsonData.Email,
// 	})

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, response{
// 			Status:  http.StatusInternalServerError,
// 			Message: "Error creating user",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, response{
// 		Status:  http.StatusOK,
// 		Message: "User created successfully",
// 		Data:    user.UserID,
// 	})
// }

// func Login(c echo.Context, queries *db.Queries) error {
// 	var loginData usernameRequest

// 	// Decodificar el cuerpo de la solicitud JSON en la estructura usernameRequest
// 	if err := json.NewDecoder(c.Request().Body).Decode(&loginData); err != nil {
// 		return c.JSON(http.StatusBadRequest, response{
// 			Status:   http.StatusBadRequest,
// 			Message:  "Datos de solicitud no válidos",
// 			ErrorLog: err,
// 		})
// 	}

// 	// Obtener el usuario por su email
// 	ctx := context.Background()
// 	user, err := queries.GetUser(ctx, loginData.Email)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, response{
// 			Status:   http.StatusInternalServerError,
// 			Message:  "Error al buscar el usuario",
// 			ErrorLog: err,
// 		})
// 	}

// 	// Verificar si el usuario existe y si la contraseña coincide
// 	if user == nil || user.Password != loginData.Password {
// 		return c.JSON(http.StatusUnauthorized, response{
// 			Status:  http.StatusUnauthorized,
// 			Message: "Credenciales inválidas",
// 		})
// 	}

// 	// Aquí puedes generar un token de autenticación si lo necesitas

// 	return c.JSON(http.StatusOK, response{
// 		Status:  http.StatusOK,
// 		Message: "Inicio de sesión exitoso",
// 		Data:    user.UserID,
// 	})
// }

// func updatePaswoord(c echo.Context) error {
// 	// User ID from path `users/:id
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

// // 1. Longitud mínima de 8 caracteres.
// // 2. Utilizar mínimo una mayúscula.
// // 3. Utilizar mínimo una minúscula.
// // 4. Utilizar mínimo un carácter especial (que no sea letra ni número).
// // 5. No permitir números consecutivos.
// // 6. No permitir letras consecutivas (con respecto al abecedario).

// func validatePassword(pass string) error {
// 	var (
// 		hasMinLen    = false
// 		hasUppercase = false
// 		hasLowercase = false
// 		hasSpecial   = false
// 		hasNoSeqNum  = true
// 		hasNoSeqLet  = true
// 	)

// 	if len(pass) >= 8 {
// 		hasMinLen = true
// 	}

// 	for i, char := range pass {
// 		switch {
// 		case unicode.IsUpper(char):
// 			hasUppercase = true
// 		case unicode.IsLower(char):
// 			hasLowercase = true
// 		case unicode.IsPunct(char) || unicode.IsSymbol(char):
// 			hasSpecial = true
// 		}

// 		// Verificar números consecutivos
// 		if i > 0 && unicode.IsDigit(char) && unicode.IsDigit(rune(pass[i-1])) {
// 			if char == rune(pass[i-1])+1 {
// 				hasNoSeqNum = false
// 			}
// 		}

// 		// Verificar letras consecutivas
// 		if i > 0 && unicode.IsLetter(char) && unicode.IsLetter(rune(pass[i-1])) {
// 			if unicode.ToLower(char) == unicode.ToLower(rune(pass[i-1]))+1 {
// 				hasNoSeqLet = false
// 			}
// 		}
// 	}

// 	if !hasMinLen {
// 		return errors.New("la contraseña debe tener al menos 8 caracteres")
// 	}
// 	if !hasUppercase {
// 		return errors.New("la contraseña debe tener al menos una mayúscula")
// 	}
// 	if !hasLowercase {
// 		return errors.New("la contraseña debe tener al menos una minúscula")
// 	}
// 	if !hasSpecial {
// 		return errors.New("la contraseña debe tener al menos un carácter especial")
// 	}
// 	if !hasNoSeqNum {
// 		return errors.New("la contraseña no debe contener números consecutivos")
// 	}
// 	if !hasNoSeqLet {
// 		return errors.New("la contraseña no debe contener letras consecutivas")
// 	}

// 	return nil
// }
