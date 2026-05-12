import { initializeApp, getApp, getApps } from 'firebase/app'
import { getAuth } from 'firebase/auth'

export const getFirebaseAuth = () => {
  const config = useRuntimeConfig().public // pega as variáveis de ambiente

  const firebaseConfig = {
    apiKey: config.firebaseApiKey,
    authDomain: config.firebaseAuthDomain,
    projectId: config.firebaseProjectId,
    storageBucket: config.firebaseStorageBucket,
    messagingSenderId: config.firebaseMessagingSenderId,
    appId: config.firebaseAppId,
  };

  // não inicializar duas vezes
  const app = !getApps().length ? initializeApp(firebaseConfig) : getApp();
  
  return getAuth(app);
}