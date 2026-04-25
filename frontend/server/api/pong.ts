//Proxy pra facilitar acesso à API
export default defineEventHandler(async () => {
 return await $fetch('http://localhost:8080/ping')
})