package internal
import (
	"github.com/shri-acha/lookForSecrets.git/config"
	"github.com/go-resty/resty/v2"
	"encoding/json"
	"fmt"
	"os"
)


func ScanKeywordMatch(config *config.InputConfig){
	
	githubPAT := os.Getenv("GITHUB_PAT")

	client := resty.New()

	res, err := client.R().
	SetHeader("Authorization",fmt.Sprintf(" token %s",githubPAT)).	
	SetHeader("Accept","application/vnd.github.v3.text-match+json").
	Get("https://api.github.com/search/code?q=Hello")

	if err != nil {
		fmt.Println(err)
	}
	var data map[string] interface{} // response can be manually structured, but, but I won't do it :D. 
	err = json.Unmarshal(res.Body(),&data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SCAN-RESULT")
	items := data["items"].([]interface{}) 
	fmt.Println(items)

	for _,raw_item := range items{
		item := raw_item.(map[string]interface{})
		fmt.Println(fmt.Sprintf("||| Source: %s \r\n||| Path: %s \r\n||| Snippet: %s \r\n||| Confidence: %.2f", 
			item["html_url"].(string), 
			item["path"].(string), 
			item["text_matches"].([]interface{})[0].(map[string]interface{})["fragment"].(string), 
			item["score"].(float64)))
		}

}
