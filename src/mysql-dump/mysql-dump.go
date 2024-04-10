package mysqldump

import (
	"fmt"
	"os/exec"
)

func Dump(host string, user string, password string, database string, destination string) bool {

	cmd := exec.Command("mysqldump", "-h", host, "-u", user, "-p"+password, database, "--result-file="+destination)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
	}

	return true
}
