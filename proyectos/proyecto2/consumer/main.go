package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Carga las variables de entorno desde el archivo .env
func cargarVariablesEntorno() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se encuentra el archivo .env, puede estar en un entorno de producción, cargando variables de entorno del sistema...")
	}
}

// Configura la conexión a MongoDB
func configurarMongoDB() (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("error al conectar con MongoDB: %w", err)
	}

	// Verifica la conexión con MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error al verificar la conexión con MongoDB: %w", err)
	}

	fmt.Println("Conectado a MongoDB con éxito")
	return client, nil
}

// Configura la conexión a Redis
func configurarRedis() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")         // Dirección del servicio de Redis
	redisPassword := os.Getenv("REDIS_PASSWORD") // Contraseña de Redis

	// Crea un cliente Redis con la dirección y contraseña
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,     // Dirección de Redis (host:port)
		Password: redisPassword, // Contraseña de Redis
		DB:       0,             // Base de datos de Redis (por defecto es 0)
	})

	// Verifica la conexión con Redis
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error al conectar con Redis: %w", err)
	}

	fmt.Println("Conectado a Redis con éxito")
	return client, nil
}

// Inserta logs en MongoDB
func insertarEnMongoDB(client *mongo.Client, log string) error {
	// Selecciona la base de datos y la colección
	collection := client.Database("proyecto2").Collection("voto")

	// Inserta el log como un documento en MongoDB
	doc := map[string]interface{}{
		"message": log,
	}

	_, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("error al insertar log en MongoDB: %w", err)
	}

	fmt.Println("Log insertado en MongoDB con éxito")
	return nil
}
func consumeFromKafka() {
	// Carga las variables de entorno
	cargarVariablesEntorno()

	// Configura el consumidor de Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_BROKER")},
		Topic:     os.Getenv("KAFKA_TOPIC"),
		Partition: 0,
		MinBytes:  10e3, // Tamaño mínimo de bytes para solicitar datos
		MaxBytes:  10e6, // Tamaño máximo de bytes para solicitar datos
	})
	defer reader.Close()

	ctx := context.Background()
	fmt.Println("Iniciando Kafka...")

	// Configura la conexión a MongoDB
	clientMongo, err := configurarMongoDB()
	if err != nil {
		log.Fatalf("Error al configurar MongoDB: %s", err)
	}
	defer clientMongo.Disconnect(ctx)

	// Configura la conexión a Redis
	clientRedis, err := configurarRedis()
	if err != nil {
		log.Fatalf("Error al configurar Redis: %s", err)
	}
	defer clientRedis.Close()

	for {
		// Lee mensajes de Kafka
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Error al leer mensaje de Kafka: %s", err)
		}

		// Muestra la información del mensaje recibido
		msgString := string(msg.Value)
		fmt.Printf("Mensaje recibido de Kafka: %s\n", msgString)

		// Inserta el log en MongoDB
		err = insertarEnMongoDB(clientMongo, msgString)
		if err != nil {
			log.Fatalf("Error al insertar log en MongoDB: %s", err)
		}

		// Analiza el mensaje recibido para extraer los campos
		parts := strings.Split(msgString, ", ")
		fmt.Printf("Partes del mensaje separadas por ', ': %v\n", parts)
		fmt.Printf("Número de partes del mensaje: %d\n", len(parts))

		// Verifica que el mensaje contenga exactamente 4 partes
		if len(parts) != 4 {
			log.Printf("Advertencia: El mensaje no contiene la estructura esperada. Mensaje: %s", msgString)
			continue
		}

		// Extrae los campos de nombre, álbum, año y rango
		name := strings.TrimPrefix(parts[0], "name: ")
		album := strings.TrimPrefix(parts[1], "album: ")
		year := strings.TrimPrefix(parts[2], "year: ")
		rank := strings.TrimPrefix(parts[3], "rank: ")

		// Muestra los valores extraídos antes de eliminar espacios
		fmt.Printf("Valores extraídos: name='%s', album='%s', year='%s', rank='%s'\n", name, album, year, rank)

		// Elimina los espacios adicionales en los valores extraídos
		name = strings.TrimSpace(name)
		album = strings.TrimSpace(album)
		year = strings.TrimSpace(year)
		rank = strings.TrimSpace(rank)

		// Muestra los valores después de eliminar espacios
		fmt.Printf("Valores después de eliminar espacios: name='%s', album='%s', year='%s', rank='%s'\n", name, album, year, rank)

		// Verifica si los campos tienen valores válidos
		if name == "" || album == "" || year == "" || rank == "" {
			log.Printf("Advertencia: Uno o más campos están vacíos. Mensaje: %s", msgString)
			continue
		}

		// Construye la clave de Redis para el contador
		key := fmt.Sprintf("%s_%s_%s_rank_%s", name, album, year, rank)
		fmt.Printf("Clave de Redis generada: %s\n", key)

		// Incrementa el contador en Redis con la clave generada
		_, err = clientRedis.Incr(ctx, key).Result()
		if err != nil {
			log.Fatalf("Error al incrementar el contador en Redis: %s", err)
		}

		fmt.Printf("Contador incrementado para clave %s\n", key)
	}
}

func main() {
	fmt.Println("CONSUMER ACTIVADO...")
	fmt.Println("REDIS LLEVA EL CONTEO")
	fmt.Println("MONGO GUARDA LOS LOGS")
	// Inicia el consumidor de Kafka
	consumeFromKafka()
}
