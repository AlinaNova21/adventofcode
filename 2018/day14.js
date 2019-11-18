const input = '37'

const recipes = input.split('').map(v => parseInt(v))

const NUM_ELVES = 2

const elves = [{ r: 0, rv: 3, p: (v) => `(${v})` }, { r: 1, rv: 7, p: (v) => `[${v}]` }]

function loop() {
  (elves[0].rv + elves[1].rv).toString().split('').forEach(v => recipes.push(parseInt(v)))
  for(const elf of elves) {
    const recipe = recipes[elf.r]
    elf.r = (elf.r + recipe + 1) % recipes.length
    elf.rv = recipes[elf.r]
  }
}

function render() {
  const inds = elves.map(e => e.r)
  console.log(recipes.map((r,i) => {
    const ind = inds.indexOf(i)
    if (ind !== -1) return elves[ind].p(r)
    return r
  }).join(' '))
}

// const target = '01245'
const target = '919901'

render()
let i=0

const len = (parseInt(target) + 10)

while(recipes.length < len) { //(let i=0;i<10;i++) {
  loop()
  //render()
  if(recipes.length - len % 100 === 0) {
    render()
  }
}

console.log('p1',recipes.slice(target, len).join(''))
const neg = target.length + 3
while(true) {
  loop()
  const str = recipes.slice(-neg).join('')
  const ind = str.indexOf(target)
  if(ind > 0) {
    console.log("\nIND " + (ind + recipes.length - neg))
    console.log(ind)
    console.log(recipes.length)
    console.log(neg)
    break
  }
  if(recipes.length % 1000000 === 0) {
    console.log(recipes.length)
  }
}