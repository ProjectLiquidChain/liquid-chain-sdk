package main

import (
	"fmt"
	"os"
	"strings"

	tool "github.com/QuoineFinancial/vertex-sdk/vertex-cdt/tools"
	utils "github.com/QuoineFinancial/vertex-sdk/vertex-cdt/utils"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "smart contract development CLI"
	app.Usage = "vertex-cdt [language option] compile [file]"
	app.Version = "0.0.1"
	app.Author = "vertex team"
}
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "c++",
			Aliases: []string{"c++"},
			Usage:   "compile c++ language file",
			Action: func(c *cli.Context) {
				compile := tool.Compile{c.Args().First(), "c++"}
				result := compile.Clang()
				if tool.CheckImportFunction(result) {
					tool.ABIgen(c.Args().First(), "c++")
					fmt.Println("compile completed!")
				} else {
					utils.DeleteFile(result) // remove file .wasm
				}
			},
		},
		{
			Name:    "c",
			Aliases: []string{"c"},
			Usage:   "compile c language file",
			Action: func(c *cli.Context) {
				compile := tool.Compile{c.Args().First(), "c"}
				result := compile.Clang()
				if tool.CheckImportFunction(result) {
					tool.ABIgen(c.Args().First(), "c")
					fmt.Println("compile completed!")
				} else {
					utils.DeleteFile(result) //remove file .wasm
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
				compile := tool.Compile{c.Args().First(), "rust"}
				compile.Rust()
				if tool.CheckImportFunction(c.Args().First() + "/target/wasm32-wasi/debug/" + file[len(file)-1] + ".wasm") {
					fmt.Println("compile completed!")
				} else {
					utils.DeleteFolder(c.Args().First() + "/target") // ? remove file .wasm
				}
			},
		},
	}
}
func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
