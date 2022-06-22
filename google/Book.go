package google

// Book es la estructura que representa un libro de Bucherliste
type Book struct {
	GID       string `json:"gid"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	Publisher string `json:"publisher"`
}
