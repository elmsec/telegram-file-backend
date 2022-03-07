# Telegram\* File Backend

**Please use it at your own risk.**

This proxy backend helps you serve a file sent to your Telegram bot on your website by using Telegram servers directly. Instead of your bot token in plain text, it uses an AES encrypted version of the token. So your bot token doesn't get exposed to the outside world while you'll still be able to serve a file using Telegram's servers as a file backend.

As an example of encrypting a bot token in programming languages other than Go, I included a Python script to encrypt the given bot token. Since you will need an encrypted token to request files from the proxy server, you can use this to generate encrypted bot tokens.

## Set up

1. Encrypt your bot token

```bash
python encryptor/encrypt.py 'YOUR_BOT_TOKEN'
# 9444e7c7861ef9dcdec4174225f2f77e
```

2. Run a redis instance
3. Run the proxy
4. Visit the proxy server. E.g `http://localhost:4627/<ENCRYPTED_BOT_TOKEN>/<FILE_ID>`

## \* Disclosure and warning

- **Please use it at your own risk.**
- We are not affiliated with, funded, or in any way associated with Telegram Messenger™.
