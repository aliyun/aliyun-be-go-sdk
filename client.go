package be

type Client struct {
	Endpoint string
	UserName string
	PassWord string
}

func (c *Client) read(request ReadRequest) (*Response, error)  {
	return nil, nil
}

func (c *Client) write(request WriteRequest) (*Response, error)  {
	return nil, nil
}
