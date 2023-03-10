package invoice

import (
	"github.com/nahueldev23/composition/pkg/customer"
	"github.com/nahueldev23/composition/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// return a  new invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {

	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
 func (i *Invoice) SetTotal() {
   i.total = i.items.Total()
 }
