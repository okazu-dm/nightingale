package cmd

type Config struct {
	TemplatePath string `mapstructure:"template_path"`
	Type         string
	Slack        slackConfig `mapstructure:"slack"`
}

type slackConfig struct {
	WebhookURL string `mapstructure:"url"`
}
