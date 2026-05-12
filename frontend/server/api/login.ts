//Exporta logica de requisiçao de token a API

export default defineEventHandler(async (event) => {
  const body = await readBody(event)

  try {
    const response = await $fetch('http://localhost:8080/auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'institutional_key': 'ChaveInstitucional123' //Chave muda a depender da instituicao
      },
      body: {
        usuario: body.usuario,
        senha: body.senha
      }
    })

    return response //Retorna token
  } 

  catch (error) {
    throw createError({
      statusCode: 401, //Unauthorized
      statusMessage: 'Login invalido'
    })
  }
})