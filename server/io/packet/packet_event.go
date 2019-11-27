package packet

type PacketEvent interface {
	EventCode() int32
}
