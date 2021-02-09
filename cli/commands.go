package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleCreateCmd(args []string) error {
	if args[0] == "" {
		return errors.New("path is required")
	}
	name, args := args[0], args[1:]

	body, err := json.Marshal(map[string]interface{}{
		"path": name,
		"args": args,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, baseUrl+"/process/create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := getClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(responseBody))

	return nil
}

func HandleCancelCmd(args []string) error {
	if args[0] == "" {
		return errors.New("pid is required")
	}

	req, err := http.NewRequest(http.MethodPost, baseUrl+"/process/cancel/"+args[0], nil)
	if err != nil {
		return err
	}

	client := getClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(responseBody))
	return nil
}

func HandleGetStatusCmd(args []string) error {
	if args[0] == "" {
		return errors.New("pid is required")
	}

	req, err := http.NewRequest(http.MethodGet, baseUrl+"/process/"+args[0], nil)
	if err != nil {
		return err
	}

	client := getClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(responseBody))
	return nil
}
