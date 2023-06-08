package controller

import (
	"fmt"

	"github.com/pavlelee/gopls-references/models"
)

func CreateVM() error {
	vm := &models.VM{
		ID:   0,
		Name: "test",
	}

	fmt.Printf("vm: %+v\n", vm)

	return nil
}
