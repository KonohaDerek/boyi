package db

// Config for db config
type Config struct {
	Read             *Database
	Write            *Database
	Secrets          string `yaml:"secrets" mapstructure:"secrets"`
	WithColor        bool   `yaml:"with_color" mapstructure:"with_color"`
	WithCaller       bool   `yaml:"with_caller" mapstructure:"with_caller"`
	MigratePath      string `yaml:"migrate_path" mapstructure:"migrate_path"`
	MigrateTableName string `yaml:"migrate_table_name" mapstructure:"migrate_table_name"`
}

// NewInjection ...
func (c *Config) NewInjection() *Config {
	return c
}
