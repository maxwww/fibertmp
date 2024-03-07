package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		method1 := ctx.Method()
		method2 := fmt.Sprintf("%s", ctx.Method())

		if err := ctx.Next(); err != nil {
			return err
		}

		go func(id, method1, method2 string) {
			time.Sleep(1 * time.Second)
			fmt.Printf("==========\n")
			fmt.Printf("id -> %s\n", id)
			fmt.Printf("method1 -> %s\n", method1)
			fmt.Printf("method2 -> %s\n", method2)
		}(id, method1, method2)
		return nil
	})
	app.Post("/", func(ctx *fiber.Ctx) error {
		log.Println("in GET")
		return ctx.SendString("hello from post")
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		log.Println("in POST")
		return ctx.SendString("hello from get")
	})

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		log.Panic(err)
		return
	}

	_, err = app.Test(req, 1)
	if err != nil {
		log.Panic(err)
	}

	req, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Panic(err)
		return
	}

	_, err = app.Test(req, 1)
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(2 * time.Second)
}
