package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	/*
		req := &telegramLoginRequest{
			OpenID:    "7067365366",
			FirstName: "zhengyi",
			UserName:  "zhengyi_C",
			Timestamp: 1732336698078,
			Signature: "5f9e25121d1eb3d2e505c6a357707d86",
		}
	*/
	req := &telegramLoginRequest{
		OpenID:    "123",
		FirstName: "A",
		UserName:  "ABC",
		Timestamp: 1732284099,
		Signature: "5f9e25121d1eb3d2e505c6a357707d86",
	}

	key := "n&KKRSub#4z@RsRc"
	sign := generateSignature(req, key)
	fmt.Println(sign)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	//b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	//b, err := bot.New("7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58", opts...) // test pid bot
	b, err := bot.New("7457971759:AAEYDs74IUlcFhaCeGFj7mm6ue8GU30JGU0", opts...) // luck bot
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, helloHandler)

	b.Start(ctx)
}

type telegramLoginRequest struct {
	OpenID    string `json:"open_id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	UserName  string `json:"user_name" validate:"required"`
	Timestamp int64  `json:"timestamp" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}

func generateSignature(req *telegramLoginRequest, secretKey string) string {
	// Collect fields into a map
	fieldMap := map[string]string{
		"open_id":    req.OpenID,
		"first_name": req.FirstName,
		"user_name":  req.UserName,
		"timestamp":  strconv.FormatInt(req.Timestamp, 10),
	}

	// Sort the keys alphabetically
	keys := make([]string, 0, len(fieldMap))
	for key := range fieldMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Concatenate the fields in sorted order
	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(fieldMap[key])
	}
	builder.WriteString(secretKey) // Append the secret key

	// Compute the MD5 hash
	hash := md5.Sum([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	/*
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
			ParseMode: models.ParseModeMarkdown,
		})
	*/
	/*
		inlineKeyboard := [][]models.InlineKeyboardButton{
			{
				models.InlineKeyboardButton{
					Text: "Button 1",
				},
				models.InlineKeyboardButton{
					Text: "Button 2",
				},
			},
		}
	*/
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Open Web App", WebApp: &models.WebAppInfo{URL: "https://test-pid-tg.polyflow.tech/"}},
				{Text: "Join PID community", CallbackData: "button_2", URL: "https://t.me/pid_channel"},
			},
			{
				{Text: "Button 3", CallbackData: "button_3", URL: "https://t.me/pid_channel"},
			},
		},
	}
	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		//Photo:   &models.InputFileString{Data: "AgACAgIAAxkDAAIBOWJimnCJHQJiJ4P3aasQCPNyo6mlAALDuzEbcD0YSxzjB-vmkZ6BAQADAgADbQADJAQ"},
		Photo:       &models.InputFileString{Data: "https://test-pid-tg.polyflow.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2FhomeCat.e5cffe34.png&w=384&q=75"},
		Caption:     "Preloaded Facebook logo\nluck test",
		ReplyMarkup: kb,
	})
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Say /hello",
	})
}

