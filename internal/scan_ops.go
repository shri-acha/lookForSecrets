package internal
import (
	"github.com/shri-acha/lookForSecrets.git/config"
	"github.com/shri-acha/lookForSecrets.git/utils"
	"github.com/go-resty/resty/v2"
	"encoding/json"
	"fmt"
	"os"
)


func ScanKeywordMatch(cfg *config.InputConfig){
	
	githubPAT := os.Getenv("GITHUB_PAT")

	key,err := utils.GetCSVValueAtIndex(cfg.FilePath,cfg.ScanIdx)

	if err != nil {
	fmt.Println(err)
	}

	client := resty.New()

	// config := config.EmailConfig{
	// 	SMTPHost: "smtp.gmail.com",
	// 	SMTPPort: "587",
	// 	Username: "happy.irhs@gmail.com",
	// 	Password: os.Getenv("GMAIL_APP_PASSWORD"), // Use App Password, not regular password
	// }
	
	// // Create email message
	// msg := config.EmailMessage{
	// 	From:    "happy.irhs@gmail.com",
	// 	To:      []string{"happy.irhs@gmail.com"},
	// 	Subject: "Leaked Key Found!",
	// 	Body:    "A key has been detected in public repositories",
	// }

	res, err := client.R().
	SetHeader("Authorization",fmt.Sprintf(" token %s",githubPAT)).	
	SetHeader("Accept","application/vnd.github.v3.text-match+json").
	Get(fmt.Sprintf("https://api.github.com/search/code?q=%s",key))

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

	for _,raw_item := range items{
		item := raw_item.(map[string]interface{})
		fmt.Println(fmt.Sprintf("||| Source: %s \r\n||| Path: %s \r\n||| Snippet: %s \r\n||| Confidence: %.2f", 
			item["html_url"].(string), 
			item["path"].(string), 
			item["text_matches"].([]interface{})[0].(map[string]interface{})["fragment"].(string), 
			item["score"].(float64)))


	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}
	
	fmt.Println("Email sent successfully!")
	}

}
