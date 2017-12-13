package aex

func (ths *ReaderEx) rdOrders() bool {
	ths.orderClt.SetClient(ths.Next().Get())

	return ths.Reader.rdOrders()
}
