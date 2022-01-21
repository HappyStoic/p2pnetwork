package config

import (
	"fmt"
	"github.com/pkg/errors"
)

type Config struct {
	Identity      IdentityConfig
	PeerDiscovery PeerDiscovery
	Server        Server
	Redis         Redis
}

type Server struct {
	Port uint32
}

type PeerDiscovery struct {
	UseDns               bool
	UseRedisCache        bool
	ListOfMultiAddresses []string
}

type IdentityConfig struct {
	GenerateNewKey  bool
	LoadKeyFromFile string
	SaveKeyToFile   string
}

type Redis struct {
	Host     string
	Port     uint
	Db       int
	Username string
	Password string
	Channels RedisChannels
}

type RedisChannels struct {
	Tl2nlAlarmChannel string
	Nl2tlAlarmChannel string
}

// Addr constructs address from host and port
func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func (c *Config) Check() error {
	// validity check
	if c.Identity.GenerateNewKey && c.Identity.LoadKeyFromFile != "" {
		return errors.New("cannot generate new key and load one from file at the same time")
	}
	if !c.Identity.GenerateNewKey && c.Identity.LoadKeyFromFile == "" {
		return errors.New("specify either to generate a new key or load one from a file")
	}
	if c.Redis.Host == "" {
		return errors.New("redis host must be specified")
	}
	if c.Redis.Channels.Tl2nlAlarmChannel == "" ||
		c.Redis.Channels.Nl2tlAlarmChannel == "" {
		return errors.New("redis channels must be specified")
	}
	// default values
	if c.Redis.Port == 0 {
		c.Redis.Port = 6379 // Use default redis port if port is not specified
	}
	return nil
}
