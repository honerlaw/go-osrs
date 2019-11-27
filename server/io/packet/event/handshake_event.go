package event

const HandshakeEventCode = -2

type HandshakeEvent struct {
	status         byte
	serverIsaacKey int64
}

func NewHandshakeEvent(status byte, serverIsaacKey int64) *HandshakeEvent {
	var data = &HandshakeEvent{
		status:         status,
		serverIsaacKey: serverIsaacKey,
	}
	return data
}

func (data *HandshakeEvent) EventCode() int32 {
	return HandshakeEventCode
}
