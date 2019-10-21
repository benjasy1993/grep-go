package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/benjasy1993/grep-go/grep"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	app.Name = "Distributed Grep CLI"
	app.Usage = "Put all servers (id, ip:port) into a file named cluster.config. The command will find all lines that contain the keyword specified in the specifed directories."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "keyword",
			Required: true,
			Usage:    "the keyword to grep",
		},
		cli.StringSliceFlag{
			Name:     "directories",
			Required: true,
			Usage:    "the directories to search, split by comma",
		},
		cli.BoolFlag{
			Name:  "line_number",
			Usage: "whether to include line number",
		},
		cli.IntSliceFlag{
			Name:  "server id",
			Usage: "the servers to search, will search all servers if unspecified",
		},
		cli.StringFlag{
			Name:     "config_path",
			Required: false,
			Usage:    "the file path of config path",
		},
		cli.BoolFlag{
			Name:     "distrubuted",
			Required: false,
			Usage:    "whether grep in distributed mode",
		},
	}
	app.Action = executeGrep
	app.Run(os.Args)
}

func executeGrep(c *cli.Context) {
	keyword := c.Args().Get(0)
	dirs := strings.Split(c.Args().Get(1), ",")
	includeLineNum, nil := strconv.ParseBool(c.Args().Get(2))
	isDistributed, nil := strconv.ParseBool(c.Args().Get(5))
	// serverID, nil := strconv.Atoi(c.Args().Get(3))
	// configPath := c.Args().Get(4)

	file, err := os.Open("cluster.config")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	idByServerAddr := make(map[int]string)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "\t")
		id, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		serverAddr := strs[1]
		idByServerAddr[id] = serverAddr
	}

	var g grep.Grepper

	if isDistributed {
		g = grep.NewDistributedGrepper(0, idByServerAddr)
	} else {
		g = grep.NewLocalGrepper()
	}

	fmt.Println(g.Grep(keyword, dirs, includeLineNum))
}
