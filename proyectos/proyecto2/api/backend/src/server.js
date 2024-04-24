// Parches
const { inject, errorHandler } = require('express-custom-error');
inject(); // Parchea Express para usar sintaxis async/await

// Requiere dependencias
const express = require('express');
const cookieParser = require('cookie-parser');
const cors = require('cors');
const helmet = require('helmet');
const logger = require('./util/logger');
require('mandatoryenv').load([
    'DB_URL',
    'PORT',
    'MONGO_DATABASE',
    'MONGO_COLLECTION'
]);

const { PORT } = process.env;

// Instancia una aplicación Express
const app = express();

// Configura la instancia de la aplicación Express
app.use(express.json({ limit: '50mb' }));
app.use(express.urlencoded({ extended: true, limit: '10mb' }));

// Configura el middleware de registro personalizado
app.use(logger.dev, logger.combined);
app.use(cookieParser());
const corsOptions = {
    origin: '*', // Cambia esto por la URL de tu frontend
    methods: ['GET', 'POST', 'PUT', 'DELETE'], // Métodos HTTP permitidos
    allowedHeaders: ['Content-Type', 'Authorization'], // Encabezados permitidos
    credentials: true // Si deseas permitir el uso de cookies, tokens, etc.
};
app.use(cors(corsOptions));

app.use(helmet());

// Configura el encabezado JSON para cada respuesta
app.use('*', (req, res, next) => {
    res.setHeader('Content-Type', 'application/json');
    next();
});

// Asigna rutas
app.use('/', require('./routes/router'));

// Manejador de errores
app.use(errorHandler());

// Ruta no válida
app.use('*', (req, res) => {
    res.status(404).json({ status: false, message: 'Endpoint no encontrado' });
});

// Abre el servidor en el puerto seleccionado
app.listen(PORT, () => console.info('Servidor escuchando en el puerto', PORT));
