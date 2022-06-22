package google

import (
	"strings"
)

// GBook es la estructura que representa un libro de google
type GBook struct {
	ID         string `json:"id"`
	VolumeInfo struct {
		Title     string   `json:"title"`
		Authors   []string `json:"authors"`
		Publisher string   `json:"publisher"`
	} `json:"volumeInfo"`
}

// ToBook regresa un objeto de tipo Book para trabajar en el proyecto
func (b *GBook) ToBook() Book {
	return Book{
		GID:       strings.Trim(b.ID, " "),
		Title:     strings.Trim(b.VolumeInfo.Title, " "),
		Authors:   strings.Trim(strings.Join(b.VolumeInfo.Authors, ","), ","),
		Publisher: strings.Trim(b.VolumeInfo.Publisher, " "),
	}
}
