package bots

import (
	"flag"
	"fmt"
	bot "github.com/edouard127/mc-go-1.12.2"
	"log"
)

var address = flag.String("address", "127.0.0.1", "The server address")

func Join(c *bot.Auth) {
	flag.Parse()
	//log.SetOutput(colorable.NewColorableStdout())

	//Login
	game, err := c.JoinServer(*address, 25565)
	if err != nil {
		log.Fatal(err)
	}

	//Handle game
	events := game.GetEvents()
	go func() {
		err := game.HandleGame()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Waiting for events...")
	for e := range events { //Receiving events
		switch e.(type) {
		case bot.JoinGameEvent:
			fmt.Printf("Joined game as %v\n", c.Name)
		case bot.ChatMessageEvent:
			fmt.Printf("Chat message: %v\n", e.(bot.ChatMessageEvent).Msg)
		}
	}
}
