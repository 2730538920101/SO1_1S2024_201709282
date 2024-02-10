const express = require('express');
const app = express();
const mongoose = require('mongoose');
const cors = require('cors');
const multer = require('multer');
const path = require('path');

// Configurar MongoDB
mongoose.connect('mongodb://database:27017/tarea2?connectTimeoutMS=30000&socketTimeoutMS=30000');

// Configurar CORS
app.use(cors());
app.use(express.json());

// Definir el modelo de la colección (puedes ajustarlo según tus necesidades)
const imageSchema = new mongoose.Schema({
  filename: String,
  date: { type: Date, default: Date.now },
});

const Image = mongoose.model('Image', imageSchema);

// Configurar Multer para manejar la subida de archivos
const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    cb(null, './uploads'); // Ruta donde se guardarán las imágenes en el servidor
  },
  filename: function (req, file, cb) {
    cb(null, Date.now() + path.extname(file.originalname)); // Nombre de archivo único
  },
});

const upload = multer({ storage: storage });

// Ruta para manejar la carga de imágenes
app.post('/upload', upload.single('image'), async (req, res) => {
  try {
    // Obtenemos el nombre del archivo subido
    const filename = req.file.filename;
    const date = req.body.date || new Date();
    // Lógica para almacenar la ruta de la imagen en MongoDB
    const newImage = new Image({ filename: filename, date: date });
    await newImage.save();

    res.status(200).json({ message: 'Imagen almacenada exitosamente.' });
  } catch (error) {
    console.error('Error al manejar la carga de imágenes:', error);
    res.status(500).json({ error: 'Error interno del servidor.' });
  }
});

app.get('/getImages', async (req, res) => {
    try {
      // Consulta la base de datos para obtener la lista de imágenes
      const images = await Image.find({}, "filename date"); 
  
      res.status(200).json(images);
    } catch (error) {
      console.error('Error al obtener la lista de imágenes:', error);
      res.status(500).json({ error: 'Error interno del servidor.' });
    }
  });

// Iniciar el servidor
const PORT = 5000;
app.listen(PORT, () => {
  console.log(`Servidor iniciado en http://localhost:${PORT}`);
});
