package config 

type InputConfig struct {
	FilePath string   `clap:"--input,-I"`
	Scan bool `clap:"--scan,-S"`
	ScanIdx int `clap: "--scan-index,-X"`
}
// email config
type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}

// email to be sent
type EmailMessage struct {
	From    string
	To      []string
	Subject string
	Body    string
}


