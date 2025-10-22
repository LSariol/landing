package router

import "github.com/mobasity-web-landing/internal/envs"

type headerInformation struct {
	TitleName string
}

type PageData struct {
	Title      string
	Navigation envs.NavLinks
}
