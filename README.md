# Telegram File Backend

This proxy backend helps you serve a file sent to your Telegram bot on your website by using Telegram servers directly. Instead of your bot token in plain text, it uses an AES encrypted version of the token. So your bot token doesn't get exposed to the outside world while you'll still be able to serve a file using Telegram's servers as a file backend.

As an example, I added a Python script to encrypt given bot token. You can use it to create encrypted bot tokens as you'll need an encrypted token to request files from the proxy server.

## Set up

1. Encrypt your bot token

```bash
python encryptor/encrypt.py 'YOUR_BOT_TOKEN'
# 9444e7c7861ef9dcdec4174225f2f77e
```

2. Run a redis instance
3. Run the proxy
4. Visit the proxy server. E.g `http://localhost:4627/<ENCRYPTED_BOT_TOKEN>/<FILE_ID>`
