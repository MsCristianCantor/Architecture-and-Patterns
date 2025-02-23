# Feedback sobre la implementación del Memento Pattern en un editor de texto

## ✅ Aspectos positivos:
1. **Cumple con la estructura del patrón Memento**: Se han definido correctamente las tres partes esenciales del patrón:
   - `Memento`: almacena el estado del `Editor`.
   - `Editor`: crea y usa los `Memento`.
   - `Caretaker`: gestiona el historial de estados.
   
2. **Uso adecuado de structs y métodos**: Se ha encapsulado correctamente la lógica dentro de las estructuras respectivas.

3. **Flujo lógico correcto**:
   - Se guarda un estado del `Editor`.
   - Se modifica el texto.
   - Se restaura el estado anterior desde el `Caretaker`.

4. **Eficiencia en la recuperación del estado**: `GetMemento` verifica que el índice esté dentro del rango, evitando errores por acceso fuera de los límites.

---

## 🛠️ Posibles mejoras:
1. **Agregar una estructura de datos más robusta para gestionar el historial**  
   Actualmente, el `Caretaker` usa un slice (`history []*Memento`), pero no hay forma de deshacer (`Undo`) múltiples veces o eliminar un estado.  
   ➡️ **Sugerencia**: Implementar una pila (stack) para un mejor control de `Undo` y `Redo`.

2. **Mejor manejo de errores en `GetMemento`**  
   Si el índice es inválido, la función devuelve `nil`, lo que puede causar un `nil pointer dereference` si no se maneja bien.  
   ➡️ **Sugerencia**: Agregar un `log.Println("Índice fuera de rango")` o devolver un `error`.

3. **Encapsulación de la variable `texto` en `Editor`**  
   Actualmente, `texto` es accesible directamente, lo que rompe el encapsulamiento.  
   ➡️ **Sugerencia**: Usar métodos `SetText()` y `GetText()` en `Editor`.

4. **Agregar funcionalidad de Redo**  
   ➡️ **Sugerencia**: Implementar una estructura separada para almacenar los `Memento` que han sido deshechos, permitiendo "rehacer" (`Redo`) cambios si es necesario.

---

## 🏆 Versión mejorada con Undo y Redo:

```go
package main

import "fmt"

// Memento almacena el estado del editor
type Memento struct {
	texto string
}

// Editor crea y usa Mementos
type Editor struct {
	texto string
}

func (e *Editor) Save() *Memento {
	return &Memento{texto: e.texto}
}

func (e *Editor) Restore(m *Memento) {
	if m != nil {
		e.texto = m.texto
	}
}

func (e *Editor) SetText(text string) {
	e.texto = text
}

func (e *Editor) GetText() string {
	return e.texto
}

// Caretaker administra los Mementos
type Caretaker struct {
	history []*Memento
	redo    []*Memento
}

func (c *Caretaker) SaveMemento(m *Memento) {
	c.history = append(c.history, m)
	c.redo = nil // Se borra la pila de Redo después de un nuevo cambio
}

func (c *Caretaker) Undo() *Memento {
	if len(c.history) == 0 {
		return nil
	}
	// Obtener el último estado
	m := c.history[len(c.history)-1]
	// Moverlo a Redo antes de eliminarlo de la pila de Undo
	c.redo = append(c.redo, m)
	c.history = c.history[:len(c.history)-1]
	return m
}

func (c *Caretaker) Redo() *Memento {
	if len(c.redo) == 0 {
		return nil
	}
	// Obtener el último estado eliminado
	m := c.redo[len(c.redo)-1]
	c.redo = c.redo[:len(c.redo)-1]
	c.history = append(c.history, m) // Guardar en Undo nuevamente
	return m
}

func main() {
	editor := &Editor{}
	caretaker := &Caretaker{}

	editor.SetText("Texto inicial")
	fmt.Println("Texto actual:", editor.GetText())

	// Guardar estado
	caretaker.SaveMemento(editor.Save())

	editor.SetText("Texto modificado")
	fmt.Println("Texto modificado:", editor.GetText())

	// Undo
	editor.Restore(caretaker.Undo())
	fmt.Println("Después de Undo:", editor.GetText())

	// Redo
	editor.Restore(caretaker.Redo())
	fmt.Println("Después de Redo:", editor.GetText())
}
```
## 🔥 Conclusión:

Tu implementación es funcional y clara. Sin embargo, con pequeñas mejoras como el manejo de errores, encapsulación y soporte para Undo/Redo, se puede hacer más robusta y flexible. 🚀