<template>
    <Header />
    <div class="page">
        <div class="top-section">
            <div>
                <h1>Turmas</h1>
                <p>Gerencie todas as turmas cadastradas no sistema</p>
            </div>

            <button class="new-class-btn">
                + Nova Turma
            </button>
        </div>

        <div class="filters">
            <input
                v-model="filtro"
                type="text"
                placeholder="Buscar turma, matéria, professor, horário, local..."
            />
        </div>

        <div class="results-info">
            {{ turmasFiltradas.length }} turmas encontradas
        </div>

        <div v-if="loading">
            Carregando turmas...
        </div>

        <div v-else-if="error">
            {{ error }}
        </div>

        <div v-else>
            <section class="subject-section">
                <div class="subject-header">
                    <h2>Todas as Turmas</h2>
                    <span>{{ turmasFiltradas.length }} turmas</span>
                </div>

                <table class="classes-table">
                    <thead>
                        <tr>
                            <th>Matéria</th>
                            <th>Código</th>
                            <th>Professor</th>
                            <th>Horário</th>
                            <th>Local</th>
                            <th>Semestre</th>
                            <th>Vagas</th>
                            <th>Status</th>
                            <th>Ações</th>
                        </tr>
                    </thead>

                    <tbody>
                        <tr v-for="turma in turmasFiltradas":key="turma.codigoTurma">
                            <td>{{ turma.nomeMateria }}</td>

                            <td>{{ turma.codigoTurma }}</td>

                            <td>
                                {{
                                    turma.professorNome &&
                                    turma.professorNome.trim()
                                        ? turma.professorNome
                                        : 'Não informado'
                                }}
                            </td>

                            <td>{{ turma.horario }}</td>

                            <td>{{ turma.local }}</td>

                            <td>{{ turma.semestre }}</td>

                            <td>
                                {{ turma.vagasOcupadas }}/{{ turma.vagasTotais }}
                            </td>

                            <td>
                                <span
                                    class="status"
                                    :class="{
                                        active: turma.vagasOcupadas < turma.vagasTotais,
                                        full: turma.vagasOcupadas >= turma.vagasTotais
                                    }"
                                >
                                    {{
                                        turma.vagasOcupadas >= turma.vagasTotais
                                            ? 'Lotada'
                                            : 'Disponível'
                                    }}
                                </span>
                            </td>

                            <td class="actions">
                                <button class="edit" @click="editarTurma(turma)">
                                    Editar
                                </button>

                                <button class="students" @click="verDetalhes(turma)">
                                    Detalhes
                                </button>

                                <button class="delete" @click="deletarTurma(turma)">
                                    Excluir
                                </button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </section>
        </div>

        <button class="new-class-btn" @click="index">
            Voltar
        </button>
    </div>

    <div v-if="modalAberto" class="modal-overlay" @click="modalAberto = false">
        <div class="modal" @click.stop>
            <div class="modal-header">
                <div>
                    <h2>
                        Detalhes da Matéria
                    </h2>

                    <p v-if="materiaSelecionada" class="modal-subtitle">
                        <strong>{{ turmaSelecionada?.nomeMateria }}</strong>
                    </p>
                </div>

                <button class="close-btn" @click="modalAberto = false">
                    ✕
                </button>
            </div>

            <div v-if="carregandoMateria">
                Carregando...
            </div>

            <div
                v-else-if="materiaSelecionada"
                class="modal-content"
            >
                <p>
                    <strong>Código:</strong>
                    {{ materiaSelecionada.codigo }}
                </p>

                <p>
                    <strong>Departamento:</strong>
                    {{ materiaSelecionada.departamentoId }}
                </p>

                <p>
                    <strong>Carga Horária:</strong>
                    {{ materiaSelecionada.cargaHoraria }}h
                </p>

                <p>
                    <strong>Nível:</strong>
                    {{ materiaSelecionada.nivelAcademico }}
                </p>

                <p>
                    <strong>Pré-requisitos:</strong>

                    <span
                        v-if="
                            materiaSelecionada.preRequisitos &&
                            materiaSelecionada.preRequisitos.length
                        "
                    >
                        {{
                            materiaSelecionada.preRequisitos
                                .flatMap(pr => pr.disciplinas)
                                .join(', ')
                        }}
                    </span>

                    <span v-else>
                        Nenhum
                    </span>
                </p>

                <p>
                    <strong>Co-requisitos:</strong>

                    <span
                        v-if="
                            materiaSelecionada.coRequisitos &&
                            materiaSelecionada.coRequisitos.length
                        "
                    >
                        {{
                            materiaSelecionada.coRequisitos
                                .flatMap(cr => cr.disciplinas)
                                .join(', ')
                        }}
                    </span>

                    <span v-else>
                        Nenhum
                    </span>
                </p>

                <p>
                    <strong>Equivalências:</strong>

                    {{
                        materiaSelecionada.equivalencias?.join(', ') ||
                        'Nenhuma'
                    }}
                </p>

                <div class="conteudo-box">
                    <strong>Conteúdo:</strong>

                    <p>
                        {{ materiaSelecionada.conteudo }}
                    </p>
                </div>
            </div>

            <div v-else>
                Matéria não encontrada.
            </div>
        </div>
    </div>

    <Footer />
</template>

