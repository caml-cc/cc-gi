# cc-gi

Command-line gitignore generator from [caml.cc](https://caml.cc).

Generates gitignore files locally or fetches templates from [gi.caml.cc](https://gi.caml.cc).

## Usage
```bash
cc-gi python,go        # generate gitignore
cc-gi -o python,go     # only use local templates
cc-gi -v python,go     # verbose
```

## Install
```bash
curl -sSL https://gitlab.com/caml.cc/cc-tools/raw/main/install.sh | bash
```

[MIT][LICENSE]

> G.I., your government has abandoned you. They have ordered you to die. Don’t trust them. They lied to you, G.I.s, you know you cannot win this war.
