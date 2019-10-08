module github.com/dicksontung/yaml-extract

go 1.12

replace github.com/dicksontung/yaml-extract/cmd => ./cmd

require (
	github.com/ghodss/yaml v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
)
