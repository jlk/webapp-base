package fields // import "github.com/jlk/webapp-base/server/queries/fields"

import (
	"github.com/jlk/webapp-base/server/data"
	types "github.com/jlk/webapp-base/server/types"

	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

// GetNotTodos - graphql query
var GetNotTodos = &graphql.Field{
	Type:        graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		notTodoCollection := data.MongoClient.DB("medium-app").C("Not_Todos")

		var todosList []todoStruct

		err := notTodoCollection.Find(nil).All(&todosList)

		if err != nil {
			logrus.Fatal("Error while querying database: " + err.Error())
		}
		return todosList, nil
	},
}
