package controller_test

import (
	"bytes"
	"testing"

	"gitee.com/infraboard/go-course/extra/prometheus/script_exportor/controller"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	sc *controller.ScriptCollector
)

func TestScriptExec(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	if err := sc.Exec("test.py", "localhost", b); err != nil {
		t.Fatal(err)
	}

	t.Log(b.String())
}

func init() {
	zap.DevelopmentSetup()

	sc = controller.NewScriptCollector("../modules")
}
