package main

import (
	"encoding/json"
	"fmt"
	"log"

	"sigs.k8s.io/yaml"
)

type Printer interface {
	Print(interface{}) error
}

func choosePrinter(p string) Printer {
	switch p {
	case "json":
		return JSONPrinter{}
	case "yaml", "yml":
		return YAMLPrinter{}
	}

	log.Printf("warning: unknown printer '%s'. Falling back to 'json'", p)
	return JSONPrinter{}
}

type JSONPrinter struct{}

func (j JSONPrinter) Print(i interface{}) error {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

type YAMLPrinter struct{}

func (y YAMLPrinter) Print(i interface{}) error {
	data, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	fmt.Print(string(data))
	return nil
}
