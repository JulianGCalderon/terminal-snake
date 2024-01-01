# Terminal Snake

Implementación del clásico juego de SNAKE en la terminal. Para el renderizado, se utilizaron los codigos de escape de ASCII, y se establecio la terminal en modo TTY.

![demo](https://github.com/JulianGCalderon/terminal-snake/assets/60768809/5875a6df-40e9-4b3e-959b-fde1da710a1e)

## Compilación

Para ejecutarlo, podemos utilizar:

```bash
$ go run .
```

## Jugabilidad

La serpiente se controla con 'WASD'. El juego no tiene menu, al ejecutarlo comenzará automáticamente el juego, el cual finalizará en cuanto el jugador pierda (mostrando el puntaje).

Las dimensiones de la pantalla se actualizan automaticamente en cuanto se cambien las dimensiones del emulador de la terminal.
