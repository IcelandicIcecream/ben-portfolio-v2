package model

import (
	"github.com/a-h/templ"
	"github.com/rs/zerolog"
)

type State struct {
	Logger *zerolog.Logger
}

type TechCard struct {
	Title       string
	Description string
	Rating      string
	Link        templ.SafeURL
}

type ProjectCard struct {
	Title       string
	Tags        []string
	Date        string
	Description string
	Link        templ.SafeURL
}
