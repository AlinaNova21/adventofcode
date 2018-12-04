const fs = require('fs')

const argv = process.argv.slice(2)
const input = fs.readFileSync(argv[0], 'utf8').split("\n")

let two = 0
let three = 0

for(const id of input) {
  const letters = {}
  id.split('').forEach(l => {
    letters[l] = (letters[l] || 0) + 1
  })
  const has2 = !!Object.values(letters).find(v => v == 2)
  const has3 = !!Object.values(letters).find(v => v == 3)
  if (has2) two++
  if (has3) three++
  console.log(id, has2, has3)
}
console.log(two * three)

input.sort()

let last = ''
for(const a of input) {
  for(const b of input) {
    let numdiff = 0
    let common = ''
    for(let i = 0;i < a.length;i++) {
      if(a[i] === b[i]) {
        common += a[i]
      } else {
        numdiff ++
      }      
    }
    if(numdiff === 1) console.log('common',common)
  }
}