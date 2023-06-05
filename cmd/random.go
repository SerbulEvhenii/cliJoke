/*
Copyright © 2023 Serbul Yevhenii <serbulwork@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke",
	Long:  `This command fetches a random joke from the icanhazdadjoke API`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.PersistentFlags().String("term", "", "A search term for a joke.")
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseByte := getJokeData(url)
	joke := Joke{}
	err := json.Unmarshal(responseByte, &joke)
	if err != nil {
		fmt.Printf("Could note unmarshal responseByte. %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	if err != nil {
		log.Printf("Could not request a joke. %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "cliJoke (https://github.com/SerbulEvhenii/cliJoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}
	return responseByte
}

func getRandomJokeWithTerm(jokeTerm string) {
	log.Printf("You searched for a joke with the term: %v", jokeTerm)
}
