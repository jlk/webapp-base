package fields // import "github.com/jlk/webapp-base/server/mutations/fields"

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

// CreateNotTodo - graphql insert mutation
var CreateNotTodo = &graphql.Field{
	Type:        types.NotTodo,
	Description: "Create a not Todo",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// get our params
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)

		mongoSession := data.MongoClient.Copy()
		defer mongoSession.Close()
		notTodoCollection := mongoSession.DB("medium-app").C("Not_Todos")

		err := notTodoCollection.Insert(map[string]string{"name": name, "description": description})

		if err != nil {
			logrus.Fatal("Error inserting db record:" + err.Error())
		}

		return todoStruct{name, description}, nil
	},
}
