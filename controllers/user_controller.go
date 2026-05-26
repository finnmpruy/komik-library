package controllers

import (
	"komik-library/config"
	"komik-library/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DashboardUser(c *gin.Context) {
	var allBuku []models.Buku
	var filteredBuku []models.Buku

	config.DB.Find(&allBuku)

	searchQuery := c.Query("search")

	// Membersihkan query dari spasi berlebih, huruf besar, dan tanda titik
	searchQueryClean := strings.TrimSpace(strings.ToLower(searchQuery))
	searchQueryClean = strings.ReplaceAll(searchQueryClean, ".", "")
	searchQueryClean = strings.ReplaceAll(searchQueryClean, " ", "")

	if searchQueryClean != "" {
		for i := 0; i < len(allBuku); i++ {
			judulBuku := strings.ToLower(allBuku[i].Judul)
			judulBuku = strings.ReplaceAll(judulBuku, ".", "")
			judulBuku = strings.ReplaceAll(judulBuku, " ", "")

			penerbitBuku := strings.ToLower(allBuku[i].Penerbit)
			penerbitBuku = strings.ReplaceAll(penerbitBuku, ".", "")
			penerbitBuku = strings.ReplaceAll(penerbitBuku, " ", "")

			if strings.Contains(judulBuku, searchQueryClean) || strings.Contains(penerbitBuku, searchQueryClean) {
				filteredBuku = append(filteredBuku, allBuku[i])
			}
		}
	} else {
		filteredBuku = allBuku
	}

	c.HTML(http.StatusOK, "dashboard_user.html", gin.H{
		"title":       "Katalog Komik Library",
		"buku":        filteredBuku,
		"searchQuery": searchQuery,
	})
}
