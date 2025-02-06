type Parser<T> = (state: string) => [T | null, string]

type Regex = (re: RegExp) => Parser<string>
const regex: Regex = re => state => {
  const fst = state[0]

  if (re.test(fst)) {
    return [fst, state.slice(1)]
  }

  return [null, state]
}

type OneOrMore = <T>(parser: Parser<T>) => Parser<T[]>
const oneOrMore: OneOrMore = parser => state => {
  const results = []
  let curResult
  let curState = state

  for (; ;) {
    [curResult, curState] = parser(curState)

    if (curResult === null) {
      break
    }

    results.push(curResult)
  }

  if (results.length === 0) {
    return [null, state]
  }

  return [results, curState]
}

type Map = <T, R>(parser: Parser<T>, map: (v: T) => R) => Parser<R>
const map: Map = (parser, map) => state => {
  const [result, newState] = parser(state)

  if (result === null) {
    return [null, state]
  }

  return [map(result), newState]
}

type Seq = <R>(cb: (emit: Emit) => R) => Parser<R>
type Emit = <T>(parser: Parser<T>) => T | null

const seq: Seq = cb => state => {
  let curState = state
  let failed = false

  const result = cb((parser) => {
    if (failed) {
      return null
    }

    const [curResult, newState] = parser(curState)

    if (curResult === null) {
      failed = true
    }

    curState = newState
    return curResult
  })

  if (result === null) {
    return [null, state]
  }

  return [result, curState]
}

type Or = <T>(left: Parser<T>, right: Parser<T>) => Parser<T>
const or: Or = (left, right) => state => {
  const [result, newState] = left(state)

  if (result === null) {
    return right(state)
  }

  return [result, newState]
}


const openBr = regex(/\(/)
const closeBr = regex(/\)/)
const numeric = regex(/[0-9]/)
const integer = map(oneOrMore(numeric), v => Number(v.join("")))
const space = regex(/\s/)
const plus = regex(/\+/)

const sum: Parser<number | null> = seq(emit => {
  emit(openBr)
  emit(plus)
  emit(space)
  const left = emit(or(sum, integer))
  emit(space)
  const right = emit(or(sum, integer))
  emit(closeBr)

  if (left === null || right === null) {
    return null
  }

  return left + right
})

const input = "(+ (+ 1 (+ 1 1)) 1)"

console.log(sum(input))
