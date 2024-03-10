import React, { useState } from 'react';
import Graphviz from 'graphviz-react';
import '../estilos/AdministrarProcesos.css';

const AdministrarProcesos = () => {
  const serverUrl = process.env.REACT_APP_SERVER_URL;
  const [pidProceso, setPidProceso] = useState(null);
  const [buttonsDisabled, setButtonsDisabled] = useState(false);
  const [graphData, setGraphData] = useState({
    nodes: [{ id: 0, label: 'NEW', color: 'green' }],
    edges: [],
  });

  const updateGraphData = (label, color) => {
    const newNodeId = graphData.nodes.length;
    const newGraphData = {
      nodes: [
        ...graphData.nodes,
        { id: newNodeId, label, color },
      ],
      edges: [...graphData.edges, { from: newNodeId - 1, to: newNodeId, color }],
    };

    setGraphData(newGraphData);
  };

  const handleStartProcess = async () => {
    try {
      const response = await fetch(`${serverUrl}/start`, { method: 'POST' });
      const data = await response.text();
      console.log(data);
      const pid = parseInt(data.split(' ')[4]);
      setPidProceso(pid);

      updateGraphData('READY', 'blue');
      updateGraphData('RUNNING', 'green');
    } catch (error) {
      console.error('Error al iniciar el proceso:', error);
    }
  };

  const handleStopProcess = async () => {
    if (pidProceso !== null) {
      try {
        await fetch(`${serverUrl}/stop?pid=${pidProceso}`, { method: 'POST' });
        console.log('Process stopped');

        updateGraphData('READY', 'blue');
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
        await fetch(`${serverUrl}/ready?pid=${pidProceso}`, { method: 'POST' });
        console.log('Process set to READY');

        updateGraphData('RUNNING', 'green');
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
        await fetch(`${serverUrl}/kill?pid=${pidProceso}`, { method: 'POST' });
        console.log('Process killed');

        updateGraphData('TERMINATED', 'red');
        // No establecer setPidProceso(null) aquí para que la etiqueta con el PID no desaparezca

        setButtonsDisabled(true);
      } catch (error) {
        console.error('Error al terminar el proceso:', error);
      }
    } else {
      console.error('No hay proceso para terminar');
    }
  };

  const handleReset = () => {
    resetGraph();
    setButtonsDisabled(false);
    // Establecer setPidProceso(null) aquí para que la etiqueta con el PID desaparezca al presionar reset
    setPidProceso(null);
  };

  const handleNewProcess = () => {
    resetGraph();
    handleStartProcess();
  };

  const resetGraph = () => {
    setGraphData({
      nodes: [{ id: 0, label: 'NEW', color: 'green' }],
      edges: [],
    });
    // No establecer setPidProceso(null) aquí para que la etiqueta con el PID no desaparezca
  };

  const generateDot = (graphData) => {
    const { nodes, edges } = graphData;
    const dot = `digraph G {
      ${nodes.map((node) => `${node.id} [label="${node.label}", color="${node.color}"];`).join('\n')}
      ${edges.map((edge) => `${edge.from} -> ${edge.to} [color="${edge.color}"];`).join('\n')}
    }`;
    return dot;
  };

  return (
    <div className='lienzo'>
      <div className='contenedor_principal'>
        <h1>ADMINISTRACION DE ESTADOS DE LOS PROCESOS</h1>
        <div className='contenedor_arbol'>
          <h2> ESTADOS DE LOS PROCESOS </h2>
          <div className='contenedor botones'>
            <button className='btn_new' onClick={handleNewProcess} disabled={buttonsDisabled}>
              NEW
            </button>
            <button className='btn_stop' onClick={handleStopProcess} disabled={buttonsDisabled}>
              STOP
            </button>
            <button className='btn_ready' onClick={handleReadyProcess} disabled={buttonsDisabled}>
              READY
            </button>
            <button className='btn_kill' onClick={handleKillProcess} disabled={buttonsDisabled}>
              KILL
            </button>
            <button className='btn_reset' onClick={handleReset}>
              RESET
            </button>
            {pidProceso !== null && <h1 style={{ color: 'black', fontSize: '20px' }}>PID del proceso: {pidProceso}</h1>}
          </div>
          <div className='contenedor_grafica_arbol'>
            <Graphviz className='grafo' dot={generateDot(graphData)} options={{ width: 600, height: 400 }} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default AdministrarProcesos;
