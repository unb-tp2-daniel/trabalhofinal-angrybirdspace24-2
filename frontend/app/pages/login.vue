<template>
    <div>
    <Header />
    <main class="container">
      <div class="card">

        <h1 class="titulo">AUTENTICAÇÃO INTEGRADA</h1>
        <div class="linha"></div>

        <form @submit.prevent="login">
          <label for="usuario">Nome de usuário:</label>
          <input 
            v-model="usuario"
            type="text"
            id="usuario"
            placeholder="Digite seu usuário"
          >

          <label for="senha">Senha:</label>
          <input 
            v-model="senha"
            type="password"
            id="senha"
            placeholder="Digite sua senha"
          >

          <button type="submit" class="botao">ENTRAR ></button>
        </form>

        <div class="links">
          <p><strong>Aluno</strong>, <a href="#">cadastre-se aqui</a></p>
          <p><strong>Servidor</strong>, <a href="#">cadastre-se aqui</a></p>
          <p><a href="#">Esqueceu a senha?</a></p>
          <p><a href="#">Esqueceu o login?</a></p>
        </div>

      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup>
    //mudança no nuxt.config.ts
    //import Header from '~/components/layout/Header.vue'
    //import Footer from '~/components/layout/Footer.vue';
    import { ref } from 'vue'

    const usuario = ref('')
    const senha = ref('')
    const error = ref('')

    const login = async () => {
        try {
            const response = await $fetch('/api/login', {
                method: 'POST',
                body: {
                    email: usuario.value,
                    password: senha.value
                }
            })

            localStorage.setItem(
                'token',
                response.token
            )

            await navigateTo('/login')
        } 

        catch (err) {
            error.value = 'Falha no login'
        }
    }
</script>

<style scoped>

    /* ===================== CONTAINER PRINCIPAL ===================== */ 
    .container { 
        flex: 1; display: flex; 
        justify-content: center; 
        align-items: center; 
        padding: 40px 16px; 
    }

    /* ===================== CARD DE LOGIN ===================== */ 
    .card { 
        background-color: #ffffff; 
        border-radius: 6px; 
        padding: 36px 40px; 
        width: 100%; 
        max-width: 380px; 
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.12); 
    }

    .titulo { 
        text-align: center; 
        font-size: 13px; 
        font-weight: 700; 
        color: #999999; 
        letter-spacing: 2px; 
        margin-bottom: 10px; 
    }

    .linha { 
        border: none; 
        border-top: 1px solid #e0e0e0; 
        margin-bottom: 22px; 
    }

    /* ===================== FORMULÁRIO ===================== */ 
    form { 
        display: flex; 
        flex-direction: column; 
    }

    label { 
        font-size: 13px; 
        color: #333333; 
        margin-bottom: 4px; 
    }

    input { 
        width: 100%; 
        border: 1px solid #b0c4de; 
        border-radius: 4px; 
        padding: 8px 10px; 
        font-size: 14px; 
        color: #333333; 
        background-color: #f5f8fb; 
        margin-bottom: 14px; 
        outline: none; 
        transition: border-color 0.2s, 
        background-color 0.2s; 
    }

    /* ===================== BOTÃO ===================== */ 
    .botao { 
        width: 100%; 
        background-color: #1a5276; 
        color: #ffffff; 
        border: none; 
        border-radius: 24px; 
        padding: 10px; 
        font-size: 14px; 
        font-weight: 700; 
        letter-spacing: 1px; 
        cursor: pointer; 
        margin-top: 4px; 
        transition: background-color 0.2s; 
    }

    .botao:hover { 
        background-color: #154360; 
    } 
    
    .botao:active { 
        background-color: #0e2f44; 
    }

    /* ===================== LINKS ===================== */ 
    .links { 
        margin-top: 18px; 
        text-align: center; 
        display: flex; 
        flex-direction: column; 
        gap: 2px; 
    }

    .links p { 
        font-size: 13px; 
        color: #333333; 
        line-height: 1.5; 
        margin: 0;
    }

    .links a { 
        color: #1a72c4; 
        text-decoration: none; 
    } 
    
    .links a:hover { 
        text-decoration: underline; 
    }

</style>