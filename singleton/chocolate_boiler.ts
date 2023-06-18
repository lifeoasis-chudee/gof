
class ChocolateBoiler {
    private empty: boolean;
    private boiled: boolean;

    constructor() {
        this.empty = true;
        this.boiled = false;
    }

    isEmpty() {
        return this.empty;
    }
    isBoiled() {
        return this.boiled;
    }

    fill() {
        if (this.isEmpty()) {
            this.empty = false;
            this.boiled = false;
            // 보일러에 우유와 초콜릿을 혼합한 재료를 넣음
        }
    }

    drain() {
        if (!this.isEmpty() && this.isBoiled()) {
            // 끌인 재료를 다음 단계로 넘김
            this.empty = true;
        }
    }

    boil() {
        if (!this.isEmpty() && !this.isBoiled()) {
            // 재료를 끓임
            this.boiled = true;
        }
    }
}


class ChocolateBoilerSingleton {
    private empty: boolean;
    private boiled: boolean;

    private static instance = null

    private constructor() {
        this.empty = true;
        this.boiled = false;
    }

    static getInstance() {
        if (this.instance == null) {
            this.instance = new ChocolateBoilerSingleton()
        }

        return this.instance
    }

    isEmpty() {
        return this.empty;
    }
    isBoiled() {
        return this.boiled;
    }

    fill() {
        if (this.isEmpty()) {
            this.empty = false;
            this.boiled = false;
            // 보일러에 우유와 초콜릿을 혼합한 재료를 넣음
        }
    }

    drain() {
        if (!this.isEmpty() && this.isBoiled()) {
            // 끌인 재료를 다음 단계로 넘김
            this.empty = true;
        }
    }

    boil() {
        if (!this.isEmpty() && !this.isBoiled()) {
            // 재료를 끓임
            this.boiled = true;
        }
    }
}

var b = new ChocolateBoiler()
b.boil()
var a = new ChocolateBoilerSingleton()
a.
