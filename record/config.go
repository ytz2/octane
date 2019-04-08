package record

// Config ...
type Config struct {
	APPName   string  `default:"octane"`
	GRPC      Grpc    `yaml:"grpc"`
	DB        Db      `yaml:"db"`
	RateLimit []Limit `yaml:"ratelimit"`
}

type Limit struct {
	Name  string  `yaml:"name"`
	Limit float64 `yaml:"limit"`
	Burst int     `yaml:"burst"`
}

type Db struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URI      string `yaml:"uri"`
	Port     uint   `yaml:"port"`
}

type Grpc struct {
	Port string `yaml:"port"`
}
