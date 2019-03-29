PLATFORMS="linux/amd64"
PLATFORMS="$PLATFORMS darwin/amd64"

rm -rf bin/*

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_PATH="bin/$GOOS/$GOARCH"
  mkdir -p $BIN_PATH
  env GOOS=${GOOS} GOARCH=${GOARCH} go build -o $BIN_PATH/process_collector src/main.go
done


