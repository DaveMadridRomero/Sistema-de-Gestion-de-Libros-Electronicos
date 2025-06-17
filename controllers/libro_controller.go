package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DaveMadridRomero/libros-electronicos/models"
	"github.com/DaveMadridRomero/libros-electronicos/services"
)

// LibroController gestiona las peticiones HTTP relacionadas con libros.
// Se basa en una interfaz LibroService para desacoplar lógica de negocio y almacenamiento.
type LibroController struct {
	Service services.LibroService
}

// NewLibroController crea un controlador con la implementación del servicio que se le pase.
func NewLibroController(service services.LibroService) *LibroController {
	return &LibroController{Service: service}
}

// ListarLibros responde con un JSON de todos los libros guardados.
func (lc *LibroController) ListarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	libros, err := lc.Service.Listar()
	if err != nil {
		http.Error(w, "Error al obtener los libros: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(libros)
}

// AgregarLibro recibe un JSON con datos del libro y lo guarda.
func (lc *LibroController) AgregarLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = lc.Service.Validar(libro)
	if err != nil {
		http.Error(w, "Validación fallida: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = lc.Service.Agregar(&libro)
	if err != nil {
		http.Error(w, "No se pudo agregar el libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(libro)
}

// EliminarLibro elimina un libro por ID desde query param ?id=
func (lc *LibroController) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = lc.Service.Eliminar(id)
	if err != nil {
		if err == services.ErrLibroNoEncontrado {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Error al eliminar el libro: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// EditarLibro actualiza un libro existente.
func (lc *LibroController) EditarLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil || libro.GetID() < 1 {
		http.Error(w, "Datos inválidos o falta el ID", http.StatusBadRequest)
		return
	}

	err = lc.Service.Validar(libro)
	if err != nil {
		http.Error(w, "Validación fallida: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = lc.Service.Editar(&libro)
	if err != nil {
		if err == services.ErrLibroNoEncontrado {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Error al actualizar el libro: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}
