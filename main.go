package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"database/sql"

	"github.com/DaveMadridRomero/libros-electronicos/models"
	"github.com/DaveMadridRomero/libros-electronicos/utils"
)

func main() {
	err := utils.ConectarDB()
	if err != nil {
		fmt.Println("❌ No se pudo conectar:", err)
		return
	}
	defer utils.DB.Close() // usa tu variable global si está definida así

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nMenú de prueba")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Salir")
		fmt.Print("Elija una opción: ")

		scanner.Scan()
		opcion := scanner.Text()

		switch opcion {
		case "1":
			agregarLibro(utils.DB, scanner)
		case "2":
			listarLibros(utils.DB)
		case "3":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

func agregarLibro(db *sql.DB, scanner *bufio.Scanner) {
	var l models.Libro
	var err error

	fmt.Print("Título: ")
	scanner.Scan()
	l.Titulo = strings.TrimSpace(scanner.Text())

	fmt.Print("Autor: ")
	scanner.Scan()
	l.Autor = strings.TrimSpace(scanner.Text())

	fmt.Print("Año: ")
	scanner.Scan()
	anio, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Año inválido")
		return
	}
	l.Anio = anio

	fmt.Print("Categoría: ")
	scanner.Scan()
	l.Categoria = strings.TrimSpace(scanner.Text())

	_, err = db.Exec(`INSERT INTO libros (titulo, autor, anio, categoria) VALUES (@p1, @p2, @p3, @p4)`,
		l.Titulo, l.Autor, l.Anio, l.Categoria)
	if err != nil {
		fmt.Println("Error al insertar:", err)
		return
	}
	fmt.Println("✅ Libro agregado")
}

func listarLibros(db *sql.DB) {
	rows, err := db.Query("SELECT id, titulo, autor, anio, categoria FROM libros")
	if err != nil {
		fmt.Println("Error al consultar:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Libros disponibles:")
	for rows.Next() {
		var l models.Libro
		err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Anio, &l.Categoria)
		if err != nil {
			fmt.Println("Error leyendo:", err)
			continue
		}
		fmt.Printf("ID %d - %s | %s | %d | %s\n", l.ID, l.Titulo, l.Autor, l.Anio, l.Categoria)
	}
}
