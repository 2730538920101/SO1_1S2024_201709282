# PASOS PARA INSTALAR, COMPILAR Y EJECUTAR EL SERVIDOR HTTP DE RUST

## PASO 1
Instalar RUST con el comando curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

### Install wasm target 
Instalar wasm32-wasi para compilar en wasm con el comando rustup target add wasm32-wasi

### Install WasmEdge
Instalar la herramienta WasmEdge para correr nuestro wasm con el comando curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash

## PASO 2
Crear el proyecto de RUST con el comando cargo new <nombre del proyecto>

## PASO 3 
Agregar las dependencias necesarias al archivo Cargo.toml

[package]
name = "wasmedge_reqwest_demo"
version = "0.1.0"
edition = "2021"

[dependencies]
reqwest_wasi = { version = "0.11", features = ["wasmedge-tls"] }
tokio_wasi = { version = "1", features = ["rt", "macros", "net", "time"] }

## PASO 4 
Probar el servidor con el siguiente comando cargo run --release

## PASO 5
Si todo funciona bien, compilar el programa y prepararlo para produccion con el siguiente comando cargo build --target wasm32-wasi --release

## PASO 6 
Correr la compilacion de nuestro wasm con el comando wasmedge target/wasm32-wasi/release/main.wasm



