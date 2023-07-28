A simple CLI generator for strong passwords. Written in Go's stdlib with cryptographically secure RNG. Provides multiple options to customize a password and robust defaults.

### Usage
There is three options: password length, password mode (alphanumeric, numeric or alphabetic) and special characters. By default, password-gen will generate a 12 character alphanumeric password with special characters.

To customize these options, you have three flags: -l (length), -m (mode) and -s (special chars).

- `-l`: password length, 1 or higher [defaults to 12]
- `-m`: generator mode, a (alphabetic), an (alphanumeric) or n (numeric) [defaults to an]
- `-s`: adds special characters (recommended)
Example: ./pwdmix -l 18 -m n would generate an 18 character-long numeric password.

