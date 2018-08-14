package cmd

type Config struct {
	TemplatePath string      `mapstructure:"template_path"`
	Slack        slackConfig `mapstructure:"slack"`
}

type slackConfig struct {
	WebhookURL string `mapstructure:"url"`
}
