package main

type Document interface {
	Accept(Visitor)
}

type Visitor interface {
	VisitInvoice(*Invoice)
	VisitReport(*Report)
	VisitReceipt(*Receipt)
}

// Concrete Document: Invoice
type Invoice struct {
	Number string
	Amount float64
}

func (i *Invoice) Accept(v Visitor) {
	v.VisitInvoice(i)
}

// Concrete Document: Report
type Report struct {
	Title   string
	Content string
}

func (r *Report) Accept(v Visitor) {
	v.VisitReport(r)
}

// Concrete Document: Receipt
type Receipt struct {
	Number string
	Amount float64
}

func (r *Receipt) Accept(v Visitor) {
	v.VisitReceipt(r)
}

// Concrete Visitor: PrintVisitor
type PrintVisitor struct{}

func (v *PrintVisitor) VisitInvoice(i *Invoice) {
	println("Factura:", i.Number, "Monto:", i.Amount)
}

func (v *PrintVisitor) VisitReport(r *Report) {
	println("Informe:", r.Title, "Contenido:", r.Content)
}

func (v *PrintVisitor) VisitReceipt(r *Receipt) {
	println("Recibo:", r.Number, "Monto:", r.Amount)
}

// Concrete Visitor: TotalVisitor
type TotalVisitor struct{}

func (v *TotalVisitor) VisitInvoice(i *Invoice) {
	println("Total Factura:", i.Amount)
}

func (v *TotalVisitor) VisitReport(r *Report) {
	println("Informe no tiene total.")
}

func (v *TotalVisitor) VisitReceipt(r *Receipt) {
	println("Total Recibo:", r.Amount)
}

// Concrete Visitor: ExportVisitor (optional, not implemented here)
type ExportVisitor struct{}

func (v *ExportVisitor) VisitInvoice(i *Invoice) {
	println("Exportando Factura:", i.Number)
}

func (v *ExportVisitor) VisitReport(r *Report) {
	println("Exportando Informe:", r.Title)
}

func (v *ExportVisitor) VisitReceipt(r *Receipt) {
	println("Exportando Recibo:", r.Number)
}

func main() {
	documents := []Document{
		&Invoice{Number: "12345", Amount: 250.75},
		&Report{Title: "Análisis de mercado", Content: "Contenido del informe..."},
		&Receipt{Number: "54321", Amount: 75.00},
	}

	printVisitor := &PrintVisitor{}
	totalVisitor := &TotalVisitor{}
	exportVisitor := &ExportVisitor{}

	for _, doc := range documents {
		doc.Accept(printVisitor)
		doc.Accept(totalVisitor)
		doc.Accept(exportVisitor)
	}
}
