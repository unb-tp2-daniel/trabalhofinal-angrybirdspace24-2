<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
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

// Controla qual menu está aberto baseado no index ou label
const activeMenu = ref<string | null>(null)

// Guarda o recuo esquerdo (left) dinâmico para alinhar no desktop  
const menuLeftOffset = ref<string>('12px')

function toggleMenu(event: MouseEvent, label: string, hasChildren: boolean) {
  if (!hasChildren) return
  
  if (activeMenu.value === label) {
    activeMenu.value = null
  } else {
    activeMenu.value = label
    
    // No desktop, calcula a posição exata do botão clicado em relação à barra do menu
    if (window.innerWidth > 768) {
      const target = event.currentTarget as HTMLElement
      const rect = target.getBoundingClientRect()
      const navBar = target.closest('.menu_bar')
      if (navBar) {
        const navRect = navBar.getBoundingClientRect()
        // Define o 'left' do submenu exatamente onde o botão começa
        menuLeftOffset.value = `${rect.left - navRect.left}px`
      }
    } else {
      menuLeftOffset.value = '12px' // Padrão seguro para mobile
    }
  }
}

function closeMenu() {
  activeMenu.value = null
}

// Fecha o menu se o usuário clicar fora dele
onMounted(() => {
  if (typeof window !== 'undefined') {
    window.addEventListener('click', closeMenu)
  }
})

onUnmounted(() => {
  if (typeof window !== 'undefined') {
    window.removeEventListener('click', closeMenu)
  }
})

const menuListRef = ref<HTMLElement | null>(null)

// Permite rolar o menu horizontalmente usando a rodinha vertical do mouse
function handleWheel(event: WheelEvent) {
  if (!menuListRef.value) return
  // Se o usuário estiver rodando o scroll, move a lista para os lados
  if (event.deltaY !== 0) {
    event.preventDefault()
    menuListRef.value.scrollLeft += event.deltaY
  }
}
</script>

<template>
  <nav class="menu_bar" @click.stop>
    <div class="menu_scroll_wrapper">
      <ul 
        ref="menuListRef"
        class="menu_list"
        @wheel="handleWheel"
        >
        <li
          v-for="item in items"
          :key="item.label"
          class="menu_item"
          :class="{ 'item_ativo': activeMenu === item.label }"
          @click="toggleMenu($event, item.label, !!item.children)"
        >
          <span>{{ item.label }}</span>
        </li>
      </ul>
    </div>
    
    <button class="botao-sair" @click="logout">Sair</button>

    <template v-for="item in items" :key="'context-' + item.label">
      <ul 
        v-if="item.children && activeMenu === item.label" 
        class="menu_contexto"
        :style="{ left: menuLeftOffset }"
      >
        <li 
          v-for="child in item.children" 
          :key="child.label"
          @click="closeMenu"
        >
          <NuxtLink v-if="child.to" :to="child.to" class="link_menu">
            {{ child.label }}
          </NuxtLink>
          <span v-else>{{ child.label }}</span>
        </li>
      </ul>
    </template>
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
  position: relative; 
}

.menu_scroll_wrapper {
  position: relative;
  flex: 1;
  display: flex;
  overflow: hidden;
}

/* Cria uma sombra suave por cima do menu indicando que há mais conteúdo à direita */
@media (max-width: 768px) {
  .menu_scroll_wrapper::after {
    content: '';
    position: absolute;
    top: 0;
    right: 0;
    height: 100%;
    
    width: 65px; 
    
    background: linear-gradient(
      to right, 
      rgba(26, 58, 122, 0) 0%, 
      rgba(26, 58, 122, 0.4) 40%, 
      rgba(26, 58, 122, 0.95) 85%, 
      #1a3a7a 100%
    );
    
    /* Adiciona uma linha de sombra projetada vertical*/
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
  /*position: relative;*/
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

.menu_item:hover, .item_ativo {
  background-color: rgba(255, 255, 255, 0.12);
  color: #fff;
}

.menu_contexto {
  /* display: none; */
  position: absolute;
  top: 100%;
  left: 12px;
  background: white;
  width: 300px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  list-style: none;
  padding: 6px;
  margin: 0;
  z-index: 9999;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  display: block;
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

.menu_contexto li:hover {
  background-color: #589dd6;
  color: white;
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