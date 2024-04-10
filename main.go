package main

import (
	"encoding/json"
	"io/ioutil"
	copydir "java-glassfish-backup/src/copy-dir.go"
	mysqldump "java-glassfish-backup/src/mysql-dump"
	"java-glassfish-backup/src/types"
	"log"

	arg "github.com/alexflint/go-arg"
)

var (
	jsonFile      = "src/config.json"
	config        = &types.Config{}
	args          = &types.Args{}
	mysqluser     string
	mysqlpass     string
	mysqlhost     string
	mysqldumpfile string
	mysqldatabase string
	appPath       string
	appName       string
)

func init() {

	arg.MustParse(args)

	if args.User == "" && args.Password == "" {
		log.Fatalln("MySQL user and password are required")
	}

	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(data, config)

	mysqluser = args.User
	mysqlpass = args.Password
	mysqlhost = config.MySQL.Host
	mysqldumpfile = config.MySQL.Dumpfile
	mysqldatabase = config.MySQL.Database
	appPath = config.Application.Path
	appName = config.Application.Name
}

func main() {

	log.Println("\033[1;32m[+]\033[0m MySQL| Dumping database:", mysqldatabase, "in:", mysqldumpfile)
	mysqldump.Dump(mysqlhost, mysqluser, mysqlpass, mysqldatabase, mysqldumpfile)
	log.Println("\033[1;32m[+]\033[0m App| Backing up application:", appPath, "in:", appName+".zip")
	copydir.Zip(appPath, appName, "/tmp")
}
