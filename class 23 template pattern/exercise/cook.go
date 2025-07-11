package main

type ICook interface {
	prepareIngredients()
	cook()
	serve()
}

type Cook struct {
	iCook ICook
}

func (c *Cook) prepareMeal() {
	c.iCook.prepareIngredients()
	c.iCook.cook()
	c.iCook.serve()
}

// ---- Pizza Implementacion concreta ----
type Pizza struct {
	Cook
}

func (p *Pizza) prepareIngredients() {
	println("Pizza: Preparing dough, sauce, cheese, and toppings.")
}

func (p *Pizza) cook() {
	println("Pizza: Baking in the oven at 220°C for 15 minutes.")
}

func (p *Pizza) serve() {
	println("Pizza: Slicing and serving hot with extra cheese.")
}

// ---- Pasta Implementacion concreta ----
type Pasta struct {
	Cook
}

func (p *Pasta) prepareIngredients() {
	println("Pasta: Boiling water, adding pasta, and preparing sauce.")
}

func (p *Pasta) cook() {
	println("Pasta: Cooking pasta for 10 minutes and mixing with sauce.")
}

func (p *Pasta) serve() {
	println("Pasta: Serving with grated cheese and herbs on top.")
}

// ---- Ensalada Implementacion concreta ----
type Ensalada struct {
	Cook
}

func (e *Ensalada) prepareIngredients() {
	println("Ensalada: Washing and chopping vegetables, preparing dressing.")
}

func (e *Ensalada) cook() {
	println("Ensalada: No cooking needed, just mixing ingredients.")
}

func (e *Ensalada) serve() {
	println("Ensalada: Serving fresh with a drizzle of olive oil and vinegar.")
}

func main() {
	pizza := &Pizza{}
	c := Cook{
		iCook: pizza,
	}
	println("Preparing Pizza:")
	c.prepareMeal()
	println()

	pasta := &Pasta{}
	c = Cook{
		iCook: pasta,
	}
	println("Preparing Pasta:")
	c.prepareMeal()
	println()

	ensalada := &Ensalada{}
	c = Cook{
		iCook: ensalada,
	}
	println("Preparing Ensalada:")
	c.prepareMeal()
}
