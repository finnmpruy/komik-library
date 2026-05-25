package models

type Buku struct {
	ID          uint `gorm:"primaryKey"`
	Judul       string
	Penulis     string
	Penerbit    string
	TahunTerbit string
	Stok        int
	Cover       string
	Deskripsi   string
}

func (Buku) TableName() string {
	return "buku"
}
