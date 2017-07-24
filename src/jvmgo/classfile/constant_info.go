package classfile

const (
	CONSTAT_class = 7
	CONSTAT_Fieldref = 9
	CONSTAT_Methodref = 10
	CONSTAT_InterfaceMethodref = 11
	CONSTAT_String = 8
	CONSTAT_Integer = 3
	CONSTAT_Float = 4
	CONSTAT_Long = 5
	CONSTAT_Double = 6
	CONSTAT_NameAndType = 12
	CONSTAT_Utf8 = 1
	CONSTAT_MethodHandle = 15
	CONSTAT_MethodType = 16
	CONSTAT_InvokeDynamic = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
