package util_enum

type BaseEnum[T any] struct {
	Code  T
	Title string
}
