package main

import (
	_ "embed"
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/urfave/cli/v2"
	"github.com/zing-dev/soft-version/soft"
	"log"
	"os"
	"simple-ims/web"
	"sync"
)

//go:embed version.json
var src []byte

func main() {
	prg := &Program{}
	app := soft.NewCli(&cli.App{
		Name: "ims",
		Action: func(context *cli.Context) error {
			return svc.Run(prg)
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run the ims server soft",
				Action: func(c *cli.Context) error {
					return svc.Run(prg)
				},
			},
		},
	}, src)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type Program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func (p *Program) Init(env svc.Environment) error {
	log.Println("服务正在初始化...")
	fmt.Printf("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (p *Program) Start() error {
	log.Println("服务正在运行...")

	p.wg.Add(1)
	web.NewOnceWeb()

	p.quit = make(chan struct{})
	go func() {
		<-p.quit
		p.wg.Done()
	}()

	return nil
}

func (p *Program) Stop() error {
	log.Println("服务正在关闭...")
	close(p.quit)
	p.wg.Wait()
	log.Println("服务已经关闭...")
	return nil
}
