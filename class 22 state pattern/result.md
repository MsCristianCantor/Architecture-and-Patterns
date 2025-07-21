# Feedback - Implementación del patrón State (Máquina Expendedora)

## 🧠 General

Tu implementación es sólida y demuestra un buen entendimiento del patrón **State**. Modelaste correctamente los estados (`HasItemState`, `HasMoneyState`, `ProductRequestedState`) y la lógica fluye bien dentro del `VendingMachineContext`.

---

## ✅ Cosas que hiciste bien

- ✅ Separaste correctamente los estados como structs que implementan la interfaz `State`.
- ✅ El `Context` (máquina expendedora) actúa como intermediario entre el cliente y los estados concretos.
- ✅ El cambio de estado (`setState`) está bien encapsulado.
- ✅ Implementaste validaciones útiles en cada estado para prevenir acciones no válidas.
- ✅ El ejemplo en `main()` cubre varios flujos importantes del sistema.

---

## 🧹 Sugerencias de mejora

### 1. Corrección ortográfica y naming

- `pruduct` → `product` en la interfaz `State`.
- `productSeleced` → `productSelected` (varias veces).
- `Producty already requested` → `Product already requested`.

Usar nombres claros y sin errores mejora mucho la mantenibilidad.

---

### 2. Simplificación de structs de producto

El campo `selected` dentro de `Product` no se utiliza. Podrías removerlo si no tiene función:

```go
type Product struct {
	name     string
	price    int
	quantity int
}
```

### 3. Inmutabilidad y seguridad del mapa de productos

Estás reasignando productos del mapa con esta línea:

```go
tmpProduct := i.vendingMachine.products[i.vendingMachine.productSeleced]
tmpProduct.quantity--
i.vendingMachine.products[i.vendingMachine.productSeleced] = tmpProduct
```

Podrías encapsular esta operación dentro de un método de `VendingMachineContext`, algo como:

```go
func (v *VendingMachineContext) reduceProductQuantity(name string) {
	product := v.products[name]
	product.quantity--
	v.products[name] = product
}
```

Así mantienes la lógica aislada.

### 4. Estado de error redundante

En `ProductRequestedState.SelectProduct`, el error:

```go
return fmt.Errorf("Producty already requested")
```
no solo tiene un typo, también podría mejorarse a algo como:
```go
return fmt.Errorf("Product already selected; dispense or cancel")
```

### 5. Validaciones preventivas

En `ProductRequestedState.Dispense()` accedes directamente al producto seleccionado sin validar si existe en el mapa. Podrías agregar una validación extra como:

```go
product, ok := i.vendingMachine.products[i.vendingMachine.productSeleced]
if !ok {
    return fmt.Errorf("Invalid product selected")
}
```

### 🌟 Extra sugerencia: manejo de estado OutOfStock

Podrías implementar un nuevo estado concreto `OutOfStockState` que se active cuando todos los productos se acaban. Esto le daría más robustez y realismo al modelo.

## 🧪 Test adicional sugerido

En el main, podrías incluir casos como:

* Intentar dispensar sin seleccionar producto.
* Eyectar monedas en diferentes estados.
* Insertar monedas adicionales y seleccionar productos costosos.

Esto haría tu flujo de prueba más completo.

## 🔚 Conclusión

Muy buen trabajo. Solo hay detalles menores de forma y organización que puedes mejorar. La estructura general y el uso del patrón están correctos. Si sigues construyendo sobre esta base, podrías incluso llegar a tener un framework mini de máquinas expendedoras con distintos tipos de lógicas o interfaces.

¡Sigue así! 🚀