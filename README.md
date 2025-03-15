# irremote-example
Example of how to work with IR remotes with Arduino boards

## hardware setup
Connect IR sensor on pin 9 and two push buttons on pin 10 and 11, respectively.
Also keep an IR remote handy.

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
Press buttons on the remote and you should see key info on the serial debugger
