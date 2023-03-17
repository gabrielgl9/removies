package database

type IDatabase[T any] interface {
	SetupDB() (T, error)
}