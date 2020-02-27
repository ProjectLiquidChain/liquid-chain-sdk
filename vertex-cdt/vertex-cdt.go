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
	app.Usage = "vertex-cdt [option] [file] --export-function [functions]"
	app.Version = "0.0.1"
	app.Author = "vertex team"
	app.Commands = []cli.Command{
		{
			Name:    "compile",
			Aliases: []string{"c"},
			Usage:   "compile c,c++ language file",
			Flags:   []cli.Flag{cli.StringFlag{Name: "export-function, f"}},
			Action: func(c *cli.Context) {
				compile := tool.Compile{c.Args().First()}
				wasmFile, nameFile := compile.Clang(c.String("export-function"))
				abiFile, event_names := tool.ABIC(c.Args().First(), nameFile, c.String("export-function"), wasmFile)
				if tool.CheckImportFunction(wasmFile, event_names) {
					log.Println("compile completed!")
				} else {
					utils.DeleteFile(abiFile)
					utils.DeleteFile(wasmFile)
				}
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "compile rust language folder",
			Action: func(c *cli.Context) {
				file := strings.Split(c.Args().First(), "/")
				if file[len(file)-1] == "" {
					file[len(file)-1] = file[len(file)-2]
				}
				if strings.Contains(file[len(file)-1], "-") {
					file[len(file)-1] = strings.ReplaceAll(file[len(file)-1], "-", "_")
				}
				compile := tool.Compile{c.Args().First()}
				compile.Rust(file[len(file)-1] + ".wasm")
				wasm_file := c.Args().First() + "/" + file[len(file)-1] + ".wasm"
				abiFile, event_names := tool.ABIRust(c.Args().First()+"/src/lib.rs", file[len(file)-1], c.Args().First()+"/", wasm_file)
				if tool.CheckImportFunction(wasm_file, event_names) {
					log.Println("compile completed!")
				} else {
					utils.DeleteFile(c.Args().First() + file[len(file)-1] + ".wasm")
					utils.DeleteFile(abiFile)
				}
			},
		},
		{
			Name:    "init",
			Aliases: []string{"r"},
			Flags:   []cli.Flag{cli.StringFlag{Name: "name, n"}},
			Usage:   "create rust project",
			Action: func(c *cli.Context) {
				file := tool.Create(c.String("name"))
				log.Println(file)
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
