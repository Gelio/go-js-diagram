package geolocation

import (
	"errors"
	"syscall/js"
)

type Coords struct {
	Latitude, Longitude float64
	Err                 error
}

func GetLocation() (chan Coords, error) {
	geolocation := js.Global().Get("navigator").Get("geolocation")
	coordsChan := make(chan Coords)
	releaseChan := make(chan struct{})

	if geolocation.IsUndefined() {
		return nil, errors.New("Geolocation is not supported")
	}

	successHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsCoords := args[0].Get("coords")

		coordsChan <- Coords{
			Latitude:  jsCoords.Get("latitude").Float(),
			Longitude: jsCoords.Get("longitude").Float(),
		}
		releaseChan <- struct{}{}

		return nil
	})

	errorHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		coordsChan <- Coords{
			Err: errors.New("Cannot retrieve location"),
		}
		releaseChan <- struct{}{}

		return nil
	})

	geolocation.Call("getCurrentPosition", successHandler, errorHandler)

	defer func() {
		go func() {
			<-releaseChan
			successHandler.Release()
			errorHandler.Release()
		}()
	}()

	return coordsChan, nil
}
