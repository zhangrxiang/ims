package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"log"
	"simple-ims/web"
	"sync"
)

func main() {
	prg := &Program{}
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
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
