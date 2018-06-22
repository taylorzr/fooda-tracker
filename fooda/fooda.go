package fooda

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/taylorzr/fooda-tracker/db"
)

var hipchatToken string

func init() {
	hipchatToken = os.Getenv("HIPCHAT_TOKEN")
}

func Remind() ([]string, error) {
	emails, err := db.GetUsers()

	if err != nil {
		return nil, err
	}

	fmt.Printf("All emails: %s\n", emails)

	orderers, err := db.GetTodaysOrders()

	fmt.Printf("All orderers: %s\n", orderers)

	if err != nil {
		return nil, err
	}

	forgetters := []string{}

	for _, email := range emails {
		found := false

		fmt.Printf("Email: %s\n", email)

		for _, orderer := range orderers {
			fmt.Printf("Orderer: %s\n", orderer)
			if email == orderer {
				found = true
			}
		}

		if !found {
			forgetters = append(forgetters, email)
		}
	}

	url := fmt.Sprintf("http://api.hipchat.com/v2/room/test/notification?auth_token=%s", hipchatToken)

	var json = []byte(fmt.Sprintf(`{"color": "purple", "message_format": "text", "message":"You forgot to order fooda!\n%s"}`, forgetters))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println("almost done")

	return forgetters, nil
}
