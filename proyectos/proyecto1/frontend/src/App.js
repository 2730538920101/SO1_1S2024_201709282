import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './componentes/Navbar';
import GraficasTiempoReal from './componentes/GraficasTiempoReal';
import GraficasHistorico from './componentes/GraficasHistorico';
import ArbolProcesos from './componentes/ArbolProcesos'
import AdministrarProcesos from './componentes/AdministrarProcesos';

const App = () => {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/GraficasTiempoReal" element={<GraficasTiempoReal/>} />
        <Route path="/GraficasHistorico" element={<GraficasHistorico/>} />
        <Route path="/ArbolProcesos" element={<ArbolProcesos/>} />
        <Route path="/AdministrarProcesos" element={<AdministrarProcesos/>} />
        <Route path="/" element={<div className='ContenedorTitutlo'><h1>PROYECTO 1 LABORATORIO DE SISTEMAS OPERATIVOS 1</h1> <br></br> <h2>CARLOS JAVIER MARTINEZ POLANCO</h2> <br></br> <h3>201709282</h3></div>} />
      </Routes>
    </Router>
  );
};
export default App;
