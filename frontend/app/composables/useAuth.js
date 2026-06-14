import { initializeApp, getApps, getApp } from 'firebase/app'
import { getAuth } from 'firebase/auth'

export const useAuth = () => {
  const currentUser = ref(null)
  const authInstance = ref(null)

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

    const app = !getApps().length ? initializeApp(firebaseConfig) : getApp()
    const auth = getAuth(app)
    authInstance.value = auth

    currentUser.value = auth.currentUser

    auth.onAuthStateChanged((user) => {
      currentUser.value = user
    })
  }

  const matriculaUsuario = computed(() => {
    return currentUser.value ? currentUser.value.uid : null
  })

  return {
    auth: authInstance,
    user: currentUser,
    matriculaUsuario
  }
}