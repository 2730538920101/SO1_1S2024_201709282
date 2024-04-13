package main

import (
	"context"
	"fmt"
	pb "grpc_service/proto" // Importa el paquete generado a partir de tu archivo .proto
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type BandsData struct {
	name  string
	album string
	year  string
	rank  string
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	voto := BandsData{
		name:  data["name"],
		album: data["album"],
		year:  data["year"],
		rank:  data["rank"],
	}

	go sendServer(voto)
	return nil
}

func sendServer(voto BandsData) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewBandServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.SendBandInfo(ctx, &pb.Band{
		Name:  voto.name,
		Album: voto.album,
		Year:  voto.year,
		Rank:  voto.rank,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetMessage())
}

func getMessage(c *fiber.Ctx) error {
	// Aquí puedes realizar la lógica necesaria para recuperar mensajes
	// Por ejemplo, podrías consultar una base de datos o recuperar mensajes de alguna otra fuente.

	// Devuelve una respuesta apropiada
	return c.JSON(fiber.Map{
		"message": "Se recupero la informacion desde el servidor",
	})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "todo bien",
		})
	})
	app.Post("/insert", insertData)

	// Endpoint GET para recibir mensajes
	app.Get("/receive", getMessage)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
