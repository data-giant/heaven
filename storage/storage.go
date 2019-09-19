package storage

type Storage interface {
	Open()
	Write(map[string]string)
}
