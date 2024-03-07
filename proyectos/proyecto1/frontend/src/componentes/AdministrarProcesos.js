import React, { useState } from 'react';
import '../estilos/AdministrarProcesos.css';

const AdministrarProcesos = () => {
  const [pidProceso, setPidProceso] = useState(null);

  const handleStartProcess = async () => {
    try {
      const response = await fetch('http://localhost:5000/start', { method: 'POST' });
      const data = await response.text();

      // Extraer el PID del texto de la respuesta (puede variar según la respuesta real)
      const pid = parseInt(data.split(' ')[4]);

      // Actualizar el estado con el PID del proceso creado
      setPidProceso(pid);
    } catch (error) {
      console.error('Error al iniciar el proceso:', error);
    }
  };

  const handleStopProcess = async () => {
    if (pidProceso !== null) {
      try {
        const response = await fetch(`http://localhost:5000/stop?pid=${pidProceso}`, { method: 'POST' });
        const data = await response.text();
        console.log(data);
      } catch (error) {
        console.error('Error al detener el proceso:', error);
      }
    } else {
      console.error('No hay proceso para detener');
    }
  };

  const handleReadyProcess = async () => {
    if (pidProceso !== null) {
      try {
        const response = await fetch(`http://localhost:5000/ready?pid=${pidProceso}`, { method: 'POST' });
        const data = await response.text();
        console.log(data);
      } catch (error) {
        console.error('Error al poner el proceso en estado Ready:', error);
      }
    } else {
      console.error('No hay proceso para poner en estado Ready');
    }
  };

  const handleKillProcess = async () => {
    if (pidProceso !== null) {
      try {
        const response = await fetch(`http://localhost:5000/kill?pid=${pidProceso}`, { method: 'POST' });
        const data = await response.text();
        console.log(data);
      } catch (error) {
        console.error('Error al terminar el proceso:', error);
      }
    } else {
      console.error('No hay proceso para terminar');
    }
  };

  return (
    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>ADMINISTRACION DE ESTADOS DE LOS PROCESOS</h1>
        <div className='contenedor_arbol'>
          <h2> ESTADOS DE LOS PROCESOS </h2>
          <div className='contenedor botones'>
            <button className='btn_new' onClick={handleStartProcess}>
              NEW
            </button>
            <button className='btn_stop' onClick={handleStopProcess}>
              STOP
            </button>
            <button className='btn_ready' onClick={handleReadyProcess}>
              READY
            </button>
            <button className='btn_kill' onClick={handleKillProcess}>
              KILL
            </button>
            {pidProceso !== null && <h1 style={{ color: 'black', fontSize: '20px' }}>PID del proceso: {pidProceso}</h1>}
          </div>  
          <div className='contenedor_grafica_arbol'>
            
          </div>
        </div>
      </div>
    </div>
  );
};

export default AdministrarProcesos;
