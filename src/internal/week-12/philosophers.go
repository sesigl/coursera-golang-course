package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const philosopherAndChopstickCount = 5

type chopstick struct {
	inUse             sync.Mutex
	name              string
	activePhilosopher *philosopher

	// philosopher to the left side
	lPhilosopher *philosopher
	// philosopher to the right side
	rPhilosopher *philosopher
}

func (c *chopstick) PickUpStick(p *philosopher) {
	if c.rPhilosopher != p && c.lPhilosopher != p {
		panic(fmt.Errorf("philosopher '%s' tried to pickup a chopstick '%s', which is not next to him", p.name, c.name))
	}

	fmt.Printf("[%s]: waiting to pick up stick '%s'\n", p.name, c.name)
	c.inUse.Lock()
	fmt.Printf("[%s]: picked up stick '%s'\n", p.name, c.name)
}

func (c *chopstick) LayDownStick(p *philosopher) {
	if c.rPhilosopher != p && c.lPhilosopher != p {
		panic(fmt.Errorf("philosopher '%s' tried to laydown a chopstick '%s', which is not next to him", p.name, c.name))
	}

	fmt.Printf("[%s]: Laying down stick '%s'\n", p.name, c.name)
	c.inUse.Unlock()
}

type philosopher struct {
	name   string
	eating sync.Mutex
	// in use chopsticks
	// nil when no chopsticks in use
	lChopStick            *chopstick
	rChopStick            *chopstick
	allowedToEat          bool
	pastAllowedToEatCount uint8
	maxAllowedToEatCount  uint8
}

func AllowToEat(p *philosopher, wg *sync.WaitGroup) {
	p.AllowToEat()

	p.DisallowToEat()

	wg.Done()
}

func (p *philosopher) canEatAnotherTime() bool {
	return p.pastAllowedToEatCount < p.maxAllowedToEatCount
}

func (p *philosopher) PickUpSticks() {
	if rand.Intn(2) == 0 {
		p.lChopStick.PickUpStick(p)
		p.rChopStick.PickUpStick(p)
	} else {
		p.rChopStick.PickUpStick(p)
		p.lChopStick.PickUpStick(p)
	}

}

func (p *philosopher) LayDownSticks() {
	if rand.Intn(2) == 0 {
		p.lChopStick.LayDownStick(p)
		p.rChopStick.LayDownStick(p)
	} else {
		p.rChopStick.LayDownStick(p)
		p.lChopStick.LayDownStick(p)
	}
}

func (p *philosopher) AllowToEat() {
	p.eating.Lock()
	fmt.Printf("[%s]: start eating\n", p.name)

	p.pastAllowedToEatCount += 1
	p.allowedToEat = true
	p.PickUpSticks()
}

func (p *philosopher) DisallowToEat() {
	p.LayDownSticks()
	p.allowedToEat = false

	fmt.Printf("[%s]: finish eating\n", p.name)
	p.eating.Unlock()
}

type host struct {
	allowedPhilosopherToEat1 *philosopher
	allowedPhilosopherToEat2 *philosopher
	philosophers             []*philosopher
}

