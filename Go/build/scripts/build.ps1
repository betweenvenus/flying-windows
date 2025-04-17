$Env:GOOS="js"
$Env:GOARCH="wasm"
cd $PSScriptRoot
go build -o ..\..\\website\bin\main.wasm