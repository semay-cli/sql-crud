package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config is the interface exposed for environment configuration.
type Config interface {
	Get(key string) string
	GetOrDefault(key, defaultValue string) string
}

// EnvConfig implements the Config interface.
type EnvConfig struct {
	configPath string
	env        string
}

func NewEnvConfig() (*EnvConfig, error) {
	// Step 1: Load .env file (if exists)
	godotenv.Load()

	// Step 2: Read environment variables
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs" // default
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // default
	}

	// Step 3: Init config
	cfg := &EnvConfig{
		configPath: configPath,
		env:        env,
	}

	// Step 4: Load additional config if needed
	if err := cfg.load(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// load reads environment files in the correct order and overrides accordingly.
func (e *EnvConfig) load() error {
	filesToLoad := []string{
		filepath.Join(e.configPath, ".env"), // base config
	}

	// Determine override file: .<env>.env or fallback to .dev.env
	overrideFile := ".dev.env"
	if e.env != "" {
		overrideFile = fmt.Sprintf(".%s.env", e.env)
	}
	filesToLoad = append(filesToLoad, filepath.Join(e.configPath, overrideFile))

	for _, file := range filesToLoad {
		if err := godotenv.Overload(file); err != nil {
			fmt.Printf("WARNING: Could not load file %s: %v\n", file, err)
		} else {
			fmt.Printf("INFO: Loaded file: %s\n", file)
		}
	}

	return nil
}

// Get returns the value for the given environment key.
func (e *EnvConfig) Get(key string) string {
	return os.Getenv(key)
}

// GetOrDefault returns the value for the key or a default if not set.
func (e *EnvConfig) GetOrDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func (e *EnvConfig) SetEnv(env string) error {
	if env == "" {
		return fmt.Errorf("environment cannot be empty")
	}

	// Update the internal env and system env var
	e.env = env
	// Reload env files
	if err := e.load(); err != nil {
		return fmt.Errorf("failed to reload env config: %w", err)
	}

	return nil

}
