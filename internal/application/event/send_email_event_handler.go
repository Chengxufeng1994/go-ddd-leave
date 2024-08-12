package event

import (
	"encoding/json"
	"fmt"

	commonevent "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	leaveevent "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
)

type SendEmailHandler struct{}

func NewSendEmailHandler() commonevent.EventHandler {
	return &SendEmailHandler{}
}

// Handle implements event.EventHandler.
func (h *SendEmailHandler) Handle(evt commonevent.Event) error {
	fmt.Println("send email")
	switch e := evt.(type) {
	case *leaveevent.LeaveCreatedEvent:
		// fmt.Printf("%s\n", e.Data())

		var leaveDo entity.LeaveAggregate
		_ = json.Unmarshal(e.Data(), &leaveDo)
		fmt.Printf("id %s\n", e.ID())
		fmt.Printf("name %s\n", e.Name())
		fmt.Printf("type %s\n", e.Type())
		fmt.Printf("status %s\n", e.Status())
	}
	return nil
}
