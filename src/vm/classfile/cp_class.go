package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *ConstantClassInfo) Cp() ConstantPool { return self.cp }
func (self *ConstantClassInfo) NameIndex() uint16 { return self.nameIndex }
