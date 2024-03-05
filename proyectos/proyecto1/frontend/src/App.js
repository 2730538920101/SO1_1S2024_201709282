// App.js
import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './componentes/Navbar';
import GraficasTiempoReal from './componentes/GraficasTiempoReal';
import GraficasHistorico from './componentes/GraficasHistorico';
import ArbolProcesos from './componentes/ArbolProcesos';
import AdministrarProcesos from './componentes/AdministrarProcesos';
import Inicio from './componentes/Inicio';

const App = () => {
  return (
    <Router>
        <div className='contenedor_nav'>
          <Navbar />
        </div>
          <Routes>
            <Route path="/" element={<Inicio />} />
            <Route path="/GraficasTiempoReal" element={<GraficasTiempoReal />} />
            <Route path="/GraficasHistorico" element={<GraficasHistorico />} />
            <Route path="/ArbolProcesos" element={<ArbolProcesos />} />
            <Route path="/AdministrarProcesos" element={<AdministrarProcesos />} />
          </Routes>
    </Router>
  );
};

export default App;
