package main

// APICaller interface is for making an API call and returning the result
type APICaller interface {
	GetBytes(api string) ([]byte, error)
}
