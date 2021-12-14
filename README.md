# Laboratorio 3 - Sistemas Distribuidos
# Star Wars

# Integrantes
* Nicolás Puente      201873618-K
* Francisco Andrades  201673584-4
* Lucas Díaz Aravena  201673524-0
----------------------

# Información
1) Los procesos están distribuidos según:
	
a) Máquina Virtual 10.6.40.181
- Proceso: Fulcrum 1 y Leia Organa
- Puertos: Fulcrum en puerto 50023

b) Máquina Virtual 10.6.40.182
- Proceso: Fulcrum 2 y Ahsoka Tano
- Puertos: El Fulcrum 2 escucha en el puerto 50023

c) Máquina Virtual 10.6.40.183
- Proceso: Broker Mos Eisley
- Puertos: Broker escucha en puerto 50000

d) Máquina Virtual 10.6.40.184
- Proceso: Fulcrum 3 y Almirante Thrawn
- Puertos: Fulcrum en puerto 50023


# Instrucciones ejecución

Importante: Los informantes y la Princesa Leia deben ejecutarse luego de haber iniciado al Broker y los Servidores Fulcrum

* dist41
Abrir 2 consolas .

(i)
´´´cd sd_lab2
make fulcrum´´´

(ii)
´´cd sd_lab2
make Leia´´

* dist43
Abrir 1 consola.

(i)
´cd sd_lab2
make broker´

* dist44
Abrir 2 consolas.

(i)
´cd sd_lab2
make fulcrum´

(ii)
´cd sd_lab2
make informantx´

* dist42
Abrir 2 consolas.

(i)
´cd sd_lab2
make fulcrum´

(ii)
´cd sd_lab2
make informantx´