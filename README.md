# irremote-example
Example of how to work with IR remotes with Arduino boards.

The setup involves connecting two push buttons, one IR transmitter 
and IR receiver to Arduino Nano 33 IoT board.
* Green button, when pressed, sends IR code via IR transmitter
* Red button, when pressed, halts the entire program
* IR receiver is always listening and decoding any data received

## hardware setup
Connect IR sensor to pin 9 and two push buttons on pin 10 and 11, respectively.
Connect IR transmitter to pin 12. Also keep an IR remote handy for testing.

## software setup
Install
[tinygo](https://github.com/tinygo-org/tinygo/releases),
[Arduino IDE](https://www.arduino.cc/en/software),
[bossac](https://github.com/shumatech/BOSSA/releases), and
[avrdude](https://github.com/avrdudes/avrdude/releases) on your computer.

Make sure that the setup is working by running
[this](c-code/main.c) code in the `Arduino IDE`

Identify the port on which `arduino` is connected
```bash
tinygo ports
```

Start serial debugger by first configuring it:
```bash
stty -f <your-port> raw 9600 cs8 clocal -cstopb
```

Start serial debugger:
```bash
cat <your-port>
```

## test
Press buttons on the remote, and you should see key info on the serial debugger
