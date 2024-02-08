package model

import (
	"github.com/a-h/templ"
)

type State struct {
	IsMobileNavBarOpen bool
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
