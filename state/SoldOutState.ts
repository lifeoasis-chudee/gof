import {GumballMachineState, State} from "./GumballMachineState";
import {SoldState} from "./SoldState";

export class SoldOutState implements State {
    private readonly gumballMachine: GumballMachineState

    constructor(gumballMachine: GumballMachineState) {
        this.gumballMachine = gumballMachine;
    }

    getStatus() {
        return "SoldOut"
    }

    // 동전 투입
    insertQuarter() {
        console.log("매진되었습니다. 다음 기회에 이용해 주세요.")
    }

    // 동전 반환
    ejectQuarter() {
        console.log("동전을 넣지 않으셨습니다. 동전이 반환되지 않습니다.")
    }

    // 손잡이 돌리기
    turnCrank() {
        console.log("매진되었습니다.")
    }

    // 알맹이 내보내기
    dispense() {
        console.log("알맹이를 내보낼 수 없습니다.")
    }

    refill() {
        this.gumballMachine.setState(this.gumballMachine.noQuarterState)
    }
}