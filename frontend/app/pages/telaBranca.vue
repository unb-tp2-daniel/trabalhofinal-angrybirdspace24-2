<template>
  <div>
    <Header />
    <main class="container">
      <div class="card">
        <h1 class="titulo">TELA BRANCA</h1>
        <div class="linha"></div>
        
        <p style="text-align: center; color: #333; margin-bottom: 20px;">
          Cadastro simulado com sucesso! 🎉
        </p>
        
        <div class="links">
          <p><a @click="telaLogin">Voltar para o Login</a></p>
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
    title: 'Tela Branca - SIGAA UnB'
  })
</script>

<style scoped>
    .container { 
        flex: 1; display: flex; 
        justify-content: center; 
        align-items: center; 
        padding: 40px 16px; 
    }
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
    .links { 
        text-align: center; 
    }
    .links a { 
        color: #1a72c4; 
        text-decoration: none; 
        font-size: 14px;
        font-weight: bold;
    } 
    .links a:hover { 
        text-decoration: underline; 
        cursor: pointer;
    }
</style>