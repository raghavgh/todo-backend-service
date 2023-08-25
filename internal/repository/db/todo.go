package db

type DBTodoRepository struct {
	dbRepo *Repository
}

func NewDbTodoRepository() *DBTodoRepository {
	return &DBTodoRepository{dbRepo: NewRepository()}
}
