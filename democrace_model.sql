-- Creación de la tabla Regimen
CREATE TABLE Regimen (
    id_regimen INTEGER PRIMARY KEY,
    nombre_regimen VARCHAR(255) NOT NULL,
    descripcion_regimen TEXT NOT NULL
);

-- Creación de la tabla Pais
CREATE TABLE Pais (
    id_pais INTEGER PRIMARY KEY,
    nombre_pais VARCHAR(255) NOT NULL,
    capital VARCHAR(255) NOT NULL,
    poblacion INTEGER NOT NULL,
    area FLOAT NOT NULL,
    moneda VARCHAR(255) NOT NULL,
    idioma VARCHAR(255) NOT null,
    id_regimen INTEGER NOT NULL,
    FOREIGN KEY (id_regimen) REFERENCES Regimen(id_regimen)
);

-- Creación de la tabla Gobernante
CREATE TABLE Gobernante (
    id_gobernante INTEGER PRIMARY KEY,
    nombre_gobernante VARCHAR(255) NOT NULL,
    fecha_nacimiento DATE NOT NULL,
    fecha_fallecimiento DATE,
    id_pais INTEGER NOT NULL,
    FOREIGN KEY(id_pais) REFERENCES Pais(id_pais)
);


-- Creación de la tabla Periodo
CREATE TABLE Periodo (
    id_periodo INTEGER PRIMARY KEY,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE,
    id_gobernante INTEGER NOT NULL,
    id_regimen INTEGER NOT NULL,
    FOREIGN KEY(id_gobernante) REFERENCES Gobernante(id_gobernante),
    FOREIGN KEY(id_regimen) REFERENCES Regimen(id_regimen)
);


-- hasta aqui
CREATE TABLE TipoCategoria (
    id_tipo SERIAL PRIMARY KEY,
    nombre_tipo VARCHAR(255) NOT NULL
);


CREATE TABLE Categoria (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Categoria_padre INTEGER REFERENCES Categoria(ID),
  id_tipo INTEGER REFERENCES TipoCategoria(id_tipo)
);

CREATE TABLE Cargo (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Categoria INTEGER REFERENCES Categoria(ID)
);
-- Creación de la tabla Funcionario
CREATE TABLE Funcionario (
    id_funcionario INTEGER PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    fecha_nacimiento DATE NOT NULL,
    genero CHAR(1) NOT NULL,
    fecha_inicio_cargo DATE NOT NULL,
    id_padre INTEGER,
    id_madre INTEGER,
    ID_Cargo INTEGER REFERENCES Cargo(ID),
    id_periodo INTEGER NOT NULL,
    FOREIGN KEY(id_periodo) REFERENCES Periodo(id_periodo),
    FOREIGN KEY(id_padre) REFERENCES Funcionario(id_funcionario),
    FOREIGN KEY(id_madre) REFERENCES Funcionario(id_funcionario)
);

-- Creación de la tabla Educacion
CREATE TABLE Educacion (
    id_educacion INTEGER PRIMARY KEY,
    institucion VARCHAR(255) NOT NULL,
    nivel VARCHAR(255) NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    id_funcionario INTEGER NOT NULL,
    FOREIGN KEY(id_funcionario) REFERENCES Funcionario(id_funcionario)
);

-- Creación de la tabla Experiencia Laboral
CREATE TABLE Experiencia_Laboral (
    id_experiencia INTEGER PRIMARY KEY,
    empresa VARCHAR(255) NOT NULL,
    puesto VARCHAR(255) NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    id_funcionario INTEGER NOT NULL,
    FOREIGN KEY(id_funcionario) REFERENCES Funcionario(id_funcionario)
);

-- Creación de la tabla Contacto
CREATE TABLE Contacto (
    id_contacto INTEGER PRIMARY KEY,
    ID_tipo_contacto INTEGER REFERENCES Categoria(ID),
    valor VARCHAR(255) NOT NULL,
    id_funcionario INTEGER NOT NULL,
    FOREIGN KEY(id_funcionario) REFERENCES Funcionario(id_funcionario)
);

-- Creación de la tabla Parentesco
CREATE TABLE Parentesco (
    id_parentesco INTEGER PRIMARY KEY,
    nombre_parentesco VARCHAR(255) NOT NULL,
    grado_parentesco INTEGER NOT NULL
);

