package player

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/joptim/awesome-ost/backend/server/media"
	"github.com/joptim/beep-utils/buffers"
	"github.com/joptim/beep-utils/streamer"
	"os"
	"time"
)

type IPlayer interface {
	List() (map[string][]string, error)
	Add(id string) error
	Remove(id string) error
	RemoveAll() error
	Play() error
	Stop()
}

type Player struct {
	assets      *media.Media
	buffers     buffers.IBuffers
	mixer       *streamer.Synchronized
	initialised bool
}

func (p *Player) List() (map[string][]string, error) {
	path := os.Getenv("ASSETS_DIR")
	return p.assets.List(path)
}

func (p *Player) Add(id string) error {
	path, err := p.assets.GetPath(id)
	if err != nil {
		return err
	}
	if err = p.buffers.Load(path); err != nil {
		return err
	}
	st, err := p.buffers.GetStreamSeeker(path)
	if err != nil {
		return err
	}
	if err = p.mixer.Add(st, path); err != nil {
		return err
	}
	return nil
}

func (p *Player) Remove(id string) error {
	path, err := p.assets.GetPath(id)
	if err != nil {
		return err
	}
	p.mixer.Remove(path)
	return nil
}

func (p *Player) RemoveAll() error {
	return fmt.Errorf("Player.RemoveAll not implemented yet")
}

func (p *Player) Play() error {
	if err := p.initialise(); err != nil {
		return err
	}
	return nil
}

func (p *Player) Stop() {
	speaker.Clear()
}

func (p *Player) initialise() error {
	if p.initialised {
		return nil
	}
	format, err := p.buffers.GetFormat()
	if err != nil {
		return err
	}
	if err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return err
	}
	speaker.Play(beep.Loop(-1, p.mixer))
	return nil
}

func New() (IPlayer, error) {
	player := &Player{
		assets:  media.New(),
		buffers: buffers.New(nil),
		mixer:   streamer.NewMixer(),
	}
	if _, err := player.List(); err != nil {
		return nil, fmt.Errorf("cannot instantiate player: %v", err)
	}
	return player, nil
}
