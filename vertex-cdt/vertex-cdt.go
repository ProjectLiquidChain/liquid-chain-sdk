package main

import (
	"log"
	"os"
	"strings"

	tool "github.com/QuoineFinancial/vertex-sdk/vertex-cdt/tools"
	utils "github.com/QuoineFinancial/vertex-sdk/vertex-cdt/utils"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Name = "smart contract development CLI"
	app.Usage = "vertex-cdt [language option] [file] --export-function [functions]"
	app.Version = "0.0.1"
	app.Author = "vertex team"
	app.Commands = []cli.Command{
		{
			Name:    "compile",
			Aliases: []string{"c"},
			Usage:   "compile c,c++ language file",
			Flags:   []cli.Flag{cli.StringFlag{Name: "export-function, ef"}},
			Action: func(c *cli.Context) {
				compile := tool.Compile{c.Args().First()}
				wasmFile, nameFile := compile.Clang(c.String("export-function"))
				abiFile, event_names := tool.ABIgen(c.Args().First(), nameFile, c.String("export-function"), wasmFile)
				if tool.CheckImportFunction(wasmFile, event_names) {
					log.Println("compile completed!")
				} else {
					utils.DeleteFile(abiFile)
					utils.DeleteFile(wasmFile)
				}
			},
		},
		{
			Name:    "rust",
			Aliases: []string{"r"},
			Usage:   "compile rust language file",
			Action: func(c *cli.Context) {
				file := strings.Split(c.Args().First(), "/")
				if file[len(file)-1] == "" {
					file[len(file)-1] = file[len(file)-2]
				}
				compile := tool.Compile{c.Args().First()}
				compile.Rust()
				if tool.CheckImportFunction(c.Args().First()+"/target/wasm32-wasi/debug/"+file[len(file)-1]+".wasm", []string{}) {
					log.Println("compile completed!")
				} else {
					utils.DeleteFolder(c.Args().First() + "/target")
				}
			},
		},
	}
}
func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
