# ASCII Art Web

## 📜 Description

**Ascii-art-web** is a web-based version of the ascii-art command-line project. It allows users to input custom text and render it in ASCII art format using a selected banner (font). The user interface is accessible through a web browser, providing a simple and interactive way to generate ASCII art.

Supported banner styles:
- `standard`
- `shadow`
- `thinkertoy`

The application handles form submissions via HTTP and returns formatted ASCII art using Go templates.

---

## 👨‍💻 Authors

- Yassine Bourazza
- Mohamed Hilli
- Abd-el-kafy Bourazza

---

## ▶️ Usage

### 1. Run the server

### bash
 - go run .

2. Open your browser and visit:
http://localhost:8080



### ⚙️ Implementation Details
- The server handles two main routes: / renders the homepage with a form, and /ascii-art processes POST requests containing user input and selected banner style. The handler reads the corresponding banner file, maps each character in the input to its ASCII art equivalent based on line offsets, assembles the result line by line, and returns it to the user using Go HTML templates. Errors like missing files or bad input are handled with appropriate HTTP status codes and a custom error page.

