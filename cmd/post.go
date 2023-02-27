package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"slk/config"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Send a message to Slack",
	Long:  `Send a message to a Slack channel or user.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		channel := config.GetChannel()

		channelFlg, _ := cmd.Flags().GetString("channel")

		if channelFlg != "" {
			channel = channelFlg
		}

		api := slack.New(token)

		message, _ := cmd.Flags().GetString("message")

		if message == "" {
			file, err := ioutil.TempFile("", "message.*.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer os.Remove(file.Name())

			cmd := exec.Command("vi", file.Name())
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			message = string(content)
		}

		_, _, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
		if err != nil {
			fmt.Printf("Error sending message: %s", err)
			return
		}
		fmt.Println("Message sent successfully!")
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringP("message", "m", "", "The message to send to Slack")
	postCmd.Flags().StringP("channel", "c", "", "The channel to send the message to")
}
