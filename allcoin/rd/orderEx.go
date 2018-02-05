package rd

func (ths *ReaderEx) rdOrders() bool {
	ths.orderClt.SetClient(ths.cltCle.Next().Get())

	return ths.Reader.rdOrders()
}
