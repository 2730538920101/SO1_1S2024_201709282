const database = require('../models/database'); // Asegúrate de que la ruta es correcta
const { MONGO_COLLECTION } = process.env;
// Controlador para obtener logs desde MongoDB
const getLogs = async (req, res, next) => {
    try {
        console.log('Intentando obtener la conexión a la base de datos...');
        const db = await database.getDb(); // Obtén la conexión a la base de datos
        console.log('Conexión a la base de datos establecida correctamente.');

        console.log('Obteniendo la colección');
        const collection = db.collection(MONGO_COLLECTION); // Especifica la colección que contiene los logs
        console.log('Colección "voto" obtenida.');

        console.log('Obteniendo todos los registros de la colección...');
        const logs = await collection.find().toArray();
        console.log('Datos obtenidos de la colección', logs);


        if (logs.length === 0) {
            console.log('No se encontraron registros en la colección.');
        }

        // Responde con los logs obtenidos
        res.json({ status: true, logs });
        console.log('Respuesta enviada con éxito.');

    } catch (error) {
        console.error('Error al obtener logs:', error);
        next(error); // Pasa el error al manejador de errores
    }
};

module.exports = {
    getLogs
};
