package model

import "time"

type File struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type FileChunk struct {
	Name    string
	Content []byte
}
type Message struct {
	Message string
}