/*
package main

import "github.com/NicoNex/echotron/v3"

const token = "7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58"

func main() {
	api := echotron.NewAPI(token)

	for u := range echotron.PollingUpdates(token) {
		if u.Message.Text == "/start" {
			api.SendMessage("Hello world", u.ChatID(), nil)
		}
	}
}
*/
/*
package main

import (
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	apiToken = "7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58"

	pollingIntervalSeconds = 1
	typingDelaySeconds     = 1

	verbose = true
)

// handle '/start' command
func startCommandHandler(b *bot.Bot, update bot.Update, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, "Starting chat...")
	}
}

// handle '/help' command
func helpCommandHandler(b *bot.Bot, update bot.Update, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, "Help message here.")
	}
}

// handle non-supported commands
func noSuchCommandHandler(b *bot.Bot, update bot.Update, cmd, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, fmt.Sprintf("No such command: %s", cmd))
	}
}

// handle non-command updates
func updateHandler(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		if update.HasMessage() {
			// 'is typing...'
			b.SendChatAction(update.Message.Chat.ID, bot.ChatActionTyping, nil)
			time.Sleep(typingDelaySeconds * time.Second)

			// send a reply,
			message := fmt.Sprintf("Received your message: %s", *update.Message.Text)
			send(b, update.Message.Chat.ID, update.Message.MessageID, message)

			// and add a reaction on the received message
			react(b, update.Message.Chat.ID, update.Message.MessageID, "👍")
		}
	} else {
		log.Printf(
			"*** error while receiving update (%s)",
			err.Error(),
		)
	}
}

// send a message
func send(b *bot.Bot, chatID, messageID int64, message string) {
	if sent := b.SendMessage(
		chatID,
		message,
		bot.OptionsSendMessage{}.
			SetReplyParameters(bot.NewReplyParameters(messageID)), // show original message
	); !sent.Ok {
		log.Printf(
			"*** failed to send a message: %s",
			*sent.Description,
		)
	}
}

// leave a reaction on a message
func react(b *bot.Bot, chatID, messageID int64, emoji string) {
	if reacted := b.SetMessageReaction(chatID, messageID, bot.NewMessageReactionWithEmoji(emoji)); !reacted.Ok {
		log.Printf(
			"*** failed to leave a reaction on a message: %s",
			*reacted.Description,
		)
	}
}

// generate bot's name
func botName(bot *bot.User) string {
	if bot != nil {
		if bot.Username != nil {
			return fmt.Sprintf("@%s (%s)", *bot.Username, bot.FirstName)
		} else {
			return bot.FirstName
		}
	}

	return "Unknown"
}

func main() {
	client := bot.NewClient(apiToken)
	client.Verbose = verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf("Bot information: %s", botName(me.Result))

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(true); unhooked.Ok {
			// add command handlers
			client.AddCommandHandler("/start", startCommandHandler)
			client.AddCommandHandler("/help", helpCommandHandler)
			client.SetNoMatchingCommandHandler(noSuchCommandHandler)

			// wait for new updates
			client.StartPollingUpdates(
				0,
				pollingIntervalSeconds,
				updateHandler,
			)
		} else {
			panic("failed to delete webhook")
		}
	} else {
		panic("failed to get info of the bot")
	}
}
*/
/*
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nickname76/telegrambot"
)

func main() {
	api, me, err := telegrambot.NewAPI("7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	stop := telegrambot.StartReceivingUpdates(api, func(update *telegrambot.Update, err error) {
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		msg := update.Message
		if msg == nil {
			return
		}

		_, err = api.SendMessage(&telegrambot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   fmt.Sprintf("Hello %v, I am %v", msg.From.FirstName, me.FirstName),
			ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
				Keyboard: [][]*telegrambot.KeyboardButton{{
					{
						Text: "Hello",
					},
				}},
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
			},
		})

		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
	})

	log.Printf("Started on %v", me.Username)

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt)

	<-exitCh

	// Waits for all updates handling to complete
	stop()
}
*/
/*
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// client := tg.New(os.Getenv("BOT_TOKEN"))
	client := tg.New("7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58")

	router := tgb.NewRouter().
		// handles /start and /help
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			return msg.Answer(
				tg.HTML.Text(
					tg.HTML.Strike("https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png"),
					tg.HTML.Bold("👋 Hi, I'm echo bot!"),
					"",
					//tg.HTML.Italic("🚀 Powered by", tg.HTML.Spoiler(tg.HTML.Link("go-tg", "github.com/mr-linch/go-tg"))),
				),
			).ParseMode(tg.HTML).DoVoid(ctx)
		}, tgb.Command("start", tgb.WithCommandAlias("help"))).
		// handles gopher image
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			if err := msg.Update.Reply(ctx, msg.AnswerChatAction(tg.ChatActionUploadPhoto)); err != nil {
				return fmt.Errorf("answer chat action: %w", err)
			}

			// emulate thinking :)
			time.Sleep(time.Second)

			return msg.AnswerPhoto(
				tg.NewFileArgURL("https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png"),
			).DoVoid(ctx)

		}, tgb.Regexp(regexp.MustCompile(`(?mi)(go|golang|gopher)[$\s+]?`))).
		// handle other messages
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			return msg.Copy(msg.Chat).DoVoid(ctx)
		}).
		MessageReaction(func(ctx context.Context, reaction *tgb.MessageReactionUpdate) error {
			// sets same reaction to the message
			answer := tg.NewSetMessageReactionCall(reaction.Chat, reaction.MessageID).Reaction(reaction.NewReaction)
			return reaction.Update.Reply(ctx, answer)
		})

	return tgb.NewPoller(
		router,
		client,
		tgb.WithPollerAllowedUpdates(
			tg.UpdateTypeMessage,
			tg.UpdateTypeMessageReaction,
		),
	).Run(ctx)
}

*/

