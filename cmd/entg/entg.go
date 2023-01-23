package main

import (
	"log"

	// _ "github.com/lib/pq"

	// _ "entdemo/ent/runtime"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var err error

	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		// entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("./ent.graphql"),
		// entgql.WithConfigPath("./apigql/gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		// entc.TemplateDir("./tmpl"),
		entc.FeatureNames("sql/execquery", "intercept", "schema/snapshot"),

		entc.Extensions(ex),
	}

	err = entc.Generate("./ent/schema",
		&gen.Config{},
		opts...,
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
