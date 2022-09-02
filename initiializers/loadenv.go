package initiializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load("go.env")

	if err != nil {
		log.Println(err)
	}
}
