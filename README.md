# Telegram Bot for IP Server Check
***
*This is a Telegram bot designed for checking the IP of a server.*


## Setup Instructions

1. Edit/Create the .env file for your environment with the following variables:
   - IP_REQUEST_ADDRESS - The URL for retrieving the IP address. Recommended: https://ifconfig.co
   - TELEGRAM_TOKEN=`Your Telegram API token (obtain from https://t.me/BotFather)`
   - ERROR_RESPONSE_MESSAGE=`Standard error message`
   - ACCESS_PASSWORD=`Your password for adding users`
   - SQLITE_PATH=`Path to your local SQLite database`

2. Run the application using the following command:

```sh
docker-compose --env-file .env up --build
```