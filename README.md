## Keylogger in Go

This Go program is a simple keylogger that monitors keyboard input and logs it to a file. Below is an explanation of its components:

### How It Works

- **GetAsyncKeyState**: Checks the state of a specified virtual key. Used to detect key presses.
- **GetKeyboardState**: Retrieves the status of all virtual keys. Used to check the current state of the keyboard.
- **MapVirtualKeyW**: Translates a virtual-key code into a scan code or character value. Used to translate virtual key codes to Unicode.
- **ToUnicode**: Translates the specified virtual-key code and keyboard state to the corresponding Unicode character or characters.

The program continuously loops to monitor key presses and writes the corresponding Unicode characters to a log file located at `C:\temp\keylogger.txt`.

#### License

This program is released under the Unlicense, which allows anyone to use, modify, and distribute the code freely, without restrictions.
