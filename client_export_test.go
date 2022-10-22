package postcodesio

// NewTestClient encapsulates New and makes possible to create a Client with a custom API URL for testing purposes.
func NewTestClient(url string, opts ...ClientOption) *Client {
	c := New(opts...)
	c.baseURL = url

	return c
}
