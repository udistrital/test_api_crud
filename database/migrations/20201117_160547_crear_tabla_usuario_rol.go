package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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
	file, err := ioutil.ReadFile("../scripts/20201117_160547_crear_tabla_usuario_rol_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablaUsuarioRol_20201117_160547) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
