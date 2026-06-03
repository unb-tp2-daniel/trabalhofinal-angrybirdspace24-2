import http from 'k6/http';
import { check, sleep } from 'k6';

/* criado com auxilio do gemini para os testes */

export const options = {
  scenarios: {
    onda_de_alunos: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '10s', target: 150 }, // simula 150 alunos entrando juntos nos primeiros 10 segundos
        { duration: '20s', target: 150 }, // mantém esse pico de 150 acessos simultâneos por 20 segundos
        { duration: '5s', target: 0 },   // alunos vão fechando o site
      ],
    },
  },
};

export default function () {
  const urlFrontend = 'http://192.168.0.11:3000/aluno/matricula'; 
  const resFrontend = http.get(urlFrontend);

  check(resFrontend, {
    'HTML do Nuxt carregou (200)': (r) => r.status === 200,
  });

  sleep(0.1); 

  const urlBuscarTurmas = 'https://southamerica-east1-matriculas242.cloudfunctions.net/ListarTurmas'; 
  const resBackend = http.get(urlBuscarTurmas);

  check(resBackend, {
    'Lista de turmas carregou (200)': (r) => r.status === 200,
  });

  sleep(Math.random() * 3 + 2); 
}