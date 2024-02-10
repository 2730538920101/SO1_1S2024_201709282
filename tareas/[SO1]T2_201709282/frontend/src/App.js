// src/App.js
import './App.css';
import CameraCapture from './components/CameraCapture';
import ImageList from './components/ImageList';


function App() {
  const handleCapture = async (imageFormData) => {
    try {
      // Enviar la imagen a la API usando fetch
      const response = await fetch('http://localhost:5000/upload', {
        method: 'POST',
        body: imageFormData, // Usar FormData directamente
      });

      if (response.ok) {
        const data = await response.json();
        console.log('Respuesta del servidor:', data.message);
      } else {
        const error = await response.json();
        console.error('Error al enviar la imagen:', error.message);
      }
    } catch (error) {
      console.error('Error al enviar la imagen:', error);
    }
  };

  return (
    <div>
      <h1>Captura de Foto</h1>
      <CameraCapture onSend={handleCapture} />
      <br></br>
      <ImageList />
    </div>
  );
}

export default App;
