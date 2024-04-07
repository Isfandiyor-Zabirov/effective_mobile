package env

type Settings struct {
	AppParams   AppParams
	Database    Database
	Log         Log
	ExternalApi ExternalApi
}

type Log struct {
	LogInfo       string
	LogDebug      string
	LogError      string
	LogCompress   bool
	LogMaxSize    int
	LogMaxAge     int
	LogMaxBackups int
}

type AppParams struct {
	Host string
	Port string
}

type Database struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

type ExternalApi struct {
	GetCarInfoUrl string
}
