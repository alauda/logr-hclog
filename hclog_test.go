package logrhclog

import (
	"bytes"
	"github.com/hashicorp/go-hclog"
	"strings"
	"testing"
	"fmt"
)

func TestLog(t *testing.T) {
	output := &bytes.Buffer{}
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     output,
		JSONFormat: false,
		TimeFormat: " ",
	})

	logr := NewLogger(logger)
	logr.Info("info msg", "a", 1)
	logr.Error(nil, "error msg", "b", 2)
	logr.V(2).Info("debug msg", "c", 3)
	logr.V(3).Info("trace msg", "c", 3)

	content := output.String()
	fmt.Println("output", content)

	if !logr.Enabled() {
		t.Fail()
	}

	if !strings.Contains(content, "INFO") {
		t.Fail()
	}

	if !strings.Contains(content, "DEBUG") {
		t.Fail()
	}

	if !strings.Contains(content, "TRACE") {
		t.Fail()
	}

	if !strings.Contains(content, "ERROR") {
		t.Fail()
	}
}
