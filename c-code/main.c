#include <Arduino.h>
#include <IRremote.h>

#define IR_RECEIVE_PIN 9

// Constants representing remote key codes
#define RemotePower 0x2
#define RemoteSource 0x1
#define RemoteKey1 0x4
#define RemoteKey2 0x5
#define RemoteKey3 0x6
#define RemoteKey4 0x8
#define RemoteKey5 0x9
#define RemoteKey6 0xA
#define RemoteKey7 0xC
#define RemoteKey8 0xD
#define RemoteKey9 0xE
#define RemoteKey0 0x11
#define RemoteDash 0x23
#define RemotePreCh 0x13
#define RemoteMute 0xF
#define RemoteChList 0x6B
#define RemoteVolUp 0x7
#define RemoteVolDn 0xB
#define RemoteChUp 0x12
#define RemoteChDown 0x10
#define RemoteMenu 0x1A
#define RemoteSmartHub 0x79
#define RemoteGuide 0x4F
#define RemoteTools 0x4B
#define RemoteInfo 0x1F
#define RemoteUp 0x60
#define RemoteDown 0x61
#define RemoteLeft 0x65
#define RemoteRight 0x62
#define RemoteEnter 0x68
#define RemoteReturn 0x58
#define RemoteExit 0x2D
#define RemoteA 0x6C
#define RemoteB 0x14
#define RemoteC 0x15
#define RemoteD 0x16
#define RemoteEManual 0x3F
#define RemoteSports 0xB8
#define RemoteCC 0x25
#define RemoteStop 0x46
#define RemoteReverse 0x45
#define RemotePlay 0x47
#define RemotePause 0x4A
#define RemoteForward 0x48

const char* getRemoteKeyName(uint16_t key) {
  switch (key) {
  case RemotePower: return "Power";
    case RemoteSource: return "Source";
    case RemoteKey0: return "0";
    case RemoteKey1: return "1";
    case RemoteKey2: return "2";
    case RemoteKey3: return "3";
    case RemoteKey4: return "4";
    case RemoteKey5: return "5";
    case RemoteKey6: return "6";
    case RemoteKey7: return "7";
    case RemoteKey8: return "8";
    case RemoteKey9: return "9";
    case RemoteDash: return "Dash";
    case RemotePreCh: return "PreCh";
    case RemoteMute: return "Mute";
    case RemoteChList: return "ChList";
    case RemoteVolUp: return "VolUp";
    case RemoteVolDn: return "VolDn";
    case RemoteChUp: return "ChUp";
    case RemoteChDown: return "ChDown";
    case RemoteMenu: return "Menu";
    case RemoteSmartHub: return "SmartHub";
    case RemoteGuide: return "Guide";
    case RemoteTools: return "Tools";
    case RemoteInfo: return "Info";
    case RemoteUp: return "Up";
    case RemoteDown: return "Down";
    case RemoteLeft: return "Left";
    case RemoteRight: return "Right";
    case RemoteEnter: return "Enter";
    case RemoteReturn: return "Return";
    case RemoteExit: return "Exit";
    case RemoteA: return "A";
    case RemoteB: return "B";
    case RemoteC: return "C";
    case RemoteD: return "D";
    case RemoteEManual: return "EManual";
    case RemoteSports: return "Sports";
    case RemoteCC: return "CC";
    case RemoteStop: return "Stop";
    case RemoteReverse: return "Reverse";
    case RemotePlay: return "Play";
    case RemotePause: return "Pause";
    case RemoteForward: return "Forward";
    default: return "Invalid Key";  // Or handle the invalid key appropriately
  }
}

const int buttonRed = 10;
const int buttonGreen = 11;
const int ledPin = LED_BUILTIN;  // Use the built-in LED

bool togglerState = false;

bool isButtonPressed(int pin) {
  return !digitalRead(pin);
}

void setup() {
  Serial.begin(9600);
  pinMode(buttonRed, INPUT_PULLUP);
  pinMode(buttonGreen, INPUT_PULLUP);
  pinMode(ledPin, OUTPUT);
  IrReceiver.begin(IR_RECEIVE_PIN);
}

void stop() {
  while (1)
    ;
}


void loop() {
  static unsigned int count = 0;
  static bool running = true;

  if (IrReceiver.decode()) {
    IrReceiver.resume();
    Serial.println(getRemoteKeyName(IrReceiver.decodedIRData.command));
  }

  // Simulate OnButtonPressedOnce (polling)
  if (isButtonPressed(buttonRed)) {
    Serial.println("interrupted: exiting main");
    running = false;  // Set a flag to stop the main loop
    stop();
  }


  if (running) {
    // Simulate OnButtonPressedRecurring (polling)
    if (isButtonPressed(buttonGreen)) {
      Serial.println("button pressed");

      count++;
      togglerState = !togglerState;
      digitalWrite(ledPin, togglerState ? HIGH : LOW);
      delay(50);  // Debounce delay – adjust as needed
    }

    delay(100);
  } else {
    // Do nothing effectively stopping the loop
  }
}