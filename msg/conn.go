package msg

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"ngrok/conn"
)

func readMsgShared(c conn.Conn) (buffer []byte, err error) {

	var sz int64
	err = binary.Read(c, binary.BigEndian, &sz)
	if err != nil {
		return
	}

	if sz < 0 || sz > 102400 {
		err = errors.New("message length exceed the limit")
		return
	}

	buffer = make([]byte, sz)
	n, err := io.ReadFull(c, buffer)

	if err != nil {
		return
	}

	c.Debug("Read message with length: %d msg: %s", sz, buffer)

	if int64(n) != sz {
		err = errors.New(fmt.Sprintf("Expected to read %d bytes, but only read %d", sz, n))
		return
	}

	return
}

func ReadMsg(c conn.Conn) (msg Message, err error) {
	buffer, err := readMsgShared(c)
	if err != nil {
		return
	}

	return Unpack(buffer)
}

func ReadMsgInto(c conn.Conn, msg Message) (err error) {
	buffer, err := readMsgShared(c)
	if err != nil {
		return
	}
	return UnpackInto(buffer, msg)
}

func WriteMsg(c conn.Conn, msg interface{}) (err error) {
	buffer, err := Pack(msg)
	if err != nil {
		return
	}

	c.Debug("Writing message: %s", string(buffer))
	err = binary.Write(c, binary.BigEndian, int64(len(buffer)))

	if err != nil {
		return
	}

	if _, err = c.Write(buffer); err != nil {
		return
	}

	return nil
}
