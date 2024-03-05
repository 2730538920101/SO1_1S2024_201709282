#!/bin/bash

# Navegar a la carpeta del primer módulo
cd /home/omen/Escritorio/SO1_1S2024_201709282/proyectos/proyecto1/modulo_ram

# Limpiar y compilar el primer módulo
make clean
make all

# Cargar el primer módulo
sudo insmod modulo_ram.ko

# Navegar a la carpeta del segundo módulo
cd /home/omen/Escritorio/SO1_1S2024_201709282/proyectos/proyecto1/modulo_cpu

# Limpiar y compilar el segundo módulo
make clean
make all

# Cargar el segundo módulo
sudo insmod modulo_cpu.ko
