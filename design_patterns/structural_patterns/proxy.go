package structural_patterns

import "fmt"

/* Provide a surrogate or placeholder for another object to control access to it.
 */

type WeaponManager interface {
	SetDamage(int)
	GetDamage() int
}

type Weapon struct {
	name string
	WeaponManager
}

type SwordManager struct {
	power int
}

func (swordManager *SwordManager) SetDamage(power int) {
	swordManager.power = power
}

func (swordManager *SwordManager) GetDamage() int {
	return swordManager.power
}

type ProxySwordManager struct {
	effect       string
	swordManager SwordManager
}

func (proxy *ProxySwordManager) SetDamage(power int) {
	if power < 100 {
		proxy.swordManager.SetDamage(power + 100)
	} else {
		proxy.swordManager.SetDamage(power)
	}
}

func (proxy *ProxySwordManager) GetDamage() int {
	return proxy.swordManager.GetDamage()
}

func ProxyExample() {
	weapon := Weapon{
		"Silver Sword",
		&ProxySwordManager{
			effect: "AOE damage",
		},
	}
	weapon.WeaponManager.SetDamage(70)
	fmt.Println(weapon.WeaponManager.GetDamage())
}
