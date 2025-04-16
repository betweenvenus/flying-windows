$Env:GOOS="js"
$Env:GOARCH="wasm"
go build -o ./website/bin/main.wasm