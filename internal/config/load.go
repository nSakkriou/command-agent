package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nSakkriou/command-agent/internal/types"
	"github.com/nSakkriou/command-agent/internal/util"
)

var globalConf *types.Config
var isLoad = false

// Charger la config
func Load() error {
	byteValue, err := os.ReadFile(util.END_COMMAND_FILE)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return err
	}

	if err := json.Unmarshal(byteValue, &globalConf); err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	isLoad = true
	fmt.Println(globalConf.EndCommands[0].EndpointName)

	return nil
}

func IsLoad() bool {
	return isLoad
}
