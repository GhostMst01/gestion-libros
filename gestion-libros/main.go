package main

import (
	"fmt"
	"log"

	"gestion-libros/domain"
	"gestion-libros/fp"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. DEFINE TU CADENA DE CONEXIÓN
	dsn := "root:test@tcp(127.0.0.1:3306)/BaseBibliotecaria?charset=utf8mb4&parseTime=True&loc=Local"
	// 2. INTENTAMOS LA CONEXIÓN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("\n[ERROR DE CONEXIÓN] No se pudo acceder. Revisa si DBeaver/MySQL está encendido.\nDetalle técnico: %v", err)
	}

	// 3. CONSULTAMOS TUS RECOLECCIONES DE DATA
	var books []domain.Book
	db.Find(&books)

	fmt.Println("   PROCESAMIENTO FUNCIONAL CON DATA DE DBEAVER")
	fmt.Println("-------------------------------------------------")

	// FILTER: Filtrar los libros de la base de datos que sean de la categoría "Novela"
	categoriaObjetivo := "Novela"
	novelas := fp.Filter(books, func(b domain.Book) bool {
		return b.Categoria == categoriaObjetivo
	})

	fmt.Printf("\n[Catálogo] Filtrando por categoría '%s':\n", categoriaObjetivo)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("| %-35s | %-25s |\n", "TÍTULO DEL LIBRO", "AÑO DE PUBLICACIÓN")
	fmt.Println("----------------------------------------------------------------------")
	for _, b := range novelas {
		fmt.Printf("| %-35s | Año: %-20s |\n", b.Titulo, b.Anio_Publicacion)
	}
	fmt.Println("----------------------------------------------------------------------")

	// REDUCE: Contar de forma funcional cuántos registros totales procesó el sistema
	totalLibros := fp.Reduce(books, 0, func(accumulator int, b domain.Book) int {
		return accumulator + 1
	})

	fmt.Printf("\n[Reportes] Registros procesados: %d libros.\n", totalLibros)
	fmt.Println("----------------------------------------------------")
}
