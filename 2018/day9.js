const input = require('./input').map(v => {
  const [players,,,,,,lastPoints,,,,,targetHighScore] = v.split(' ')
  return [parseInt(players), parseInt(lastPoints)*100,parseInt(targetHighScore),v]
})


class CircularLinkedList {
  constructor(value) {
    this.node = { value }
    this.node.next = this.node
    this.node.prev = this.node
  }
  prev(cnt=1) {
    while(cnt--) {
      this.node = this.node.prev
    }
  }
  next(cnt=1) {
    while(cnt--) {
      this.node = this.node.next
    }
  }
  insert(value) {
    const newNode = { value }
    newNode.prev = this.node
    newNode.next = this.node.next
    this.node.next.prev = newNode
    this.node.next = newNode
    this.node = newNode
  }
  remove() {
    const value = this.node.value
    const nextNode = this.node.next
    this.node.next.prev = this.node.prev
    this.node.prev.next = this.node.next
    this.node.prev = null
    this.node.next = null
    this.node = nextNode
    return value
  }
  toArray() {
    const arr = []
    arr.push(this.node.value)
    let nextNode = this.node.next
    while(nextNode != this.node) {
      arr.push(nextNode.value)
      nextNode = nextNode.next
    }
    return arr
  }
}

for(const [players, lastPoints, targetHighScore, raw] of input) {
  // if (players !== 13) continue
  // if (players !== 9) continue
  const circle = new CircularLinkedList(0)
  const scores = Array(players).fill(0).map(v => 0)
  let curInd = 0
  let player = 0
  let highScore = 0
  let nextMarble = 1
  // renderLine('-', curInd, circle)
  while(nextMarble <= lastPoints) {
    if (nextMarble % 23 === 0) {
      circle.prev(7)
      const marble = circle.remove()
      scores[nextMarble % players] += nextMarble + marble
    } else {
      circle.next(1)
      circle.insert(nextMarble)
      // console.log(`placed ${nextMarble} at ${ind}`)
    }
    // renderLine(player + 1, curInd, circle)
    // player++
    // player %= players
    if(nextMarble % 10000 === 0)
      console.log(`${nextMarble}/${lastPoints} ${Math.floor((nextMarble/lastPoints) * 10000)/100}%`)
    nextMarble++
  }
  highScore = scores.reduce((a,b) => Math.max(a,b),0)
  const diff = targetHighScore - highScore
  console.log(`${raw}: high score is ${highScore}; ${targetHighScore} diff ${diff} ${diff % 23}`)
  // console.log(scores)
}

function nextLowest(circle,kept,max) {
  for(let i = 0; i <= max; i++) {
    if (!circle.includes(i) && !kept.includes(i)) return i
  }
  return -1
}
function renderLine(player, curInd, circle) {
  const out = circle.toArray().map((c,i) => i === curInd ? `(${c})` : `${c} `).map(c => ('    '+c).slice(-4))
  console.log(`[${player}]`, out.join(''))
}
