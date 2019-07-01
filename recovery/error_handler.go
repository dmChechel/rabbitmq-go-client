package recovery

import (
	"github.com/kismia/rabbitmq-go-client"
)

type ForgivingErrorHandler struct{}

func (ForgivingErrorHandler) HandleConnectionRecoveryError(c rabbitmq.Connection, err error) {
	rabbitmq.Logger.Println("connection recovery error", err)
}

func (ForgivingErrorHandler) HandleChannelRecoveryError(ch rabbitmq.Channel, err error) {
	rabbitmq.Logger.Println("channel recovery error", err)
}

func (ForgivingErrorHandler) HandleTopologyRecoveryError(c rabbitmq.Connection, err error) {
	rabbitmq.Logger.Println("topology recovery error", err)
}
