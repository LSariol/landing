package envs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load(vars []string) error {

	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	if err := validate(vars); err != nil {
		return err
	}

	return nil
}

func validate(vars []string) error {

	for _, v := range vars {
		value := os.Getenv(v)
		if value == "" {
			log.Printf("%s: %\n", v, value)
			return fmt.Errorf("unable to load %s", v)
		}
	}
	return nil
}
