package main


import (
	"os"

	"github.com/n0mansky/countbeat/cmd"

	_ "github.com/n0mansky/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
