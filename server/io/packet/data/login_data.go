package data

type LoginData struct {
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

func NewLoginData(
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
) *LoginData {
	return &LoginData{
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

func (data *LoginData) EventCode() int32 {
	return -1;
}
