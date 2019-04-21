package main

// Container for message and client
type MCMessage struct {
	sender  *Client
	message []byte
}
