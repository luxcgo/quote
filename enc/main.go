package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() (err error) {
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			newErr := errors.New("engine timeout, force quit")
			if err != nil {
				err = fmt.Errorf("%v\n%v", err, newErr)
			} else {
				err = newErr
			}
		}
	}()

	return errors.New("fuck up")
}
