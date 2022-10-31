package _5_abstract_factory

import "fmt"

type MainOrder interface {
	SaveMainOrder()
}

type DetailOrder interface {
	SaveDetailOrder()
}

type DAOFactory interface {
	CreateMainOrder() MainOrder
	CreateDetailOrder() DetailOrder
}

type RDBMainFactory struct{}

func (f *RDBMainFactory) SaveMainOrder() {
	fmt.Println("rdb main save")
}

type RDBDetailOrder struct{}

func (d *RDBDetailOrder) SaveDetailOrder() {
	fmt.Println("rdb detail save")
}

type RDBDAOFactory struct{}

func (r *RDBDAOFactory) CreateMainOrder() MainOrder {
	return &RDBMainFactory{}
}

func (r *RDBDAOFactory) CreateDetailOrder() DetailOrder {
	return &RDBDetailOrder{}
}

type XMLMainFactory struct{}

func (f *XMLMainFactory) SaveMainOrder() {
	fmt.Println("xml main save")
}

type XMLDetailOrder struct{}

func (d *XMLDetailOrder) SaveDetailOrder() {
	fmt.Println("xml detail save")
}

type XMLDAOFactory struct{}

func (x *XMLDAOFactory) CreateMainOrder() MainOrder {
	return &XMLMainFactory{}
}

func (x *XMLDAOFactory) CreateDetailOrder() DetailOrder {
	return &XMLDetailOrder{}
}
