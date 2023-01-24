package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; true; i++ {

		// print start
		fmt.Printf("\nStart: %d\n", i)
		tick := 20
		trick := 23
		track := 25
		fmt.Printf("tick: %d, trick: %d, track: %d\n", tick, trick, track)

		var tickC, trickC, trackC bool
		for i := 0; true; i++ {
			// check if one has no ticket
			one, name0 := oneWithNoTicket(tick, trick, track)
			two, name1 := twoWithNoTicket(tick, trick, track)
			none, name := allTicketsGone(tick, trick, track)
			if one || two || none {
				if none {
					fmt.Println("all tickets gone", name)
					break
				}
				if two {
					fmt.Println("two with no ticket: ", name1)
					break
				}
				if one {
					fmt.Println("one with no ticket: ", name0)

					break
				}

			}
			tick, trick, track = ride(tick, trick, track)
			tickC, trickC, trackC = false, false, false
			if tick%2 == 0 || trick%2 == 0 || track%2 == 0 {
				if tick%2 == 0 {
					tickC = true
				}
				if trick%2 == 0 {
					trickC = true
				}
				if track%2 == 0 {
					trackC = true
				}
				// if all false skip ticket change
				if tickC == false && trickC == false && trackC == false {
					continue
				} else {
					// check how many are allowed to change
					// all three
					if tickC == true && trickC == true && trackC == true {
						// ranomly choose one donor
						// 0 = tick, 1 = trick, 2 = track
						r := rand.Intn(3)
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r1 := rand.Intn(3)
						// the receiver looses the half of the ticket
						// the donor gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 && r1 == 0 {
							fmt.Printf("No change\n")
						}
						if r == 0 && r1 == 1 {
							tick = tick / 2
							trick = trick + tick
							fmt.Printf("tick -> trick\n")
						}
						if r == 0 && r1 == 2 {
							tick = tick / 2
							track = track + tick
							fmt.Printf("tick -> track\n")
						}
						if r == 1 && r1 == 0 {
							trick = trick / 2
							tick = tick + trick
							fmt.Printf("trick -> tick\n")
						}
						if r == 1 && r1 == 1 {
							fmt.Printf("No change\n")
						}
						if r == 1 && r1 == 2 {
							trick = trick / 2
							track = track + trick
							fmt.Printf("trick -> track\n")
						}
						if r == 2 && r1 == 0 {
							track = track / 2
							tick = tick + track
							fmt.Printf("track -> tick\n")
						}
						if r == 2 && r1 == 1 {
							track = track / 2
							trick = trick + track
							fmt.Printf("track -> trick\n")
						}
						if r == 2 && r1 == 2 {
							fmt.Printf("No change\n")
						}
					}
					// just tick and trick are donor, all three are receiver
					if tickC == true && trickC == true && trackC == false {
						// ranomly choose one donor
						// 0 = tick, 1 = trick, 2 = track
						r := rand.Intn(2)
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r1 := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 && r1 == 0 {
							fmt.Printf("No change\n")
						}
						if r == 0 && r1 == 1 {
							tick = tick / 2
							trick = trick + tick
							fmt.Printf("tick -> trick\n")
						}
						if r == 0 && r1 == 2 {
							tick = tick / 2
							track = track + tick
							fmt.Printf("tick -> track\n")
						}
						if r == 1 && r1 == 0 {
							trick = trick / 2
							tick = tick + trick
							fmt.Printf("trick -> tick\n")
						}
						if r == 1 && r1 == 1 {
							fmt.Printf("No change\n")
						}
						if r == 1 && r1 == 2 {
							trick = trick / 2
							track = track + trick
							fmt.Printf("trick -> track\n")
						}
					}
					// just tick and track are donor, all three are receiver
					if tickC == true && trickC == false && trackC == true {
						// ranomly choose one donor (0 = tick, 2 = track)
						r := rand.Intn(2)

						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r1 := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 && r1 == 0 {
							fmt.Printf("No change\n")
						}
						if r == 0 && r1 == 1 {
							tick = tick / 2
							trick = trick + tick
							fmt.Printf("tick -> trick\n")
						}
						if r == 0 && r1 == 2 {
							tick = tick / 2
							track = track + tick
							fmt.Printf("tick -> track\n")
						}
						if r == 1 && r1 == 0 {
							track = track / 2
							tick = tick + track
							fmt.Printf("track -> tick\n")
						}
						if r == 1 && r1 == 1 {
							fmt.Printf("No change\n")
						}
						if r == 1 && r1 == 2 {
							track = track / 2
							trick = trick + track
							fmt.Printf("track -> trick\n")
						}
					}
					// just trick and track are donor, all three are receiver
					if tickC == false && trickC == true && trackC == true {
						// ranomly choose one donor (1 = trick, 2 = track)
						r := rand.Intn(2)
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r1 := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 && r1 == 0 {
							trick = trick / 2
							tick = tick + trick
							fmt.Printf("trick -> tick\n")
						}
						if r == 0 && r1 == 1 {
							fmt.Printf("No change\n")
						}
						if r == 0 && r1 == 2 {
							trick = trick / 2
							track = track + trick
							fmt.Printf("trick -> track\n")
						}
						if r == 1 && r1 == 0 {
							track = track / 2
							tick = tick + track
							fmt.Printf("track -> tick\n")
						}
						if r == 1 && r1 == 1 {
							track = track / 2
							trick = trick + track
							fmt.Printf("track -> trick\n")
						}
						if r == 1 && r1 == 2 {
							fmt.Printf("No change\n")
						}
					}
					// just tick is donor, all are receiver
					if tickC == true && trickC == false && trackC == false {
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 {
							fmt.Printf("No change\n")
						}
						if r == 1 {
							tick = tick / 2
							trick = trick + tick
							fmt.Printf("tick -> trick\n")
						}
						if r == 2 {
							tick = tick / 2
							track = track + tick
							fmt.Printf("tick -> track\n")
						}
					}
					// just trick is donor, all are receiver
					if tickC == false && trickC == true && trackC == false {
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 {
							trick = trick / 2
							tick = tick + trick
							fmt.Printf("trick -> tick\n")
						}
						if r == 1 {
							fmt.Printf("No change\n")
						}
						if r == 2 {
							trick = trick / 2
							track = track + trick
							fmt.Printf("trick -> track\n")
						}
					}
					// just track is donor, all are receiver
					if tickC == false && trickC == false && trackC == true {
						// randomly choose one receiver
						// 0 = tick, 1 = trick, 2 = track
						r := rand.Intn(3)
						// the donor looses the half of the ticket
						// the receiver gets the half of the ticket
						// if the receiver and the donor are the same nothing happens
						if r == 0 {
							track = track / 2
							tick = tick + track
							fmt.Printf("track -> tick\n")
						}
						if r == 1 {
							track = track / 2
							trick = trick + track
							fmt.Printf("track -> trick\n")
						}
						if r == 2 {
							fmt.Printf("No change\n")
						}
					}

				}
			}

		}
		one, _ := oneWithNoTicket(tick, trick, track)
		two, _ := twoWithNoTicket(tick, trick, track)
		none, _ := allTicketsGone(tick, trick, track)
		if one || two || none {
			println("End")
			if none {
				//fmt.Println("all tickets gone", name)
				break
			}
			if two {
				//fmt.Println("two with no ticket: ", name1)
				//break
			}
			if one {
				//fmt.Println("one with no ticket: ", name0)
				//break
			}

		}
	}
}

