package main

import (
	"fmt"

	"github.com/nahueldev23/composition/pkg/customer"
	"github.com/nahueldev23/composition/pkg/invoice"
	"github.com/nahueldev23/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"colombia",
		"Popaya",
		customer.New("Nahuel", "cl 1111", "1234567"),
		invoiceitem.NewItems(

			invoiceitem.New(1, "curso de go", 12.34),
			invoiceitem.New(2, "curso de js", 12.34),
			invoiceitem.New(3, "curso de react", 12.34),
		),
	)

	i.SetTotal()

	fmt.Printf("%+v", i)

}
