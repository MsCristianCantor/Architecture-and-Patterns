package main

import "fmt"

// sub-sistema de sonido
type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("Sistema de sonido prendido")
}

func (s *SoundSystem) ReproduceSound(sound string) {
	fmt.Printf("Reproduciendo sonido %s \n", sound)
}

// sub-sistema de proyector
type ProyectorSystem struct{}

func (p *ProyectorSystem) TurnOn() {
	fmt.Println("Proyector prendido")
}

func (p *ProyectorSystem) ReproduceVideo(video string) {
	fmt.Printf("Reproduciendo video %s \n", video)
}

// sub-sistema de reproductor blu-ray
type BluRaySystem struct {
	movie string
}

func (b *BluRaySystem) TurnOn() {
	fmt.Println("Sistema blue ray prendido")
}

func (b *BluRaySystem) SetMovie(movie string) {
	fmt.Printf("Pelicula insertada %s \n", movie)
	b.movie = movie
}

func (b *BluRaySystem) ReproduceMedia() (string, string) {
	return fmt.Sprintf("%s video", b.movie), fmt.Sprintf("%s audio", b.movie)
}

// HomeTheater Facade
type HomeTheaterFacade struct {
	sound     *SoundSystem
	proyector *ProyectorSystem
	bluRay    *BluRaySystem
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		sound:     &SoundSystem{},
		proyector: &ProyectorSystem{},
		bluRay:    &BluRaySystem{},
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	h.sound.TurnOn()
	h.proyector.TurnOn()
	h.bluRay.TurnOn()
	h.bluRay.SetMovie(movie)
	video, audio := h.bluRay.ReproduceMedia()
	h.sound.ReproduceSound(audio)
	h.proyector.ReproduceVideo(video)
}

func main() {
	facade := NewHomeTheaterFacade()
	facade.WatchMovie("star wars I")
}
