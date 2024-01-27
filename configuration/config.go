package configuration

import (
	"boyi/pkg/Infra/db"
	"boyi/pkg/Infra/errors"
	"boyi/pkg/Infra/gin"
	"boyi/pkg/Infra/qqzeng_ip"
	"boyi/pkg/Infra/redis"
	"boyi/pkg/Infra/storage"
	"boyi/pkg/Infra/zlog"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var (
	// configOnce sync.Once
	config = AppConfig{}
)

// AppConfig APP設定
type AppConfig struct {
	fx.Out

	App       *App        `mapstructure:"app"`
	Http      *gin.Config `mapstructure:"http"`
	Log       *zlog.Config
	Databases *db.Config         `mapstructure:"databases"`
	Redis     *redis.Config      `mapstructure:"redis"`
	Storage   *storage.Config    `mapstructure:"storage"`
	Sentry    *zlog.SentryConfig `mapstructure:"sentry"`
	QQZeng    *qqzeng_ip.Config  `mapstructure:"qqzeng"`
	Jwt       *Jwt               `mapstructure:"jwt"`
	// Slack     *log.SlackConfig

}

// Init 初始化 config & log & Global Setting
func Init() (AppConfig, error) {
	viper.AutomaticEnv()
	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		configPath = viper.GetString("PROJ_DIR")

		if viper.GetString("PROJ_DIR") == "" {
			return AppConfig{}, errors.New("PROJ_DIR is required")
		}
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "app"
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("properties")

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, err
	}

	if config.Sentry != nil {
		config.Sentry.IgnoreErrors = []string{
			errors.ErrSpinachAllowed.Error(),
		}
	}

	return config, nil
}

// Get get config
func Get() *AppConfig {
	return &config
}

type App struct {
	MenuFilePath             string            `mapstructure:"menu_file_path"`
	MenuDefaultAdminFilePath string            `mapstructure:"menu_default_admin_file_path"`
	MenuDefaultCSFilePath    string            `mapstructure:"menu_default_cs_file_path"`
	IpTablePath              string            `mapstructure:"ip_table_path"`
	BusinessSystems          map[string]string `mapstructure:"business_systems"` // key: app_id, value: app_key
	Origin                   Origin            `mapstructure:"origin"`
}

type Origin struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
}

type SchedulerConfig struct {
	Jobs map[string]string `mapstructure:"jobs"` // key spec
}

type Jwt struct {
	Issuer         string `mapstructure:"issuer"`
	Audience       string `mapstructure:"audience"`
	SignKey        string `mapstructure:"sign_key"`
	ExpiresMinutes int16  `mapstructure:"expires_minutes"`
}
