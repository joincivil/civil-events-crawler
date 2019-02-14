package eth

// Service provides tools that help with interacting with the Ethereum blockchain
type Service struct {
	TxListener *TxListener
}

// NewService creates a new Service instance
func NewService(txListener *TxListener) *Service {
	return &Service{TxListener: txListener}
}
