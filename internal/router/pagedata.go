package router

import "github.com/mobasity-web-landing/internal/envs"

func LoadPageData() PageData {

	var pd PageData

	pd.Navigation = envs.GetLinks()

	return pd
}
