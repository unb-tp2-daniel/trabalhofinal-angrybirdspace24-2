/*
// https://nuxt.com/docs/api/configuration/nuxt-config 
// export default defineNuxtConfig({ compatibilityDate: '2025-07-15', devtools: { enabled: true } })
*/

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  app: {
  head: {
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  }
},

  modules: ['@nuxt/icon'],

  components: [
    {
      path: '~/components',
      pathPrefix: false
    }
  ]
})