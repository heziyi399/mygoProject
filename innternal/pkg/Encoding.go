package pkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	var pkg = new(bytes.Buffer)
	//读取消息长度转化为int32
	var length = int32(len(message))
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//这段代码的目的是将 message 中的数据以小端字节序的方式写入到 pkg 中，[]byte(message)：将 message 转换为字节数组。
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4) //读取前四个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
