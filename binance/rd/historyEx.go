package rd

// rdHistory : readout histroy datas from aex.com, datas saved in History datas
func (ths *ReaderEx) rdHistory() bool {
	ths.historyClt.SetClient(ths.cltCle.Next().Get())

	return ths.Reader.rdHistory()
}
