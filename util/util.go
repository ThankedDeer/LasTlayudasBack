package util

import (
	"math/rand"
	"time"
)

// RandomString genera una cadena aleatoria de longitud n.
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// RandomOwner genera un correo electrónico aleatorio.
func RandomOwner() string {
	domains := []string{"example.com", "test.com", "email.com"}
	names := []string{"user", "admin", "test"}
	rand.Seed(time.Now().UnixNano())

	name := names[rand.Intn(len(names))]
	domain := domains[rand.Intn(len(domains))]

	return name + "@" + domain
}
