package buffer_test

import (
	"testing"

	"github.com/Jeffail/benthos/v3/internal/log"
	"github.com/Jeffail/benthos/v3/internal/manager/mock"
	"github.com/Jeffail/benthos/v3/internal/old/buffer"
	"github.com/Jeffail/benthos/v3/internal/old/metrics"
	yaml "gopkg.in/yaml.v3"

	_ "github.com/Jeffail/benthos/v3/public/components/all"
)

func TestConstructorBadType(t *testing.T) {
	conf := buffer.NewConfig()
	conf.Type = "not_exist"

	if _, err := buffer.New(conf, mock.NewManager(), log.Noop(), metrics.Noop()); err == nil {
		t.Error("Expected error, received nil for invalid type")
	}
}

func TestConstructorConfigYAMLInference(t *testing.T) {
	conf := []buffer.Config{}

	if err := yaml.Unmarshal([]byte(`[
		{
			"memory": {
				"value": "foo"
			},
			"none": {
				"query": "foo"
			}
		}
	]`), &conf); err == nil {
		t.Error("Expected error from multi candidates")
	}

	if err := yaml.Unmarshal([]byte(`[
		{
			"memory": {
				"limit": 10
			}
		}
	]`), &conf); err != nil {
		t.Error(err)
	}

	if exp, act := 1, len(conf); exp != act {
		t.Errorf("Wrong number of config parts: %v != %v", act, exp)
		return
	}
}