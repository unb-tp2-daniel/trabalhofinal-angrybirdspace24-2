<script setup lang="ts">
import { signOut } from 'firebase/auth'
import { useAuth } from '~/composables/useAuth'

const { auth } = useAuth()

async function logout() {
  try {
    if (!auth.value) return

    await signOut(auth.value)
    await navigateTo('/login')
  } catch (error) {
    console.error('Erro ao deslogar usuário:', error)
  }
}

interface MenuItem {
  label: string
  to?: string
  children?: MenuItem[]
}

defineProps<{
  items: MenuItem[]
}>()
</script>

<template>
  <nav class="menu_bar">
    <div class="menu_scroll_wrapper">
      <ul class="menu_list">
        <li
          v-for="item in items"
          :key="item.label"
          class="menu_item"
        >
          <span>{{ item.label }}</span>

          <ul v-if="item.children" class="menu_contexto">
            <li v-for="child in item.children" :key="child.label">
              <NuxtLink v-if="child.to" :to="child.to" class="link_menu">
                {{ child.label }}
              </NuxtLink>
              <span v-else>{{ child.label }}</span>
            </li>
          </ul>
        </li>
      </ul>
    </div>
    <button class="botao-sair" @click="logout">Sair</button>
  </nav>
</template>

<style scoped>
.link_menu {
  text-decoration: none;
  color: inherit;
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.menu_bar {
  background-color: #1a3a7a;
  border-bottom: 1px solid #122d62;
  padding: 0 12px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  box-sizing: border-box;
  gap: 10px;
}

/* --- NOVO: GERENCIADOR DO FADE-OUT GRADIENT --- */
.menu_scroll_wrapper {
  position: relative;
  flex: 1;
  overflow: hidden;
  display: flex;
}

/* Cria uma sombra suave por cima do menu indicando que há mais conteúdo à direita */
@media (max-width: 768px) {
  .menu_scroll_wrapper::after {
    content: '';
    position: absolute;
    top: 0;
    right: 0;
    height: 100%;
    
    /* Aumentamos a largura para o efeito cobrir mais espaço horizontal */
    width: 65px; 
    
    /* TÉCNICA DE CONTRAS-TE:
      1. Começa totalmente invisível (0%)
      2. No meio do caminho (40%), já injeta o azul com 40% de opacidade
      3. Na borda final (100%), entrega o azul totalmente sólido 
    */
    background: linear-gradient(
      to right, 
      rgba(26, 58, 122, 0) 0%, 
      rgba(26, 58, 122, 0.4) 40%, 
      rgba(26, 58, 122, 0.95) 85%, 
      #1a3a7a 100%
    );
    
    /* Adiciona uma linha de sombra projetada vertical para simular profundidade real */
    box-shadow: inset -15px 0 15px -10px rgba(13, 29, 61, 0.6);
    
    pointer-events: none;
    z-index: 2;
  }
}

.menu_list {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 2px;
  overflow-x: auto;
  white-space: nowrap;
  -webkit-overflow-scrolling: touch;
  width: 100%;
}

.menu_list::-webkit-scrollbar {
  display: none;
}

.menu_item {
  position: relative;
  padding: 11px 14px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.82);
  display: flex;
  align-items: center;
  gap: 6px;
  border-radius: 6px;
  transition: background 0.15s, color 0.15s;
  user-select: none;
  flex-shrink: 0;
}

.menu_item:hover {
  background-color: rgba(255, 255, 255, 0.12);
  color: #fff;
  border-radius: 10px;
}

.menu_item:hover .menu_contexto {
  display: block;
}

.menu_contexto li:hover {
  background-color: #589dd6;
  color: #1a3a7a;
}

.menu_contexto {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  width: 300px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  list-style: none;
  padding: 6px;
  margin: 0;
  z-index: 1000;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.menu_contexto li {
  white-space: normal;
  padding: 8px;
  font-size: 14px;
  border-bottom: 1px solid #eee;
  color: #333;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
}

.botao-sair {
  background: transparent;
  border: none;
  color: rgba(255, 255, 255, 0.82);
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 11px 14px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 500;
  flex-shrink: 0;
}

.botao-sair:hover {
  background-color: rgba(255, 255, 255, 0.12);
  color: white;
  text-decoration: underline;
}
</style>