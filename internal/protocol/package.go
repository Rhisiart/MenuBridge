package protocol

import (
	"encoding/binary"
	"fmt"
)

const HEADER_SIZE = 4
const VERSION = 1

type Package struct {
	Command byte
	Data    []byte
}

// Marshal serializes the Package into a byte slice for transmission or storage.
// It follows a specific format by prepending the header with version, command,
// and data length information.
//
// The format of the output byte slice is:
// [VERSION | COMMAND | DATA_LENGTH (2 bytes) | DATA]
//
// Returns:
// - []byte: A byte slice representing the serialized package, with the structure:
//   - VERSION (1 byte): The version of the protocol or format.
//   - COMMAND (1 byte): A command identifier indicating the action or type of data.
//   - DATA_LENGTH (2 bytes): The length of the DATA segment in big-endian format.
//   - DATA (variable length): The actual data payload.
//
// - error: An error object, which is always nil in this implementation.
//
// The function uses big-endian encoding to store the length of the DATA segment.
func (p *Package) Marshal() ([]byte, error) {
	ulen := uint16(len(p.Data))

	dataLength := make([]byte, 2)
	pack := make([]byte, 0, HEADER_SIZE+ulen)

	binary.BigEndian.PutUint16(dataLength, ulen)

	pack = append(pack, VERSION)
	pack = append(pack, p.Command)
	pack = append(pack, dataLength...)
	pack = append(pack, p.Data...)

	return pack, nil
}

// Unmarshal deserializes a byte slice into a Package object.
//
// Parameters:
//   - bytes []byte: The byte slice containing the serialized package data.
//     The expected format is:
//     [VERSION | COMMAND | DATA_LENGTH (2 bytes) | DATA]
//
// Returns:
// - error: An error if the byte slice cannot be deserialized correctly. Possible errors include:
//   - Version mismatch error if the VERSION in the byte slice does not match the expected VERSION.
//   - Length validation error if the length of the provided byte slice is insufficient to contain the data.
//
// The function performs the following checks and operations:
//
//   - It compares the first byte with the expected VERSION. If they do not match, it returns an error.
//   - It extracts the expected length of the data from the DATA_LENGTH header segment.
//   - It checks if the actual length of the package is equal or greater when header size plus the DATA_LENGTH.
//   - It sets the Command and Data fields of the Package struct with the values extracted from the byte slice.
func (p *Package) Unmarshal(bytes []byte) error {
	if VERSION != bytes[0] {
		return fmt.Errorf("unable to unmarshal package, because the versions doesnt match")
	}

	validLength := int(binary.BigEndian.Uint16(bytes[2:4]))
	end := HEADER_SIZE + validLength
	packageLength := len(bytes)

	if packageLength < end {
		return fmt.Errorf("invalid length of data, the data has %d should be equal or above %d", packageLength, end)
	}

	p.Command = bytes[1]
	p.Data = bytes[HEADER_SIZE:end]

	return nil
}
