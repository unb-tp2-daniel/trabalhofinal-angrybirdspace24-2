<template>
    <Header />
    <div class="page">
        <div class="top-section">
            <div>
                <h1>Matérias</h1>
                <p>Gerencie todas as matérias cadastradas no sistema</p>
            </div>

            <div class="top-actions">
                <button class="new-class-btn" @click="index">
                    Voltar
                </button>

                <button class="new-class-btn">
                    + Nova Turma
                </button>
            </div>
        </div>

        <div class="filters">
            <input
                v-model="filtro"
                type="text"
                placeholder="Buscar código, departamento, conteúdo, pré-requisito, equivalência..."
            />
        </div>

        <div class="results-info">
            {{ totalFiltrado }} matérias encontradas
        </div>

        <div v-if="loading">
            Carregando matérias...
        </div>

        <div v-else-if="error">
            {{ error }}
        </div>

        <section
            v-for="(listaMaterias, departamento) in materiasAgrupadas"
            :key="departamento"
            class="subject-section"
        >
            <div class="subject-header">
                <h2>{{ departamento }}</h2>
                <span>{{ listaMaterias.length }} matérias</span>
            </div>

            <table class="classes-table">
                <thead>
                    <tr>
                        <th>Código</th>
                        <th>Departamento</th>
                        <th>Carga Horária</th>
                        <th>Nível</th>
                        <th>Pré-requisitos</th>
                        <th>Co-requisitos</th>
                        <th>Equivalências</th>
                        <th>Conteúdo</th>
                    </tr>
                </thead>

                <tbody>
                    <tr
                        v-for="materia in listaMaterias"
                        :key="materia.codigo"
                    >
                        <td>
                            {{ materia.codigo }}
                        </td>

                        <td>
                            {{ materia.departamentoId }}
                        </td>

                        <td>
                            {{ materia.cargaHoraria }}h
                        </td>

                        <td>
                            {{ materia.nivelAcademico }}
                        </td>

                        <td>
                            <span
                                v-if="
                                    materia.preRequisitos &&
                                    materia.preRequisitos.length
                                "
                            >
                                {{
                                    materia.preRequisitos
                                        .flatMap(pr => pr.disciplinas)
                                        .join(', ')
                                }}
                            </span>

                            <span v-else>
                                Nenhum
                            </span>
                        </td>

                        <td>
                            <span
                                v-if="
                                    materia.coRequisitos &&
                                    materia.coRequisitos.length
                                "
                            >
                                {{
                                    materia.coRequisitos
                                        .flatMap(cr => cr.disciplinas)
                                        .join(', ')
                                }}
                            </span>

                            <span v-else>
                                Nenhum
                            </span>
                        </td>

                        <td>
                            {{
                                materia.equivalencias &&
                                materia.equivalencias.length
                                    ? materia.equivalencias.join(', ')
                                    : 'Nenhuma'
                            }}
                        </td>

                        <td class="conteudo-cell">
                            {{
                                materia.conteudo?.length > 120
                                    ? materia.conteudo.substring(0, 120) + '...'
                                    : materia.conteudo
                            }}
                        </td>
                    </tr>
                </tbody>
            </table>
        </section>

    </div>

    <Footer />
</template>

<script setup>
    import { ref, computed, onMounted } from 'vue'

    const filtro = ref('')
    const materias = ref([])
    const loading = ref(false)
    const error = ref(null)

    async function carregarMaterias() {
        loading.value = true

        try {
            const response = await $fetch(
                'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarMaterias'
            )

            materias.value = response
        } catch (err) {
            console.error(err)
            error.value = err.message
        } finally {
            loading.value = false
        }
    }

    const materiasAgrupadas = computed(() => {
        const grupos = {}

        for (const materia of materias.value) {
            if (!materiaContemFiltro(materia, filtro.value)) {
                continue
            }

            const departamento = materia.departamentoId

            if (!grupos[departamento]) {
                grupos[departamento] = []
            }

            grupos[departamento].push(materia)
        }

        return grupos
    })

    function materiaContemFiltro(materia, termo) {
        if (!termo) return true

        termo = termo.toLowerCase()

        const textoPesquisa = [
            materia.codigo,
            materia.departamentoId,
            materia.cargaHoraria,
            materia.nivelAcademico,
            materia.conteudo,

            ...(materia.equivalencias || []),

            ...(materia.preRequisitos || []).flatMap(
                pr => pr.disciplinas || []
            ),

            ...(materia.coRequisitos || []).flatMap(
                cr => cr.disciplinas || []
            )
        ]
            .join(' ')
            .toLowerCase()

        return textoPesquisa.includes(termo)
    }

    const totalFiltrado = computed(() =>
        Object.values(materiasAgrupadas.value)
            .reduce((acc, lista) => acc + lista.length, 0)
    )

    function index() {
        navigateTo('/coordenador')
    }

    onMounted(() => {
        carregarMaterias()
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

.conteudo-cell{
    max-width: 500px;
    min-width: 350px;
    white-space: normal;
    line-height: 1.4;
}

.results-info{
    margin-bottom:20px;
    color:#1a5276;
    font-size:14px;
    font-weight:600;
}

.top-actions{
    display:flex;
    gap:12px;
}
</style>