package structural_patterns

import "fmt"

/* Attach additional responsibilities to an object dynamically.
Decorators provide a flexible alternative to subclassing for extending functionality
*/

type Character struct {
	cloth, shoe string
}

func (character Character) DescribeEquipments() string {
	return fmt.Sprintf("I'm wearing %s, %s", character.cloth, character.shoe)
}

type EquipmentDecorator interface {
	SetMinorDecorator(EquipmentDecorator) EquipmentDecorator
	DescribeEquipments() string
}

type GlassDecorator struct {
	decorator EquipmentDecorator
}

func (glass GlassDecorator) SetMinorDecorator(decorator EquipmentDecorator) EquipmentDecorator {
	glass.decorator = decorator
	return glass
}

func (glass GlassDecorator) DescribeEquipments() string {
	if glass.decorator == nil {
		return ", a glass"
	}
	return fmt.Sprintf("%s, %s", glass.decorator.DescribeEquipments(), "a glass")
}

type HatDecorator struct {
	decorator EquipmentDecorator
}

func (hat HatDecorator) SetMinorDecorator(decorator EquipmentDecorator) EquipmentDecorator {
	hat.decorator = decorator
	return hat
}

func (hat HatDecorator) DescribeEquipments() string {
	if hat.decorator == nil {
		return ", a hat"
	}
	return fmt.Sprintf("%s, %s", hat.decorator.DescribeEquipments(), "a hat")
}

type WingDecorator struct {
	decorator EquipmentDecorator
}

func (wing WingDecorator) SetMinorDecorator(decorator EquipmentDecorator) EquipmentDecorator {
	wing.decorator = decorator
	return wing
}

func (wing WingDecorator) DescribeEquipments() string {
	if wing.decorator == nil {
		return ", a wing"
	}
	return fmt.Sprintf("%s, %s", wing.decorator.DescribeEquipments(), "a wing")
}

type HumanWithEquipments struct {
	character          Character
	equipmentDecorator EquipmentDecorator
}

func (human *HumanWithEquipments) DescribeEquipments() string {
	if human.equipmentDecorator == nil {
		return human.character.DescribeEquipments()
	}
	return human.character.DescribeEquipments() + human.equipmentDecorator.DescribeEquipments()
}

func (human *HumanWithEquipments) AddDecorator(decorator EquipmentDecorator) {
	human.equipmentDecorator = decorator.SetMinorDecorator(human.equipmentDecorator)
}

func DecoratorExample() {
	human := HumanWithEquipments{
		character: Character{"a suit", "slippers"},
	}
	human.AddDecorator(HatDecorator{})
	human.AddDecorator(HatDecorator{})
	human.AddDecorator(WingDecorator{})
	human.AddDecorator(GlassDecorator{})
	fmt.Println(human.DescribeEquipments())
}
