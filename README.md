# 🚀 Welcome to the One Commit Every Day Project! 🕒

Welcome, programmer! This project is designed to generate daily commits with custom messages using AI. It’s all automated, but you’re free to modify everything to suit your needs. Let’s explore the world of fun, automation, and creativity! 🌟

## 🌐 Multilingual Instructions

### 🇧🇷 Como executar o projeto:
1. Certifique-se de ter o Docker e o Docker Compose instalados.
2. Configure sua chave SSH:
   - Copie sua chave pública e privada usando o comando do Makefile: `make build`.
3. Para iniciar o projeto, execute:
   ```bash
   make up
   ```
4. Para executar os testes:
   ```bash
   make test
   ```
5. Desligue tudo:
   ```bash
   make down
   ```
6. **Aviso:** Este projeto usa sua chave SSH para se comunicar com o GitHub. Edite e personalize conforme necessário. 

---

### 🇺🇸 How to run the project:
1. Ensure Docker and Docker Compose are installed.
2. Set up your SSH key:
   - Use the Makefile command to copy your public and private keys: `make build`.
3. Start the project with:
   ```bash
   make up
   ```
4. Run tests:
   ```bash
   make test
   ```
5. Shut down:
   ```bash
   make down
   ```
6. **Note:** This project uses your SSH key to interact with GitHub. Feel free to edit and customize it as needed.

---

### 🇲🇽 Cómo ejecutar el proyecto:
1. Asegúrate de tener Docker y Docker Compose instalados.
2. Configura tu clave SSH:
   - Usa el comando del Makefile para copiar tus claves pública y privada: `make build`.
3. Inicia el proyecto con:
   ```bash
   make up
   ```
4. Ejecuta las pruebas:
   ```bash
   make test
   ```
5. Apaga todo:
   ```bash
   make down
   ```
6. **Nota:** Este proyecto usa tu clave SSH para interactuar con GitHub. Edítalo y personalízalo según sea necesario.

---

## 🛠️ Things You Can Edit 📝

### Configurations in `main.go`
All configurations for this project are located in the `const` block:

```go
const (
    ollamaURL      = "http://ollama:11434/api/generate" // URL for the Ollama API
    outputDir      = "generated_files"                // Directory for saving generated files
    promptOllama   = "I'm going to write a simple readme.md that explains how to write a simple hello world in Python:" // The AI prompt
    cronSchedule   = "0 8 * * *"                       // Cron schedule for generating commits
    fileExtension  = "md"                             // File extension for generated files
)
```

#### What You Can Change:
1. **`ollamaURL`:** Update the URL if your Ollama API is hosted elsewhere.
2. **`outputDir`:** Change the directory where generated files are saved.
3. **`promptOllama`:** Modify the prompt to generate messages tailored to your needs.
4. **`cronSchedule`:** Adjust the cron schedule to generate commits at different times.
5. **`fileExtension`:** Change the file type for the generated files (e.g., `txt`, `json`).

### Documentation
- [Ollama Documentation](https://github.com/ollama/ollama): Learn more about Ollama and its capabilities.

### AI Model
This project uses the **llama3.2:1b** model from Ollama. It’s lightweight, fast, and efficient for generating meaningful messages. Feel free to experiment with other models available in the Ollama suite.

---

## 🌟 Additional Features
- **Fork-Friendly:** Clone, fork, and modify this project as you like.
- **Customizable Cron Jobs:** Edit the `cronSchedule` to generate commits as frequently as desired.
- **Extensible Code:** The code is clean and modular, so you can easily add new features or modify existing ones.

---

## 💡 Notes for Developers
- Always ensure your SSH keys are correctly set up.
- Test modifications locally using `make test` before deploying changes.
- Keep the fun alive by exploring new prompts and tweaking automation settings.

---

### 🎉 Have fun and automate your GitHub commits like a pro! 🚀
