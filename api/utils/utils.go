package utils

import (
	"fmt"
	env "github.com/joho/godotenv"
	"os"
)

func LoadEnv()  {
	env.Load()
	fmt.Println(os.Getenv("PORT"))
}
