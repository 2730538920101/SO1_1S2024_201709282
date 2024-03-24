require('dotenv').config();
const redis = require('redis');

// Conectarse a la instancia de Redis en Memory Store
const client = redis.createClient({
   host: process.env.REDIS_HOST,
   port: process.env.REDIS_PORT
});

// Publicar un mensaje en el canal 'test'
function publishMessage() {
   const message = JSON.stringify({ msg: "Hola a todos" });
   client.publish('test', message);
}

// Publicar un mensaje cada segundo (solo para demostración)
setInterval(publishMessage, 1000);