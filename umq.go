package umq

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Options struct {
	Client                mqtt.Client
	Broker                string
	ClientID              string
	UserName              string
	PassWord              string
	DefaultPublishHandler mqtt.MessageHandler
	OnConn                mqtt.OnConnectHandler
}

func (t *Options) Connect() error {
	ops := mqtt.NewClientOptions().AddBroker(t.Broker).SetClientID(t.ClientID)
	ops.SetUsername(t.UserName)
	ops.SetPassword(t.PassWord)
	ops.SetKeepAlive(60 * time.Second)
	ops.SetPingTimeout(5 * time.Second)
	ops.CleanSession = true
	ops.SetDefaultPublishHandler(t.DefaultPublishHandler)
	ops.SetOnConnectHandler(t.OnConn)
	t.Client = mqtt.NewClient(ops)
	if token := t.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
