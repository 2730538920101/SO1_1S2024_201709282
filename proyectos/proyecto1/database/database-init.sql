CREATE TABLE PROCESO (
    ID_PROCESO INT PRIMARY KEY AUTO_INCREMENT,
    PID INT,
    NOMBRE VARCHAR(255),
    ESTADO INT,
    RSS INT,
    UID INT,
    TIEMPO DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE PROCESO_HIJO (
    ID_PROCESO_HIJO INT PRIMARY KEY AUTO_INCREMENT,
    ID_PROCESO INT,
    PID INT,
    NOMBRE VARCHAR(255),
    ESTADO INT,
    RSS INT,
    UID INT,
    TIEMPO DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ID_PROCESO) REFERENCES PROCESO(ID_PROCESO) ON DELETE CASCADE
);

CREATE TABLE RAM (
    ID_RAM INT PRIMARY KEY AUTO_INCREMENT,
    TOTAL_MEMORIA INT,
    MEMORIA_LIBRE INT,
    MEMORIA_UTILIZADA INT,
    PORCENTAJE_UTILIZADO FLOAT,
    TIEMPO DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE CPU (
    ID_CPU INT PRIMARY KEY AUTO_INCREMENT,
    PORCENTAJE_UTILIZADO FLOAT,
    TIEMPO DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ESTADO (
    ID_ESTADO INT PRIMARY KEY AUTO_INCREMENT,
    ID_PROCESO INT,
    NOMBRE VARCHAR(50),
    TIEMPO DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ID_PROCESO) REFERENCES PROCESO(ID_PROCESO) ON DELETE CASCADE
)