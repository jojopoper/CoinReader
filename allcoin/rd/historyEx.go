package rd

// rdHistory : readout histroy datas from allcoin.com, datas saved in History datas
func (ths *ReaderEx) rdHistory() bool {
	ths.historyClt.SetClient(ths.cltCle.Next().Get())

	return ths.Reader.rdHistory()
}
