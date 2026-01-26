package configs

import (
	"testing"

	"github.com/spf13/viper"
)

func TestRace(t *testing.T) {
	done := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			Get()
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			configLock.Lock()
			if err := viper.Unmarshal(config); err != nil {
				// ignore error
			}
			configLock.Unlock()
		}
		done <- true
	}()

	<-done
	<-done
}
