<script setup>
defineProps({
  visivel: Boolean,
  turma: Object
})

const emit = defineEmits(['fechar', 'confirmar'])
</script>

<template>
  <Transition name="modal">
    <div v-if="visivel" class="ver" @click.self="emit('fechar')">
      <div class="modal">

        <div class="modal-icone">
          <i class="ti ti-school" aria-hidden="true"></i>
        </div>

        <h2>Confirmar matrícula</h2>
        <p>Você está solicitando matrícula na turma:</p>
        
        <div class="turma-card">
          <span class="turma-codigo">{{ turma.codigoTurma }}</span>
          <span v-if="turma.professorNome" class="turma-docente">
            <i class="ti ti-user" style="font-size: 12px;" aria-hidden="true"></i>
            {{ turma.professorNome }}
          </span>
        </div>


        <div class="botoes">
          <button class="btn-cancelar" @click="emit('fechar')">
            Cancelar
          </button>
          <button class="btn-confirmar" @click="emit('confirmar', turma)">
            <i class="ti ti-check" aria-hidden="true"></i>
            Confirmar
          </button>
        </div>

      </div>
    </div>
  </Transition>
</template>

<style scoped>
.ver {
  position: fixed;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  width: 420px;
  background-color: white;
  border-radius: 16px;
  padding: 32px 28px 24px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.modal-icone {
  width: 52px;
  height: 52px;
  background: #e6f1fb;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 16px;
  font-size: 24px;
  color: #1a3a7a;
}

.modal h2 {
  font-size: 20px;
  font-weight: 600;
  color: #1a3a7a;
  margin-bottom: 8px;
}

.modal p {
  font-size: 14px;
  color: #666;
  margin-bottom: 16px;
  line-height: 1.5;
}

.turma-card {
  width: 100%;
  background: #f0f5fa;
  border: 1px solid #c5d8f0;
  border-radius: 10px;
  padding: 12px 16px;
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.turma-codigo {
  font-size: 20px;
  font-weight: 600;
  color: #1a3a7a;
  font-family: 'DM Mono', 'Courier New', monospace;
}

.turma-docente {
  font-size: 13px;
  color: #666;
  display: flex;
  align-items: center;
  gap: 4px;
}

.botoes {
  display: flex;
  justify-content: center;
  gap: 10px;
  width: 100%;
}

.btn-cancelar {
  flex: 1;
  border: 1px solid #e0e0e0;
  background: #f4f4f4;
  color: #444;
  padding: 10px 18px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: background 0.15s;
}

.btn-cancelar:hover {
  background: #e8e8e8;
}

.btn-confirmar {
  flex: 1;
  border: none;
  background: #1a3a7a;
  color: white;
  padding: 10px 18px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  transition: background 0.15s;
}

.btn-confirmar:hover {
  background: #254d9e;
}

.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}

.modal-enter-active .modal,
.modal-leave-active .modal {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal,
.modal-leave-to .modal {
  opacity: 0;
  transform: scale(0.95);
}
</style>