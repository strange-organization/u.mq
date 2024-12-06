package umq

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Options struct {
	Client           mqtt.Client
	Broker           string
	ClientID         string
	UserName         string
	PassWord         string
	DefaultPublish   mqtt.MessageHandler
	OnConn           mqtt.OnConnectHandler
	OnReconnect      mqtt.ReconnectHandler
	OnConnectionLost mqtt.ConnectionLostHandler
}

func (t *Options) Connect() error {
	ops := mqtt.NewClientOptions().AddBroker(t.Broker).SetClientID(t.ClientID)
	ops.SetUsername(t.UserName)
	ops.SetPassword(t.PassWord)
	ops.SetKeepAlive(60 * time.Second)
	ops.SetPingTimeout(5 * time.Second)
	ops.CleanSession = true
	ops.SetDefaultPublishHandler(t.DefaultPublish)
	ops.SetOnConnectHandler(t.OnConn)
	ops.SetReconnectingHandler(t.OnReconnect)
	ops.SetConnectionLostHandler(t.OnConnectionLost)

	t.Client = mqtt.NewClient(ops)
	if token := t.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (t *Options) Close() {
	if t.Client != nil {
		t.Client.Disconnect(250)
		time.Sleep(1 * time.Second)
	}
}
