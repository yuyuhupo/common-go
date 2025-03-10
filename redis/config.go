package redis

// Config is the configuration for a Redis client.
type Config struct {
	// Network is the network type to use. Default is "tcp".
	Network string

	// Addr is the address of the Redis server. Default is "localhost:6379".
	Addr string

	// Password is the password of the Redis server. Default is "".
	Password string

	// DB is the database to be selected after connecting to the server. Default is 0.
	DB int
}
