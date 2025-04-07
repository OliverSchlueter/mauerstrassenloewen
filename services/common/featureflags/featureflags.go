package featureflags

import "os"

type FeatureFlag string

var flags = make(map[FeatureFlag]bool)

func Register(flag string) FeatureFlag {
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
