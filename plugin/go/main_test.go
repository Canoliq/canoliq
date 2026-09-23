package main

import (
	"strings"
	"testing"

	"github.com/canopy-network/go-plugin/canoliq"
)

func TestApplyCanoliqEnvironment(t *testing.T) {
	cfg := canoliq.DefaultConfig()
	env := map[string]string{
		canoliqRpcAddrEnv:          ":8587",
		canoliqAlertURLEnv:         "https://alerts.example.test/canoliq",
		canoliqActivationHeightEnv: "12345",
	}
	if err := applyCanoliqEnvironment(&cfg, func(name string) string { return env[name] }); err != nil {
		t.Fatalf("applyCanoliqEnvironment: %v", err)
	}
	if cfg.RpcAddress != ":8587" {
		t.Fatalf("rpc address = %q", cfg.RpcAddress)
	}
	if cfg.Alerts == nil || cfg.Alerts.WebhookURL != env[canoliqAlertURLEnv] {
		t.Fatalf("alert override was not applied: %+v", cfg.Alerts)
	}
	if cfg.ActivationHeight != 12345 {
		t.Fatalf("activation height = %d", cfg.ActivationHeight)
	}
}

func TestApplyCanoliqEnvironmentRejectsInvalidActivationHeight(t *testing.T) {
	for _, raw := range []string{"0", "not-a-height", "-1"} {
		t.Run(raw, func(t *testing.T) {
			cfg := canoliq.DefaultConfig()
			err := applyCanoliqEnvironment(&cfg, func(name string) string {
				if name == canoliqActivationHeightEnv {
					return raw
				}
				return ""
			})
			if err == nil || !strings.Contains(err.Error(), canoliqActivationHeightEnv) {
				t.Fatalf("expected activation-height error for %q, got %v", raw, err)
			}
		})
	}
}
