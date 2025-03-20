#include <IRremote.h>

#define IR_RECEIVE_PIN 9
#define IR_SEND_PIN 3

void setup() {
  Serial.begin(9600);
  IrReceiver.begin(IR_RECEIVE_PIN);
  IrSender.begin(IR_SEND_PIN);
}

void loop() {
  IrSender.sendNEC(0x0102, 0x34, 0); // send code 0x34
  if (IrReceiver.decode()) {
    IrReceiver.resume();
    Serial.println(IrReceiver.decodedIRData.command, HEX);
  }
  delay(1000);
}