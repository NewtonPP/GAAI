#Command Similar to Cowsay 🐄
This project implements a Go command-line utility similar to the classic cowsay program. It generates a speech bubble containing text piped into it and appends an ASCII cow illustration below.

##📚 Purpose
This project was created as a learning exercise to explore Go programming concepts.

##🧠 Key Concepts Demonstrated
###Standard Input Handling
Reads piped input using os.Stdin for dynamic and interactive command-line usage.

###String Manipulation
Formats text to align and fit into a speech bubble by handling character counts, whitespace, and padding.

###Control Flow
Differentiates between single-line and multi-line inputs, wrapping them with the appropriate borders.

###ASCII Art
Integrates a visual representation (ASCII cow) with the formatted speech bubble for a complete output.

##🛠 How to Use
bash
echo -e "Hello, World!\nHow are you?" | go run main.go

###🌟 Output Example

Command Similar to Cowsay 🐄
This project implements a Go command-line utility similar to the classic cowsay program. It generates a speech bubble containing text piped into it and appends an ASCII cow illustration below.

📚 Purpose
This project was created as a learning exercise to explore Go programming concepts.

🧠 Key Concepts Demonstrated
Standard Input Handling
Reads piped input using os.Stdin for dynamic and interactive command-line usage.

String Manipulation
Formats text to align and fit into a speech bubble by handling character counts, whitespace, and padding.

Control Flow
Differentiates between single-line and multi-line inputs, wrapping them with the appropriate borders.

ASCII Art
Integrates a visual representation (ASCII cow) with the formatted speech bubble for a complete output.

🛠 How to Use
bash
Copy code
fortune | go run main.go
OR

bash
Copy code
echo -e "Hello, World!\nHow are you?" | go run main.go
🌟 Output Example
markdown
Copy code
 ________________
/ Hello, World!  \
| How are you?   |
 ----------------
         \  ^__^
          \ (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
