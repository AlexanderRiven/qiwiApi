package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Sum struct{
	Amount int `json:"amount"`
	Currency string `json:"currency"`
}
type PaymentMethod struct{
	Type string `json:"type"`
	AccountId string `json:"accountId"`
}
type Fields struct{
	Account string `json:"account"`
}

type Message struct {
	ID string `json:"id"`
	Sum Sum `json:"sum"`

	PaymentMethod PaymentMethod `json:"paymentMethod"`
	Comment string `json:"comment"`
	Fields Fields `json:"fields"`
}
type mainAcc struct{
	Accounts accounts `json:"accounts"`
}
type accounts struct{
	Alias string `json:"alias"`
	FsAlias string `json:"fsAlias"`
	BankAlias string `json:"bankAlias"`
	Title string `json:"title"`
	Type Type `json:"type"`
	HasBalance bool `json:"hasBalance"`
	Balance Balance `json:"balance"`
	Currency int `json:"currency"`
	DefaultAccount bool `json:"defaultAccount"`


}
type Type struct{
	Id string `json:"id"`
	Title string `json:"title"`
}
type Balance struct{
	Amount float32 `json:"amount"`
	Currency int `json:"currency"`
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Выберите операцию, 1 - совершить платеж, 2 - проверить баланс, 3 - последние платежи ")
	var userChoose int
	fmt.Scan(&userChoose)

	fmt.Println(userChoose)
	switch userChoose{
	case 1:

		fmt.Print("Введите получателя платежа-> ")
		acco, _ := reader.ReadString('\n')
		fmt.Print("Введите сумму платежа-> ")
		var amou int
		fmt.Scan(&amou)
		now := time.Now()
		dataToSend := Message{
			ID: string(int(now.Unix() * 1000)),
			Sum: Sum{
				Amount: amou,
				Currency: "643",
			},
			PaymentMethod: PaymentMethod{
				Type: "Account",
				AccountId: "643",
			},
			Comment: "kek",
			Fields: Fields{
				Account: acco,
			},
		}

		jsondata, _ := json.Marshal(dataToSend)
		_ = ioutil.WriteFile("qiwi.json", jsondata, 0644)



		var bearer = "Bearer " + "c5aa05b641c47de70714e8639586ac24"
		req, err := http.NewRequest("POST", "https://edge.qiwi.com/sinap/api/v2/terms/99/payments", bytes.NewBuffer(jsondata))

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", bearer)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "Android v3.2.0 MKT")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	case 2:
		var bearer = "Bearer " + "c5ac47de70714e8639586ac24"

		req, err := http.NewRequest("GET","https://edge.qiwi.com/funding-sources/v2/persons/79045150003/accounts", nil)


		req.Header.Set("Authorization", bearer)
		req.Header.Set("Accept", "application/json")


		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)
		data := mainAcc{}
		json.Unmarshal([]byte(s), &data)
		fmt.Println("Operation: ", mainAcc{Accounts: accounts{
			Alias:          "",
			FsAlias:        "",
			BankAlias:      "",
			Title:          "",
			Type:           Type{},
			HasBalance:     false,
			Balance:        Balance{},
			Currency:       0,
			DefaultAccount: false,
		}})
		fmt.Println("response Body:", string(body))

	case 3:
		var bearer = "Bearer " + "c5a1c47de70714e8639586ac24"

		req, err := http.NewRequest("GET","https://edge.qiwi.com/payment-history/v2/persons/79045150003/payments?rows=10", nil)


		req.Header.Set("Authorization", bearer)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")


		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Body:", string(body))


	}




}