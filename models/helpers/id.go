package helpers

type ModelId interface {
	Id() string
}

func IdMapper[T ModelId, U string](t T) string {
	return t.Id()
}