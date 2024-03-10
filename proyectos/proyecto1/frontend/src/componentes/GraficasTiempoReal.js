import React, { useEffect, useState, useRef } from 'react';
import Chart from 'chart.js/auto';
import '../estilos/GraficasTiempoReal.css';

const GraficasTiempoReal = () => {
  const serverUrl = process.env.REACT_APP_SERVER_URL;
  const [datosCpu, setDatosCpu] = useState(null);
  const [datosRam, setDatosRam] = useState(null);
  const cpuChartRef = useRef(null);
  const ramChartRef = useRef(null);

  const obtenerDatos = async (url) => {
    try {
      console.log(`Realizando solicitud a: ${url}`);
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`Error al obtener datos (${response.status}): ${await response.text()}`);
      }
      const datos = await response.json();
      console.log(`Datos obtenidos correctamente de ${url}:`, datos);
      return datos;
    } catch (error) {
      console.error(`Error al obtener datos de ${url}:`, error.message);
      return null;
    }
  };

  const actualizarDatos = async () => {
    console.log('Actualizando datos...');
    console.log('URL del servidor:', serverUrl);
    // Obtener datos de CPU
    const datosCpu = await obtenerDatos(`${serverUrl}/cpu`);
    if (datosCpu) {
      console.log('Datos de CPU:', datosCpu);
      setDatosCpu(datosCpu);
    }

    // Obtener datos de RAM
    const datosRam = await obtenerDatos(`${serverUrl}/ram`);
    if (datosRam) {
      console.log('Datos de RAM:', datosRam);
      setDatosRam(datosRam);
    }
  };

  useEffect(() => {
    // Actualizar datos inicialmente
    actualizarDatos();

    // Configurar la actualización de datos a intervalos regulares
    const intervaloActualizacion = setInterval(actualizarDatos, 5000);

    // Limpiar el intervalo al desmontar el componente
    return () => clearInterval(intervaloActualizacion);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    // Actualizar gráficas cuando cambian los datos de CPU y RAM
    if (cpuChartRef.current && datosCpu) {
      actualizarGrafica(
        cpuChartRef.current,
        'cpuChart',
        datosCpu.informacion_cpu.porcentaje_utilizacion_cpu
      );
    }
    if (ramChartRef.current && datosRam) {
      actualizarGrafica(
        ramChartRef.current,
        'ramChart',
        datosRam.informacion_memoria.porcentaje_utilizado
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [datosCpu, datosRam]);

  const actualizarGrafica = (chartRef, label, porcentaje) => {
    const etiquetas = [`${label} Usado (${porcentaje}%)`, `${label} No Usado (${100 - porcentaje}%)`];

    if (!chartRef.data) {
      // Si chartRef.data es undefined, inicializar el gráfico
      chartRef.data = {
        labels: etiquetas,
        datasets: [
          {
            data: [porcentaje, 100 - porcentaje],
            backgroundColor: ['#FF6384', '#36A2EB'],
            hoverBackgroundColor: ['#FF6384', '#36A2EB'],
          },
        ],
      };
    } else {
      // Si chartRef.data ya está inicializado, actualizar los datos
      chartRef.data.labels = etiquetas;
      chartRef.data.datasets[0].data = [porcentaje, 100 - porcentaje];
    }

    if (!chartRef.current) {
      // Si chartRef.current es undefined, inicializar el gráfico
      chartRef.current = new Chart(chartRef, {
        type: 'pie',
        data: chartRef.data,
      });
    } else {
      // Si chartRef.current ya está inicializado, actualizar el gráfico
      chartRef.current.update();
    }
  };

  return (
    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>GRAFICAS DE USO DE CPU Y RAM EN TIEMPO REAL</h1>

        <div className='contenedor_graficas'>
          <div className='grafica_izquierda'>
            <h2>Datos de CPU:</h2>
            <canvas ref={cpuChartRef} />
          </div>
          <div className='grafica_derecha'>
            <h2>Datos de RAM:</h2>
            <canvas ref={ramChartRef} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default GraficasTiempoReal;