func ride(tick int, trick int, track int) (int, int, int) {
	// print the current state
	tick = tick - 1
	trick = trick - 1
	track = track - 1
	fmt.Printf("tick: %d, trick: %d, track: %d\n", tick, trick, track)
	return tick, trick, track
}

// oneWithNoTicket prüft, ob nach einer Fahrt genau einer kein Ticket mehr hat
func oneWithNoTicket(tick int, trick int, track int) (bool, string) {
	if tick == 0 && trick > 0 && track > 0 {
		return true, "tick"
	}

	if trick == 0 && tick > 0 && track > 0 {
		return true, "trick"
	}

	if track == 0 && tick > 0 && trick > 0 {
		return true, "track"
	}

	return false, "none"
}

// twoWithNoTicket prüft, ob nach einer Fahrt genau zwei kein Ticket mehr haben
func twoWithNoTicket(tick int, trick int, track int) (bool, string) {
	if tick == 0 && trick == 0 && track > 0 {
		return true, "tick, trick"
	}

	if trick == 0 && track == 0 && tick > 0 {
		return true, "trick, track"
	}

	if track == 0 && tick == 0 && trick > 0 {
		return true, "track, tick"
	}

	return false, "none"
}

// allTicketsGone prüft, ob alle Tickets abgegeben sind
func allTicketsGone(tick int, trick int, track int) (bool, string) {
	if tick == 0 && trick == 0 && track == 0 {
		return true, "tick, trick, track"
	}
	return false, "none"
}
