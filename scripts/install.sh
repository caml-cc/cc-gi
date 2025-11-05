#1!/bin/bash
set -e

if ! command -v git &> /dev/null; then
    echo "Error: git is not installed."
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed."
    exit 1
fi

BIN_DIR=/usr/local/bin
RAND_STR=$(head -n 4096 /dev/urandom | openssl sha1 | awk '{print $2}')

git clone https://github.com/caml-cc/cc-gi.git $RAND_STR
cd $RAND_STR
go build -o ./bin/cc-gi ./cmd/cc-gi/main.go
mkdir -p BIN_DIR
rm /usr/local/bin/cc-gi || true
mv ./bin/cc-gi /usr/local/bin/cc-gi
cd ..
rm -rf $RAND_STR
echo "cc-gi installed to ${BIN_DIR}/cc-gi"