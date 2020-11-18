package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPais_20201118_093359 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPais_20201118_093359{}
	m.Created = "20201118_093359"

	migration.Register("CrearTablaPais_20201118_093359", m)
}

// Run the migrations
func (m *CrearTablaPais_20201118_093359) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE  public.pais(id serial NOT NULL,nombre varchar NOT NULL)")
}

// Reverse the migrations
func (m *CrearTablaPais_20201118_093359) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS public.pais")
}
