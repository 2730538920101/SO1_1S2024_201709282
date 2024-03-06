import '../estilos/GraficasHistorico.css';
import React, { useState, useEffect } from 'react';
import { Line } from 'react-chartjs-2';

const GraficasHistorico = () => {
  const [datosCpu, setDatosCpu] = useState({ lista_cpu: [] });
  const [datosRam, setDatosRam] = useState({ lista_ram: [] });
  const [cpuChart, setCpuChart] = useState(null);
  const [ramChart, setRamChart] = useState(null);

  const obtenerDatosHistoricos = async (endpoint) => {
    try {
      const response = await fetch(`http://localhost:5000/${endpoint}`);
      if (!response.ok) {
        throw new Error(`Error al obtener datos históricos (${response.status}): ${await response.text()}`);
      }
      const datos = await response.json();
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
      setDatosCpu(datosCpuaux);
      console.log(datosCpu);
      // Actualizar gráfica de CPU
    }
  };

  const handleActualizarRam = async () => {
    console.log('Clic en actualizar RAM');
    const datosRamaux = await obtenerDatosHistoricos('historico_ram');
    if (datosRamaux) {
      setDatosRam(datosRamaux);
      console.log(datosRam);
      // Actualizar gráfica de RAM
    }
  };

  useEffect(() => {
    // Configuración inicial de la gráfica de CPU
    if (datosCpu.lista_cpu && datosCpu.lista_cpu.length > 0) {
      const cpuChartConfig = {
        labels: datosCpu.lista_cpu.map((data) => data.tiempo_registro),
        datasets: [
          {
            label: 'Porcentaje CPU',
            fill: false,
            lineTension: 0.1,
            backgroundColor: 'rgba(75,192,192,0.4)',
            borderColor: 'rgba(75,192,192,1)',
            borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            borderJoinStyle: 'miter',
            pointBorderColor: 'rgba(75,192,192,1)',
            pointBackgroundColor: '#fff',
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: 'rgba(75,192,192,1)',
            pointHoverBorderColor: 'rgba(220,220,220,1)',
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: datosCpu.lista_cpu.map((data) => data.porcentaje),
          },
        ],
      };
      setCpuChart(cpuChartConfig);
    }
  }, [datosCpu]);
  
  useEffect(() => {
    // Configuración inicial de la gráfica de RAM
    if (datosRam.lista_ram && datosRam.lista_ram.length > 0) {
      const ramChartConfig = {
        labels: datosRam.lista_ram.map((data) => data.tiempo_registro),
        datasets: [
          {
            label: 'Porcentaje RAM',
            fill: false,
            lineTension: 0.1,
            backgroundColor: 'rgba(255,99,132,0.4)',
            borderColor: 'rgba(255,99,132,1)',
            borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            borderJoinStyle: 'miter',
            pointBorderColor: 'rgba(255,99,132,1)',
            pointBackgroundColor: '#fff',
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: 'rgba(255,99,132,1)',
            pointHoverBorderColor: 'rgba(220,220,220,1)',
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: datosRam.lista_ram.map((data) => data.porcentaje_utilizado),
          },
        ],
      };
      setRamChart(ramChartConfig);
    }
  }, [datosRam]);

  return (
    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>GRAFICAS DE HISTORICOS MEMORIA RAM Y CPU</h1>
        <hr />
        <h3>PORCENTAJE DE UTILIZACION EN LOS ULTIMOS 10 MINUTOS</h3>
        <div className='contenedor_graficas'>
          <div className='grafica_izquierda'>
            <div className='titulo_btn'>
              <h2 className='subtitulo_izq'>Datos de CPU:</h2>
              <button className='btn_izq' onClick={handleActualizarCpu}>
                ACTUALIZAR
              </button>
            </div>
            <div className='panel_grafica'>
              {cpuChart && <Line data={cpuChart} />}
            </div>
          </div>
          <div className='grafica_derecha'>
            <div className='titulo_btn'>
              <h2 className='subtitulo_der'>Datos de RAM:</h2>
              <button className='btn_der' onClick={handleActualizarRam}>
                ACTUALIZAR
              </button>
            </div>
            <div className='panel_grafica'>
              {ramChart && <Line data={ramChart} />}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GraficasHistorico;
