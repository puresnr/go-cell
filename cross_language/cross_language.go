package cross_language

// js 等语言中, 即使是 64 位无符号数, int 的最大值也不是 2 ^ 63 - 1, 所以定义一个跨语言交互时的最大和最小整数
const MaxInt = 9007199254740991 // math.Pow(2, 53) - 1
const MinInt = -9007199254740991
