package configuration

import "strings"

// ConfigKeyToCLIArgument is used by the cobra subsystem to convert environment
// variable names into command-line argument names
func ConfigKeyToCLIArgument(configKey string) string {
	return strings.Replace(strings.ToLower(configKey), "_", "-", -1)
}
