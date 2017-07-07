package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {}

func (self *MemberInfo) AccessFlags() uint16 {}

func (self *MemberInfo) Name() string {}

func (self *MemberInfo) Descriptor() string {}
