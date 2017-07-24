package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantPool, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantPool(reader, cp)
		switch cp[i].(type) {
		case *ConstantPool, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
}

func (self ConstantPool) getConstantPool(index uint16) ConstantPool{
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}

	panic("Invalid constant pool index!")

}

func(self ConstantPool) getNameAndType(index uint16) (string, string){
	ntInfo := self.getConstantPool(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClasName(index uint16) string{
	classInfo := self.getConstantPool(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string{
	utf8Info := self.getConstantPool(index).(*ConstantUtf8Info)
	return utf8Info.str
}



