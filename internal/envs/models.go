package envs

type EnvNavLinks struct {
	Dev  NavLinks `json:"dev"`
	Prod NavLinks `json:"prod"`
}

type NavLinks struct {
	Home            string
	Bot             string
	Dashboard       string
	Roadmap         string
	AboutMe         string
	AboutBot        string
	AboutLighthouse string
	AboutHomeServer string
	AllProjects     string
}