-- Creación de la tabla Funcionario_Parentesco
CREATE TABLE Funcionario_Parentesco (
    id_funcionario INTEGER NOT NULL,
    id_parentesco INTEGER NOT NULL,
    id_familiar INTEGER NOT NULL,
    PRIMARY KEY (id_funcionario, id_parentesco, id_familiar),
    FOREIGN KEY(id_funcionario) REFERENCES Funcionario(id_funcionario),
    FOREIGN KEY(id_parentesco) REFERENCES Parentesco(id_parentesco),
    FOREIGN KEY(id_familiar) REFERENCES Funcionario(id_funcionario)
);

-- Creación de la tabla Ley
CREATE TABLE Ley_Paren (
    id_ley INTEGER PRIMARY KEY,
    nombre_ley VARCHAR(255) NOT NULL,
    fecha_promulgacion DATE NOT NULL
);

-- Creación de la tabla Ley_Parentesco
CREATE TABLE Ley_Parentesco (
    id_ley INTEGER NOT NULL,
    id_parentesco INTEGER NOT NULL,
    FOREIGN KEY(id_ley) REFERENCES Ley_Paren(id_ley),
    FOREIGN KEY(id_parentesco) REFERENCES Parentesco(id_parentesco)
);




CREATE TABLE Trayectoria (
  ID SERIAL PRIMARY KEY,
  Anio_Inicio INTEGER NOT NULL,
  Anio_Fin INTEGER,
  ID_Cargo INTEGER REFERENCES Cargo(ID),
  ID_Funcionario INTEGER REFERENCES Funcionario(Id_funcionario)
);

CREATE TABLE Tipo_Universidad (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL
);

CREATE TABLE Universidad (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Categoria INTEGER REFERENCES Categoria(ID),
  ID_Tipo_Universidad INTEGER REFERENCES Tipo_Universidad(ID),
  Ubicacion VARCHAR(255) NOT NULL,
  Tipo_Empresa VARCHAR(255) NOT NULL
);



CREATE TABLE Ley (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  Fecha DATE NOT NULL,
  Descripcion TEXT,
  ID_Funcionario INTEGER REFERENCES Funcionario(ID_funcionario)
);

CREATE TABLE Categoria_Ley (
  ID SERIAL PRIMARY KEY,
  ID_Categoria INTEGER REFERENCES Categoria(ID),
  ID_Ley INTEGER REFERENCES Ley(ID),
  Importancia INTEGER NOT NULL,
  Calificacion INTEGER NOT NULL
);

CREATE TABLE Calificacion_Funcionario (
  ID SERIAL PRIMARY KEY,
  ID_Funcionario INTEGER REFERENCES Funcionario(ID_funcionario),
  ID_Categoria INTEGER REFERENCES Categoria(ID),
  Calificacion INTEGER NOT NULL
);

CREATE TABLE Nivel_Educativo (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL
);

CREATE TABLE Area_Estudio (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Nivel_Educativo INTEGER REFERENCES Nivel_Educativo(ID)
);

CREATE TABLE Especialidad (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Area_Estudio INTEGER REFERENCES Area_Estudio(ID)
);

CREATE TABLE Institucion (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL,
  ID_Ubicacion INTEGER REFERENCES Categoria(ID),
  ID_Tipo_Empresa INTEGER REFERENCES Categoria(ID),
  ID_Pais INTEGER REFERENCES Pais(ID_pais)
);


CREATE TABLE Titulo (
  ID SERIAL PRIMARY KEY,
  Nombre VARCHAR(255) NOT NULL
);



CREATE TABLE Estudio (
  ID SERIAL PRIMARY KEY,
  ID_Nivel_Educativo INTEGER REFERENCES Nivel_Educativo(ID),
  ID_Titulo INTEGER REFERENCES Titulo(ID),
  ID_Institucion INTEGER REFERENCES Institucion(ID),
  Anio_Inicio INTEGER NOT NULL,
  Anio_Fin INTEGER NOT NULL,
  ID_Funcionario INTEGER REFERENCES Funcionario(ID_funcionario)
);
