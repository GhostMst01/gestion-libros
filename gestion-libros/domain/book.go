package domain

// Book representa la estructura de un libro tal como está en DBeaver
type Book struct {
	ID_Libro         uint   `gorm:"column:ID_Libro;primaryKey" json:"id_libro"`
	Titulo           string `gorm:"column:Titulo" json:"titulo"`
	Categoria        string `gorm:"column:Categoria" json:"categoria"`
	Anio_Publicacion string `gorm:"column:Anio_Publicacion" json:"anio_publicacion"`
	ID_Editorial     uint   `gorm:"column:ID_Editorial" json:"id_editorial"`
}

// TableName le indica a GORM que busque la tabla exactamente como "LIBRO" en mayúsculas
func (Book) TableName() string {
	return "LIBRO"
}