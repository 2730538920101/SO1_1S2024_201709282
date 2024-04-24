const express = require('express');
const router = express.Router();
const { getLogs } = require('../controllers/logsController'); // Importa el controlador de logs

// Define la ruta para obtener logs
router.get('/logs', getLogs);

module.exports = router;
