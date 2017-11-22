package main

type BTC struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	TotalItems int    `json:"totalItems"`
	From       string `json:"from"`
	To         string `json:"to"`
	Items      []item `json:"items"`
}

type item struct {
Txid     string `json:"txid"`
Version  int    `json:"version"`
Locktime int    `json:"locktime"`
Vin      []vin `json:"vin"`
Vout 	 []vout `json:"vout"`
Blockhash     string  `json:"blockhash"`
Blockheight   int     `json:"blockheight"`
Confirmations int     `json:"confirmations"`
Time          int     `json:"time"`
Blocktime     int     `json:"blocktime"`
ValueOut      float32 `json:"valueOut"`
Size          int     `json:"size"`
ValueIn       float32 `json:"valueIn"`
Fees          float32 `json:"fees"`
}

type vin struct {
	Txid      string `json:"txid"`
	Vout      int    `json:"vout"`
	ScriptSig []scriptSig `json:"scriptSig"`
	Sequence        int     `json:"sequence"`
	N               int     `json:"n"`
	Addr            string  `json:"string"`
	ValueSat        float32 `json:"valueSat"`
	Value           float32 `json:"value"`
	DoubleSpentTxID string  `json:"doubleSpentTxID"`
}

type scriptSig struct {
	asm string `json:"hex"`
	hex string `json:"asm"`
}

type vout struct {}
	Value        int `json:"value"`
	N            int `json:"n"`
	ScriptPubkey struct `json:"scriptPubKey"`
	SpentTxId   string `json:"spentTxId"`
	SpentIndex  int    `json:"spentIndex"`
	SpentHeight int    `json:"spentHeight"`
}

type scriptPubkey struct {
	Hex       string   `json:"hex"`
	Asm       string   `json:"asm"`
	Addresses []string `json:"adresses"`
	Keytype   string   `json:"type"`
}