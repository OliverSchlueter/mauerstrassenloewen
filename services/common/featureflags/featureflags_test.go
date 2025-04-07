package featureflags

import (
	"os"
	"testing"
)

func TestFeatureFlag_Enable(t *testing.T) {
	flag := Register("TEST_FEATURE_FLAG")
	flag.Enable()

	if !flag.IsEnabled() {
		t.Errorf("Expected feature flag %s to be enabled, but it is not", flag)
	}
}

func TestFeatureFlag_Disable(t *testing.T) {
	flag := Register("TEST_FEATURE_FLAG")
	flag.Enable()
	flag.Disable()

	if flag.IsEnabled() {
		t.Errorf("Expected feature flag %s to be disabled, but it is not", flag)
	}
}

func TestFeatureFlag_Reset(t *testing.T) {
	flag := Register("TEST_FEATURE_FLAG")
	flag.Enable()
	flag.Reset()

	if flag.IsEnabled() {
		t.Errorf("Expected feature flag %s to be reset, but it is not", flag)
	}
}

func TestFeatureFlag_IsEnabled(t *testing.T) {
	flag := Register("TEST_FEATURE_FLAG")
	err := os.Setenv("TEST_FEATURE_FLAG", "true")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	if !flag.IsEnabled() {
		t.Errorf("Expected feature flag %s to be enabled, but it is not", flag)
	}
}
