# PutPut
Multitool in golang by hampterl

# In early development
Base 64, Hex, Aes256 implementation

# Getting Started
This is a Command Line Tool, you have to open it in a terminal. E.g. cmd or powershell.

When opened, you can use the commands:
Base64: base64 [encode|decode] [string]
Hex: hex [encode|decode] [string]
Aes256: aes256 [encrypt|decrypt] [string]

After you choose a command, you have to choose between encoding or decoding.
Then you have to type in the string you want to encode or decode.
Afterwards you can save it into a file, with the filename you want. It gets saved in the same directory as the main.

In Aes256 you have to type in a password.
It has to be either 16, 24 or 32 characters long.
If you encrypt a string, it gets encrypted and printed in hex format.
If you want to decrypt, you have to type in the same password as the original encryption or it won't print the same string.
Afterwards you can save it into a file, same as in Base64 and Hex.

Use '0' to exit the program while in the menu.

## Authors
hampterl

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
