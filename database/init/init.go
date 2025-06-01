package init

import (
	"context"
	"server/database/repository/repo"

	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Printf("Logg %v", help)

		}
	})
	return instance
}

func NewClient(ctx context.Context, sc StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil

}

func NewRepository(client repo.Client) repo.Repository1 {
	return &repo.Repository{
		Client: client,
	}
}
