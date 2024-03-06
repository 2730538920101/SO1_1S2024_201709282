import React, { useState, useEffect } from 'react';
import '../estilos/ArbolProcesos.css';
import { saveAs } from 'file-saver';
import Viz from 'viz.js';
import { Module, render } from 'viz.js/full.render.js';
import Graphviz from 'graphviz-react';

const ArbolProcesos = () => {
  const [procesos, setProcesos] = useState([]);
  const [arbolDot, setArbolDot] = useState('');
  const [arbolSVG, setArbolSVG] = useState('');

  const handleSelectChange = async (event) => {
    const selectedPid = event.target.value;

    try {
      const response = await fetch(`http://localhost:5000/generarArbol/${selectedPid}`);
      const arbolDotResponse = await response.text();

      // Verifica si el árbol DOT se obtuvo exitosamente
      if (response.ok) {
        // Actualiza el estado con el árbol DOT
        setArbolDot(arbolDotResponse);

        // Convierte el DOT a SVG usando Viz.js
        const viz = new Viz({ Module, render });
        const svg = await viz.renderString(arbolDotResponse, { format: 'svg' });
        setArbolSVG(svg);

        console.log('Árbol DOT obtenido exitosamente:', arbolDotResponse);
      } else {
        console.error('Error al obtener el árbol DOT:', arbolDotResponse);
      }
    } catch (error) {
      console.error('Error al obtener el árbol DOT:', error);
    }
  };

  const handleDownload = () => {
    // Verifica si hay un árbol SVG para descargar
    if (arbolSVG) {
      const blob = new Blob([arbolSVG], { type: 'image/svg+xml' });
      saveAs(blob, 'arbol_proceso.svg');
    } else {
      console.error('No hay árbol SVG para descargar');
    }
  };

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
        <h1>MOSTRAR ÁRBOLES DE LOS PROCESOS</h1>
        <div className='contenedor_arbol'>
          <div className='contenedor_select'>
            <h2> SELECCIONA EL PID DEL PROCESO PARA GENERAR SU ARBOL </h2>
            <select className='select_procesos' onChange={handleSelectChange}>
              {procesos.map((pid) => (
                <option key={pid} value={pid}>
                  Proceso {pid}
                </option>
              ))}
            </select>
            {/* Agrega el botón de descarga aquí */}
            <button className='boton_descarga' onClick={handleDownload}>Descargar SVG</button>
          </div>
          <div className='contenedor_grafica_arbol' style={{ textAlign: 'center' }}>
            {/* Renderiza el árbol DOT con Graphviz */}
            {arbolDot && (
              <Graphviz
                dot={arbolDot}
                options={{
                  zoom: true, // Permite hacer zoom
                  fit: true, // Hace que el gráfico se ajuste al contenedor
                  center: true, // Centra el gráfico en el contenedor
                }}
                style={{ width: '100%', height: '100%' }} // Ocupa todo el espacio disponible
              />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ArbolProcesos;
