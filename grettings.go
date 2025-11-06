package modularizacion

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Nombre vacío")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Recibir varios nombres y darles la bienvenida con distintos saludos.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludo
// se selecciona de forma aleatoria.
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
