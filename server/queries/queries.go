package queries // import "github.com/jlk/webapp-base/server/queries"

import (
  "github.com/graphql-go/graphql"
  fields "github.com/jlk/webapp-base/server/queries/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    "getNotTodos": fields.GetNotTodos,
  },
})
