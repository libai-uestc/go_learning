package filter

import "libai/go/basic/phase-one/oip/common"

type AddressFilter struct {
	Tag  string
	City string
}

func (self AddressFilter) Name() string {
	return self.Tag
}

func (self AddressFilter) Filter(products []*common.Product) []*common.Product {
	rect := make([]*common.Product, 0, len(products))
	for _, product := range products { // range返回的第一个是序号，不需要使用，所以 _
		if product.ShipAddress == self.City {
			rect = append(rect, product)
		}
	}
	return rect
}
