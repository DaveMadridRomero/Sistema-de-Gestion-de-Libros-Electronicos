# Sistema de Gestión de Libros Electrónicos

## Autor

**Madrid Romero Dave**  
Universidad Internacional de Ecuador (UIDE)  
Curso: Programación Orientada a Objetos  

## Descripción

Sistema de Gestión de Libros Electrónicos desarrollado en Go (Golang) como proyecto educativo para la asignatura de Programación Orientada a Objetos en la Universidad Internacional de Ecuador (UIDE). 

Este sistema facilita la administración de una biblioteca digital, permitiendo agregar, editar, eliminar, visualizar, buscar y descargar libros electrónicos. Está diseñado para ser una herramienta práctica que introduce conceptos clave del desarrollo de software modular, manejo de bases de datos, interfaces web y gestión de archivos.

## Introducción

El sistema busca ofrecer una solución sencilla pero robusta para gestionar una colección digital de libros electrónicos, orientada tanto a fines académicos como a posibles aplicaciones empresariales. 

Su estructura modular incluye funcionalidades esenciales para el manejo completo del ciclo de vida de los libros, así como herramientas de búsqueda avanzada y acceso a los contenidos digitales mediante visualización y descarga.

## Funcionalidades Principales

- Registro de libros con detalles completos (título, autor, editorial, año, categoría, idioma, páginas, archivo digital).  
- Edición y actualización de registros existentes.  
- Eliminación de libros obsoletos o duplicados.  
- Consulta y visualización clara de la biblioteca.  
- Búsquedas eficientes con filtros por autor, año y categoría.  
- Visualización y descarga de los archivos electrónicos.

## Módulos del Sistema

- **Gestión de Libros:** CRUD completo para la administración de los libros electrónicos.  
- **Búsqueda y Filtros:** Optimización del acceso a libros mediante búsquedas y filtros dinámicos.  
- **Visualización y Descarga:** Acceso y descarga directa de los archivos digitales.

## Paquetes Utilizados

- `net/http`, `html/template`, `net/url` para el servidor web y manejo HTTP.  
- `encoding/json` para intercambio de datos en formato JSON.  
- `os`, `io`, `mime/multipart` para manipulación de archivos.  
- `database/sql` para conexión y manejo de bases de datos SQL.

## Funciones del Sistema

- Agregar, editar y eliminar libros a través de formularios web.  
- Visualizar detalles y archivos de libros.  
- Buscar y filtrar libros con facilidad.  
- Descargar libros para lectura offline.

## Alcance del Proyecto

El proyecto cubre la implementación de un sistema funcional y educativo para la gestión básica de libros electrónicos, con enfoque en:

- Manejo completo del ciclo de vida del libro digital.  
- Interfaz web amigable y modular.  
- Operaciones con bases de datos y archivos.  
- Aplicación práctica de conceptos de programación orientada a objetos en Go.


