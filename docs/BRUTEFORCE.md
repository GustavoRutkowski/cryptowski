# 💥 O Algoritmo de Força Bruta

Esse documento explica de maneira detalhada como funciona o algoritmo de força bruta.

Tenha em mente que o algoritmo de força bruta foi desenvolvido apenas para fins didáticos e não é exatamente igual a um bruteforce profissional, mas serve para os nossos objetivos.

---

Vamos analisar profundamente o que acontece por trás dos panos quando você roda:
```bash
bruteforce encoded.enc --size 6
```

Se você leu a [documentação do CLI de bruteforce](/README.md#-bruteforce) já sabe que esse comando irá aplicar força bruta no arquivo `encoded.enc` utilizando chaves de tamanho 6.


Mas... **Como isso funciona?**

## 📌 Sumário

1. [Entendendo o CHARSET do Algoritmo](#1-entendendo-o-charset-do-algoritmo)
2. [Núcleos do Processador](#2-núcleos-do-processador)
3. [Dividindo as Tarefas Entre Núcleos](#3-dividindo-as-tarefas-entre-núcleos)
    1. [Calculando o Total de Chaves](#31-calculando-o-total-de-chaves)
    2. [Dividindo as Chaves Igualmente](#32-dividindo-as-chaves-igualmente)
    3. [Como Converter um Índice para Chave?](#33-como-converter-um-índice-para-chave)
4. [Testando as Chaves](#4-testando-as-chaves)
    1. [Testando Caracteres Válidos](#41-testando-caracteres-válidos)
    2. [Testando Palavras Conhecidas](#42-testando-palavras-conhecidas)
5. [Considerações Finais](#5-considerações-finais)


## 1. Entendendo o CHARSET do Algoritmo

O CHARSET é o conjunto de caracteres que vai ser testado na hora de tentar quebrar uma chave, nele estão todos os caracteres que o algoritmo reconhece. Ex.:

Imagine que seu CHARSET é: `abcd`

Neste caso, as chaves testadas vão usar apenas esses caracteres.

```txt
aaaaaa -> Tentativa 1
aaaaab -> Tentativa 2
aaaaac -> Tentativa 3
...
ddddda
dddddb
dddddc
dddddd -> Última tentativa
```

> Caso a chave possua qualquer caractere diferente dos que estão no CHARSET, o algoritmo **NUNCA** vai conseguir quebrá-la pois ele não vai testar esses caracteres na hora de gerar as chaves.

O CHARSET oficial utilizado neste projeto é mais abrangente. Veja abaixo:
```go
// Possui letras minúsculas, letras maiúsculas e números
// Totalizando 62 caracteres
const KEY_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
```

Mesmo sendo maior, esse CHARSET ainda é bem modesto, possuindo apenas 62 caracteres.

Algoritmos profissionais costumam mapear símbolos, pontos e espaços também, além de permitir usar CHARSETs customizados e aplicar regras na força bruta.

---

É interessante entender um trade-off acerca do tamanho dos CHARSETs:

**Quanto mais caracteres mapeados:**
* Mais abrangente o algoritmo fica
* Mais lento o algoritmo fica (afinal ele vai precisar testar mais combinações)

A escolha do CHARSET depende da sua necessidade.

Optei inicialmente por um CHARSET alfanumérico básico para fins acadêmicos, mas pretendo tornar isso configurável via flags no CLI do bruteforce.


## 2. Núcleos do Processador

Quando você roda o programa de força bruta, a primeira coisa que ele faz é descobrir quantos [núcleos](https://tecnoblog.net/responde/o-que-e-nucleo-core-do-processador/) o seu processador possui.

Os núcleos vão determinar quantos *workers* estarão trabalhando simultaneamente no seu computador para tentar quebrar uma chave.

Quanto mais núcleos seu processador tem, mais rápido o bruteforce termina e mais longas são as senhas que ele é capaz de quebrar em tempos aceitáveis.


## 3. Dividindo as Tarefas Entre Núcleos

Apenas o fato do seu computador possuir vários núcleos não significa nada se o desenvolvedor criar algo que rode em uma única ***thread***/***goroutine***.

Por isso precisamos dividir as tarefas do nosso programa em diferentes unidades de trabalho chamadas ***workers***.


### 3.1. Calculando o Total de Chaves

Para calcular o número total de chaves que vão ser processadas com o tamanho `size`, precisamos utilizar a seguinte fórmula:

```bash
total = CHARSET^size
total = 62^6           # CHARSET de 62 caracteres e chave de tamanho 6
total = 56.800.235.584 # 56,8 bilhões de chaves
```

Quando informamos `--size 6`, o comando pode processar até 56 bilhões de chaves no total.

É claro que nem sempre ele vai realmente processar isso tudo, afinal quando ele encontra a chave ele encerra a busca. Se ele encontrar a chave logo no início da busca, ele não precisará executar todas as 56,8 bilhões de tentativas possíveis.


### 3.2. Dividindo as Chaves Igualmente

Agora que o algoritmo sabe que precisa processar até 56,8 bilhões de chaves, ele precisa distribuir isso de maneira uniforme entre os workers da aplicação.

Primeiramente precisamos dividir as chaves pela quantidade de núcleos:
```bash
keys_per_core = total / cores
# Presumindo que temos 4 núcleos:
keys_per_core = 56.800.235.584 / 4
keys_per_core = 14.200.058.896 # 14,2 bilhões de chaves para cada worker processar
```

Mas não basta apenas saber QUANTAS chaves cada worker deve processar, devemos organizar QUAIS chaves cada worker deve processar.

Vamos utilizar um exemplo mais simples:

Temos 1000 chaves para processar e cada um dos 4 workers vai processar 250 dessas chaves.

Neste caso, deveremos dividir em fatias (ou intervalos) iguais para cada worker:
```
Worker 1 --> Processa os índices de 0 a 249
Worker 2 --> Processa os índices de 250 a 499
Worker 3 --> Processa os índices de 500 a 749
Worker 4 --> Processa os índices de 750 a 999
```

Chamamos de índice pois esse número ainda não é a chave em si, apenas um identificador numérico que podemos converter em uma chave real posteriormente.

Repare que ficou tudo bem distribuído. Agora sabemos não só que cada worker processa 250 chaves como também quais são os intervalos de chave que cada worker processa.

> Repare que nem sempre essa divisão é exata.
> Quando a divisão deixa resto, o algoritmo sempre manda o último worker processar esse resto.


### 3.3. Como Converter um Índice para Chave?

Para converter um índice numérico para uma chave (texto) o algoritmo utiliza um simples cálculo de conversão de bases numéricas.

* Os índices numéricos seguem a base decimal (base 10)
* As chaves geradas seguem a base CHARSET (base 62)

Essa propriedade permite que a gente converta bases transformando um índice em uma chave legível. Ex.:
```
CHARSET = abcd

a = 0
b = 1
c = 2
d = 3

Índice 0 -> aaa
Índice 1 -> aab
Índice 2 -> aac
Índice 3 -> aad
Índice 4 -> aba
```


## 4. Testando as Chaves

Agora só nos resta entender como o algoritmo testa a chave e sabe se o teste deu certo.

Em algoritmos de força bruta tradicionais, onde testamos uma senha, as coisas são mais simples:

* Testa a senha (faz login)
* Se o login der certo --> Descobriu a senha
* Se o login deu errado --> Continua procurando pela senha

Entretanto, no nosso caso não estamos tentando descobrir uma senha de login armazenada como hash. Estamos tentando recuperar um texto criptografado utilizando a chave correta.

No Cryptowski, não existe uma maneira simples e eficiente de saber se um bloco de texto descriptografado é certo ou errado, você teria que ler ele para ter certeza, e por isso os algoritmos de força bruta em criptografia simétrica usam uma abordagem diferente.

Vamos entender o contexto atual. Cada worker:

1. Processa seus índices
2. Converte o índice para uma chave *K* válida
3. Chama a função ***`Decode`*** passando como argumentos:
    1. O arquivo `encoded.enc`
    2. A chave *K* que nós acabamos de converter
4. A função ***`Decode`*** retorna o bloco descriptografado.

O problema é: não sabemos se esse bloco descriptografado está correto ou não

Devido à lógica interna do algoritmo de criptografia, utilizar uma chave errada não causa um erro explícito, apenas devolve outro bloco ilegível.

Entretanto, é possível validar se esse bloco descriptografado está correto utilizando duas etapas de verificação sobre o bloco.


### 4.1. Testando Caracteres Válidos

A primeira etapa é pegar o bloco descriptografado e avaliar se os caracteres são válidos.

Em blocos criptografados é muito comum ter caracteres inválidos como `�`.

Em um texto potencialmente correto, os caracteres são letras, números, espaços, pontuação, quebras de linhas, etc.

Basta calcular quantos % dos caracteres do bloco descriptografado são caracteres válidos (letras, números, pontos, espaço, etc...)

O resultado será algo como:

`94% do texto é composto por caracteres válidos`

O algoritmo determina que se pelo menos 80% do texto for composto por caracteres válidos, então ele pode passar a próxima etapa de validação.

> Essa etapa sozinha não é suficiente para validar o texto, pois uma chave incorreta pode ocasionalmente produzir caracteres válidos por acaso.


### 4.2. Testando Palavras Conhecidas

A segunda etapa consiste em medir a taxa de **palavras conhecidas** presentes no texto.

Temos um arquivo de texto [google-10000-english.txt](/internal/bruteforce/google-10000-english.txt) contendo as 10 mil palavras mais usadas no inglês.

O algoritmo avalia palavra por palavra do bloco descriptografado e mede a taxa de palavras válidas.

No final dessa medição teremos um resultado como:

`78% das palavras do texto são compostas por palavras conhecidas`

Caso 50% ou mais do texto seja composto por palavras conhecidas, então ele é considerado válido.

Se um bloco de texto passar nessa segunda etapa: **parabéns, você descobriu a chave!**

Quando ocorre de um bloco passar nessas duas etapas, ele retorna duas coisas:

1. O bloco descriptografado em si
2. A chave responsável por gerar esse bloco

E são essas duas informações que o bruteforce retorna quando termina (se bem-sucedido).


## 5. Considerações Finais

O algoritmo possui uma abordagem relativamente simples para realizar a quebra de chaves.

De forma resumida, o algoritmo:

1. Descobre a quantidade N de núcleos no processador do usuário e cria N workers, um para cada núcleo.
2. Divide igualmente as chaves entre os workers
3. Cada worker itera sobre as suas chaves testando uma por uma
    1. Decodifica um bloco usando a chave gerada
    2. Checa se o bloco decodificado faz sentido
    3. Interrompe a busca e escreve a chave na sua tela

Você pode estar se perguntando como que o algoritmo funciona quando você usa `--minsize`/`--maxsize` ao invés de um tamanho fixo.

Na verdade ele faz exatamente a mesma coisa, a única diferença é que ele vai aumentando o tamanho da chave conforme necessário.

Rodar usando `--minsize` e `--maxsize` é o mesmo que rodar várias vezes usando `--size` com valores diferentes.
