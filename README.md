### PutPut Multitool
Multitool in golang by hampterl

## Download
Download latest release -> https://github.com/hampterl/Multitool/releases

## On Windows
1. Download `multi-tool-windows.exe`
2. Open Command Prompt (Win + R → "cmd")
3. Navigate to download folder: `cd Downloads`
4. Run: `multi-tool-windows.exe `

## Build from source
# 1. Clone repository
git clone https://github.com/hampterl/multi-tool
cd multi-tool

# 2. Build
go build -ldflags="-s -w" -o putput_v1.0.exe

# 3. Run
./putput_v1.0.exe



# In early development
Base 64, Hex, Aes256 implementation

## Getting Started
This is a Command Line Tool, you have to open it in a terminal. E.g. cmd or powershell.
When opened, you can choose between encoding/decoding and encryption/decryption. More will be added soon.

# Features

- Encoding/Decoding (Base64, Hex)
- Encryption/Decryption (Aes256)

After you choose encoding/decoding, you can choose between Base64 and Hex. More options will be added soon.

Then you have to choose the mode you want to use. Encode or Decode.
Until now, only text can be encoded/decoded. Files and pictures are not supported yet.
Result will be printed in the terminal.
Afterwards you can save the result into a file, with the filename and kind of file you want. E.g. .txt .doc...
The file gets saved in the same directory as the main. So Multitool/cmd/multi-tool/<here your file> and the main

Press '0' to exit the program while in the menu and to go back to the main menu while in other modes.

After you choose encryption/decryption, you can only choose Aes256 right now, more options will be added soon.

Then you have to choose the mode you want to use. Encrypt text (it will encrypt to hex so it's readable) or decrypt text or file encryption.

For encryption and decryption to text you have to first type in a password.
It has to be either 16, 24 or 32 characters long. Important! It has to be that long else  
it won't work and it will just print an error.

Then you type in the text you want to encrypt/decrypt.
The result will be printed in the terminal.
Afterwards you can save the result into a file aswell, with the filename and kind of file you want. E.g. .txt .doc...
The file gets saved in the same directory as the main. So Multitool/cmd/multi-tool/<here your file> and the main.

If you want to decrypt, you have to type in the same password as the original encryption or it won't print the same string.
Then you can save the result into a file aswell.

# File encryption
DISCLAIMER and WARNING: This WILL encrypt the file you choose! IF you lose your key and overwrite your original, 
you will NOT be able to decrypt the file! Never share your key with anyone! ALWAYS make a backup before encrypting! 
IMPORTANT! Only encryption is supported right now! Without another tool, you won't be able to decrypt the file!

## Legal Disclaimer
USE FOR EDUCATIONAL PURPOSES ONLY!

NEVER USE THIS TO ENCRYPT SENSITIVE DATA!

THE DEVELOPER OF THIS TOOL IS NOT RESPONSIBLE FOR ANY DAMAGE BY MISSUSE! SEE LICENSE FOR MORE INFORMATION!

## Author
hampterl on github

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
