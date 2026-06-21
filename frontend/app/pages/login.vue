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

          <button type="submit" class="botao">ENTRAR ></button>
          
        </form>

        <div class="links">
          <!--<p><a @click="loginMicrosoft">ENTRAR COM MICROSOFT</a></p> -->
          <p><strong>Aluno</strong>, <a @click="cadastroAluno">cadastre-se aqui</a></p>
          <p><strong>Coordenador</strong>, <a @click="cadastroCoordenador">cadastre-se aqui</a></p>
          <p><a @click="recuperarSenha">Esqueceu a senha?</a></p>
          <p><a @click="telaBranca">Esqueceu o login?</a></p>
        </div>

      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup>

    function cadastroAluno() {
        navigateTo('/cadastroAluno')
    }
    function cadastroCoordenador() {
        navigateTo('/cadastroServidor')
    }
    function recuperarSenha() {
        navigateTo('/recuperarSenha')
    }
    function telaBranca() {
        navigateTo('/telaBranca')
    }

    /*
    function loginMicrosoft() {
        //Não tá funcionando
        const provider = new OAuthProvider('microsoft.com');
        const auth = getFirebaseAuth();
        signInWithPopup(auth, provider)
            .then((result) => {
                const credential = OAuthProvider.credentialFromResult(result);
                const accessToken = credential.accessToken;
                const idToken = credential.idToken;
                navigateTo('/aluno/')
            })
            .catch((error) => {
                console.log(error)
            });
    }
    */

    function passa() {
        navigateTo('/aluno/')
    }
    
    useHead({
        title: 'Login - SIGAA UnB'
    })

    //mudança no nuxt.config.ts
    //import Header from '~/components/layout/Header.vue'
    //import Footer from '~/components/layout/Footer.vue';
    
import { ref } from 'vue'
import { signInWithEmailAndPassword, setPersistence, browserLocalPersistence } from "firebase/auth"
import { useAuth } from '~/composables/useAuth' // Importa o nosso composable unificado

const { auth } = useAuth() // Resgata a instância do auth inicializada com segurança

    const usuario = ref('')
    const senha = ref('')
    const mostrarSenha = ref(false)
    const error = ref('')

    const togglePassword = () => {
        mostrarSenha.value = !mostrarSenha.value
    }

    const login = async () => {
        console.log("foi")
        error.value = '' // limpa erros anteriores
        
        try {
            if (!auth.value) {
                error.value = 'Erro ao carregar o módulo de autenticação.'
                return
            }

            await setPersistence(auth.value, browserLocalPersistence)
            
            const userCredential = await signInWithEmailAndPassword(auth.value, usuario.value, senha.value)
            
            console.log("Usuário logado:", userCredential.user.uid)
            
            await navigateTo('/aluno/') 
        } 
        catch (err) {
            console.error("Erro detalhado do login:", err)
            error.value = 'Falha no login: verificar usuário e senha.'
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