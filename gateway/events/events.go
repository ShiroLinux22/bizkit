package events

import (
	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/util"
)

var (
	log = util.Logger {
		Name: "EventHandler",
	}
)

type Handler struct {
	handler.EventHandlerR;
}
