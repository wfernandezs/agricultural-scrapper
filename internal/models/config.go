package models

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Crawler  CrawlerConfig  `mapstructure:"crawler"`
	API      APIConfig      `mapstructure:"api"`
	Analyzer AnalyzerConfig `mapstructure:"analyzer"`
}

// DatabaseConfig contains database connection settings
type DatabaseConfig struct {
	Engine        string        `mapstructure:"engine"`    // sqlite, postgres, etc.
	DSN           string        `mapstructure:"dsn"`       // Data Source Name
	MaxOpenConns  int           `mapstructure:"max_open_conns"`
	MaxIdleConns  int           `mapstructure:"max_idle_conns"`
	ConnLifetime  time.Duration `mapstructure:"conn_lifetime"`
	MigrationsDir string        `mapstructure:"migrations_dir"`
}

// CrawlerConfig contains settings for the web crawler
type CrawlerConfig struct {
	UserAgent         string        `mapstructure:"user_agent"`
	Concurrency       int           `mapstructure:"concurrency"`
	RequestDelay      time.Duration `mapstructure:"request_delay"`
	RespectRobotsTxt  bool          `mapstructure:"respect_robots_txt"`
	ScheduleInterval  time.Duration `mapstructure:"schedule_interval"`
	MaxRetries        int           `mapstructure:"max_retries"`
	RetryDelay        time.Duration `mapstructure:"retry_delay"`
	DisallowedDomains []string      `mapstructure:"disallowed_domains"`
}

// APIConfig contains settings for the API server
type APIConfig struct {
	ListenAddress string        `mapstructure:"listen_address"`
	JWTSecret     string        `mapstructure:"jwt_secret"`
	JWTExpiry     time.Duration `mapstructure:"jwt_expiry"`
	RateLimit     int           `mapstructure:"rate_limit"`
	CORSOrigins   []string      `mapstructure:"cors_origins"`
	ReadTimeout   time.Duration `mapstructure:"read_timeout"`
	WriteTimeout  time.Duration `mapstructure:"write_timeout"`
}

// AnalyzerConfig contains settings for data analysis
type AnalyzerConfig struct {
	ForecastWindow     int           `mapstructure:"forecast_window"`      // Days to forecast ahead
	HistoryWindow      int           `mapstructure:"history_window"`       // Days of history to use
	UpdateInterval     time.Duration `mapstructure:"update_interval"`      // How often to update analysis
	AnomalyThreshold   float64       `mapstructure:"anomaly_threshold"`    // Threshold for price anomalies
	SeasonalAdjustment bool          `mapstructure:"seasonal_adjustment"`  // Whether to apply seasonal adjustments
}

// LoadConfig loads configuration from file
func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	
	// Set defaults
	v.SetDefault("database.engine", "sqlite")
	v.SetDefault("database.dsn", "data/agricultural_data.db")
	v.SetDefault("database.max_open_conns", 20)
	v.SetDefault("database.max_idle_conns", 5)
	v.SetDefault("database.conn_lifetime", "1h")
	v.SetDefault("database.migrations_dir", "migrations")
	
	v.SetDefault("crawler.user_agent", "AgriculturalIntelligenceBot/1.0")
	v.SetDefault("crawler.concurrency", 5)
	v.SetDefault("crawler.request_delay", "2s")
	v.SetDefault("crawler.respect_robots_txt", true)
	v.SetDefault("crawler.schedule_interval", "24h")
	v.SetDefault("crawler.max_retries", 3)
	v.SetDefault("crawler.retry_delay", "5s")
	
	v.SetDefault("api.listen_address", ":8080")
	v.SetDefault("api.jwt_expiry", "24h")
	v.SetDefault("api.rate_limit", 100)
	v.SetDefault("api.cors_origins", []string{"*"})
	v.SetDefault("api.read_timeout", "15s")
	v.SetDefault("api.write_timeout", "15s")
	
	v.SetDefault("analyzer.forecast_window", 30)
	v.SetDefault("analyzer.history_window", 365)
	v.SetDefault("analyzer.update_interval", "6h")
	v.SetDefault("analyzer.anomaly_threshold", 0.15)
	v.SetDefault("analyzer.seasonal_adjustment", true)
	
	// Set config name and path
	v.SetConfigFile(path)
	
	// Read config file
	if err := v.ReadInConfig(); err != nil {
		// Config file not found, use defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}
	
	// Check for environment variables
	v.AutomaticEnv()
	
	// Unmarshal config
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}
	
	// Override JWT secret with environment variable if set
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.API.JWTSecret = jwtSecret
	}
	
	return &config, nil
}