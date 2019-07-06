package mutations // import "github.com/jlk/webapp-base/server/mutations"

import (
	"github.com/graphql-go/graphql"
	fields "github.com/jlk/webapp-base/server/mutations/fields"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateNotTodo,
	},
})
