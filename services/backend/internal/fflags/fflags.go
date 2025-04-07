package fflags

import "common/featureflags"

var (
	EndToEndEnvironment = featureflags.Register("FEATURE_FLAGS_END_TO_END_ENVIRONMENT")
	StartTestContainers = featureflags.Register("FEATURE_FLAGS_START_TEST_CONTAINERS")
	SendLogsToLoki      = featureflags.Register("FEATURE_FLAGS_SEND_LOGS_TO_LOKI")
)
