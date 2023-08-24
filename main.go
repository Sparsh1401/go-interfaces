package main

import "fmt"

type Attacker interface {
	Attack() error
}

type Defender interface {
	Defend() error
}

type StrongAttacker struct{}

func DoAttack(a Attacker) error {
	return a.Attack()
}

func DoDefending(d Defender) error {
	return d.Defend()
}

func (S StrongAttacker) Attack() error {
	fmt.Println("Strong Attack i am low")
	return nil
}

func (S StrongAttacker) Defend() error {
	fmt.Println("I can defend as well")
	return nil
}

func main() {
	DoAttack(StrongAttacker{})
	DoDefending(StrongAttacker{})
}
