<script setup lang="ts">
const router = useRouter()
const route = useRoute()

function voltarHome() {
  const partes = route.path.split('/').filter(Boolean)

  if (partes.length > 0) {
    router.push('/' + partes[0])
  } else {
    router.push('/')
  }
}
</script>

<template>
    <header>
      <div class="header-conteudo">
        <img 
          src="/images/logo.png" 
          alt="Logo UnB" 
          class="logo-unb" 
          @click="voltarHome()" 
        />
        <span class="separador">|</span>
        <span class="nome-universidade">Universidade de Brasília</span>
      </div>
    </header>
</template>

<style scoped>
    header { 
        background-color: #ffffff; 
        border-bottom: 3px solid #1a5276; 
        padding: 12px 24px; 
    }

    .header-conteudo { 
        display: flex; 
        align-items: center; 
        gap: 10px; 
    }

    .logo-unb { 
        font-size: 22px; 
        font-weight: 700; 
        color: #1a5276; 
        height: 25px;    
        width: auto;
        cursor: pointer;
        margin-left: -23px;
    } 

    .logo-unb:hover {
        opacity: 0.7;
    }

    
    .separador { 
        color: #cccccc; 
        font-size: 20px; 
    } 
    
    .nome-universidade { 
        font-size: 13px; 
        color: #555555; 
        text-transform: uppercase; 
        letter-spacing: 1px; 
    }
    
</style>