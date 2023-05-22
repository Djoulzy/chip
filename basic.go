package chip

import (
	"io/ioutil"

	"github.com/Djoulzy/mmu"
)

type Basic struct {
	Name string
	Buff []byte
	mmu.ChipAccess
}

func (b *Basic) GetName() string {
	return b.Name
}

func (b *Basic) LoadData(file string, memStart uint16) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	for i, val := range content {
		// mem.Write(memStart+uint16(i), val)
		b.Buff[memStart+uint16(i)] = val
	}
	return nil
}
