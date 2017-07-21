package classfile

import "fmt"

//类文件结构
type ClassFile struct {
	magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*Memberinfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error){
	defer func(){
		if r := recover(); r != nil{
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckmagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

//getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

//getter
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName((self.thisClass))
}

func (self *ClassFile) SuperClassName()  string {
	if self.superClass > 0{
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}
