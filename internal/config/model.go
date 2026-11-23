package config

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
	Email    Email
	Redis    Redis
	Midtrans Midtrans
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
	Tz   string
}

type Jwt struct {
	Key string
	Exp int
}

type Email struct {
	Host     string
	User     string
	Password string
	Port     string
}

type Redis struct {
	Addr string
	Pass string
}

type Midtrans struct {
	Key    string
	IsProd bool
}
