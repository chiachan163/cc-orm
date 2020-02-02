package main

import (
	"fmt"
	"os"

	"github.com/chiachan163/cc-orm/orm/create"
	"github.com/chiachan163/cc-orm/orm/info"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Micro project orm"
	app.Version = "1.0.0"
	app.Usage = "a deployment tools of tp-micro model"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "chiachan163",
			Email: "chiachan@163.com",
		},
	}

	newHello := &cli.Command{
		Name:  "hello",
		Usage: "hello world！",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "s",
				Value: "world",
				Usage: "Say hello",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("s") != "" {
				fmt.Println(fmt.Sprintf("hello,%s!", c.String("s")))
			} else {
				fmt.Println("gen test")
			}
			return nil
		},
	}

	// 创建model文件
	newCom := &cli.Command{
		Name:  "gen",
		Usage: "create model file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "app_path, p",
				Usage: "The path(relative/absolute) of the project",
			},
			&cli.BoolFlag{
				Name:  "force, f",
				Value: false,
				Usage: "Forced to rebuild the whole project",
			},
			&cli.BoolFlag{
				Name:  "newdoc",
				Value: false,
				Usage: "Rebuild the README.md",
			},
		},
		Before: initProject,
		Action: func(c *cli.Context) error {
			create.CreateModel(c.Bool("force"), c.Bool("newdoc"))
			return nil
		},
	}

	app.Commands = []*cli.Command{newCom, newHello}
	app.Run(os.Args)
}

func initProject(c *cli.Context) error {
	appPath := c.String("app_path")
	if len(appPath) == 0 {
		appPath = c.Args().First()
	}
	if len(appPath) == 0 {
		appPath = "./"
	}
	return info.Init(appPath)
}
