package migration

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	dbHelper "github.com/gauravlad21/ecommerce-golang/user-auth/dbhelper"
)

func setEnv() {

	oldDB := dbHelper.GetDbUrl()

	os.Setenv("GOOSE_DBSTRING", oldDB)
	os.Setenv("GOOSE_DRIVER", "postgres")
}

func Migrate() {
	setEnv()

	t := time.After(10 * time.Second)
	select {
	case <-t:
		runMigration()
	}
}

func runMigration() {
	cmd := exec.Command("goose", "up")
	cmd.Dir = "./migrations"
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		panic(err)
	}
}
