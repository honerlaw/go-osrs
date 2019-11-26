package data

type HandshakeData struct {
	status         byte
	serverIsaacKey int64
}

func NewHandshakeData(status byte, serverIsaacKey int64) *HandshakeData {
	var data = &HandshakeData{
		status:         status,
		serverIsaacKey: serverIsaacKey,
	}
	return data
}

func (data *HandshakeData) EventCode() int32 {
	return -2;
}
