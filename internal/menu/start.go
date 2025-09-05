package menu

import (
	"fmt"
	"go-todo/internal/logger"
	"go-todo/internal/todo"
)

type StartMenu struct {
	UserChoice int8
}

func (menu *StartMenu) MenuStart(todoStruct *todo.Todo) {

	// s := bufio.NewScanner(os.Stdin)
	// for s.Scan() {
	// 	log.Println("line", s.Text())
	// }

	fmt.Println(todoStruct)

	for menu.UserChoice != 4 {
		var input string

		fmt.Println("Press 1 to add \nPress 0 to exit")
		var _, scanErr = fmt.Scanln(&input)
		if scanErr != nil {
			logger.LogError(scanErr)
			menu.UserChoice = 4
		}

		switch input {
		case "1":
			fmt.Println("You pressed 1")
			menu.UserChoice = 1
		case "0":
			fmt.Println("time to exit")
			menu.UserChoice = 4
		default:
			fmt.Println("You pressed incorrect value")
		}
	}
	// Show todo
	//
	// Move to  new submenu
	//
	// Quit
}
