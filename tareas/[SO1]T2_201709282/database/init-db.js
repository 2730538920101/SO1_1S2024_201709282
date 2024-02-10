// init-db.js

// Define el nombre de la base de datos
var dbName = 'tarea2';

// Crear el usuario administrador con un nombre y contraseña personalizados
db.createUser({
  user: 'admin',
  pwd: 'admin',  // Cambia 'TuPasswordSeguro' por tu contraseña segura
  roles: ['readWrite', 'dbAdmin'],
  passwordDigestor: 'server',
});

// Usar la base de datos especificada
db = db.getSiblingDB(dbName);

// Crear una colección de ejemplo
db.createCollection('images');
