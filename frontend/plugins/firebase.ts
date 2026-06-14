import { initializeApp, getApp, getApps } from 'firebase/app'
import { getAuth } from 'firebase/auth'

export default defineNuxtPlugin(() => {
  // Só inicializa se estiver no navegador
  if (process.client) {
    const config = useRuntimeConfig().public

    const firebaseConfig = {
      apiKey: config.firebaseApiKey,
      authDomain: config.firebaseAuthDomain,
      projectId: config.firebaseProjectId,
      storageBucket: config.firebaseStorageBucket,
      messagingSenderId: config.firebaseMessagingSenderId,
      appId: config.firebaseAppId,
    }

    if (!getApps().length) {
      initializeApp(firebaseConfig)
    }
  }
})