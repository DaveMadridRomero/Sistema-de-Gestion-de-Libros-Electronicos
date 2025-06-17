package utils

import "errors"

// Errores comunes para el sistema de gestión de libros electrónicos
var (
	ErrLibroNoEncontrado = errors.New("libro no encontrado")
	ErrDatosInvalidos    = errors.New("datos inválidos")
	ErrIDInvalido        = errors.New("ID inválido")
)
