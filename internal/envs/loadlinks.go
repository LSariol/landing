package envs

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetLinks() NavLinks {
	data, err := os.ReadFile("internal/envs/links.json")
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	var nL EnvNavLinks
	if err := json.Unmarshal(data, &nL); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	env := os.Getenv("ENV")
	if env == "DEV" {
		return nL.Dev
	}
	return nL.Prod
}
