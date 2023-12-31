package producer

type IProductProducer interface {
	Publish(value []byte)
}

type IInventoryProducer interface {
	Publish(value []byte)
}
