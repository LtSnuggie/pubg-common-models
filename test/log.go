package common_test

import (
	"strings"
	"testing"
)

var testLog []string

type TestWriter struct {
}

func NewTestWriter() (t TestWriter) {
	purgeTestLog()
	return TestWriter{}
}

func (tw TestWriter) Write(p []byte) (n int, err error) {
	testLog = append(testLog, string(p))
	return
}

func purgeTestLog() { testLog = make([]string, 0) }
func (tw TestWriter) AuditLogWithoutPurge(msg string, t *testing.T) {
	if msg == "" && len(testLog) > 0 {
		t.Error("received non-empty log during audit")
		return
	} else {
		return
	}
	for _, item := range testLog {
		if strings.Contains(item, msg) {
			return
		}
	}
	t.Errorf("no message matching \"%s\" was found during log audit", msg)
}
func (tw TestWriter) AuditLog(msg string, t *testing.T) {
	tw.AuditLogWithoutPurge(msg, t)
	purgeTestLog()
}
