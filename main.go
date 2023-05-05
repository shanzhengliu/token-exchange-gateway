package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

type Response struct {
	AccessToken string `json:"access_token"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// var client = &http.Client{
		// 	Timeout: 5 * time.Second,
		// }

		ctx, oktaClient, err := okta.NewClient(
			context.TODO(),
			okta.WithOrgUrl("https://dev-38565349.okta.com"),
			// okta.WithClientId("Ffw8nsliSv9s4SUgNTem9RkV0jnGeseW"),
			// okta.WithAuthorizationMode("Bearer"),
			okta.WithToken("001uIr659fI5QYddGPDnlEzmLmfVgRl2-VPlvYlEMs"),
		)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		applicationList, resp, err := oktaClient.Application.ListApplications(ctx, nil)
		if err != nil {
			fmt.Printf("Error listing applications: %v\n", err)
		}

		fmt.Printf("ApplicationList: %+v\n Response: %+v\n\n", applicationList, resp)

		fmt.Println("request in the sidecar ----- handler token, and send the request to main")

		// req, err := http.NewRequest(r.Method, "http://localhost:8081"+r.RequestURI, r.Body)

		// req.Header.Add("Authorization", "hello")
		// fmt.Println(fmt.Println("request in the sidecar ----- use new token"))
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// rsps, err := client.Do(req)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer rsps.Body.Close()
		// body, _ := ioutil.ReadAll(rsps.Body)
		// w.Header().Set("Succeed", "application/json")
		// w.Write([]byte(body))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getApi() string {
	url := "https://dev-wyd4tdwx0z8gyxil.us.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"Ffw8nsliSv9s4SUgNTem9RkV0jnGeseW\",\"client_secret\":\"VwdvvZ6CLF_gZJ3noKiAYqDQeI4rW4ADVIu6tJNi6vf5okAIwN1T_TnHZ0s9VOva\",\"audience\":\"https://dev-wyd4tdwx0z8gyxil.us.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response Response
	err := json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	return response.AccessToken
}
