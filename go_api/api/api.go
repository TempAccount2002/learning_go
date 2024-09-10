package api

// Endpoint paramaters/response structs

// Paramters that our API endpoint will take
type CoinBalanceParams struct {
	Username string
}

// Represents the successful response from the server containing the coin balance and the success code
type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

// Error struct representing the response when the server encounters and error processing the request
type Error struct {
	Code    int
	Message string
}
