package aex

// rdHistory : readout histroy datas from aex.com, datas saved in History datas
func (ths *ReaderEx) rdHistory() bool {
	ths.historyClt.SetClient(ths.Next().Get())

	return ths.Reader.rdHistory()
}
