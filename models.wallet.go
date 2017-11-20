package main

type wallet struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	TotalItems int    `json:"totalItems"`
	From       string `json:"from"`
	To         string `json:"to"`
	Items      []struct {
		Txid     string `json:"txid"`
		Version  int    `json:"version"`
		Locktime int    `json:"locktime"`
		Vin      []struct {
			Txid      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig struct {
				asm string `json:"hex"`
				hex string `json:"asm"`
			} `json:"scriptSig"`
			Sequence        int     `json:"sequence"`
			N               int     `json:"n"`
			Addr            string  `json:"string"`
			ValueSat        float32 `json:"valueSat"`
			Value           float32 `json:"value"`
			DoubleSpentTxID string  `json:"doubleSpentTxID"`
		} `json:"vin"`
		Vout []struct {
			Value        int `json:"value"`
			N            int `json:"n"`
			ScriptPubkey struct {
				Hex       string   `json:"hex"`
				Asm       string   `json:"asm"`
				Addresses []string `json:"adresses"`
				Keytype   string   `json:"type"`
			} `json:"scriptPubKey"`
			SpentTxId   string `json:"spentTxId"`
			SpentIndex  int    `json:"spentIndex"`
			SpentHeight int    `json:"spentHeight"`
		} `json:"vout"`
		Blockhash     string  `json:"blockhash"`
		Blockheight   int     `json:"blockheight"`
		Confirmations int     `json:"confirmations"`
		Time          int     `json:"time"`
		Blocktime     int     `json:"blocktime"`
		ValueOut      float32 `json:"valueOut"`
		Size          int     `json:"size"`
		ValueIn       float32 `json:"valueIn"`
		Fees          float32 `json:"fees"`
	} `json:"items"`
}
