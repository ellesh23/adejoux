// nmon2influxdb
// import nmon data in InfluxDB
// author: adejoux@djouxtech.net

package main

import (
	"github.com/codegangsta/cli"
	"strings"
)

//
//helper functions
//
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReplaceComma(s string) string {
	return "<tr><td>" + strings.Replace(s, ",", "</td><td>", 1) + "</td></tr>"
}

type Params struct {
	Filepath string
	NoDisks  bool
	CpuAll   bool
	Server   string
	User     string
	Port     string
	File     bool
	Guser    string
	Gpass    string
	Gurl     string
	Db       string
	DS       string
	Password string
	Template string
	Debug    bool
}

func ParseParameters(c *cli.Context) (params *Params) {
	return &Params{Filepath: c.Args()[0],
		NoDisks:  c.Bool("nodisks"),
		CpuAll:   c.Bool("cpus"),
		File:     c.Bool("file"),
		Guser:    c.String("guser"),
		Gpass:    c.String("gpassword"),
		Gurl:     c.String("gurl"),
		DS:       c.String("datasource"),
		Debug:    c.GlobalBool("debug"),
		Server:   c.GlobalString("server"),
		User:     c.GlobalString("user"),
		Port:     c.GlobalString("port"),
		Db:       c.GlobalString("db"),
		Password: c.GlobalString("pass"),
		Template: c.String("template"),
	}
}