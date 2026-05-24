<template>
  <nav class="menu_bar">
    <ul class="menu_list">
        <li class="menu_item">
          <span>Ensino</span>
          <ul class="menu_contexto">
            <li><NuxtLink to="aluno/matricula" class="link_menu">Realizar Matrícula Extraordinária</NuxtLink></li>
            <li>Retirar Matrícula</li>
            <li>Trancamento de Matrícula</li>
            <li>Consultar Minhas Notas</li>
          </ul>
        </li>

        <li class="menu_item">
          <span>Pesquisa</span>
        </li>

        <li class="menu_item">
          <span>Bolsas</span>
        </li>

        <li class="menu_item">
          <span>Estágio</span>
        </li>

        <li class="menu_item">
          <span>Outros</span>
        </li>
    </ul>
  </nav>
</template>

<style scoped>

.link_menu{
  text-decoration: none;
  color: inherit;
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.menu_bar{
  background-color: #1a3a7a;
  border-bottom: 1px solid #122d62;
  padding: 0 12px;
}

.menu_list{
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 2px;
}

.menu_item{
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
}

.menu_item:hover{
  background-color: rgba(255, 255, 255, 0.12);
  color: #fff;
  border-radius: 10px;
}

.menu_item:hover .menu_contexto{
  display: block;
}

.menu_contexto li:hover{
  background-color: #589dd6;
  color: #1a3a7a;
}

.menu_contexto{
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  width: auto;
  width: 300px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  list-style: none;
  padding: 6px;
  margin: 0;
  z-index: 1000;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.menu_contexto li{
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

</style>