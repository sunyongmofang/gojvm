package classfile

type ConstantNameAndTypeInfo struct {
	namdIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.namdIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
