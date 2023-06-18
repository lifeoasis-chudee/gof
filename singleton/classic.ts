
class Singleton {
    private static uniqueInstance = null;

    // 기타 인스턴스 변수

    // 생성자를 private로 선언했으므로 Singleton에서만 클래스의 인스턴스를 만들 수 있다.
    private constructor() {}

    public static getInstance(): Singleton  {
        if (this.uniqueInstance == null) {
            this.uniqueInstance = new Singleton()
        }

        return this.uniqueInstance
    }

    // 기타 메서드
}