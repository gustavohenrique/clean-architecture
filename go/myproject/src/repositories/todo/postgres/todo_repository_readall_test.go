package postgres_test

import (
	"context"
	"testing"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/infra/postgres"
	db "{{ .ProjectName }}/src/repositories/todo/postgres"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/shared/uuid"
	"{{ .ProjectName }}/src/valueobjects"
)

func TestTodoItemRepositoryReadAll(ts *testing.T) {
	todoItem := entities.NewTodoItemEntity()
	todoItem.Title = "My todoitem"
	insertTodoItemQuery := "insert into todo_items (id, title) values ($1, $2)"

	test.WithPostgres(ts, "Should return all", func(t *testing.T, store *postgres.PostgresStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, todoItem.ID, todoItem.Title))

		item := valueobjects.TodoItemTable{}
		item.ID = uuid.NewV4()
		item.Title = "TODO 1"
		item.IsDone = true
		query := "insert into todo_items (id, title, is_done) values ($1, $2, $3)"
		assert.Nil(t, store.Exec(
			query,
			item.ID, item.Title, item.IsDone,
		), "Cannot insert todo item")

		repo := db.NewRepository(infra.New())
		founds, err := repo.ReadAll(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(founds), 2)
		assert.Equal(t, founds[1].ID, item.ID)
		assert.Equal(t, founds[1].Title, item.Title)
		assert.Equal(t, founds[1].IsDone, item.IsDone)
	})
}