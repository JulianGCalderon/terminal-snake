# Terminal Snake

Implementación del clásico juego de SNAKE en la terminal. Para el renderizado, se utilizaron los codigos de escape de ASCII, y se establecio la terminal en modo TTY.

## Compilación

Para ejecutarlo, podemos utilizar:

```bash
$ go run .
```

## Jugabilidad

La serpiente se controla con 'WASD'. El juego no tiene menu, al ejecutarlo comenzará automáticamente el juego, el cual finalizará en cuanto el jugador pierda (mostrando el puntaje).

Las dimensiones de la pantalla se actualizan automaticamente en cuanto se cambien las dimensiones del emulador de la terminal.
