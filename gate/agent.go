package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	PlayerData() interface{}
	SetPlayerData(data interface{})
	ConfigData() interface{}
	SetConfigData(data interface{})
}
