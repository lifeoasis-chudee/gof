
enum GumballMachineState {
    SOLD_OUT,
    NO_QUARTER,
    HAS_QUARTER,
    SOLD
}

export class GumballMachine {
    private state = GumballMachineState.SOLD_OUT;
    private count = 0;

    constructor(count: number) {
        this.count = count;
        if (count > 0) {
            this.state = GumballMachineState.NO_QUARTER
        }
    }

    toString() {
        const isSoldOut = this.state === GumballMachineState.SOLD_OUT;
        const log = `
>주식회사 왕뽑기
 자바로 돌아가는 최신형 뽑기 기계
 남은 개수: ${this.count}개
 ${isSoldOut ? "매진" : "동전 투입 대기중"}
 `
        console.log(log)
    }

    // 동전 투입
    insertQuarter() {
        switch (this.state) {
            case GumballMachineState.HAS_QUARTER:
                console.log("동전은 한 개만 넣어주세요.")
                break;

            case GumballMachineState.NO_QUARTER:
                this.state = GumballMachineState.HAS_QUARTER;
                console.log("동전을 넣으셨습니다.")
                break;

            case GumballMachineState.SOLD_OUT:
                console.log("매진되었습니다. 다음 기회에 이용해 주세요.")
                break;

            case GumballMachineState.SOLD:
                console.log("알맹이를 내보내고 있습니다.")
                break;
        }
    }

    // 동전 반환
    ejectQuarter() {
        switch (this.state) {
            case GumballMachineState.HAS_QUARTER:
                console.log("동전이 반환됩니다.")
                this.state = GumballMachineState.NO_QUARTER;
                break;

            case GumballMachineState.NO_QUARTER:
                console.log("동전을 넣어주세요.")
                break;

            case GumballMachineState.SOLD_OUT:
                console.log("동전을 넣지 않으셨습니다. 동전이 반환되지 않습니다.")
                break;

            case GumballMachineState.SOLD:
                console.log("이미 알맹이를 뽑으셨습니다.")
                break;
        }
    }

    // 손잡이 돌리기
    turnCrank() {
        switch (this.state) {
            case GumballMachineState.HAS_QUARTER:
                console.log("손잡이를 돌리셨습니다.")
                this.state = GumballMachineState.SOLD;
                this.dispense()
                break;

            case GumballMachineState.NO_QUARTER:
                console.log("동전을 넣어주세요.")
                break;

            case GumballMachineState.SOLD_OUT:
                console.log("매진되었습니다.")
                break;

            case GumballMachineState.SOLD:
                console.log("손잡이는 한 번만 돌려주세요.")
                break;
        }
    }

    // 알맹이 내보내기
    dispense() {
        switch (this.state) {
            case GumballMachineState.HAS_QUARTER:
                console.log("알맹이를 내보낼 수 없습니다.")
                break;

            case GumballMachineState.NO_QUARTER:
                console.log("동전을 넣어주세요.")
                break;

            case GumballMachineState.SOLD_OUT:
                console.log("매진되었습니다.")
                break;

            case GumballMachineState.SOLD:
                console.log("알맹이를 내보내고 있습니다.")
                this.count -= 1;
                if (this.count === 0) {
                    console.log("더 이상 알맹이가 없습니다.")
                    this.state = GumballMachineState.SOLD_OUT;
                } else {
                    this.state = GumballMachineState.NO_QUARTER;
                }
                break;
        }
    }
}