package run

import (
	"fmt"

	"github.com/Mmmakskl/task-3/internal/config"
	"github.com/Mmmakskl/task-3/internal/logic"
	strct "github.com/Mmmakskl/task-3/internal/structures"
)

func Run(configPath string) error {

	conf := strct.Config{}
	if err := config.LoadConfig(configPath, &conf); err != nil {
		return fmt.Errorf("Error execution LoadConfig: %w", err)
	}

	fmt.Printf("Running with config %v\n", conf)

	valutes := strct.ValCurs{}
	if err := logic.Parser(conf.InputFile, &valutes); err != nil {
		return fmt.Errorf("Error execution Parser: %w", err)
	}

	if err := logic.WriteJSON(conf.OutputFile, &valutes); err != nil {
		return fmt.Errorf("Error execution WriteJSON: %w", err)
	}

	fmt.Println("\nData successfully written to JSON")
	return nil
}
