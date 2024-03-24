require('dotenv').config();
const redis = require('redis');

// Conectarse a la instancia de Redis en Memory Store
const client = redis.createClient({
   host: process.env.REDIS_HOST,
   port: process.env.REDIS_PORT
});

let counter = 0;
const maxMessages = 10; // Número máximo de mensajes a publicar

// Publicar un mensaje en el canal 'test' y cerrar el cliente después de completar todas las publicaciones
function publishMessage() {
   const message = JSON.stringify({ msg: "Hola a todos" });
   client.publish('test', message, () => {
      counter++;
      if (counter >= maxMessages) {
         client.quit(); // Cerrar el cliente después de completar todas las publicaciones
      }
   });
}

// Publicar un mensaje cada segundo (solo para demostración)
const intervalId = setInterval(publishMessage, 1000);

// Detener el intervalo después de un cierto tiempo
setTimeout(() => {
   clearInterval(intervalId);
}, maxMessages * 1000); // Detener el intervalo después de que se haya completado el número máximo de mensajes
