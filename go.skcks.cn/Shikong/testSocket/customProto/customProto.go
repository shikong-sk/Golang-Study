package customProto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	length := new(int32)
	err := binary.Read(lengthBuff, binary.LittleEndian, length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < *length+4 {
		return "", err
	}

	// 读取真实数据
	packet := make([]byte, int(*length+4))
	_, err = reader.Read(packet)
	if err != nil {
		return "", err
	}
	return string(packet[4:]), nil
}
