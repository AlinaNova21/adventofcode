const input = require('./input').slice(1).map(v => {
  const [pattern,willHave] = v.split(' => ')
  return [pattern,willHave]
}).reduce((ret,[pattern,willHave]) => {
  ret[pattern] = willHave
  return ret
},{})

const start = -10
const state = require('./input')[0].split(' ').slice(-1)[0].split('')

for(let i=start;i< 0;i++) {
  state.unshift('.')
}

const GENERATIONS =  50000000000//20
// console.log(input)
console.log(('   '+0).slice(-3), state.join(''))
for (let gen=1;gen<=GENERATIONS;gen++) {
  if(state.slice(-1)[0] !== '.') state.push('.','.','.','.','.')
  const prev = Array.from(state)
  const g = (i) => prev[i] || '.'
  const s = (i,v) => {
    while (i >= state.length) state.push('.')
    state[i] = v
  }
  for(let i=0;i<state.length;i++) {
    const p = [i-2,i-1,i,i+1,i+2].map(g).join('')
    // console.log(('   '+(i+start)).slice(-3),p,input[p] || '.')
    s(i, input[p] || '.')
  }
  console.log(('   '+gen).slice(-3), state.join(''))
  if (state.join('').slice(1) == prev.join('').slice(0,-1)) {
    console.log(`Repeat found at gen ${gen}`)
    const offset = GENERATIONS - gen
    const ret = state.reduce((cum, val, ind) => cum + (val === '#' ? ind + start + offset : 0), 0)
    console.log(ret)
    break;
  }
}