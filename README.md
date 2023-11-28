A simple CLI generator for strong passwords. Written in Go's stdlib with cryptographically secure RNG. Provides multiple options to customize a password. It shadows the generated password in memory after execution.

![pwdmix-demo](https://github.com/parsec501/pwdmix/assets/105080989/e105ab71-e2c4-444d-a6d6-5d0ed54a9d9d)

### Usage (`--help`)

There is three options: password length, password mode (alphanumeric, numeric or alphabetic) and special characters. By default, pwdmix will generate a 12 character alphanumeric password without special characters.

To customize these options, you have three flags: -l (length), -m (mode) and -s (special chars).

- `-l`: password length, 1 or higher [defaults to 12]
- `-m`: generator mode, a (alphabetic), an (alphanumeric) or n (numeric) [defaults to an]
- `-s`: adds special characters (recommended)
- 
Example: `./pwdmix -l 18 -m n` would generate an 18 character-long numeric password.

### Installation
Download a precompiled binary (available for Windows, Linux and MacOS in multiple architectures) from the releases page. I recommend renaming it to just pwdmix or pwdmix.exe since the OS and architecture info makes the filename relatively long.

Alternatively, compile it yourself with `go build -ldflags="-s -w" -o outputname cmd/pwdmix/main.go`.
