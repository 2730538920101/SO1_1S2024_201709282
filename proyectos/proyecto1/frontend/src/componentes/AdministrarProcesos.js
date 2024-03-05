import React, { useState, useEffect } from 'react';
import '../estilos/AdministrarProcesos.css'

const AdministrarProcesos = () => {

  const [procesos, setProcesos] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:5000/lista_pid_procesos');
        const data = await response.json();
        setProcesos(data.pids);

        // Agrega un console.log para mostrar la ejecución de la petición
        console.log('Petición ejecutada a las', new Date());
      } catch (error) {
        console.error('Error al obtener datos:', error);
      }
    };

    const intervalId = setInterval(fetchData, 500);

    // Limpia el intervalo al desmontar el componente
    return () => clearInterval(intervalId);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []); // El segundo parámetro [] asegura que useEffect se ejecute solo una vez al montar el componente


  return (
    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>ADMINISTRACION DE ESTADOS DE LOS PROCESOS</h1>
        <div className='contenedor_arbol'>
          <div className='contenedor_select'>
            <h2> SELECCIONA EL PID DEL PROCESO PARA GENERAR SU DIAGRAMA DE ESTADOS </h2>
            <select className='select_procesos'>
              {procesos.map((pid) => (
                <option key={pid} value={pid}>
                  Proceso {pid}
                </option>
              ))}
            </select>
          </div>
          <div className='contenedor_grafica_arbol'></div>
        </div>
      </div>
    </div>
  );
};

export default AdministrarProcesos;