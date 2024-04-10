package types

type MySQL struct {
	Dumpfile string `json:"dumpfile"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

type Application struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Config struct {
	Application Application `json:"application"`
	MySQL       MySQL       `json:"mysql"`
}

type Args struct {
	User     string `arg:"-u,--user" help:"MySQL username"`
	Password string `arg:"-p,--password" help:"MySQL password"`
}
