package resolvers

//go:generate go run github.com/99designs/gqlgen generate
import datasources "github.com/hwoodall30/sqlite-gql/graph/dataSources"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DataSources *datasources.DataSource
}
