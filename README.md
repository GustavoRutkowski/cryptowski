# 🔑 Criptowski

O ***Cryptowski*** é um conjunto de algoritmos de criptografia simétrica desenvolvidos para fins acadêmicos e disponibilizados por uma interface de CLI.

O pacote `cryptowski` também conta com um CLI para aplicação de um algoritmo de força bruta, que também é meramente acadêmico, mas bastante útil para testar os algoritmos de criptografia.


## 📌 Sumário

* [🚀 Como Rodar o Projeto](#-como-rodar-o-projeto)
    * [Downloads](#-downloads)
    * [Instalação via Build](#️-instalação-via-build)
* [📚 Documentação dos CLIs](#-documentação-dos-clis)
    * [cryptowski](#cryptowski)
    * [bruteforce](#bruteforce)
    * [Flag de Versão](#flag-de-versão)
* [💥 O Algoritmo de Força Bruta](#-o-algoritmo-de-força-bruta)
* [🔐 Os Algorimos de Criptografia](#-os-algoritmos-de-criptografia)


## 🚀 Como Rodar o Projeto

No momento existem duas alternativas para rodar o Cryptowski na sua máquina.

Você pode instalar alguma das versões pré-compiladas na nossa [página de downloads](https://github.com/GustavoRutkowski/cryptowski/releases).

Ou, se preferir, você mesmo pode compilar o projeto na sua máquina se tiver o ***Golang*** instalado.

---


### 📁 Downloads

[Clique aqui para ir para a aba de downloads!](https://github.com/GustavoRutkowski/cryptowski/releases)


Após instalar o pacote, basta descompacta-lo e então [finalizar a instalação](#-finalizando-a-instalação)

### 🛠️ Instalação via Build

Caso você prefira compilar os arquivos para binário manualmente, então precisará buildar isso usando o **compilador do Go**

Caso seja necessário, [instale o Golang](https://go.dev/dl/).

**Passo-a-passo:**

1. [Instale o zip do projeto](https://github.com/GustavoRutkowski/cryptowski/archive/refs/heads/main.zip);

2. Crie uma pasta **`/cryptowski`** e extraia o zip dentro dela;

3. Abra a pasta do extraída com o terminal;

4. Rode os comandos abaixo no terminal:

    ```shell
    GOOS=<os> GOARCH=<arch> go build -o cryptowski ./cmd/cryptowski

    GOOS=<os> GOARCH=<arch> go build -o bruteforce ./cmd/bruteforce 
    ```
    * **`GOOS`** indica o Sistema Operacional (`linux` | `windows` | `darwin`)

    * **`GOARCH`** indica a arquitetura (`amd64` | `arm64`)

---

### ✅ Finalizando a Instalação

Independente de qual método você escolheu para instalar o projeto (download ou build manual), ainda pode ser necessário vincular os arquivos a um diretório específico dependendo do seu Sistema Operacional:

* [Windows](#-windows)
* [Linux](#-linux)
* [MacOS](#-macos)

--- 

#### 🪟 Windows:

Crie um diretório para armazenar os executáveis do Cryptowski:
```bash
mkdir "$env:USERPROFILE\bin"
```

Mova os executáveis para esse diretório:
```bash
move cryptowski.exe "$env:USERPROFILE\bin"
move bruteforce.exe "$env:USERPROFILE\bin"
```

Adicione o diretório ao PATH do usuário:
```bash
setx PATH "$env:USERPROFILE\bin;$env:PATH"
```
> Reinicie o terminal após rodar o comando acima

Teste se a instalação funcionou:
```bash
cryptowski --help
bruteforce --help
```

#### 🐧 Linux:

Garanta que `~/.local/bin` esteja no PATH:
```bash
mkdir -p ~/.local/bin
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Mova os arquivos compilados para o diretório correto:
```bash
mv cryptowski ~/.local/bin
mv bruteforce ~/.local/bin
chmod +x ~/.local/bin/cryptowski
chmod +x ~/.local/bin/bruteforce
```

Teste se a instalação funcionou:
```bash
cryptowski --help
bruteforce --help
```
> Em muitas distribuições modernas `~/.local/bin` já está configurado no PATH por padrão.


#### 🍎 MacOS:

Garanta que ~/.local/bin esteja no PATH:
```bash
mkdir -p ~/.local/bin
```

Se estiver usando zsh (padrão nas versões atuais do macOS):
```zsh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Se estiver usando bash:
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bash_profile
source ~/.bash_profile
```

Mova os executáveis para o diretório correto:
```bash
mv cryptowski ~/.local/bin
mv bruteforce ~/.local/bin
chmod +x ~/.local/bin/cryptowski
chmod +x ~/.local/bin/bruteforce
```

Teste se a instalação funcionou:
```bash
cryptowski --help
bruteforce --help
```


## 📚 Documentação dos CLIs


### 📗 `cryptowski`

Responsável por fornecer uma interface para os comandos encode/decode dos algoritmos desenvolvidos.

> A documentação completa pode ser lida utilizando `cryptowski --help`

**Principais comandos:**


#### `cryptowski encode` — Criptografa um arquivo.

```bash
cryptowski encode plaintext.txt -o encoded.enc -k my_s3cret_key
```

Lê o arquivo `plaintext.txt` e gera um arquivo criptografado `encoded.enc` no ***CWD*** atual.

---


#### `cryptowski decode` — Descriptografa um arquivo.

```
cryptowski decode encoded.enc -o plaintext.dec -k my_s3cret_key
```

Lê o arquivo `encoded.enc` e gera um arquivo descriptografado `plaintext.dec` no ***CWD*** atual.

---


### 📕 `bruteforce`

Responsável por fornecer uma interface para os comandos de força bruta. Muito útil para testar a eficácia dos algoritmos de criptografia.

> A documentação completa pode ser lida utilizando `bruteforce --help`

**Quebrar chaves de tamanho fixo:**
```bash
# Busca por uma chave com 6 caracteres que consiga transformar encoded.enc em um texto legível.
# Caso não encontre uma chave com exatos 6 caracteres, informa "key not found".
bruteforce encoded.enc --size 6
```

**Quebrar chaves de tamanho variado:**
```bash
# Busca por uma chave de 4 a 9 caracteres que consiga transformar encoded.enc em um texto legível.
bruteforce encoded.enc --minsize 4 --maxsize 9
```
> **Obs.**: `--minsize` pode ser omitida, neste caso o *bruteforce* começa procurando senhas de tamanho 1

---


### 🚩 Flag de Versão

Ambos os CLIs possuem flags como: `--v1`, `--v2`, `--v3`, etc. Que indicam a versão do algoritmo que está sendo usada para criptografar/descriptografar.

`cryptowski --v1 encode input.txt -o output.enc -k admin1234`
> Se nenhuma versão for informada ele pega a mais moderna disponível

---


## [💥 O Algoritmo de Força Bruta](/docs/BRUTEFORCE.md)

Como o algoritmo de força bruta funciona de uma maneira relativamente complexa, deixei-o em um arquivo separado. [Clique aqui](/docs/BRUTEFORCE.md) para entender como funciona o algoritmo de força bruta.

## 🔐 Os Algoritmos de Criptografia

Todos os algoritmos de criptografia seguem a seguinte interface:

```go
type ICrypto interface {
	Encode(data []byte, key string) []byte
	Decode(data []byte, key string) []byte
}
```

Porém, cada versão do algoritmo implementa uma lógica diferente em `Encode` e `Decode`.

Para mais detalhes, leia:
* [Documentação do algoritmo da V1](/docs/V1.md)
