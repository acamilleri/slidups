package model

// File - Represent uploaded file model
type File struct {
	Name        string
	Size        int // Octets
	Content     []byte
	Destination string
}

// NewFile - Init file
func NewFile(name string, size int, content []byte, destination string) File {
	return File{
		Name:        name,
		Size:        size,
		Content:     content,
		Destination: destination,
	}
}
