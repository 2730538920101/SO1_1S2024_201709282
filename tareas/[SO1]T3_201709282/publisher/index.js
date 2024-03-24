const { createClient } = require("redis");
(async () => {
    const client = createClient({
        socket: {
            host: process.env.REDIS_HOST,
            port: process.env.REDIS_PORT,
        }
    });
    client.on("error", (error) => console.log("ERROR CON EL CLIENTE REDIS", error));
    await client.connect();
    console.log("CONEXION EXITOSA");
    setInterval(async () => {
        const msg = JSON.stringify({ msg: "HOLA MUNDO, SALE 100 :) " });
        console.log("PUBLICANDO...", msg);
        try {
            const result = await client.publish("test", msg);
            console.log("MENSAJE PUBLICADO EXITOSAMENTE", result);
        } catch (error) {
            console.log("ERROR AL PUBLICAR EL MENSAJE", error);
        }
    }, 3000);
})();