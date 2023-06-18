import {GumballMachineState, State} from "./GumballMachineState";

export class HasQuarterState implements State {
    private readonly gumballMachine: GumballMachineState

    constructor(gumballMachine: GumballMachineState) {
        this.gumballMachine = gumballMachine;
    }

    getStatus() {
        return "HasQuarter"
    }

    // 동전 투입
    insertQuarter() {
        console.log("동전은 한 개만 넣어주세요.")
    }

    // 동전 반환
    ejectQuarter() {
        console.log("동전이 반환됩니다.")
        this.gumballMachine.setState(this.gumballMachine.noQuarterState)
    }

    // 손잡이 돌리기
    turnCrank() {
        console.log("손잡이를 돌리셨습니다.")

        // 1 ~ 100
        const isWinner = Math.floor(Math.random() * 100) + 1 >= 90;

        // 당첨되었고, 남은 알맹이 갯수가 2개 이상이어야 함.
        if (isWinner && this.gumballMachine.getCount() > 1) {
            this.gumballMachine.setState(this.gumballMachine.winnerState)
        } else {
            this.gumballMachine.setState(this.gumballMachine.soldState)
        }
    }

    // 알맹이 내보내기
    dispense() {
        console.log("알맹이를 내보낼 수 없습니다.")
    }

    refill() {}
}