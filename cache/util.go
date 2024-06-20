package cache

import (
	"bytes"
	"encoding/gob"
)

// Serialize 将对象序列化为字节数组
func Serialize(value interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	return buf, err
}

// Deserialize 将字节数组反序列化为对象
func Deserialize(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}

func DeserializeWithStr(data string, value interface{}) error {
	buf := bytes.NewBufferString(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}
