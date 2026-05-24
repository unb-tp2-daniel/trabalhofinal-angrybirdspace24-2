<script setup>
defineProps({
  visivel: Boolean,
  mensagem: String,
  tipo: String
})
</script>

<template>
  <div v-if="visivel" class="toast" :class="tipo"> {{mensagem }} </div>
</template>

<style scoped>
.toast{
  position: fixed;
  top: 30px;
  right: 30px;
  padding: 16px 22px;
  border-radius: 10px;
  color: white;
  font-weight: bold;
  z-index: 99999;
  animation: aparecer 0.3s ease;
}

.success{
  background-color: #16a34a;
}

.error{
  background-color: #dc2626;
}

@keyframes aparecer{
  from{
    opacity: 0;
    transform: translateY(-10px);
  }

  to{
    opacity: 1;
    transform: translateY(0);
  }
}
</style>