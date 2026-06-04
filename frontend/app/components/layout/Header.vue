<script setup lang="ts">
const router = useRouter()

</script>

<template>
    <header>
      <div class="header-conteudo">
        <img 
          src="/images/logo.png" 
          alt="Logo UnB" 
          class="logo-unb" 
          @click="router.back()" 
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