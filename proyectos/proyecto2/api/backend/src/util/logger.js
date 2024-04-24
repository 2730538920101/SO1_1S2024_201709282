const morgan = require('morgan');
const rfs = require('rotating-file-stream');
const path = require('path');
const fs = require('fs');

// Directorio de logs
const logDirectory = path.resolve(__dirname, '../../log');
fs.existsSync(logDirectory) || fs.mkdirSync(logDirectory);

// Crea un flujo de escritura rotativo
const accessLogStream = rfs('access.log', {
    interval: '1d',
    path: logDirectory
});

module.exports = {
    dev: morgan('dev'),
    combined: morgan('combined', { stream: accessLogStream })
};
