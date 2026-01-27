package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	// Verbose enables verbose logging output.
	Verbose = field.BoolField(
		"verbose",
		field.WithDescription("Enable verbose logging output."),
	)
)

//go:generate go run ./gen
var Config = field.NewConfiguration([]field.SchemaField{
	Verbose,
})
