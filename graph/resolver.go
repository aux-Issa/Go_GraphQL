package graph

import "github.com/aux-Issa/Go_GraphQL/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// graph/resolver.goからリクエストに対する適切なメソッドがこのファイルから呼ばれる

type Resolver struct {
	todos []*model.Todo // *はポインタの指す実態
}
