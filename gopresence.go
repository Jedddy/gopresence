package gopresence


func New(clientId string) (*Client, error) {
	c := &Client{}

	if err := c.Connect(clientId); err != nil {
		return nil, err
	}

	return c, nil
}