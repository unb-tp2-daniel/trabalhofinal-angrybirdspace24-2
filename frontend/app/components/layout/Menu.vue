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
}

.menu_bar{
  background-color: #e7e6df;
  border-bottom: 1px solid #c8c8c8;
  padding: 1px;
}

.menu_list{
  display: flex;
  list-style: none;
  margin: 1px;
  padding: 1px;
}

.menu_item{
  position: relative;
  padding: 10px;
  cursor: pointer;
  font-size: 15px;
  font-weight: normal;
  color: #222;
}

.menu_item:hover{
  background-color: #d8d8d8;
  border-radius: 10px;
}

.menu_item:hover .menu_contexto{
  display: block;
}

.menu_contexto li:hover{
  background-color: #589dd6;
  
}

.menu_contexto{
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  width: auto;
  width: 300px;
  border: 1px solid #dcdfe4;
  list-style: none;
  padding: 0;
  margin: 0;
  z-index: 1000;
}

.menu_contexto li{
  padding: 8px;
  font-size: 14px;
  border-bottom: 1px solid #eee;
}

</style>