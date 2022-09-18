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
	port    = flag.Int("port", 25565, "The server port")
)

func Join(c *Auth) {
	flag.Parse()

	//Login
	game, err := c.JoinServer(*address, *port)
	if err != nil {
		log.Fatal(err)
	}

	//Handle game
	queue := utils.Queue{
		Position: 0,
	}
	ticks := 0
	events := game.GetEvents()
	motions := game.Motion
	go func() {
		err := game.HandleGame()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for e := range events {
		switch event := e.(type) {
		case JoinGameEvent:
			fmt.Printf("Joined game as %v\n", c.Name)
		case ChatMessageEvent:
			if game.Server.Addr == "connect.2b2t.org" {
				queue.ParseString(event.Content)
				fmt.Printf("Queue position: %v\n", queue.Position)
			}
			if len(event.Sender) > 0 {
				fmt.Printf("<%v> %v\n", event.Sender, event.Content)
				ParseCommands(game, event.Content)
			} else {
				fmt.Printf("%v\n", event.Content)
			}
			utils.LogFile(event.RawString)
		case DisconnectEvent:
			fmt.Printf("Disconnected: %v\n", event.Text)
		case TitleEvent:
			//fmt.Printf("Title: %v\n", e.(bot.TitleEvent).Text)
		case TimeUpdateEvent:
			// https://static.wikia.nocookie.net/minecraft_gamepedia/images/b/bc/Day_Night_Clock_24h.png
			switch event.Time.TimeOfDay {
			case 22500:
				game.Chat("Good morning cubic world!")
			case 12000:
				game.Chat("It's time to eat dinner! It's 6:00 PM!")
			case 13000:
				game.Chat("If you are muslim, you can eat now! It's 9:00 PM!")
			case 18000:
				game.Chat("It's midnight, If you want to be pounded by 14 werewolves, It's the time!")
			}
		case TickEvent:
			// Execute at every 100 ticks
			ticks++
			// Switch ticks%100 == 0
			if ticks%100 == 0 {
				game.SwingHand(true)
			}
			if ticks%6000 == 0 {
				game.Chat("This bot is running on golang, and it's open source! https://github.com/Edouard127/mc-go-1.12.2")
			}
		}
	}
	for m := range motions {
		m()
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
