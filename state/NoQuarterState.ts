import {GumballMachineState, State} from "./GumballMachineState";

export class NoQuarterState implements State {
    private readonly gumballMachine: GumballMachineState

    constructor(gumballMachine: GumballMachineState) {
        this.gumballMachine = gumballMachine;
    }

    getStatus() {
        return "NoQuarter"
    }

    // 동전 투입
    insertQuarter() {
        console.log("동전을 넣으셨습니다.")
        this.gumballMachine.setState(this.gumballMachine.hasQuarterState)
    }

    // 동전 반환
    ejectQuarter() {
        console.log("동전을 넣어주세요.")
    }

    // 손잡이 돌리기
    turnCrank() {
        console.log("동전을 넣어주세요.")
    }

    // 알맹이 내보내기
    dispense() {
        console.log("동전을 넣어주세요.")
    }

    refill() {}
}