package seClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
	"fmt"
)

type Client struct {
	endpoint   string
	bot_key    string
	bot_secret string
	client     *http.Client
}

func New(endpoint string, optional_bot_key string, optional_bot_secret string) *Client {
	if endpoint[len(endpoint)-2:len(endpoint)-1] != "/" {
		endpoint += "/"
	}
	client := Client{
		endpoint:   endpoint,
		bot_key:    optional_bot_key,
		bot_secret: optional_bot_secret,
		client:     &http.Client{},
	}

	return &client
}

/*
 *	Public: Returns all markets
 */
func (self *Client) ListMarkets() ([]*Market, error) {
	data, err := self.requestPublic("markets", 0)
	if err != nil {
		return nil, err
	}
	markets := []*Market{}
	err = json.Unmarshal(data, &markets)
	return markets, err
}

/*
 *	Public: Returns latest trade history of a single market
 */
func (self *Client) ListMarketHistory(market_id int64, limit int) (*Market_trade, error) {
	data, err := self.requestPublic(fmt.Sprintf("trades/%d", market_id), limit)
	if err != nil {
		return nil, err
	}
	trades := &Market_trade{}
	err = json.Unmarshal(data, &trades)
	return trades, err
}

/*
 *	Public: Returns all orders in the market
 */
func (self *Client) ListMarketOrders(market_id int64, limit int) (*Market_data, error) {
	data, err := self.requestPublic(fmt.Sprintf("market/%d", market_id), limit)
	if err != nil {
		return nil, err
	}
	market := &Market_data{}
	err = json.Unmarshal(data, &market)
	return market, err
}

/*
 *	Private: Returns all balances of the user
 */
func (self *Client) ListBalances() (map[string]*Balance, error) {
	data, err := self.requestPrivate("balances", "")
	if err != nil {
		return nil, err
	}
	balances := map[string]*Balance{}
	err = json.Unmarshal(data, &balances)
	return balances, err
}

/*
 *	Private: Returns balance for a single coin
 */
func (self *Client) GetBalance(coin_code string) (*Balance, error) {
	data, err := self.requestPrivate("balance", coin_code)
	if err != nil {
		return nil, err
	}
	balances := &Balance{}
	err = json.Unmarshal(data, &balances)
	return balances, err
}

/*
 *	Private: Returns all active orders of a user
 */
func (self *Client) ListOrders(market_id int) ([]*Order, error) {
	data, err := self.requestPrivate("orders", market_id)
	if err != nil {
		return nil, err
	}
	orders := []*Order{}
	err = json.Unmarshal(data, &orders)
	return orders, err
}

/*
 *	Private: Returns all withdraw addresses
 */
func (self *Client) ListWithdrawAddresses(coin_code string) ([]*Address_withdraw, error) {
	data, err := self.requestPrivate("withdraw/address", coin_code)
	if err != nil {
		return nil, err
	}
	addresses := []*Address_withdraw{}
	err = json.Unmarshal(data, &addresses)
	return addresses, err
}

/*
 *	Private: Withdraw to a saved address
 */
func (self *Client) Withdraw(address_id int64, amount uint64) (bool, error) {
	_, err := self.requestPrivate("withdraw", struct {
		Address_id int64  `json:"address_id"`
		Amount     uint64 `json:"amount"`
	}{
		Amount:     amount,
		Address_id: address_id,
	})
	if err != nil {
		return false, err
	}

	return true, err
}

/*
 *	Private: Get a deposit address of a coin
 */
func (self *Client) Deposit(coin_code string) (*Address_deposit, error) {
	data, err := self.requestPrivate("deposit", coin_code)
	if err != nil {
		return nil, err
	}
	address := &Address_deposit{}
	err = json.Unmarshal(data, &address)
	return address, err
}

/*
 *	Private: Creates a new order
 */
func (self *Client) CreateOrder(market_id int64, price int64, amount int64, order_type int) (bool, error) {
	_, err := self.requestPrivate("orders/create", struct {
		Market_id  int64 `json:"market_id"`
		Price      int64 `json:"price"`
		Amount     int64 `json:"amount"`
		Order_type int   `json:"order_type"`
	}{
		Market_id:  market_id,
		Price:      price,
		Amount:     amount,
		Order_type: order_type,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

/*
 *	Private: Cancels an order by id
 */
func (self *Client) CancelOrder(order_id int64) (bool, error) {
	_, err := self.requestPrivate("orders/cancel", order_id)
	if err != nil {
		return false, err
	}

	return true, nil
}

//public requests
func (self *Client) requestPublic(path string, limit int) ([]byte, error) {
	limitTxt := ""
	if limit > 0{
		limitTxt = fmt.Sprintf("?limit=%d", limit)
	}
	req, err := http.Get(self.endpoint+path + limitTxt)
	// Process response
	if err != nil {
		return []byte(nil), err
	}
	defer req.Body.Close()

	b, err := ioutil.ReadAll(req.Body)
	if err != nil{
		return []byte(nil), err
	}

	if req.StatusCode >= 400 {
		return []byte(nil), errors.New(string(b))
	}

	return b, nil
}

//private requests
func (self *Client) requestPrivate(path string, data interface{}) ([]byte, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return []byte(nil), err
	}
	payload_encrypted, err := Encrypt(payload, self.bot_secret)

	if err != nil {
		return []byte(nil), err
	}

	encoded_str := base64.StdEncoding.EncodeToString(payload_encrypted)
	final_payload, _ := json.Marshal(map[string]interface{}{"payload": encoded_str})

	req, err := http.NewRequest("POST", self.endpoint+path, bytes.NewBuffer(final_payload))
	if err != nil {
		return []byte(nil), err
	}

	req.Header.Set("Content-Type", "application/json")
	if self.bot_key != "" {
		req.Header.Set("KEY", self.bot_key)
	}

	resp, err := self.client.Do(req)
	if err != nil {
		return []byte(nil), err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte(nil), err
	}

	if resp.StatusCode >= 400 {
		return []byte(nil), errors.New(string(b))
	}

	return b, nil
}
