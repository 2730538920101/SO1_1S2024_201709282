<template>
  <div class="log-viewer">
    <h1>PROYECTO 2 LABORATORIO DE SISTEMAS OPERATIVOS 1</h1>
    <h1>Logs</h1>
    <button @click="fetchLogs">Actualizar Logs</button>
    <div v-if="logs.length" class="logs-container">
      <h2>Registros:</h2>
      <ul class="logs-list">
        <li v-for="(log, index) in logs" :key="index" class="log-item">
          {{ log.message }}
        </li>
      </ul>
    </div>
    <div v-else class="no-logs">
      <p>No hay registros disponibles.</p>
    </div>
    <h1>CARLOS JAVIER MARTINEZ POLANCO 201709282</h1>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LogViewer',
  data() {
    return {
      logs: [],
    };
  },
  methods: {
    async fetchLogs() {
      try {
        // Utiliza la variable de entorno VUE_APP_API_URL
        const response = await axios.get(process.env.VUE_APP_API_URL);
        if (response.data.status) {
          this.logs = response.data.logs;
        }
      } catch (error) {
        console.error('Error al obtener los logs:', error);
      }
    }
  },
  mounted() {
    // Cargar logs cuando el componente se monte
    this.fetchLogs();
  },
};
</script>

<style scoped>
body {
  height: 100%;
}

.log-viewer {
  background-color: #282c34; /* Fondo oscuro */
  color: #abb2bf; /* Texto claro */
  font-family: 'Courier', monospace; /* Fuente de consola */
  padding: 20px; /* Relleno alrededor del contenido */
  height: 100%;
}

h1, h2 {
  color: #61dafb; /* Color del encabezado */
}

button {
  background-color: #61dafb;
  color: #282c34;
  padding: 10px;
  margin-bottom: 20px;
  border: none;
  cursor: pointer;
  font-weight: bold;
}

button:hover {
  background-color: #50c0e8;
}

.logs-container {
  background-color: #1e2227; /* Fondo oscuro para la sección de logs */
  padding: 10px;
  border-radius: 4px;
  max-height: 400px; /* Altura máxima de la sección de logs */
  overflow-y: auto; /* Permite el desplazamiento si hay muchos logs */
}

.logs-list {
  list-style: none; /* Elimina los puntos de la lista */
  padding: 0;
  margin: 0;
}

.log-item {
  margin-bottom: 10px; /* Espacio entre cada registro */
  white-space: pre-wrap; /* Permite que las líneas largas se envuelvan */
}

.no-logs {
  color: #e06c75; /* Color para el mensaje de no hay registros */
}
</style>
