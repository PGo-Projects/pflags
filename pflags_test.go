package pflags

import (
	"fmt"
	"testing"
)

func TestParseFeatureOne(t *testing.T) {
	DebugMode = true
	cfg, err := Parse("test.pflags", "feature_one")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cfg)
}

func TestParseFeatureTwo(t *testing.T) {
	DebugMode = true
	cfg, err := Parse("test.pflags", "feature_two")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cfg)
}

func TestParseFeatureOneAndTwo(t *testing.T) {
	DebugMode = true
	cfg, err := Parse("test.pflags", "feature_one", "feature_two")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cfg)
}
