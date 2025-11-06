# cc-gi

Command-line gitignore generator from [caml.cc](https://caml.cc).

Generates gitignore files locally or fetches templates from [gi.caml.cc](https://gi.caml.cc).

## Usage
```bash
cc-gi generate python go        # generate gitignore
cc-gi -o generate python go     # only use local templates
cc-gi -v generate python go     # verbose
cc-gi clean                     # delete all local templates
```

## Installation
### Requirements for installation
- git
- golang
```bash
curl -sSL https://github.com/caml-cc/cc-gi/raw/main/scripts/install.sh | sudo bash
```

[MIT](LICENSE)

> G.I., your government has abandoned you. They have ordered you to die. Don’t trust them. They lied to you, G.I.s, you know you cannot win this war.
