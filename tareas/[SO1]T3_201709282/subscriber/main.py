import json
import os
import redis
from dotenv import load_dotenv

# Cargar variables de entorno desde el archivo .env
load_dotenv()

# Conectarse a la instancia de Redis en Memory Store
client = redis.StrictRedis(
    host=os.getenv('REDIS_HOST'),
    port=os.getenv('REDIS_PORT'),
    decode_responses=True
)

# Suscribirse al canal 'test' y manejar los mensajes recibidos
def handle_message(message):
    print("Mensaje recibido:", message['msg'])

pubsub = client.pubsub()
pubsub.subscribe('test')

for message in pubsub.listen():
    if message['type'] == 'message':
        handle_message(json.loads(message['data']))