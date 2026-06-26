export default defineNuxtRouteMiddleware(async (to) => {
  if (process.server) return

  const { user, restaurarSessao } = useAuth()

  await restaurarSessao()

  const isAuthRoute = to.path === '/login' || to.path === '/'
  const isProtectedRoute =
    to.path.startsWith('/aluno') ||
    to.path.startsWith('/coordenador')

  if (!user.value) {
    if (isProtectedRoute) {
      return navigateTo('/login')
    }
    return
  }

  if (isAuthRoute) {
    const isCoordinator = user.value.uid?.startsWith('coord')
    return navigateTo(isCoordinator ? '/coordenador/' : '/aluno/')
  }
})