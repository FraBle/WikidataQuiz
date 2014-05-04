package arduino

import (
	"github.com/distributed/sers"
	"log"
)

func SetColor(color string) (err error) {
	s, err := sers.Open("COM6")
	if err != nil {
		log.Printf("Error connecting to Arduino: %v", err)
		return
	}
	n, err := s.Write([]byte(color))
	if err != nil {
		log.Printf("Error setting LED to green: %v", err)
	}
	log.Printf("Bytes written: %d", n)
	s.Close()
	return
}
