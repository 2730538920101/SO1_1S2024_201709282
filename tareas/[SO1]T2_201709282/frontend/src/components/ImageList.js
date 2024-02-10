import React, { useState, useEffect } from 'react';

const ImageList = () => {
  const [imageList, setImageList] = useState([]);

  const fetchImageList = async () => {
    try {
      // Realiza una solicitud GET para obtener la lista de imágenes desde tu servidor
      const response = await fetch('http://localhost:5000/getImages');
      if (response.ok) {
        const data = await response.json();
        setImageList(data);
      } else {
        console.error('Error al obtener la lista de imágenes');
      }
    } catch (error) {
      console.error('Error al obtener la lista de imágenes:', error);
    }
  };

  useEffect(() => {
    // Llama a la función de obtención de la lista de imágenes cuando el componente se monta
    fetchImageList();
  }, []);

  const handleUpdate = () => {
    // Actualiza la lista de imágenes cuando se hace clic en el botón
    fetchImageList();
  };

  return (
    <div>
      <h2>Lista de Imágenes</h2>
      <ul>
        {imageList.map((image, index) => (
          <li key={index}>{image.filename + "fecha: " + image.date}</li>
        ))}
      </ul>
      <button onClick={handleUpdate}>Actualizar Datos</button>
    </div>
  );
};

export default ImageList;
