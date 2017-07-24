package classfile

type MemberInfo struct{
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo{
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i:= range members{
		members[i] = readMembers(reader, cp)
	}
	return members
}
