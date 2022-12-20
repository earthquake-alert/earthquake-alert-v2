package src_test

import (
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

func TestInitPublishConfig(t *testing.T) {
	t.Run("正しく読み取れる", func(t *testing.T) {
		publishConfig := src.C.PublishConfig

		require.Len(t, publishConfig.Twitter, 2)
		require.Len(t, publishConfig.Discord, 3)
		require.Len(t, publishConfig.Slack, 1)
		require.Len(t, publishConfig.LineNotify, 1)

		t.Run("twitter1", func(t *testing.T) {
			twitter := publishConfig.Twitter[0]

			require.Equal(t, twitter.ClientName, "twitter1")
			require.Equal(t, twitter.Token, "twitter-token")

			require.Len(t, twitter.Areas, 0)
			require.Equal(t, twitter.IsTsunami, false)
			require.Equal(t, twitter.MinInt, jma.IntUnknown)
		})

		t.Run("twitter2", func(t *testing.T) {
			twitter := publishConfig.Twitter[1]

			require.Equal(t, twitter.ClientName, "twitter2")
			require.Equal(t, twitter.Token, "twitter-token-2")

			require.Len(t, twitter.Areas, 0)
			require.Equal(t, twitter.IsTsunami, false)
			require.Equal(t, twitter.MinInt, jma.Int5l)
		})

		t.Run("discord1", func(t *testing.T) {
			discord := publishConfig.Discord[0]

			require.Equal(t, discord.ClientName, "discord1")
			require.Equal(t, discord.WebhookURL, "https://example.com/discord/webhook/1")

			require.Equal(t, discord.Areas, []string{"茨城県", "埼玉県", "東京都"})
			require.Equal(t, discord.IsTsunami, true)
			require.Equal(t, discord.MinInt, jma.Int4)
		})

		t.Run("discord2", func(t *testing.T) {
			discord := publishConfig.Discord[1]

			require.Equal(t, discord.ClientName, "discord2")
			require.Equal(t, discord.WebhookURL, "https://example.com/discord/webhook/2")

			require.Len(t, discord.Areas, 0)
			require.Equal(t, discord.IsTsunami, false)
			require.Equal(t, discord.MinInt, jma.IntUnknown)
		})

		t.Run("discord3", func(t *testing.T) {
			discord := publishConfig.Discord[2]

			require.Equal(t, discord.ClientName, "discord3")
			require.Equal(t, discord.WebhookURL, "https://example.com/discord/webhook/3")

			require.Len(t, discord.Areas, 0)
			require.Equal(t, discord.IsTsunami, false)
			require.Equal(t, discord.MinInt, jma.IntUnknown)
		})

		t.Run("slack1", func(t *testing.T) {
			slack := publishConfig.Slack[0]

			require.Equal(t, slack.ClientName, "slack1")
			require.Equal(t, slack.WebhookURL, "https://example.com/slack/webhook/1")
			require.Equal(t, slack.Channel, "general")

			require.Len(t, slack.Areas, 0)
			require.Equal(t, slack.IsTsunami, false)
			require.Equal(t, slack.MinInt, jma.IntUnknown)
		})

		t.Run("line1", func(t *testing.T) {
			line := publishConfig.LineNotify[0]

			require.Equal(t, line.ClientName, "line1")
			require.Equal(t, line.Token, "line-token-1")

			require.Len(t, line.Areas, 0)
			require.Equal(t, line.IsTsunami, false)
			require.Equal(t, line.MinInt, jma.IntUnknown)
		})
	})
}
