package model

type Player struct {
	Namehash   byte
	MemoryType byte
	ClientId   uint32
	Username   string
	Password   string
	Position   *Position
}

func NewPlayer(namehash byte) *Player {
	return &Player{
		Namehash: namehash,
		Position: NewPosition(0, 0, 0),
	}
}

func (p *Player) SetLoginInformation(memoryType byte, clientId uint32, username string, password string) {
	p.MemoryType = memoryType
	p.ClientId = clientId
	p.Username = username
	p.Password = password
}
