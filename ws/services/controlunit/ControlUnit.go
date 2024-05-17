package controlunit

import (
	"architecture/ws/services/alu"
	"architecture/ws/services/memory"
	"errors"
	"fmt"
)

type ControlUnit struct {
	Registers *memory.Registers
	Memory    *memory.Memory
	ALU       *alu.ALU
}
type Memory struct {
	Data map[int]int
}

func NewMemory() *Memory {
	return &Memory{
		Data: make(map[int]int),
	}
}

func NewControlUnit(registers *memory.Registers, memory *memory.Memory, alu *alu.ALU) *ControlUnit {
	return &ControlUnit{
		Registers: registers,
		Memory:    memory,
		ALU:       alu,
	}
}

func (cu *ControlUnit) Fetch() error {
	instruction, err := cu.Memory.Read(cu.Registers.PC)
	if err != nil {
		return err
	}
	cu.Registers.IR = instruction
	cu.Registers.PC++
	return nil
}

func (cu *ControlUnit) Decode() (string, int, error) {
	opcode := (cu.Registers.IR & 0xF000) >> 12
	address := cu.Registers.IR & 0x0FFF
	return fmt.Sprintf("%X", opcode), address, nil
}

func (cu *ControlUnit) Execute(opcode string, address int) error {
	switch opcode {
	case "1":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.DR = value
		cu.Registers.AC = cu.ALU.Add(cu.Registers.AC, cu.Registers.DR)
	case "2":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.DR = value
		cu.Registers.AC = cu.ALU.Subtract(cu.Registers.AC, cu.Registers.DR)
	case "3":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.AC = value
	case "4":
		return cu.Memory.Write(address, cu.Registers.AC)
	case "5":
		cu.Registers.PC = address
	case "6":
		err := cu.Memory.Write(address, cu.Registers.PC)
		if err != nil {
			return err
		}
		cu.Registers.PC = address + 1
	case "7":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		value++
		err = cu.Memory.Write(address, value)
		if err != nil {
			return err
		}
		if value == 0 {
			cu.Registers.PC++
		}
	case "8":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.DR = value
		cu.Registers.AC = cu.ALU.And(cu.Registers.AC, cu.Registers.DR)
	case "9":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.DR = value
		cu.Registers.AC = cu.ALU.Or(cu.Registers.AC, cu.Registers.DR)
	case "A":
		value, err := cu.Memory.Read(address)
		if err != nil {
			return err
		}
		cu.Registers.DR = value
		cu.Registers.AC = cu.ALU.Xor(cu.Registers.AC, cu.Registers.DR)
	case "B":
		cu.Registers.AC = cu.ALU.Not(cu.Registers.AC)
	default:
		return errors.New("unsupported opcode: " + opcode)
	}
	return nil
}

func (cu *ControlUnit) RunCycle() error {
	cu.Fetch()
	opcode, address, err := cu.Decode()
	if err != nil {
		return err
	}
	return cu.Execute(opcode, address)
}

// import (
// 	"architecture/ws/services/bus"
// 	"architecture/ws/services/memory"
// 	ws "architecture/ws/services/memory"
// )

// type ControlUnit struct {
// 	Register    *ws.Registers
// 	AddressBus  *bus.AddressBus
// 	DataBus     *bus.DataBus
// 	ControlBus  *bus.ControlBus
// 	MainMemory  *memory.Memory
// 	CacheMemory *memory.Cache
// }

// func NewControlUnit(register *ws.Registers, adrressBus *bus.AddressBus, dataBus *bus.DataBus, controlBus *bus.ControlBus, mainMemory *memory.Memory, cacheMemory *memory.Cache) *ControlUnit {
// 	return &ControlUnit{
// 		Register:    register,
// 		AddressBus:  adrressBus,
// 		DataBus:     dataBus,
// 		ControlBus:  controlBus,
// 		MainMemory:  mainMemory,
// 		CacheMemory: cacheMemory,
// 	}
// }

// func (cu *ControlUnit) FetchIntstruction() {
// 	cu.AddressBus.Write(cu.Register.PC)
// 	instruction := cu.MainMemory.Read(byte(cu.Register.PC))
// 	cu.Register.IR = instruction
// 	cu.Register.PC++
// }

// func (cu *ControlUnit) DecodeInstruction() {
// 	opcode := cu.Register.IR >> 4
// 	operand := cu.Register.IR & 0x0f

// 	switch opcode {
// 	case 0x0:
// 		cu.Register.MAR = int(operand)
// 		cu.ControlBus.WriteSignal(0, 1)
// 	case 0x1:
// 		cu.Register.MAR = int(operand)
// 		cu.ControlBus.WriteSignal(1, 1)
// 	case 0x2:
// 		cu.Register.ACC += cu.MainMemory.Read(byte(operand))
// 	case 0x3:
// 		cu.Register.ACC -= cu.MainMemory.Read(byte(operand))
// 		///// بعدا پیاده سازی کن بقیه دستورات رو

// 	}
// }

// func (cu *ControlUnit) Execute() {
// 	switch cu.ControlBus.ReadSignal(0) {
// 	case 1:
// 		// خواندن حافظه و ذخیره در MBR
// 		value := cu.CacheMemory.Read(cu.Register.MAR)
// 		cu.Register.MBR = value
// 		// cu.DataBus.Read()
// 		// cu.Register.MBR = cu.DataBus.Read()
// 	case 2:
// 		// نوشتن در حافظه از MBR
// 		cu.DataBus.Write(cu.Register.MBR)
// 	}
// 	// پاک کردن سیگنال های کنترلی
// 	cu.ControlBus.WriteSignal(0, 0)
// 	cu.ControlBus.WriteSignal(1, 0)
// }
