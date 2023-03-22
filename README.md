#  DemocrACE
# Introducción

DemocrACE es una plataforma digital de código abierto que busca mejorar la participación ciudadana en procesos políticos y electorales, a través del acceso y evaluación de información relevante sobre los funcionarios y procesos políticos.

Este proyecto se enfoca en proporcionar información actualizada sobre los funcionarios públicos, leyes y normas, obras públicas y casos de corrupción. Además, permite a los ciudadanos denunciar casos de corrupción y evaluar a los funcionarios públicos mediante calificaciones y opiniones compartidas.

# Objetivos
#  Objetivo general:

- Crear una plataforma digital de código abierto llamada DemocrACE que permita a los ciudadanos acceder y evaluar información relevante sobre sus funcionarios y procesos políticos.
#  Objetivos específicos:

- Proporcionar información actualizada sobre los funcionarios y sus actividades, incluyendo su historial académico y profesional, sus proyectos y su rendimiento en el cargo.
- Ofrecer información detallada sobre las leyes y normas a nivel nacional y local que rigen la actividad de los funcionarios públicos, así como sus implicaciones para la ciudadanía.
- Facilitar la visualización de obras públicas y su presupuesto, con la finalidad de que se pueda hacer seguimiento a los proyectos financiados por el estado.
- Permitir a los ciudadanos denunciar casos de corrupción y seguimiento de los procesos judiciales asociados.
- Proporcionar una herramienta para que los ciudadanos puedan evaluar a los funcionarios públicos y compartir sus opiniones y calificaciones.
# Indicadores de éxito:

- Número de visitas y usuarios activos en la plataforma.
- Número de denuncias de corrupción presentadas por los ciudadanos y su resolución.
- Número de funcionarios evaluados por los ciudadanos y el promedio de sus calificaciones.

# Alcance
#  Áreas de trabajo del proyecto:

- Desarrollo de la plataforma digital DemocrACE.
- Implementación de las funcionalidades descritas anteriormente.
- Integración de la plataforma con fuentes de datos relevantes para garantizar la actualización y precisión de la información.
#  Funcionalidades y características del producto:

- Acceso a información actualizada sobre funcionarios públicos, leyes y normas, obras públicas y casos de corrupción.
- Funcionalidad de denuncia de casos de corrupción y seguimiento de procesos judiciales.
- Herramienta de evaluación de funcionarios públicos con calificaciones y opiniones compartidas por los ciudadanos.
- Interfaz amigable y fácil de usar para los usuarios.
# Usuarios y público objetivo:

- Ciudadanos interesados en la transparencia y la rendición de cuentas de los funcionarios públicos.
- Grupos de defensa de los derechos ciudadanos y anticorrupción.
- Medios de comunicación y periodistas que buscan información precisa y actualizada sobre los temas políticos y sociales.
- Académicos e investigadores que deseen utilizar la plataforma como fuente de datos.

# Contribución

Este proyecto está abierto a contribuciones de cualquier persona interesada en mejorar la transparencia y la rendición de cuentas en procesos políticos y electorales. Si deseas contribuir, por favor



# governance

## validar usuario:
```
mutation {
  registrar(input: {
    nombres: "Nombre"
    apellidos: "Apellido"
    nombreUsuario: "you"
    email: "you@example.com"
    password: "???"
  }) {
    id
    nombres
    apellidos
    email
    nombreUsuario
    status
  }
}
```

```
mutation {
  enviarCodigo(email: "you@example.com")
}
```

```
mutation {
  verificarUsuario(email: "you@example.com", code: "123456")  
}

```

## Cambiar contrasena
```
mutation {
  enviarCodigo(email: "you@example.com")
}
```

```
mutation {
  changePassword(input: {
    email:"you@example.com"
    code: "123456"
    newPassword: "supersegura"
  })
}
```