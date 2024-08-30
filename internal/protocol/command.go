package protocol

const (
	RESERVATION = iota
	PLACE
)

var commandMap = map[string]byte{
	"reservation": RESERVATION,
	"place":       PLACE,
}

var commandMapLookup = map[byte]string{
	RESERVATION: "reservation",
	PLACE:       "place",
}

type Command struct {
	extensions map[string]byte
	size       byte
}
