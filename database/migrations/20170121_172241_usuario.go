package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Usuario_20170121_172241 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Usuario_20170121_172241{}
	m.Created = "20170121_172241"
	migration.Register("Usuario_20170121_172241", m)
}

// Run the migrations
func (m *Usuario_20170121_172241) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE usuario(`id` int(11) NOT NULL AUTO_INCREMENT,`nome` varchar(128) NOT NULL,`idade` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Usuario_20170121_172241) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `usuario`")
}
