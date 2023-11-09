package config

type config struct {
	Database      database
	WebServerPort string
}

type database struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}
