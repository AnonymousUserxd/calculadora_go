#!/usr/bin/python2

import os
import time
import platform
from sys import stdout, exit

RED, WHITE, CYAN, GREEN, END = '\033[91m', '\33[46m', '\033[36m', '\033[1;32m', '\033[0m'

def limpiar():
    os.system("clear")




def banner():
    print '''
 
                                                     
    
  ___                    _          
 | _ \__ _ __ _ _  _ ___| |_ ___ ___
 |  _/ _` / _` | || / -_)  _/ -_|_-<
 |_| \__,_\__, |\_,_\___|\__\___/__/
             |_|                    



       {2}Instalador de los paquetes
       Gracias por usar esta calculadora xD{3}'''.format(GREEN, END, CYAN, RED)

limpiar()
banner()
time.sleep(0.7)
limpiar()
time.sleep(0.3)
banner()
time.sleep(0.3)
limpiar()
time.sleep(0.3)
banner()
time.sleep(0.3)
limpiar()
time.sleep(0.3)
banner()
time.sleep(1.5) 

def paquetes():
    print("------------------------------------")
    print("Instalando los paquetes necesarios")
    print("------------------------------------")
    time.sleep(2)
    os.system("sudo apt-get update && sudo apt-get upgrade")
    os.system("apt-get install golang")
    os.system("apt-get install golang-doc")
    os.system("go get github.com/fatih/color")
    print("Los paquetes han sido instalados ejecute la herramienta con go run calculadora.go")

paquetes()
