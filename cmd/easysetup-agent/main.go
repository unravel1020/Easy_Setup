package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/unravel1020/Easy_Setup/internal/agent"
	"github.com/unravel1020/Easy_Setup/internal/catalog"
)

func main() {
	catalogPath := flag.String("catalog", "src/catalog/catalog.json", "path to catalog JSON")
	host := flag.String("host", "127.0.0.1", "listen host")
	port := flag.String("port", "17772", "listen port")
	platform := flag.String("platform", "windows", "target platform")
	allowExecute := flag.Bool("allow-execute", false, "enable execution endpoint")
	flag.Parse()

	catalogData, err := catalog.Load(*catalogPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	address := *host + ":" + *port
	fmt.Printf("Easy_Setup Go Agent listening on http://%s/\n", address)
	fmt.Printf("AllowExecute: %v\n", *allowExecute)

	server := agent.New(catalogData, *platform, *allowExecute)
	if err := http.ListenAndServe(address, server.Handler()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
