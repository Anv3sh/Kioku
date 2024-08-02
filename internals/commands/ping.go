package commands


func PingCommand(args []string)[]byte{
	return []byte("PONG!\n")
}