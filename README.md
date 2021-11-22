# [SNiC MobilIT](https://mobilit.snic.nl/) availability checker

## Environment variables
- TELEGRAM_BOT_TOKEN: The token you got from the [BotFather](https://core.telegram.org/bots#6-botfather).
- TELEGRAM_CHAT_ID: Your chat ID with the bot. If you send <code>/id</code> to the bot, it will respond with your chat ID.
- INTERESTING_TRACKS: A comma seperated list with the tracks you are interested in. For example <code>ns,anwb</code>.

## Heroku
To run this code on one of the [Heroku servers](https://www.heroku.com/), we listen on a specific port, so it will act like a webserver. However with the free approach on Heroku, your app will go offline if it isn't accessed within an hour. Therefore [kaffeine](https://kaffeine.herokuapp.com/) is used to keep this app online.

## Telegram bot
### Automatic update
The backend checks every minute if there is an empty spot at one of the talks where you are interested in.<br/>
If it has found such a spot, it sends you a message via a [Telegram bot](https://core.telegram.org/bots). For example:<br/>
>There is 1 place left for the talk: ns<br/>

### Manual update
If you send <code>/update</code> to the Telegram bot, it will first respond with <code>Running...</code> to indicate that it's checking for empty spots. If it has found empty spots for the talks where you are interested in, it will send a new message, for example:<br/>
>There are 2 places left for the talk: ns<br/>
>There is 1 place left for the talk: anwb<br/>

It will always finish with a message <code>Done</code> to indicate it finished your request.

### Get info about all talks
If you send <code>/info</code> to the Telegram bot, it will first respond with <code>Running...</code> to indicate that it's checking all spots for all talks. Then it will send a new message with the number of occupied spots for every talk, for example:<br/>
>cofano: 103 of 230 places are gone.<br/>
computest: 230 of 230 places are gone.<br/>
xomnia: 82 of 230 places are gone.<br/>
han: 228 of 230 places are gone.<br/>
speeddates1: 7 of 8 places are gone.<br/>
projectmarch: 225 of 230 places are gone.<br/>
martin: 122 of 230 places are gone.<br/>
ns: 230 of 230 places are gone.<br/>
anwb: 189 of 230 places are gone.<br/>
aivd: 230 of 230 places are gone.<br/>
speeddates2: 8 of 8 places are gone.<br/>
speeddates3: 8 of 8 places are gone.<br/>

Then it will finish with a message <code>Done</code> to indicate it finished your request.

### Online / offline
The telegram bot will also send you a message if the server starts up (<code>I'm back online!</code>), or shuts down (<code>Going offline...</code>). This message should always be sent, even if a part of the program crashed. Note that if the Telegram bot crashes, you will obviously not get a message.<br/>
By sending those messages, you will always be notified about the current status of the server.
(Keep in mind that you can always access the corresponding website on Heroku to get your server back online.)