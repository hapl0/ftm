package main

type wallet struct {
	ID         int           `json:"id"`
	Address    string        `json:"address"`
	TotalItems int           `json:"totalItems"`
	From       string        `json:"from"`
	To         string        `json:"to"`
	Items      []transaction `json:"items"`
}

type transaction struct {
	txid          string  `json:"txid"`
	version       int     `json:"version"`
	locktime      int     `json:"locktime"`
	vin           []vin   `json:"vin"`
	vout          []vout  `json:"vout"`
	blockhash     string  `json:"blockhash"`
	blockheight   int     `json:"blockheight"`
	confirmations int     `json:"confirmations"`
	time          int     `json:"time"`
	blocktime     int     `json:"blocktime"`
	valueOut      float32 `json:"valueOut"`
	size          int     `json:"size"`
	valueIn       float32 `json:"valueIn"`
	fees          float32 `json:"fees"`
}

type vout struct {
	value        int          `json:"value"`
	n            int          `json:"n"`
	scriptPubkey scriptPubKey `json:"scriptPubKey"`
	spentTxId    string       `json:"spentTxId"`
	spentIndex   int          `json:"spentIndex"`
	spentHeight  int          `json:"spentHeight"`
}

type scriptPubKey struct {
	hex       string   `json:"hex"`
	asm       string   `json:"asm"`
	addresses []string `json:"adresses"`
	keytype   string   `json:"type"`
}

type vin struct {
	txid            string    `json:"txid"`
	vout            int       `json:"vout"`
	scriptSig       scriptSig `json:"scriptSig"`
	sequence        int       `json:"sequence"`
	n               int       `json:"n"`
	addr            string    `json:"string"`
	valueSat        float32   `json:"valueSat"`
	value           float32   `json:"value"`
	doubleSpentTxID string    `json:"doubleSpentTxID"`
}

type scriptSig struct {
	asm string `json:"hex"`
	hex string `json:"asm"`
}
