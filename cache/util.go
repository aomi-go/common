package cache

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
)

// Serialize 将对象序列化为字节数组
func Serialize(value interface{}) (bytes.Buffer, error) {
	enc := msgpack.GetEncoder()
	var buf bytes.Buffer
	enc.Reset(&buf)
	err := enc.Encode(value)
	return buf, err
}

// Deserialize 将字节数组反序列化为对象
func Deserialize(data []byte, value interface{}) error {
	return msgpack.Unmarshal(data, value)
}

func DeserializeWithStr(data string, value interface{}) error {
	buf := bytes.NewBufferString(data)
	return Deserialize(buf.Bytes(), value)
}
