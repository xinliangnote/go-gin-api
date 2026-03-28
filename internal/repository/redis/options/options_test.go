package options

import (
	"testing"
)

func TestBuildUniversalOptions_Standalone(t *testing.T) {
	cfg := RedisConfig{
		Addr:         "127.0.0.1:6379",
		Pass:         "secret",
		Db:           1,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 5,
	}

	opts := BuildUniversalOptions(cfg)

	if len(opts.Addrs) != 1 || opts.Addrs[0] != "127.0.0.1:6379" {
		t.Errorf("expected Addrs=[127.0.0.1:6379], got %v", opts.Addrs)
	}
	if opts.MasterName != "" {
		t.Errorf("expected empty MasterName, got %q", opts.MasterName)
	}
	if opts.Password != "secret" {
		t.Errorf("expected password 'secret', got %q", opts.Password)
	}
	if opts.DB != 1 {
		t.Errorf("expected DB=1, got %d", opts.DB)
	}
	if opts.MaxRetries != 3 {
		t.Errorf("expected MaxRetries=3, got %d", opts.MaxRetries)
	}
	if opts.PoolSize != 10 {
		t.Errorf("expected PoolSize=10, got %d", opts.PoolSize)
	}
	if opts.MinIdleConns != 5 {
		t.Errorf("expected MinIdleConns=5, got %d", opts.MinIdleConns)
	}
}

func TestBuildUniversalOptions_Cluster(t *testing.T) {
	cfg := RedisConfig{
		Addrs:        []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
		Pass:         "clusterpass",
		MaxRetries:   5,
		PoolSize:     20,
		MinIdleConns: 10,
	}

	opts := BuildUniversalOptions(cfg)

	if len(opts.Addrs) != 3 {
		t.Fatalf("expected 3 addrs, got %d", len(opts.Addrs))
	}
	if opts.Addrs[0] != "127.0.0.1:7000" || opts.Addrs[1] != "127.0.0.1:7001" || opts.Addrs[2] != "127.0.0.1:7002" {
		t.Errorf("unexpected addrs: %v", opts.Addrs)
	}
	if opts.MasterName != "" {
		t.Errorf("expected empty MasterName for cluster, got %q", opts.MasterName)
	}
}

func TestBuildUniversalOptions_Sentinel(t *testing.T) {
	cfg := RedisConfig{
		Addrs:      []string{"127.0.0.1:26379", "127.0.0.1:26380"},
		MasterName: "mymaster",
		Pass:       "sentinelpass",
		Db:         0,
	}

	opts := BuildUniversalOptions(cfg)

	if opts.MasterName != "mymaster" {
		t.Errorf("expected MasterName='mymaster', got %q", opts.MasterName)
	}
	if len(opts.Addrs) != 2 {
		t.Fatalf("expected 2 addrs, got %d", len(opts.Addrs))
	}
}

func TestBuildUniversalOptions_AddrsOverridesAddr(t *testing.T) {
	cfg := RedisConfig{
		Addr:  "127.0.0.1:6379",
		Addrs: []string{"10.0.0.1:7000", "10.0.0.2:7000"},
	}

	opts := BuildUniversalOptions(cfg)

	// Addrs should take precedence; Addr should be ignored
	if len(opts.Addrs) != 2 {
		t.Fatalf("expected 2 addrs (Addrs takes precedence), got %d", len(opts.Addrs))
	}
	if opts.Addrs[0] != "10.0.0.1:7000" {
		t.Errorf("expected first addr '10.0.0.1:7000', got %q", opts.Addrs[0])
	}
}

func TestBuildUniversalOptions_EmptyConfig(t *testing.T) {
	cfg := RedisConfig{}

	opts := BuildUniversalOptions(cfg)

	if len(opts.Addrs) != 0 {
		t.Errorf("expected empty Addrs for empty config, got %v", opts.Addrs)
	}
}
