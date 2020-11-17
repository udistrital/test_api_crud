package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaUsuarioRol_20201117_160547 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaUsuarioRol_20201117_160547{}
	m.Created = "20201117_160547"

	migration.Register("CrearTablaUsuarioRol_20201117_160547", m)
}

// Run the migrations
func (m *CrearTablaUsuarioRol_20201117_160547) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CrearTablaUsuarioRol_20201117_160547) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
