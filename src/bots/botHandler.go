package bots

import (
	"flag"
	"fmt"
	. "github.com/edouard127/mc-go-1.12.2/struct"
	"kamigen/2fucked2pay/src/utils"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	address = flag.String("address", "127.0.0.1", "The server address")
	port    = flag.Int("port", 25565, "The server port")
)

func Join(c *Auth) {
	var wg sync.WaitGroup
	wg.Add(1)
	flag.Parse()

	game, err := c.JoinServer(*address, *port)
	if err != nil {
		log.Fatal(err)
	}

	//Handle game
	queue := utils.Queue{
		Position: 0,
	}
	events := game.GetEvents()
	motion := game.Motion
	go func() {
		err := game.HandleGame()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		ticks := 0
		for {
			time.Sleep(50 * time.Millisecond)
			// Execute at every 100 ticks
			ticks++
			// Switch ticks%100 == 0
			if ticks%60 == 0 {
				// Random yaw and pitch
				yaw := RandomFloat64(-180, 180)
				pitch := RandomFloat64(-90, 90)
				game.LookYawPitch(float32(yaw), float32(pitch), true)
				game.TweenJump()
				game.SwingHand(true)
			}
			if ticks%6000 == 0 {
				game.Chat(RandomStr())
			}
			/*case EntityRelativeMoveEvent:
			e := game.World.ClosestEntity(game.Player.GetPosition(), 50)
			if e != nil {
				game.LookAt(e.Position)
				game.SwingHand(true)
			}*/
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
				fmt.Printf("%s%s\n", event.Sender, event.Content)
				ParseCommands(game, event.Content)
			} else {
				fmt.Printf("%s\n", event.Content)
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
		}
	}
	for f := range motion {
		f()
	}
}

func GetMemoryUsage() (uint64, uint64, uint32) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	/*fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)*/
	return bToMb(m.Alloc), bToMb(m.TotalAlloc), m.NumGC
}
func alloc() uint64 {
	alloc, _, _ := GetMemoryUsage()
	return alloc
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomOwO() string {
	owo := []string{"OwO", "UwU", "OvO", "UvU", "OwU", "UwO", "OvU", "UvO"}
	return owo[rand.Intn(len(owo))]
}

func RandomStr() string {
	s := []string{
		fmt.Sprintf("> This bot is running on golang, using %v MiB of RAM and it's open source! https://github.com/Edouard127/mc-go-1.12.2 Contribute today %s !", alloc(), RandomOwO()),
		fmt.Sprintf("> Are you a Go developer? Contribute to this project! https://github.com/Edouard127/mc-go-1.12.2 %s", RandomOwO()),
		fmt.Sprintf("> Are you a Minecraft expert? Contribute to this project! https://github.com/Edouard127/mc-go-1.12.2 %s", RandomOwO()),
		fmt.Sprintf("> Would you like to have bots building highways for you? Contribute to this project! https://github.com/Edouard127/mc-go-1.12.2 %s", RandomOwO()),
	}
	return s[rand.Intn(len(s))]
}
