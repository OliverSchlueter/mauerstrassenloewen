package fflags

import "github.com/OliverSchlueter/mauerstrassenloewen/common/featureflags"

var (
	EndToEndEnvironment = featureflags.Register("FEATURE_FLAGS_END_TO_END_ENVIRONMENT")
	SendLogsToLoki      = featureflags.Register("FEATURE_FLAGS_SEND_LOGS_TO_LOKI")
)
