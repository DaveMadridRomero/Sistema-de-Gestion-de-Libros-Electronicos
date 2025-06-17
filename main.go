package main

import (
	"fmt"
	"net/http"

	"github.com/DaveMadridRomero/libros-electronicos/controllers"
	"github.com/DaveMadridRomero/libros-electronicos/services"
	"github.com/DaveMadridRomero/libros-electronicos/utils"
)

func main() {
	// Conexión a la base de datos SQL Server
	if err := utils.ConectarDB(); err != nil {
		fmt.Println("❌ Error al conectar a la base de datos:", err)
		return
	}

	// Se crea la instancia del servicio y se pasa al controlador
	libroService := services.NewLibroService()
	libroController := controllers.NewLibroController(libroService)

	// Registro de rutas y controladores
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bienvenido al Sistema de Gestión de Libros Electrónicos")
	})

	http.HandleFunc("/libros", libroController.ListarLibros)    // GET
	http.HandleFunc("/agregar", libroController.AgregarLibro)   // POST
	http.HandleFunc("/eliminar", libroController.EliminarLibro) // DELETE con ?id=
	http.HandleFunc("/editar", libroController.EditarLibro)     // PUT

	// Iniciar el servidor web en el puerto 8080
	fmt.Println("✅ Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("❌ Error al iniciar el servidor:", err)
	}
}
