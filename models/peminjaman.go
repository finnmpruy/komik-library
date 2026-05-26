package models

import (
	"time"
)

type Peminjaman struct {
	ID             uint       `gorm:"primaryKey"`
	BukuID         uint       `gorm:"not null"` // ID Komik yang dipinjam
	Buku           Buku       `gorm:"foreignKey:BukuID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID         uint       `gorm:"not null"` // ID User/Member yang meminjam
	User           User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TanggalPinjam  time.Time  `gorm:"not null"`
	TanggalKembali *time.Time // Menggunakan pointer (*) supaya nilainya bisa NULL sebelum komik dikembalikan
	BatasKembali   time.Time  `gorm:"not null"` // Tanggal maksimal harus kembali (misal +7 hari)
	Denda          int        `gorm:"default:0"`
	Status         string     `gorm:"type:varchar(20);default:'Dipinjam'"` // "Dipinjam" atau "Kembali"
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (Peminjaman) TableName() string {
	return "Peminjaman"
}
