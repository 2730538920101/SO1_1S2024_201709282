import React, {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import 'bootstrap/dist/css/bootstrap.min.css'
import {
  CircularProgressbar,
  buildStyles
} from "react-circular-progressbar";
import "react-circular-progressbar/dist/styles.css";

function App() {
    const [total_memoria, setTotal_memoria] = useState("");
    const [memoria_libre, setMemoria_libre] = useState("");
    const [memoria_utilizada, setMemoria_utilizada] = useState("");
    const [porcentaje_utilizado, setPorcentaje_utilizado] = useState(""); 
    
    const updateResultText = (result) => {
        console.log("Received JSON from backend:", result);
    
        try {
            const objetoJSON = JSON.parse(result);
            console.log("Parsed JSON object:", objetoJSON);
    
            // Check if the field names match the actual structure of the JSON object
            console.log("Total memoria:", objetoJSON.informacion_memoria.total_memoria);
            console.log("Memoria libre:", objetoJSON.informacion_memoria.memoria_libre);
            console.log("Memoria utilizada:", objetoJSON.informacion_memoria.memoria_utilizada);
            console.log("Porcentaje utilizado:", objetoJSON.informacion_memoria.porcentaje_utilizado);
    
            // Update state
            setTotal_memoria(objetoJSON.informacion_memoria.total_memoria);
            setMemoria_libre(objetoJSON.informacion_memoria.memoria_libre);
            setMemoria_utilizada(objetoJSON.informacion_memoria.memoria_utilizada);
            setPorcentaje_utilizado(objetoJSON.informacion_memoria.porcentaje_utilizado);
        } catch (error) {
            console.error("Error parsing JSON:", error);
        }
    };

    useEffect(()=>{
        const intervalID = setInterval( () => {
            greet();
        }, 500);
        return () => clearInterval(intervalID);
    }, []);

    function greet() {
        Greet().then(updateResultText);
    }

    return (
        <div className='App-header'>
            <div className='container text-center'>
                <h1 className='text-white'>Uso de memoria RAM</h1>
                <div className='row'>
                <div className='col'>
                    <div className='ram-info mb-3'>
                    <div className='ram-title' style={{ color: "black", fontSize: "20px" }}>RAM Total</div>
                    <div className='ram-value' style={{ color: "black", fontSize: "20px" }}>{total_memoria} MB</div>
                    </div>
                </div>
                <div className='col'>
                    <div className='ram-info mb-3'>
                    <div className='ram-title' style={{ color: "black", fontSize: "20px" }}>RAM Usada</div>
                    <div className='ram-value' style={{ color: "black", fontSize: "20px" }}>{memoria_utilizada} MB</div>
                    </div>
                </div>
                <div className='col'>
                    <div className='ram-info mb-3'>
                    <div className='ram-title' style={{ color: "black", fontSize: "20px" }}>RAM Libre</div>
                    <div className='ram-value' style={{ color: "black", fontSize: "20px" }}>{memoria_libre} MB</div>
                    </div>
                </div>
            </div> 
            <div className='circular-progress mx-auto' style={{ width: '250px', height: '250px' }}>
              <CircularProgressbar
                value={porcentaje_utilizado}
                text={`${porcentaje_utilizado}%`}
                background
                backgroundPadding={6}
                styles={buildStyles({
                  backgroundColor: "blue",
                  textColor: "white",
                  pathColor: "white",
                  trailColor: "transparent",
                  Width: "5px",
                  Height: "5px",
                })}
              />
            </div>   
            </div>
        </div>
        )
    
}

export default App
