package models

// Libro representa un libro electrónico con sus datos encapsulados.
type Libro struct {
	ID        int
	Titulo    string
	Autor     string
	Anio      int
	Categoria string
}

// Getters

func (l *Libro) GetID() int {
	return l.ID
}

func (l *Libro) GetTitulo() string {
	return l.Titulo
}

func (l *Libro) GetAutor() string {
	return l.Autor
}

func (l *Libro) GetAnio() int {
	return l.Anio
}

func (l *Libro) GetCategoria() string {
	return l.Categoria
}

// Setters

func (l *Libro) SetID(id int) {
	l.ID = id
}

func (l *Libro) SetTitulo(titulo string) {
	l.Titulo = titulo
}

func (l *Libro) SetAutor(autor string) {
	l.Autor = autor
}

func (l *Libro) SetAnio(anio int) {
	l.Anio = anio
}

func (l *Libro) SetCategoria(categoria string) {
	l.Categoria = categoria
}
