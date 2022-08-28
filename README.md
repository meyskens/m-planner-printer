# ESCPOS Printer for M-Planner

This is a ticket printer module for [M-Planner](https://github.com/mplanner/M-Planner).
It is built using TinyGo and currently focusses on the Arduino Nano RP2040 Connect (with WifiNINA).

This module fetches print jobs from the API every 0.5 seconds and prints them. Part of the ESCPS commands are done serverside (font options, text alignment, etc.) while the rest is done clientside on in TinyGo (initializing the printer, cutting, etc.).
