CURR_DIR=$(pwd)
cd ~
CONFIG_PATH=$(pwd)/.config/pump
cd "$CURR_DIR"

go build \
  --ldflags "-X main.ConfigPath=$CONFIG_PATH/config.json -s -w" \
  -o /usr/local/bin/pump cmd/main.go

