import { initializeApp, getApps, getApp } from 'firebase/app'
import { getAuth } from 'firebase/auth'

const currentUser = ref(null)
const authInstance = ref(null)
const isInited = ref(false)

let authReadyPromise = null

export const useAuth = () => {
  if (process.client && !authInstance.value) {
    const config = useRuntimeConfig().public
    const firebaseConfig = {
      apiKey: config.firebaseApiKey,
      authDomain: config.firebaseAuthDomain,
      projectId: config.firebaseProjectId,
      storageBucket: config.firebaseStorageBucket,
      messagingSenderId: config.firebaseMessagingSenderId,
      appId: config.firebaseAppId,
    }

    const app = !getApps().length ? initializeApp(firebaseConfig) : getApp()
    const auth = getAuth(app)
    authInstance.value = auth
    currentUser.value = auth.currentUser

    // segura o fluxo até o onAuthStateChanged disparar
    authReadyPromise = new Promise((resolve) => {
      auth.onAuthStateChanged((user) => {
        currentUser.value = user
        isInited.value = true
        resolve(user)
      })
    })
  }

  const matriculaUsuario = computed(() => {
    return currentUser.value ? currentUser.value.uid : null
  })

  return {
    auth: authInstance,
    user: currentUser,
    matriculaUsuario,
    ready: isInited, // se quiser fazer loading
    
    restaurarSessao: () => authReadyPromise || Promise.resolve(currentUser.value)
  }
}