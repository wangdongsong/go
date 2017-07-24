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


