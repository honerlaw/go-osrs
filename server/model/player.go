package model

type Player struct {
	namehash   byte
	memoryType byte
	clientId   uint32
	username   string
	password   string
}

func NewPlayer(namehash byte) *Player {
	return &Player{
		namehash: namehash,
	}
}

func (p *Player) SetLoginInformation(memoryType byte, clientId uint32, username string, password string) {
	p.memoryType = memoryType
	p.clientId = clientId
	p.username = username
	p.password = password
}
