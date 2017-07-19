package classfile

type ClassReader struct {
	data []byte
}

//读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}