<script setup>
    import { ref, onMounted } from 'vue'

    const filtro = ref('')
    const turmaSelecionada = ref(null)
    const modalAberto = ref(false)
    const materiaSelecionada = ref(null)
    const carregandoMateria = ref(false)
    const turmas = ref([])
    const loading = ref(false)
    const error = ref(null)

    const turmasFiltradas = computed(() =>
    turmas.value.filter(turma =>
        turmaContemFiltro(turma, filtro.value)
    )
)

    async function carregarTurmas() {
        loading.value = true

        try {
            const response = await $fetch(
                'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas'
            )

            turmas.value = response
            console.log(response)
            console.log(JSON.stringify(response, null, 2))
        } catch (err) {
            console.error(err)
            error.value = err.message
        } finally {
            loading.value = false
        }
    }

    async function verDetalhes(turma) { 
        turmaSelecionada.value = turma
        modalAberto.value = true
        carregandoMateria.value = true

        try {
            const materias = await $fetch(
                'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarMaterias'
            )

            materiaSelecionada.value = materias.find(
                m => m.codigo === turma.materiaId
            )
        } catch (err) {
            console.error(err)
            materiaSelecionada.value = null
        } finally {
            carregandoMateria.value = false
        }
    }

    function turmaContemFiltro(turma, termo) {
        if (!termo) return true

        termo = termo.toLowerCase()

        const status =
            turma.vagasOcupadas >= turma.vagasTotais
                ? 'Lotada'
                : 'Disponível'

        const textoPesquisa = [
            turma.codigoTurma,
            turma.materiaId,
            turma.nomeMateria,
            turma.professorId,
            turma.professorNome,
            turma.semestre,
            turma.local,
            turma.horario,
            turma.vagasTotais,
            turma.vagasOcupadas,
            turma.ativo ? 'ativa' : 'inativa',
            status
        ]
            .filter(Boolean)
            .join(' ')
            .toLowerCase()

        return textoPesquisa.includes(termo)
    }

    function index() {
        navigateTo('/coordenador')
    }

    onMounted(() => {
        carregarTurmas()
    })
</script>

<style scoped>

.page{
    padding:30px;
    background:#f1f5f9;
    min-height:100vh;
}

.top-section{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:30px;
}

.top-section h1{
    font-size:36px;
    margin-bottom:6px;
}

.top-section p{
    color:#1a5276;
}

.new-class-btn{
    background:#1a5276;
    color:white;
    border:none;
    padding:14px 20px;
    border-radius:10px;
    cursor:pointer;
    font-weight:700;
    transition:0.2s;
}

.new-class-btn:hover{
    opacity:0.9;
}

.filters{
    display:flex;
    gap:16px;
    margin-bottom:30px;
}

.filters input,
.filters select{
    padding:12px;
    border-radius:10px;
    border:1px solid #cbd5e1;
    background:white;
}

.filters input{
    width:320px;
}

.subject-section{
    background:white;
    border-radius:16px;
    padding:24px;
    margin-bottom:30px;
    box-shadow:0 2px 8px rgba(0,0,0,0.06);
}

.subject-header{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:20px;
}

.subject-header h2{
    color:#1a5276;
}

.subject-header span{
    color:#1a5276;
    font-size:14px;
}

.classes-table{
    width:100%;
    border-collapse:collapse;
}

.classes-table th{
    background:#f8fafc;
    padding:14px;
    text-align:left;
    color:#334155;
    font-size:14px;
}

.classes-table td{
    padding:16px 14px;
    border-top:1px solid #e2e8f0;
}

.classes-table tr:hover{
    background:#f8fafc;
}

.actions{
    display:flex;
    gap:8px;
}

.actions button{
    border:none;
    padding:8px 12px;
    border-radius:8px;
    cursor:pointer;
    font-weight:600;
}

.edit{
    background:#dbeafe;
    color:#1d4ed8;
}

.students{
    background:#dcfce7;
    color:#166534;
}

.delete{
    background:#fee2e2;
    color:#b91c1c;
}

.status{
    padding:6px 10px;
    border-radius:999px;
    font-size:12px;
    font-weight:700;
}

.active{
    background:#dcfce7;
    color:#166534;
}

.warning{
    background:#fef3c7;
    color:#92400e;
}

.full{
    background:#fee2e2;
    color:#991b1b;
}

.actions button{
    border:none;
    padding:8px 12px;
    border-radius:8px;
    cursor:pointer;
    font-weight:600;
    transition:all 0.2s ease;
}

.actions button:hover{
    transform:translateY(-2px);
    box-shadow:0 4px 10px rgba(0,0,0,0.12);
}

.edit:hover{
    background:#bfdbfe;
}

.students:hover{
    background:#bbf7d0;
}

.delete:hover{
    background:#fecaca;
}

.modal-overlay{
    position:fixed;
    inset:0;
    background:rgba(0,0,0,0.45);
    display:flex;
    align-items:center;
    justify-content:center;
    z-index:9999;
}

.modal{
    background:white;
    width:800px;
    max-width:90vw;
    max-height:85vh;
    overflow-y:auto;
    border-radius:16px;
    padding:24px;
    box-shadow:0 10px 30px rgba(0,0,0,0.2);
}

.modal-header{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:20px;
}

.modal-header h2{
    color:#1a5276;
}

.close-btn{
    border:none;
    background:none;
    font-size:22px;
    cursor:pointer;
}

.modal-content{
    display:flex;
    flex-direction:column;
    gap:12px;
}

.conteudo-box{
    margin-top:10px;
    padding:16px;
    background:#f8fafc;
    border-radius:10px;
    border:1px solid #e2e8f0;
}

.conteudo-box p{
    margin-top:10px;
    line-height:1.6;
}

.modal-subtitle{
    margin-top:4px;
    color:#000000;
    font-weight:500;
    font-size:20px;
}

.results-info{
    margin-bottom:20px;
    color:#1a5276;
    font-size:14px;
    font-weight:600;
}

</style>