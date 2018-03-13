package dao

type DatabaseHandler struct {
}

func (handler *DatabaseHandler) Connect(connStr string, maxConnections int) error {

	return nil
}

func (handler *DatabaseHandler) Disconnect() error {
	return nil
}

func (handler *DatabaseHandler) DoInitialSetup() error {

	return nil

}
