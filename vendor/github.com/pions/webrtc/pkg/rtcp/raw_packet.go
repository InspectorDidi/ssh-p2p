package rtcp

// RawPacket represents an unparsed RTCP packet. It's returned by Unmarshal when
// a packet with an unknown type is encountered.
type RawPacket []byte

// Marshal encodes the packet in binary.
func (r RawPacket) Marshal() ([]byte, error) {
	return r, nil
}

// Unmarshal decodes the packet from binary.
func (r *RawPacket) Unmarshal(b []byte) error {
	if len(b) < (headerLength) {
		return errPacketTooShort
	}
	*r = b

	var h Header
	if err := h.Unmarshal(b); err != nil {
		return err
	}

	return nil
}

// Header returns the Header associated with this packet.
func (r RawPacket) Header() Header {
	var h Header
	if err := h.Unmarshal(r); err != nil {
		return Header{}
	}
	return h
}

// DestinationSSRC returns an array of SSRC values that this packet refers to.
func (r *RawPacket) DestinationSSRC() []uint32 {
	return []uint32{}
}
