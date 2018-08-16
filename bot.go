package seClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
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
 *	Returns all balances of the user
 */
func (self *Client) ListBalances() (map[string]*Balance, error) {
	data, err := self.request("balances", "")
	if err != nil {
		return nil, err
	}
	balances := map[string]*Balance{}
	err = json.Unmarshal(data, &balances)
	return balances, err
}

/*
 *	Returns balance for a single coin
 */
func (self *Client) GetBalance(coin_code string) (*Balance, error) {
	data, err := self.request("balance", coin_code)
	if err != nil {
		return nil, err
	}
	balances := &Balance{}
	err = json.Unmarshal(data, &balances)
	return balances, err
}

/*
 *	Returns all active orders of a user
 */
func (self *Client) ListOrders(market_id int) ([]*Order, error) {
	data, err := self.request("orders", market_id)
	if err != nil {
		return nil, err
	}
	orders := []*Order{}
	err = json.Unmarshal(data, &orders)
	return orders, err
}

/*
 *	Returns all withdraw addresses
 */
func (self *Client) ListWithdrawAddresses(coin_code string) ([]*Address_withdraw, error) {
	data, err := self.request("withdraw/address", coin_code)
	if err != nil {
		return nil, err
	}
	addresses := []*Address_withdraw{}
	err = json.Unmarshal(data, &addresses)
	return addresses, err
}

/*
 *	Withdraw to a saved address
 */
func (self *Client) Withdraw(address_id int64, amount uint64) (bool, error) {
	_, err := self.request("withdraw", struct {
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
 *	Get a deposit address of a coin
 */
func (self *Client) Deposit(coin_code string) (*Address_deposit, error) {
	data, err := self.request("deposit", coin_code)
	if err != nil {
		return nil, err
	}
	address := &Address_deposit{}
	err = json.Unmarshal(data, &address)
	return address, err
}

/*
 *	Creates a new order
 */
func (self *Client) CreateOrder(market_id int64, price int64, amount int64, order_type int) (bool, error) {
	_, err := self.request("orders/create", struct {
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
 *	Cancels an order by id
 */
func (self *Client) CancelOrder(order_id int64) (bool, error) {
	_, err := self.request("orders/cancel", order_id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (self *Client) request(path string, data interface{}) ([]byte, error) {
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
