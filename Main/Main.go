package main

import (
	m "GoChannels/manager"
)

func main() {
	//Create the manager and start the game
	manager := m.GetInstance()
	manager.StartGame(5)
}
