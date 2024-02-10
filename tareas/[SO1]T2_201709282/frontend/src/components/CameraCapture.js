import React, { useRef, useState, useEffect } from 'react';

const CameraCapture = ({ onSend }) => {
  const videoRef = useRef(null);
  const [imageData, setImageData] = useState(null);
  const mediaStreamRef = useRef(null);

  const startCamera = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ video: true });
      videoRef.current.srcObject = stream;
      mediaStreamRef.current = stream;
    } catch (error) {
      console.error('Error al acceder a la cámara:', error);
    }
  };

  useEffect(() => {
    startCamera();

    return () => {
      if (mediaStreamRef.current) {
        const tracks = mediaStreamRef.current.getTracks();
        tracks.forEach(track => track.stop());
      }
    };
  }, []);

  const handleCapture = () => {
    const video = videoRef.current;
  
    if (video) {
      const canvas = document.createElement('canvas');
      canvas.width = video.videoWidth || 640;  // Valor predeterminado o un valor adecuado
      canvas.height = video.videoHeight || 480; // Valor predeterminado o un valor adecuado
  
      const context = canvas.getContext('2d');
      context.drawImage(video, 0, 0, canvas.width, canvas.height);
  
      const dataURL = canvas.toDataURL('image/png');
      setImageData(dataURL);
    } else {
      console.error('Elemento de video no definido');
    }
  };

  const handleSend = () => {
    if (imageData) {
      const formData = new FormData();
      const blobData = dataURLtoBlob(imageData, 'image/png'); // Especificar el tipo como 'image/png'
      formData.append('image', blobData);
      const currentDate = new Date().toISOString();
      formData.append('date', currentDate);
      onSend(formData);
    }
  };

  const handleNewPhoto = () => {
    setImageData(null);
    startCamera();
  };

  const dataURLtoBlob = (dataURL, type) => {
    const arr = dataURL.split(',');
    const mime = type || arr[0].match(/:(.*?);/)[1]; // Usar el tipo proporcionado o extraerlo del dataURL
    const bstr = atob(arr[1]);
    let n = bstr.length;
    const u8arr = new Uint8Array(n);

    while (n--) {
      u8arr[n] = bstr.charCodeAt(n);
    }

    return new Blob([u8arr], { type: mime });
  };

  return (
    <div>
      {imageData ? (
        <img src={imageData} alt="Captured" style={{ marginTop: '10px', maxWidth: '100%' }} />
      ) : (
        <video ref={videoRef} autoPlay />
      )}
      <button onClick={handleCapture}>Capturar foto</button>
      <button onClick={handleNewPhoto}>Tomar otra foto</button>
      <button onClick={handleSend}>Enviar foto</button>
    </div>
  );
};

export default CameraCapture;