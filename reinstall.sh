sudo apt-get remove protobuf-compiler
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.12.1/protoc-3.12.1-linux-x86_64.zip

sudo apt install unzip
unzip protoc-3.12.1-linux-x86_64.zip -d $HOME/.local

export PATH="$PATH:$HOME/.local/bin"