import {GumballMachineState, State} from "./GumballMachineState";

export class SoldState implements State {
    private readonly gumballMachine: GumballMachineState

    constructor(gumballMachine: GumballMachineState) {
        this.gumballMachine = gumballMachine;
    }

    getStatus() {
        return "Sold"
    }

    // 동전 투입
    insertQuarter() {
        console.log("알맹이를 내보내고 있습니다.")
    }

    // 동전 반환
    ejectQuarter() {
        console.log("이미 알맹이를 뽑으셨습니다.")
    }

    // 손잡이 돌리기
    turnCrank() {
        console.log("손잡이는 한 번만 돌려주세요.")
    }

    // 알맹이 내보내기
    dispense() {
        this.gumballMachine.releaseBall()

        if (this.gumballMachine.getCount() > 0) {
            this.gumballMachine.setState(this.gumballMachine.noQuarterState)
        } else {
            console.log("더 이상 알맹이가 없습니다.")
            this.gumballMachine.setState(this.gumballMachine.soldOutState)
        }
    }

    refill() {}
}