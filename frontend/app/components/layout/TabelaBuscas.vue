<template>
  <div class="filtro_container">
    <div class="titulo">BUSCAR TURMAS ABERTAS</div>

    <div class="formulario">

      <div class="campo">
        <input type="checkbox" class="caixinha" v-model="ativos.codigoTurma"/>
        <label>Código do Componente:</label>
        <input v-model="filtro.codigoTurma" type="text" class="entrada" />
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha" v-model="ativos.nome"/>
        <label>Nome do Componente:</label>
        <input v-model="filtro.nome" type="text" class="entrada" />
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha" v-model="ativos.horario"/>
        <label>Horário:</label>
        <input v-model="filtro.horario" type="text" class="entrada" />
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha" v-model="ativos.docente"/>
        <label>Nome do Docente:</label>
        <input v-model="filtro.docente" type="text" class="entrada" />
      </div>

      <div class="campo">
        <input type="checkbox" class="caixinha" v-model="ativos.unidade"/>
        <label>Unidade Responsável:</label>
        <select v-model="filtro.unidade" class="entrada">
          <option value="">Selecione...</option> <!-- tenho que dar um jeitinho de puxar isso aq tb -->
          <option>FACULDADE DE TECNOLOGIA</option>
          <option>FACULDADE DE SAÚDE</option>
          <option>FACULDADE DE AGRONOMIA</option>
          <option>INSTITUTO DE COMPUTAÇÃO</option>
        </select>
      </div>

      <div class="rodape">
        <button @click="buscarTurmas" :disabled="carregando">
          {{ carregando ? 'Buscando...' : 'Buscar' }}
        </button>
        <span v-if="erro" class="erro">{{ erro }}</span>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue'

const emit = defineEmits<{
  (e: 'resultados', turmas: any[]): void
}>()

const filtro = reactive({
  codigoTurma: '',
  nome: '',
  horario: '',
  docente: '',
  unidade: ''
})

//checkbox
const ativos = reactive({
  codigoTurma: false,
  nome: false,
  horario: false,
  docente: false,
  unidade: false
})

watch(filtro, (val) =>{
  ativos.codigoTurma = !!val.codigoTurma
  ativos.nome = !!val.nome
  ativos.horario = !!val.horario
  ativos.docente = !!val.docente
  ativos.unidade = !!val.unidade
})

const carregando =ref(false)
const erro =ref('')
/* dsfafeads
function buscarTurmas() {
  console.log(filtro)
} */

async function buscarTurmas() {
  carregando.value = true
  erro.value = ''

  try {
    const todos = await $fetch<any[]>(
      'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas'
    )

    const dados = todos.filter(turma =>{
      if (ativos.codigoTurma  && filtro.codigoTurma  && !turma.codigoTurma?.toLowerCase().includes(filtro.codigoTurma.toLowerCase()))  return false
      if (ativos.nome    && filtro.nome    && !turma.nomeMateria?.toLowerCase().includes(filtro.nome.toLowerCase()))    return false
      if (ativos.horario && filtro.horario && !turma.horario?.toLowerCase().includes(filtro.horario.toLowerCase()))     return false
      if (ativos.docente && filtro.docente && !turma.professorNome?.toLowerCase().includes(filtro.docente.toLowerCase())) return false
      if (ativos.unidade && filtro.unidade && turma.unidade !== filtro.unidade) return false
      return true
    })
    emit('resultados', dados)
  }catch(e){
    erro.value = 'Não foi possível carregar as turmas.'
  }finally{
    carregando.value = false
  }
}
</script>

<style scoped>
button:disabled{
  opacity: 0.6;
  cursor: not-allowed;
}
.erro{
  display: block;
  margin-top: 6px;
  color: #e53935;
  font-size: 0.85rem;
}

.filtro_container {
  margin-top: 10px;
  border: 1px solid #b8cfe8;
  border-radius: 10px;
  overflow: hidden;
  background-color: #f0f5fa;
}

.titulo {
  background-color: #1a3a7a;
  color: white;
  font-weight: 600;
  font-size: 13px;
  padding: 11px 16px;
  letter-spacing: 0.07em;
}

.formulario {
  padding: 16px;
}

.campo {
  display: flex;
  width: 60%;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

label {
  width: 190px;
  font-size: 14px;
  color: #333;
  flex-shrink: 0; 
}

.caixinha {
  appearance: none;
  -webkit-appearance: none;
  width: 16px;
  height: 16px;
  border: 1.5px solid #7aa3cc;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  flex-shrink: 0;
  position: relative;
  transition: background 0.15s, border-color 0.15s;
}

.caixinha:checked {
  background-color: #1a3a7a;
  border-color: #1a3a7a;
}

.caixinha:checked::after {
  content: '';
  position: absolute;
  left: 4px;
  top: 1px;
  width: 5px;
  height: 9px;
  border: 2px solid white;
  border-top: none;
  border-left: none;
  transform: rotate(45deg);
}

.caixinha:hover {
  border-color: #1a3a7a;
}

.entrada {
  height: 30px;
  flex: 1; 
  border: 1.5px solid #c8d8ea;
  border-radius: 6px;
  background: white;
  font-size: 13px;
  padding: 0 10px;
  color: #1a1a2e;
  outline: none;
  transition: border-color 0.15s, box-shadow 0.15s;
}

.entrada:focus {
  border-color: #1a3a7a;
  box-shadow: 0 0 0 3px rgba(26, 58, 122, 0.1);
}

select.entrada {
  height: 32px;
  cursor: pointer;
}

.rodape {
  margin-top: 12px;
  padding-left: 26px; 
}

button {
  background-color: #1a3a7a;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 8px 22px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.15s;
}

button:hover {
  background-color: #254d9e;
}
</style>