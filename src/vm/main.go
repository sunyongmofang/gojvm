package main

import (
	"fmt"
	"strings"
	"vm/classfile"
	"vm/classpath"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)
	fmt.Printf(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cp *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cp.MajorVersion(), cp.MinorVersion())
	fmt.Printf("constants count:%v\n", len(cp.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cp.AccessFlags())
	fmt.Printf("this class:%v\n", cp.SuperClassName())
	fmt.Printf("interfaces: %v\n", cp.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cp.Fields()))
	for _, f := range cp.Fields() {
		fmt.Printf("   %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cp.Methods()))
	for _, m := range cp.Methods() {
		fmt.Printf("   %s\n", m.Name())
	}
}
