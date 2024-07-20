package server

type Error struct {
	Error string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

type PostFile struct {
	FileId string `json:"file_id"`
}
