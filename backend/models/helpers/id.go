package helpers

type ModelId interface {
	SetId(string)
	GetGoalId() string
	GetCreatedAt() string
}

func IdMapper[T ModelId, U string](t T) string {
	return t.GetGoalId()
}
