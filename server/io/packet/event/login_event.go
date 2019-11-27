package event

const LoginEventCode = -1;

type LoginEvent struct {
	requestType            byte
	magicId                byte
	version                uint16
	memoryType             byte
	crcKeys                []uint32 // can toss on the client
	encryptedOpcodeSuccess byte
	decryptIsaacSeed       []uint32
	encryptIsaacSeed       []uint32
	clientId               uint32
	username               string
	password               string
}

func NewLoginEvent(
	requestType byte,
	magicId byte,
	version uint16,
	memoryType byte,
	crcKeys []uint32,
	encryptedOpcodeSuccess byte,
	decryptIsaacSeed []uint32,
	encryptIsaacSeed []uint32,
	clientId uint32,
	username string,
	password string,
) *LoginEvent {
	return &LoginEvent{
		requestType,
		magicId,
		version,
		memoryType,
		crcKeys,
		encryptedOpcodeSuccess,
		decryptIsaacSeed,
		encryptIsaacSeed,
		clientId,
		username,
		password,
	}
}

func (data *LoginEvent) EventCode() int32 {
	return LoginEventCode
}
