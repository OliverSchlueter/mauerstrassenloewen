package featureflags

import "os"

type FeatureFlag string

var (
	EndToEndEnvironment = register("FEATURE_FLAGS_END_TO_END_ENVIRONMENT")
	SendLogsToLoki      = register("FEATURE_FLAGS_SEND_LOGS_TO_LOKI")
)

var flags = make(map[FeatureFlag]bool)

func register(flag string) FeatureFlag {
	return FeatureFlag(flag)
}

func (ff FeatureFlag) Enable() {
	flags[ff] = true
}

func (ff FeatureFlag) Disable() {
	flags[ff] = false
}

func (ff FeatureFlag) Reset() {
	delete(flags, ff)
}

func (ff FeatureFlag) IsEnabled() bool {
	value := os.Getenv(string(ff))
	if value == "true" {
		return true
	}

	enabled, found := flags[ff]
	if found {
		return enabled
	}

	return false
}
