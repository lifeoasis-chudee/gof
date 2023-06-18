import {GumballMachineState} from "./GumballMachineState";

test("GumballMachineStateTestDrive", () => {
    const gumballMachine = new GumballMachineState(4);
    gumballMachine.toString();

    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.toString();

    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.toString();

    gumballMachine.refill(10);
    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.insertQuarter()
    gumballMachine.turnCrank()
    gumballMachine.toString();
});