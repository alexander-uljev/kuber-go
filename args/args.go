package args

import (
	"log"
	"os"
)

func Parse() (string, string) {
	if len(os.Args) < 3 {
		log.Fatal("Can't run with no parameters\nPass VM and image names")
	}
	return os.Args[1], os.Args[2]
}
