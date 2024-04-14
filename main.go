package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	zip "java-glassfish-backup/src/copy-dir.go"
	mysqldump "java-glassfish-backup/src/mysql-dump"
	s3 "java-glassfish-backup/src/s3"
	"java-glassfish-backup/src/types"
	"log"

	arg "github.com/alexflint/go-arg"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

var (
	jsonFile      = "src/config.json"
	config_custom = &types.Config{}
	args          = &types.Args{}
	mysqluser     string
	mysqlpass     string
	mysqlhost     string
	mysqldumpfile string
	mysqldatabase string
	appPath       string
	appName       string
	zipApp        string
	zipDump       string
	destination   string
	bucketName    string
	err           error
	cfg           aws.Config
)

func init() {

	arg.MustParse(args)

	if args.User == "" && args.Password == "" {
		log.Fatalln("MySQL user and password are required")
	}

	godotenv.Load()

	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(data, config_custom)

	mysqluser = args.User
	mysqlpass = args.Password
	mysqlhost = config_custom.MySQL.Host
	mysqldumpfile = config_custom.MySQL.Dumpfile
	mysqldatabase = config_custom.MySQL.Database
	appPath = config_custom.Application.Path
	appName = config_custom.Application.Name
	zipApp = config_custom.Application.ZipApp
	zipDump = config_custom.MySQL.ZipDump
	destination = config_custom.Destination
	bucketName = config_custom.AWS.S3.BucketName
}

func main() {

	log.Println("\033[1;32m[+]\033[0m MySQL  | Dumping database:", mysqldatabase, "as:", mysqldumpfile)
	mysqldump.Dump(mysqlhost, mysqluser, mysqlpass, mysqldatabase, mysqldumpfile)
	log.Println("\033[1;32m[+]\033[0m App    | Backing up application:", appPath, "as:", appName+".zip")
	zip.Dir(appPath, zipApp, destination)
	log.Println("\033[1;32m[+]\033[0m MySQL  | Backing up MySQL:", appPath, "as:", appName+".zip")
	zip.File(mysqldumpfile, zipDump, destination)
	log.Println("\033[1;32m[+]\033[0m AWS-S3 | Uploading MySQL:", zipDump, "on Bucket:", bucketName)
	s3.Upload(cfg, bucketName, zipDump, destination+"/"+zipDump)
	log.Println("\033[1;32m[+]\033[0m AWS-S3 | Uploading App:", zipApp, "on Bucket:", bucketName)
	s3.Upload(cfg, bucketName, zipApp, destination+"/"+zipApp)
}
