package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnildです")
	} else {
		tracer.Trace("こんにtは、traceパッケージ")
		if buf.String() != "こんにtは、traceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
	// t.Error("まだテストを作成していません")
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}
