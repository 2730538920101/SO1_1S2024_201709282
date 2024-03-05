import '../estilos/GraficasHistorico.css'
import React, { useState, useRef } from 'react';
const GraficasHistorico = () => {
  const [datosCpu, setDatosCpu] = useState(null);
  const [datosRam, setDatosRam] = useState(null);
  const cpuChartRef = useRef(null);
  const ramChartRef = useRef(null);
  const obtenerDatosHistoricos = async (endpoint) => {
    try {
      console.log(`Realizando solicitud a: http://localhost:5000/${endpoint}`);
      const response = await fetch(`http://localhost:5000/${endpoint}`);
      if (!response.ok) {
        throw new Error(`Error al obtener datos históricos (${response.status}): ${await response.text()}`);
      }
      const datos = await response.json();
      console.log(`Datos históricos obtenidos correctamente de ${endpoint}:`, datos);
      return datos;
    } catch (error) {
      console.error(`Error al obtener datos históricos de ${endpoint}:`, error.message);
      return null;
    }
  };

  const handleActualizarCpu = async () => {
    console.log('Clic en actualizar CPU');
    const datosCpuaux = await obtenerDatosHistoricos('historico_cpu');
    if (datosCpuaux) {
      setDatosCpu(datosCpu);
      // Actualizar gráfica de CPU
    }
  };
  
  const handleActualizarRam = async () => {
    console.log('Clic en actualizar RAM');
    const datosRamaux = await obtenerDatosHistoricos('historico_ram');
    if (datosRamaux) {
      setDatosRam(datosRam);
      // Actualizar gráfica de RAM
    }
  };



  return (

    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>GRAFICAS DE HISTORICOS MEMORIA RAM Y CPU</h1>
        <hr></hr>
        <h3>PORCENTAJE DE UTILIZACION EN LOS ULTIMOS 10 MINUTOS</h3>
        <div className='contenedor_graficas'>
          <div className='grafica_izquierda'>
            <div className='titulo_btn'>
              <h2 className='subtitulo_izq'>Datos de CPU:</h2>
              <button className='btn_izq' onClick={handleActualizarCpu}>ACTUALIZAR</button>
              
            </div>
            <div className='panel_grafica'>
              <canvas ref={cpuChartRef} />
            </div>
          </div>
        <div className='grafica_derecha'>
          <div className='titulo_btn'>
            <h2 className='subtitulo_der'>Datos de RAM:</h2>
            <button className='btn_der' onClick={handleActualizarRam} >ACTUALIZAR</button>
          </div>
          <div className='panel_grafica'>
            <canvas ref={ramChartRef} />
          </div>
        </div>
        </div>
      </div>
    </div>
  );
};

export default GraficasHistorico;