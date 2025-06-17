package services

import (
	"errors"

	"github.com/DaveMadridRomero/libros-electronicos/models"
	"github.com/DaveMadridRomero/libros-electronicos/utils"
)

var (
	ErrLibroNoEncontrado = errors.New("libro no encontrado")
	ErrDatosInvalidos    = errors.New("datos inválidos")
)

// LibroService define los métodos disponibles para gestionar libros
type LibroService interface {
	Listar() ([]models.Libro, error)
	Agregar(libro *models.Libro) error
	Editar(libro *models.Libro) error
	Eliminar(id int) error
	Validar(libro models.Libro) error
}

// libroServiceImpl es la implementación concreta
type libroServiceImpl struct{}

// NewLibroService crea una nueva instancia del servicio
func NewLibroService() LibroService {
	return &libroServiceImpl{}
}

// Listar devuelve todos los libros
func (s *libroServiceImpl) Listar() ([]models.Libro, error) {
	rows, err := utils.DB.Query("SELECT ID, Titulo, Autor, Anio, Categoria FROM Libros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []models.Libro
	for rows.Next() {
		var l models.Libro
		if err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Anio, &l.Categoria); err != nil {
			return nil, err
		}
		libros = append(libros, l)
	}
	return libros, nil
}

// Agregar inserta un nuevo libro
func (s *libroServiceImpl) Agregar(libro *models.Libro) error {
	query := "INSERT INTO Libros (Titulo, Autor, Anio, Categoria) OUTPUT INSERTED.ID VALUES (@p1, @p2, @p3, @p4)"
	return utils.DB.QueryRow(query, libro.GetTitulo(), libro.GetAutor(), libro.GetAnio(), libro.GetCategoria()).Scan(&libro.ID)
}

// Editar actualiza un libro existente
func (s *libroServiceImpl) Editar(libro *models.Libro) error {
	query := "UPDATE Libros SET Titulo = @p1, Autor = @p2, Anio = @p3, Categoria = @p4 WHERE ID = @p5"
	result, err := utils.DB.Exec(query, libro.GetTitulo(), libro.GetAutor(), libro.GetAnio(), libro.GetCategoria(), libro.GetID())
	if err != nil {
		return err
	}
	afectadas, _ := result.RowsAffected()
	if afectadas == 0 {
		return ErrLibroNoEncontrado
	}
	return nil
}

// Eliminar borra un libro por ID
func (s *libroServiceImpl) Eliminar(id int) error {
	result, err := utils.DB.Exec("DELETE FROM Libros WHERE ID = @p1", id)
	if err != nil {
		return err
	}
	afectadas, _ := result.RowsAffected()
	if afectadas == 0 {
		return ErrLibroNoEncontrado
	}
	return nil
}

// Validar comprueba los datos de un libro
func (s *libroServiceImpl) Validar(libro models.Libro) error {
	if libro.GetTitulo() == "" {
		return errors.New("el título es obligatorio")
	}
	if libro.GetAutor() == "" {
		return errors.New("el autor es obligatorio")
	}
	if libro.GetAnio() < 0 {
		return errors.New("el año debe ser válido")
	}
	if libro.GetCategoria() == "" {
		return errors.New("la categoría es obligatoria")
	}
	return nil
}
