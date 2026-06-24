import { useAuth } from '~/composables/useAuth'

export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) return

  const { user, restaurarSessao } = useAuth()

  // aguarda pacientemente que o firebase leia o localStorage
  await restaurarSessao()

  if (!user.value && to.path.startsWith('/aluno')) {
    console.warn("Acesso negado: Redirecionando para a página de login.")
    return navigateTo('/login')
  }

  if (user.value && to.path === '/login') {
    return navigateTo('/aluno/')
  }
})