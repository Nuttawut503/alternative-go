package behavioral_patterns

/* Encapsulate a request as an object,
thereby letting you parameterize clients with different requests,
queue or log requests, and support undoable operations.
*/

type (
	Command interface {
		execute()
		redo()
	}

	CertainNumber struct {
		value        int
		undos, redos []Command
	}

	IncreaseCommand struct {
		state *CertainNumber
		add   int
	}
	MultiplyCommand struct {
		state *CertainNumber
		mul   int
	}
)

func (cn *CertainNumber) GetValue() int {
	return cn.value
}

func (cn *CertainNumber) Execute(cmd Command) {
	cmd.execute()
	cn.undos = append(cn.undos, cmd)
}

func (cn *CertainNumber) Undo() {
	if len(cn.undos) == 0 {
		return
	}
	cmd := cn.undos[len(cn.undos)-1]
	cn.undos = cn.undos[:len(cn.undos)-1]
	cmd.redo()
	cn.redos = append(cn.redos, cmd)
}

func (cn *CertainNumber) Redo() {
	if len(cn.redos) == 0 {
		return
	}
	cmd := cn.redos[len(cn.redos)-1]
	cn.redos = cn.redos[:len(cn.redos)-1]
	cmd.execute()
	cn.undos = append(cn.undos, cmd)
}

func (cmd IncreaseCommand) execute() {
	cmd.state.value += cmd.add
}

func (cmd IncreaseCommand) redo() {
	cmd.state.value -= cmd.add
}

func (cmd MultiplyCommand) execute() {
	cmd.state.value *= cmd.mul
}

func (cmd MultiplyCommand) redo() {
	cmd.state.value /= cmd.mul
}

func CommandExample() {
	certainNumber := CertainNumber{value: 0}
	certainNumber.Execute(IncreaseCommand{&certainNumber, 10})
	certainNumber.Execute(MultiplyCommand{&certainNumber, 2})
	certainNumber.Execute(IncreaseCommand{&certainNumber, 5})
	certainNumber.Undo()
	certainNumber.Undo()
	certainNumber.Redo()
	println(certainNumber.GetValue()) // 20
}
