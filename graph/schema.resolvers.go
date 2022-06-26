package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// schema.graphqlから生成されたresolverの実装ファイル

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/aux-Issa/Go_GraphQL/graph/generated"
	"github.com/aux-Issa/Go_GraphQL/graph/model"
)

// reference:https://blog.y-yuki.net/entry/2017/05/05/000000
// func(レシーバー) 関数名(引数)(戻り値の型)
// 定義したメソッドは(レシーバー.メソッド)の形で呼び出すことができる
// ex) mutationResolver構造体の構造体のインスタンスaに対してa.CreateTodo(ctx, input)を呼び出す
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// レシーバーとして構造体を定義
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
