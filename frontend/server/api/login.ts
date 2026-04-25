export default defineEventHandler(async (event) => {
  const body = await readBody(event)

  try {
    const response = await $fetch('http://localhost:8080/auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'institutional_key': 'ChaveInstitucional123'
      },
      body: {
        usuario: body.usuario,
        senha: body.senha
      }
    })

    console.log('Resposta da API:', response)

    return response
  } 

  catch (error) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Login invalido'
    })
  }
})