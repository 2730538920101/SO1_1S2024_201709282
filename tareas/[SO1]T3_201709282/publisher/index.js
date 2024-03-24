require('dotenv').config();
const redis = require('redis');

// Conectarse a la instancia de Redis en Memory Store
const client = redis.createClient({
   host: process.env.REDIS_HOST,
   port: process.env.REDIS_PORT,
   password: process.env.REDIS_PASSWORD
});

let counter = 0;
const maxMessages = 10; // Cambia esto según tus necesidades

// Publicar un mensaje en el canal 'test' y cerrar el cliente después de un número determinado de repeticiones
function publishMessage() {
   const message = JSON.stringify({ msg: "Hola a todos" });
   client.publish('test', message, () => {
      counter++;
      if (counter >= maxMessages) {
         clearInterval(intervalId);
         client.quit(); // Cerrar el cliente después de completar todas las publicaciones
      }
   });
}

// Publicar un mensaje cada segundo (solo para demostración)
const intervalId = setInterval(publishMessage, 1000);
