
# 🎌 Otaku Bot (Slack Anime Quote Bot)

Otaku Bot is a fun Slack bot built in Go that delivers random inspirational, funny, or emotional anime quotes directly into your Slack channels — complete with character names, anime titles, and even anime poster images via the Kitsu API.

---

## 📸 Preview

> Example output from the `/animequote` command:

![alt text](anime-quote_slack.gif)

---

## ✨ Features

- ⚡️ Slash command: `/animequote`
- 🎭 Random anime quotes via [AnimeChan API](https://animechan.io)
- 🖼️ Anime cover images fetched from [Kitsu API](https://kitsu.docs.apiary.io/)
- 💬 Nicely formatted message using [Slack Block Kit](https://api.slack.com/block-kit)

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yash-death/otaku-bot.git
cd otaku-bot
```

### 2. Set up your `.env`

Create a `.env` file in the root directory:

```
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-level-token
```

> You can get these from your Slack App dashboard.

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the bot

```bash
go run main.go
```

> The bot will connect using **Slack Socket Mode** and post quotes using the `/animequote` command.

---

## 🧱 Project Structure

```bash
.
├── main.go             # Entry point
├── commands/
│   └── animequote.go   # Quote command logic
├── utils/
│   └── kitsu.go        # Anime image fetcher (Kitsu API)
├── .env                # Environment variables
```

---

## 🔮 Planned Features 

- 🔘 Slack interactive buttons (`Refresh`, `Save`, `More Info`)
- 🗂️ Quote history storage (JSON or DB)
- 🏆 Karma system (upvote/downvote anime)
- 🧠 Slash command autocomplete

---

## 👨‍💻 Built With

- [Go](https://golang.org/)
- [Slack Go SDK](https://github.com/slack-go/slack)
- [AnimeChan API](https://animechan.xyz/)
- [Kitsu API](https://kitsu.docs.apiary.io/)

---

## 🤝 Contributing

Pull requests are welcome! If you have ideas for fun features, feel free to open an issue or submit a PR.

---

## 📜 License

MIT License

---

## 📫 Contact

Made with ❤️ by [Yash Badoniya](https://github.com/Yash-death)
