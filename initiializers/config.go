package initiializers

import "os"

var (
	SECRET_KEY = os.Getenv("SECRET_KEY")
	dsn        = os.Getenv("DSN")
)

var Tokenn string
