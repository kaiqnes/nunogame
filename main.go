package main

import (
	"errors"
	"fmt"
	"github.com/mercadolibre/nunogame/src/entities"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clearConsole()

	// MOCK Main Menu
	intro()
	game := entities.NewGame(2, entities.NormalDungeon)

	// Get a player - DONE
	for index := 0; index < len(game.Players); {
		// LOG
		fmt.Printf("PLAYERS: %v\n", game.Players)
		fmt.Printf("Mão: %v\n", game.Players[index].Hand.Cards)
		fmt.Printf("Descarte: %v\n", game.DiscardedDoors)
		fmt.Printf("Deck: %v\n", game.Doors)
		fmt.Println()

		preCombatAction()

		// KICK THE DOOR
		game.KickTheDoor(&game.Players[index])

		// OPEN DOOR
		drawnCard, err := drawCard(game)
		if err != nil {
			fmt.Println(err)
			exit()
		}

		// TODO: IDENTIFICAR TIPO DE CARTA (MONSTRO, MALDIÇÃO, CLASSE, RAÇA)

		// TODO: MOVER A ESCOLHA DO PLAYER PARA DENTRO DA IDENTIFICAÇÃO DE CARTA
		fmt.Println("Você gostaria de guardar ou descartar a carta?")
		fmt.Println("1 - Guardar")
		fmt.Println("2 - Descartar")
		switch scanChoice() {
		case 1:
			game.Players[index].Hand.KeepCard(drawnCard, &game.DiscardedDoors)
		case 2:
			game.Doors.DiscardCard(drawnCard, &game.DiscardedDoors)
		default:
			exit("Opção inválida, finalizando o programa")
		}

		// TODO: VALIDAR A QUANTIDADE DE CARTAS AQUI, PARA PLAYER DESCARTAR O EXCEDENTE

		clearConsole()

		if index == len(game.Players)-1 {
			index = 0
		} else {
			index++
		}
	}

	fmt.Println(game.Players)
}

func preCombatAction() {
	// TODO: PRE COMBAT ACTION
	takeAction()
}

func takeAction() {
	// TODO: TAKE ACTION
}

func scanChoice() (choice int8) {
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println(err)
		choice = -1
	}
	return choice
}

func drawCard(game *entities.Game) (entities.Card, error) {
	fmt.Println("Puxe uma carta")
	if len(game.Doors.Cards) == 0 {
		if len(game.DiscardedDoors) > 0 {
			game.Doors.Cards = game.DiscardedDoors
			game.DiscardedDoors = []entities.Card{}
			game.Doors.ShuffleDeck()

			drawnCard := game.Doors.DrawCard()
			fmt.Println("a carta puxada foi: ", drawnCard)
			return drawnCard, nil
		}
		return entities.Card{}, errors.New("sem novas cartas p/ sacar")
	}

	drawnCard := game.Doors.DrawCard()
	fmt.Println("a carta puxada foi: ", drawnCard)
	return drawnCard, nil
}

func intro() {
	// PRINT INTRO HERE
}

func exit(message ...string) {
	if len(message) > 0 {
		fmt.Println("[ EXIT:", strings.Join(message, " | "), "]")
	}
	os.Exit(1)
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
