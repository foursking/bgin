package safe

import (
	"fmt"
	"github.com/foursking/bgin/helper/dingtalk"
	"github.com/foursking/golib/utils/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"runtime/debug"
)

func Defer() {
	r := recover()
	if r == nil {
		return
	}
	err := cast.ToString(r)
	stack := string(debug.Stack())
	logger.Use("panic").Error(err, zap.String("debug_stack", stack))
	dingtalk.PushSimpleMessage(fmt.Sprintf("recovery from panic:\n%s\n%s", err, stack), true)
}
