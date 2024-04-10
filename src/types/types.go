package types

type MySQL struct {
	Dumpfile string `json:"dumpfile"`
	Host     string `json:"host"`
	Database string `json:"database"`
	ZipDump  string `json:"zipdump"`
}

type Application struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	ZipApp string `json:"zipapp"`
}

type S3 struct {
	BucketName string `json:"bucketname"`
}

type AWS struct {
	S3 S3 `json:"s3"`
}

type Config struct {
	Application Application `json:"application"`
	MySQL       MySQL       `json:"mysql"`
	Destination string      `json:"destination"`
	AWS         AWS         `json:"aws"`
}

type Args struct {
	User     string `arg:"-u,--user" help:"MySQL username"`
	Password string `arg:"-p,--password" help:"MySQL password"`
}
