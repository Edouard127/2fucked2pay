package bots

import (
	"flag"
	"fmt"
	. "github.com/edouard127/mc-go-1.12.2/struct"
	"kamigen/2fucked2pay/src/utils"
	"log"
	"runtime"
)

var (
	address = flag.String("address", "127.0.0.1", "The server address")
)

func Join(c *Auth) {
	flag.Parse()

	//Login
	game, err := c.JoinServer(*address, 25565)
	if err != nil {
		log.Fatal(err)
	}

	//Handle game
	queue := utils.Queue{
		Position: 0,
	}
	events := game.GetEvents()
	go func() {
		err := game.HandleGame()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for e := range events { //Receiving events
		switch e.(type) {
		case JoinGameEvent:
			fmt.Printf("Joined game as %v\n", c.Name)
		case ChatMessageEvent:
			chat := e.(ChatMessageEvent)
			if game.Server.Addr == "connect.2b2t.org" {
				queue.ParseString(chat.Content)
				fmt.Printf("Queue position: %v\n", queue.Position)
			}
			if len(chat.Sender) > 0 {
				fmt.Printf("<%v> %v\n", chat.Sender, chat.Content)
				utils.LogFile(chat.RawString)
			} else {
				fmt.Printf("%v\n", chat.Content)
				utils.LogFile(chat.RawString)
			}
		case DisconnectEvent:
			fmt.Printf("Disconnected: %v\n", e.(DisconnectEvent).Text)
		case TitleEvent:
			//fmt.Printf("Title: %v\n", e.(bot.TitleEvent).Text)
		}
	}
}

func GetMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