/*
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

// This bot demonstrates some example interactions with commands on telegram.
// It has a basic start command with a bot intro.
// It also has a source command, which sends the bot sourcecode, as a file.
func main() {
	// Get token from the environment variable
	// token := os.Getenv("7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58")
	token := "7369039700:AAHH2Db4G2nDYI6RzsL5ttI1XWBpTJVuy58"
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	// Create bot from environment value.
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// /start command to introduce the bot
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	// /source command to send the bot source code
	dispatcher.AddHandler(handlers.NewCommand("source", source))

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

func source(b *gotgbot.Bot, ctx *ext.Context) error {
	// Sending a file by file handle
	f, err := os.Open("samples/commandBot/main.go")
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}

	m, err := b.SendDocument(ctx.EffectiveChat.Id,
		gotgbot.InputFileByReader("source.go", f),
		&gotgbot.SendDocumentOpts{
			Caption: "Here is my source code, by file handle.",
			ReplyParameters: &gotgbot.ReplyParameters{
				MessageId: ctx.EffectiveMessage.MessageId,
			},
		})
	if err != nil {
		return fmt.Errorf("failed to send source: %w", err)
	}

	// Or sending a file by file ID
	_, err = b.SendDocument(ctx.EffectiveChat.Id,
		gotgbot.InputFileByID(m.Document.FileId),
		&gotgbot.SendDocumentOpts{
			Caption: "Here is my source code, sent by file id.",
			ReplyParameters: &gotgbot.ReplyParameters{
				MessageId: ctx.EffectiveMessage.MessageId,
			},
		})
	if err != nil {
		return fmt.Errorf("failed to send source: %w", err)
	}

	return nil
}

// start introduces the bot.
func start(b *gotgbot.Bot, ctx *ext.Context) error {
	text := `<b>bold</b>, <strong>bold</strong>
<i>italic</i>, <em>italic</em>
<u>underline</u>, <ins>underline</ins>
<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
<span class="tg-spoiler">spoiler</span>, <tg-spoiler>spoiler</tg-spoiler>
<b>bold <i>italic bold <s>italic bold strikethrough <span class="tg-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
<a href="http://www.example.com/">inline URL</a>
<a href="tg://user?id=123456789">inline mention of a user</a>
<tg-emoji emoji-id="5368324170671202286">👍</tg-emoji>
<code>inline fixed-width code</code>
<pre>pre-formatted fixed-width code block</pre>
<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
<blockquote>Block quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>
<blockquote expandable>Expandable block quotation started\nExpandable block quotation continued\nExpandable block quotation continued\nHidden by default part of the block quotation started\nExpandable block quotation continued\nThe last line of the block quotation</blockquote>`
	//_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s. I <b>repeat</b> all your messages.", b.User.Username), &gotgbot.SendMessageOpts{
	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

*/
