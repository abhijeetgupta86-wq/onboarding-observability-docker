package main

import (
	"errors"
	"math/rand"
)

func maybeFail() error {
	if rand.Intn(10) < 3 {
		return errors.New("random onboarding failure")
	}
	return nil
}

func forcedFail(mode string) error {
	if mode == "fail" {
		return errors.New("forced onboarding failure")
	}
	return nil
}
