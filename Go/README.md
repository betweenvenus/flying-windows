# Flying Windows (Go version)

## How to build
Build scripts are available in `build/scripts`. Run whichever one is suitable for your environment.

## Running the demo page
Run `npm start` in `./website` to launch the dev server.

## Go complains about build constraints when I try to install `honnef.co/go/js/dom/v2`

This package targets `GOOS="js"` and `GOARCH="wasm"` so it won't install or import correctly by default if you're working with `gopls`. If you're using VS Code, you can fix this by setting `gopls`' environment variables in the workspace-level `.vscode/settings.json`:

```json
// .vscode/settings.json
{
	"gopls": {
		
		"build.env": {
			"GOOS": "js",
			"GOARCH": "wasm"
		}
	}
}
```

You'll need to restart VS Code afterwards.