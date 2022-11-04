package loggerimpl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/rezaAmiri123/microservice/pkg/logger/tag"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDefaultLogger(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	// copy the output in a separate goroutine so logging can't block indefinitely
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		assert.NoError(t, err)
		outC <- buf.String()
	}()

	zapLogger := zap.NewExample()

	logger := NewLogger(zapLogger)
	preCaller := caller(1)
	logger.WithTags(tag.Error(fmt.Errorf("test error"))).Info("test info", tag.WorkflowActionWorkflowStarted)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	sps := strings.Split(preCaller, ":")
	par, err := strconv.Atoi(sps[1])
	assert.Nil(t, err)
	lineNum := fmt.Sprintf("%v", par+1)
	fmt.Println(out, lineNum)
	assert.Equal(t, out, `{"level":"info","msg":"test info","error":"test error","wf-action":"add-workflow-started-event","logging-call-at":"logger_test.go:`+lineNum+`"}`+"\n")

}

func TestEmptyMsg(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	// copy the output in a separate goroutine so logging can't block indefinitely
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		assert.NoError(t, err)
		outC <- buf.String()
	}()

	zapLogger := zap.NewExample()

	logger := NewLogger(zapLogger)
	preCaller := caller(1)
	logger.WithTags(tag.Error(fmt.Errorf("test error"))).Info("", tag.WorkflowActionWorkflowStarted)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	sps := strings.Split(preCaller, ":")
	par, err := strconv.Atoi(sps[1])
	assert.Nil(t, err)
	lineNum := fmt.Sprintf("%v", par+1)
	fmt.Println(out, lineNum)
	assert.Equal(t, out, `{"level":"info","msg":"`+defaultMsgForEmpty+`","error":"test error","wf-action":"add-workflow-started-event","logging-call-at":"logger_test.go:`+lineNum+`"}`+"\n")

}

func TestErrorWithDetails(t *testing.T) {
	sb := &strings.Builder{}
	zapLogger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{MessageKey: "msg"}), zapcore.AddSync(sb), zap.DebugLevel))
	logger := NewLogger(zapLogger)

	err := &testError{"workflow123"}
	logger.Error("oh no", tag.Error(err))
	zapLogger.Sync()

	assert.Contains(t, sb.String(), `"msg":"oh no","error":"test error","error-details":{"workflow-id":"workflow123"}`)
}

type testError struct {
	WorkflowID string
}

func (e testError) Error() string { return "test error" }
func (e testError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("workflow-id", e.WorkflowID)
	return nil
}