func (h *host) SelectNextPhilosophers() bool {
	fmt.Printf("[host]: Selecting next philosophers\n")
	if h.allowedPhilosopherToEat1 == nil && h.allowedPhilosopherToEat2 == nil {
		h.allowedPhilosopherToEat1 = h.philosophers[0]
		h.allowedPhilosopherToEat2 = h.philosophers[2]
	} else {
		h.allowedPhilosopherToEat1 = h.allowedPhilosopherToEat1.rChopStick.rPhilosopher
		h.allowedPhilosopherToEat2 = h.allowedPhilosopherToEat2.rChopStick.rPhilosopher
	}

	fmt.Printf("[host]: Selected %s and %s\n", h.allowedPhilosopherToEat1.name, h.allowedPhilosopherToEat2.name)

	wg := sync.WaitGroup{}

	if h.allowedPhilosopherToEat1.canEatAnotherTime() {
		fmt.Printf("[host]: %s allowed to eat (%d/%d)\n", h.allowedPhilosopherToEat1.name, h.allowedPhilosopherToEat1.pastAllowedToEatCount, h.allowedPhilosopherToEat1.maxAllowedToEatCount)
		wg.Add(1)
		go AllowToEat(h.allowedPhilosopherToEat1, &wg)
	} else {
		fmt.Printf("max eat allowed count (%d) for philosopher '%s' reached\n", h.allowedPhilosopherToEat1.maxAllowedToEatCount, h.allowedPhilosopherToEat1.name)
	}

	if h.allowedPhilosopherToEat2.canEatAnotherTime() {
		fmt.Printf("[host]: %s allowed to eat (%d/%d)\n", h.allowedPhilosopherToEat2.name, h.allowedPhilosopherToEat2.pastAllowedToEatCount, h.allowedPhilosopherToEat2.maxAllowedToEatCount)
		wg.Add(1)
		go AllowToEat(h.allowedPhilosopherToEat2, &wg)
	} else {
		fmt.Printf("max eat allowed count (%d) for philosopher '%s' reached\n", h.allowedPhilosopherToEat2.maxAllowedToEatCount, h.allowedPhilosopherToEat2.name)
	}

	fmt.Println("Waiting until philosophers finish eating...")
	wg.Wait()

	if h.allowedPhilosopherToEat1.canEatAnotherTime() == false && h.allowedPhilosopherToEat2.canEatAnotherTime() == false {
		return false
	}

	return true
}

func (h *host) startDining(philosophers []*philosopher, wg *sync.WaitGroup) {

	fmt.Printf("[host]: start dining\n")

	okFailedCount := 0

	for {
		ok := h.SelectNextPhilosophers()

		if ok == false {
			okFailedCount += 1
		}

		if okFailedCount >= philosopherAndChopstickCount {
			fmt.Printf("selecting next philosphers failed %d times: %s, stopping... looks like all are done.\n", okFailedCount)

			fmt.Printf("Philosphers: \n")

			for _, philosopher := range philosophers {
				fmt.Printf("philosopher %v \n", *philosopher)
			}

			break
		}
	}

	fmt.Printf("[host]: finish dining\n")
	wg.Done()
}

func main() {

	philosophers := make([]*philosopher, 0, philosopherAndChopstickCount)
	chopsticks := make([]*chopstick, 0, philosopherAndChopstickCount)

	for i := 0; i < philosopherAndChopstickCount; i++ {
		chopsticks = append(chopsticks, &chopstick{
			inUse:             sync.Mutex{},
			name:              fmt.Sprintf("Chopstick %d", i+1),
			activePhilosopher: nil,
			lPhilosopher:      nil,
			rPhilosopher:      nil,
		})
	}

	for i := 0; i < philosopherAndChopstickCount; i++ {
		currentPhilosopher := &philosopher{
			name:                  fmt.Sprintf("Philosopher %d", i+1),
			lChopStick:            chopsticks[i],
			rChopStick:            chopsticks[modulo(i+1, philosopherAndChopstickCount)],
			allowedToEat:          false,
			pastAllowedToEatCount: 0,
			maxAllowedToEatCount:  3,
			eating:                sync.Mutex{},
		}
		philosophers = append(philosophers, currentPhilosopher)

		chopsticks[i].rPhilosopher = currentPhilosopher
	}

	for i := 0; i < philosopherAndChopstickCount; i++ {
		chopsticks[i].rPhilosopher = philosophers[i]
		chopsticks[i].lPhilosopher = philosophers[modulo(i-1, philosopherAndChopstickCount)]
	}

	host := host{
		allowedPhilosopherToEat1: nil,
		allowedPhilosopherToEat2: nil,
		philosophers:             philosophers,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go host.startDining(philosophers, &wg)
	wg.Wait()
}

func modulo(a, b int) int {
	result := a % b
	if result < 0 {
		result += b
	}
	return result
}
