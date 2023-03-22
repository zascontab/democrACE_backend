# democraceback
Golang

El modelado de la base de datos se lo realizo utilizando chat gpt, asi que puede tener algunas fallas nos gustaria que revisen el modelo

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