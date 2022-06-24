package graph

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/aux-Issa/Go_GraphQL/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// graph/resolver.goからリクエストに対する適切なメソッドがこのファイルから呼ばれる

type Resolver struct {
	todos []*model.Todo // *はポインタの指す実態
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user" + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
