package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // Driver SQL Server
)

var DB *sql.DB

// Conexión con autenticación de Windows
func ConectarDB() error {
	// Uso conexión integrada de Windows (sin usuario/contraseña)
	connectionString := "server=DESKTOP-BU1JQH3;database=LibrosDB;integrated security=true"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return fmt.Errorf("error al abrir conexión: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("no se pudo conectar a SQL Server: %v", err)
	}

	DB = db
	fmt.Println("✅ Conexión exitosa a SQL Server (autenticación Windows)")
	return nil
}
