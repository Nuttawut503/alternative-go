package creational_patterns

import "fmt"

/* The intent of the Builder design pattern is
to separate the construction of a complex object from its representation.
So, the same construction process can create different representations.
*/

type Weapon interface {
	Describe() string
}

type Sword struct {
	name, element, owner string
	physicalATK          int
}

type Wand struct {
	name, element, owner string
	magicalATK           int
}

func (sword Sword) Describe() string {
	return fmt.Sprintf("%s's %s Sword, %s atk: %d ", sword.owner, sword.name, sword.element, sword.physicalATK)
}

func (wand Wand) Describe() string {
	return fmt.Sprintf("%s's %s Wand, %s atk: %d ", wand.owner, wand.name, wand.element, wand.magicalATK)
}

type WeaponBuilder interface {
	SetWeaponName(string) WeaponBuilder
	SetATK(int) WeaponBuilder
	SetElement(string) WeaponBuilder
	SetOwner(string) WeaponBuilder
	GetWeapon() Weapon
}

type SwordBuilder struct {
	name, element, owner string
	atk                  int
}

func (builder SwordBuilder) SetWeaponName(weaponName string) WeaponBuilder {
	builder.name = weaponName
	return builder
}

func (builder SwordBuilder) SetATK(atk int) WeaponBuilder {
	builder.atk = atk
	return builder
}

func (builder SwordBuilder) SetElement(element string) WeaponBuilder {
	builder.element = element
	return builder
}

func (builder SwordBuilder) SetOwner(ownerName string) WeaponBuilder {
	builder.owner = ownerName
	return builder
}

func (builder SwordBuilder) GetWeapon() Weapon {
	return Sword{
		name:        builder.name,
		physicalATK: builder.atk,
		element:     builder.element,
		owner:       builder.owner,
	}
}

type WandBuilder struct {
	name, element, owner string
	atk                  int
}

func (builder WandBuilder) SetWeaponName(weaponName string) WeaponBuilder {
	builder.name = weaponName
	return builder
}

func (builder WandBuilder) SetATK(atk int) WeaponBuilder {
	builder.atk = atk
	return builder
}

func (builder WandBuilder) SetElement(element string) WeaponBuilder {
	builder.element = element
	return builder
}

func (builder WandBuilder) SetOwner(ownerName string) WeaponBuilder {
	builder.owner = ownerName
	return builder
}

func (builder WandBuilder) GetWeapon() Weapon {
	return Wand{
		name:       builder.name,
		magicalATK: builder.atk,
		element:    builder.element,
		owner:      builder.owner,
	}
}

func BuilderExample() {
	wandBuilder := WandBuilder{}
	wand := wandBuilder.
		SetWeaponName("Lightweaver Cras Virge").
		SetATK(2953).
		SetElement("Light").
		SetOwner("me").
		GetWeapon()
	wand.Describe()
}
