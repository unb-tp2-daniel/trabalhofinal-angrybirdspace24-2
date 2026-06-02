<template>
    <div>
    <Header />
    <main class="container">
      <div class="card">

        <h1 class="titulo">AUTENTICAÇÃO INTEGRADA</h1>
        <div class="linha"></div>

        <form @submit.prevent="cadastrar">
          <label for="usuario">Definir nome de usuário:</label>
          <input 
            v-model="usuario"
            type="text"
            id="usuario"
            placeholder="Digite seu usuário"
          >

          <label for="senha">Definir senha:</label>
          <div class="password-wrapper">
            <input 
                v-model="senha"
                :type="mostrarSenha ? 'text' : 'password'" 
                id="senha"
                placeholder="Digite sua senha"
            >
            <button type="button" class="toggle-btn" @click="togglePassword">
                <Icon :name="mostrarSenha ? 'uil:eye-slash' : 'uil:eye'" size="20px" />
            </button>
            </div>

            <label for="confirmarSenha">Confirmar senha:</label>
            <div class="password-wrapper">
                <input 
                    v-model="confirmarSenha" 
                    :type="mostrarConfirmarSenha ? 'text' : 'password'" 
                    id="confirmarSenha"
                    placeholder="Confirme sua senha"
                >
                <button type="button" class="toggle-btn" @click="toggleConfirmarPassword">
                    <Icon :name="mostrarConfirmarSenha ? 'uil:eye-slash' : 'uil:eye'" size="20px" />
                </button>
            </div>

            <p v-if="error" class="mensagem-erro">{{ error }}</p>

          <button type="submit" class="botao">Cadastrar ></button>
        </form>

        <div class="links">
          <p><a @click="telaLogin">Voltar</a></p>
        </div>

      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup>
    function telaLogin() {
        navigateTo('/login')
    }
    
    useHead({
        title: 'Login - SIGAA UnB'
    })

    //mudança no nuxt.config.ts
    //import Header from '~/components/layout/Header.vue'
    //import Footer from '~/components/layout/Footer.vue';
    
    import { ref } from 'vue'

    const usuario = ref('')
    const senha = ref('')
    const confirmarSenha = ref('') // 1. New variable for the second input
    const mostrarSenha = ref(false) // New state: false = hidden, true = visible
    const mostrarConfirmarSenha = ref(false) // 2. Separate visibility state

    const error = ref('')

    const togglePassword = () => {
        mostrarSenha.value = !mostrarSenha.value
    }

    const toggleConfirmarPassword = () => {
        mostrarConfirmarSenha.value = !mostrarConfirmarSenha.value
    }

    const cadastrar = async () => {
        error.value = '' // Reset error message

        // 5. Check if passwords match before sending to API!
        if (senha.value !== confirmarSenha.value) {
            error.value = 'As senhas não coincidem.'
            return // Stops the function here
        }

        // try {
        //     // Note: You might need to change '/api/login' to '/api/register' in the backend later
        //     const response = await $fetch('/api/cadastro', { 
        //         method: 'POST',
        //         body: {
        //             email: usuario.value,
        //             password: senha.value
        //         }
        //     })

        //     localStorage.setItem('token', response.token)
        //     await navigateTo('/login') 
        // } catch (err) {
        //     error.value = 'Falha no cadastro. Tente novamente.'
        // }

        await navigateTo('/coordenador')
    }

    // Adicionando botao de "olho" na senha
    
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

    .password-wrapper {
        position: relative;
        width: 100%;
        margin-bottom: 14px;
        /* We ensure this is a block container so the input fills it */
        display: block; 
    }

    .password-wrapper input {
        width: 100%;
        /* Standardized with your other input */
        margin-bottom: 0 !important; 
        padding: 8px 10px 8px 10px; /* Right padding is 40px to hide text under eye */
        display: block;
        /* REMOVE the fixed height: 38px; let padding decide the height */
    }

    .toggle-btn {
        position: absolute;
        right: -11px;
        top: 50%;
        /* -45% is visually more 'centered' than -50% for eye icons */
        transform: translateY(-45%); 
        
        background: none;
        border: none;
        cursor: pointer;
        color: #666;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 0;
        z-index: 2;
    }

    .toggle-btn:hover {
        color: #1a5276; 
    }

    .mensagem-erro {
        color: #d32f2f;
        font-size: 13px;
        text-align: center;
        margin-bottom: 12px;
        font-weight: bold;
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
        cursor: pointer;
    }

</style>