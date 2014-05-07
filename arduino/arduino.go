package arduino

import (
	"github.com/FraBle/WikidataQuiz/config"
	"github.com/distributed/sers"
	"log"
)

func SetColor(color string) (err error) {
	s, err := sers.Open(config.CONFIG.ComPort)
	if err != nil {
		log.Printf("Error connecting to Arduino: %v", err)
		return
	}
	_, err = s.Write([]byte(color))
	if err != nil {
		log.Printf("Error setting LED to green: %v", err)
	}
	s.Close()
	return
}
