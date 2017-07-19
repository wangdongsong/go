package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

//读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//使用Uint16读取u2类型数据
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}