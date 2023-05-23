package chip

import (
	"log"
)

type ROM struct {
	Basic
	Under *Basic
}

func NewROM(name string, size int, file string) *ROM {
	var tmp ROM

	tmp.Name = name
	tmp.Buff = make([]byte, size)
	tmp.Under = nil
	tmp.LoadData(file, 0)

	// fmt.Printf("Create new ROM %4s with size %d\n", name, size)
	return &tmp
}

func (r *ROM) SetUnderChip(chip *Basic) {
	r.Under = chip
}

func (r *ROM) Read(addr uint16) byte {
	// fmt.Printf("Reading $%02X from %4s at $%04X\n", r.Buff[addr], r.Name, addr)
	return r.Buff[addr]
}

func (r *ROM) Write(addr uint16, data byte) {
	// fmt.Printf("Writting $%02X to %4s at $%04X\n", data, r.Name, addr)
	if r.Under == nil {
		log.Fatal("No RAM under ROM")
	}
	r.Under.Write(addr, data)
}

func (r *ROM) ReadOnly() bool {
	return true
}
