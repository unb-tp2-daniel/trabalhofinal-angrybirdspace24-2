<script setup>
    defineProps({
    visivel: Boolean,
    materia: Object
    })

    defineEmits(['fechar'])
</script>

<template>
  <div v-if="visivel" class="modal-overlay" @click="$emit('fechar')">
    <div class="modal" @click.stop>
      <div class="modal-header">
        <div>
          <h2>Detalhes da Matéria</h2>

          <p v-if="materia" class="modal-subtitle">
            {{ materia.codigo }}
          </p>
        </div>

        <button class="close-btn" @click="$emit('fechar')">
          ✕
        </button>
      </div>

      <template v-if="materia">
        <p>
          <strong>Departamento:</strong>
          {{ materia.departamentoId }}
        </p>

        <p>
          <strong>Carga Horária:</strong>
          {{ materia.cargaHoraria }}h
        </p>

        <p>
          <strong>Nível:</strong>
          {{ materia.nivelAcademico }}
        </p>

        <p>
          <strong>Pré-requisitos:</strong>

          {{
            materia.preRequisitos?.length
              ? materia.preRequisitos
                  .flatMap(p => p.disciplinas)
                  .join(', ')
              : 'Nenhum'
          }}
        </p>

        <p>
          <strong>Co-requisitos:</strong>

          {{
            materia.coRequisitos?.length
              ? materia.coRequisitos
                  .flatMap(c => c.disciplinas)
                  .join(', ')
              : 'Nenhum'
          }}
        </p>

        <p>
          <strong>Equivalências:</strong>

          {{
            materia.equivalencias?.join(', ')
          }}
        </p>

        <div class="conteudo-box">
          <strong>Conteúdo</strong>

          <p>
            {{ materia.conteudo }}
          </p>
        </div>

      </template>

      <template v-else>
        Matéria não encontrada.
      </template>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5); /* Fundo escurecido semi-transparente */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999; /* Garante que fica por cima de tudo */
}

.modal {
  background: white;
  padding: 24px;
  border-radius: 12px;
  width: 90%;
  max-width: 600px;
  max-height: 85vh;
  overflow-y: auto;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  gap: 12px;
  color: #333;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  border-bottom: 1px solid #eee;
  padding-bottom: 12px;
  margin-bottom: 8px;
}

.modal-header h2 {
  margin: 0;
  color: #1a3a7a;
  font-size: 22px;
}

.modal-subtitle {
  margin: 4px 0 0 0;
  color: #666;
  font-family: monospace;
}

.close-btn {
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
  color: #999;
  transition: color 0.2s;
}

.close-btn:hover {
  color: #333;
}

.conteudo-box {
  background: #f8fafc;
  padding: 12px;
  border-radius: 6px;
  border: 1px solid #e2e8f0;
  margin-top: 8px;
}

p {
  margin: 4px 0;
  font-size: 14px;
}

strong {
  color: #4a5568;
}
</style>