const { MongoClient } = require('mongodb');

const { DB_URL, MONGO_DATABASE } = process.env;

const client = new MongoClient(DB_URL, { useUnifiedTopology: true });

module.exports = {
    async getDb() {
        if (!client.isConnected()) {
            await client.connect();
        }
        // Selecciona la base de datos "proyecto2"
        return client.db(MONGO_DATABASE);
    }
};