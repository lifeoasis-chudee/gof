import {HasQuarterState} from "./HasQuarterState";
import {SoldState} from "./SoldState";
import {SoldOutState} from "./SoldOutState";
import {NoQuarterState} from "./NoQuarterState";
import {WinnerState} from "./WinnerState";

export interface State {
    insertQuarter(): void;
    ejectQuarter(): void;
    turnCrank(): void;
    dispense(): void;
    refill(): void;

    getStatus(): string;
}

export class GumballMachineState {
    // 귀찮아서 퍼블릭으로 함
    readonly soldOutState: State;
    readonly noQuarterState: State;
    readonly hasQuarterState: State;
    readonly soldState: State;
    readonly winnerState: State;

    private state: State;
    private count = 0;

    constructor(numberGumballs: number) {
        this.soldOutState = new SoldOutState(this);
        this.noQuarterState = new NoQuarterState(this);
        this.hasQuarterState = new HasQuarterState(this);
        this.soldState = new SoldState(this);
        this.winnerState = new WinnerState(this);

        this.count = numberGumballs;
        if (numberGumballs > 0) {
            this.state = this.noQuarterState;
        } else {
            this.state = this.soldOutState;
        }
    }

    getCount() {
        return this.count
    }

    setState(state: State) {
        this.state = state
    }

    toString() {
        const isSoldOut = this.state == this.soldOutState;
        const log = `
>주식회사 왕뽑기
 자바로 돌아가는 최신형 뽑기 기계
 남은 개수: ${this.count}개
 상태: ${this.state.getStatus()}
 `

        console.log(log)
    }

    // 동전 투입
    insertQuarter() {
        this.state.insertQuarter()
    }

    // 동전 반환
    ejectQuarter() {
        this.state.ejectQuarter()
    }

    // 손잡이 돌리기
    turnCrank() {
        this.state.turnCrank()
        this.state.dispense()
    }

    // 알맹이 내보내기
    releaseBall() {
        console.log("알맹이를 내보내고 있습니다.")
        if (this.count > 0) {
            this.count -= 1
        }
    }

    refill(count: number) {
        this.count += count;
        console.log("리필됨: ", count)
        this.state.refill();
    }
}