// Navbar.js
import React from 'react';
import { Link } from 'react-router-dom';
import '../Navbar.css';

const Navbar = () => {
  return (
    <nav className="navbar">
      <ul className="navbar-list">
        <li><Link to="/">DATOS DEL ESTUDIANTE</Link></li>
        <li><Link to="/GraficasTiempoReal">GRAFICAS EN TIEMPO REAL</Link></li>
        <li><Link to="/GraficasHistorico">GRAFICAS DE HISTORICOS</Link></li>
        <li><Link to="/ArbolProcesos">ARBOL DE PROCESOS</Link></li>
        <li><Link to="/AdministrarProcesos">ESTADOS DE PROCESOS</Link></li>
      </ul>
    </nav>
  );
};

export default Navbar;
