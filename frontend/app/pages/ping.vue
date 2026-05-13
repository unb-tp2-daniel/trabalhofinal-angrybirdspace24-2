<template>
  <div>
    <h1>Teste de API</h1>

    <button @click="loginGoogle">
      Entrar com Google
    </button>
    
    <p v-if="pending">
      Carregando...
    </p>

    <p v-else-if="error">
      Erro ao conectar à API. Recarregue a página
    </p>

    <pre v-else>
        <p>Ebaaa 👅</p>
        {{ data }}
    </pre>
  </div>
</template>

<script setup>
//Usa proxy pra evitar problema de conexão
  import {
  GoogleAuthProvider,
  signInWithPopup
} from "firebase/auth"

import { getFirebaseAuth } from "../../plugins/firebase.client"

  const loginGoogle = async () => {

    try {

        const auth = getFirebaseAuth()

        const provider = new GoogleAuthProvider()

        const result =
            await signInWithPopup(auth, provider)

        const token =
            await result.user.getIdToken()

        localStorage.setItem(
            "token",
            token
        )

        console.log(result.user)

        await navigateTo("/matricula")

    } catch (err) {

        console.log(err)

        error.value = "Erro no login Google"
    }
  }
</script>