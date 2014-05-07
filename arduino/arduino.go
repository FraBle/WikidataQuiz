package arduino

import (
	// standard library
	"fmt"
	"log"

	// external packages
	"github.com/distributed/sers"

	// internal packages
	"github.com/FraBle/WikidataQuiz/config"
)

// SetColor() send the color command to the connected Arduino.
func SetColor(color string) (err error) {
	s, err := sers.Open(config.CONFIG.ComPort)
	if err != nil {
		log.Printf("Error connecting to Arduino: %v", err)
		err = fmt.Errorf("Error connecting to Arduino: %v", err)
		return
	}
	_, err = s.Write([]byte(color))
	if err != nil {
		log.Printf("Error setting LED color: %v", err)
		err = fmt.Errorf("Error setting LED color: %v", err)
	}
	s.Close()
	return
}
