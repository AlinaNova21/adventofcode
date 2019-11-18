const input = require('./input')

const BOOST = +process.argv[3] || 0

const armies = []
let lastArmy = ''
let cnt = 1
for(const line of input) {
  if(line.endsWith(':')) {
    lastArmy = line.slice(0,-1)
    cnt = 1
    continue
  }
  const [,units] = line.match(/(\d+) units/)
  const [,hits] = line.match(/(\d+) hit points/)
  const [,dmg,dmgType] = line.match(/(\d+) (\w+) damage/)
  const [,init] = line.match(/initiative (\d+)/)
  const weakTo = (line.match(/weak to ([\w, ]+)/) || ['',''])[1].split(', ').filter(Boolean)
  const immuneTo = (line.match(/immune to ([\w, ]+)/) || ['',''])[1].split(', ').filter(Boolean)
  armies.push({
    group: cnt++,
    army: lastArmy,
    units: +units,
    hits: +hits,
    dmg: +dmg,
    dmgType,
    init: +init,
    weakTo,
    immuneTo,
    get power() {
      return this.units * this.dmg
    },
    get selSort() {
      return (this.power * 100) + this.init
    },
    get attSort() {
      return (this._dmg * 100000) + this.selSort
    },
    reset() {
      this._dmg = 0
      this._att = null
      this._def = null
    }
  })
}

const teams = {
  get immune() {
    return armies.filter(v => v.army === 'Immune System' && v.units > 0)
  },
  get infection() {
    return armies.filter(v => v.army === 'Infection' && v.units > 0)
  }
}

if(BOOST) {
  teams.immune.forEach(u => {
    u.dmg += BOOST
  })
}

while(teams.immune.length && teams.infection.length) {
  armies.sort((a,b) => a.group - b.group)
  console.log('Immune System:')
  teams.immune.forEach(u => {
    console.log(`Group ${u.group} contains ${u.units}`)
  })
  console.log('Infection:')
  teams.infection.forEach(u => {
    console.log(`Group ${u.group} contains ${u.units}`)
  })
  console.log()
  // Target phase
  armies.sort((a,b) => b.selSort - a.selSort) // descending
  armies.forEach(a => a.reset())
  for(const army of armies) {
    if(army.units <= 0) continue
    const targets = enemies(army)
    targets.forEach(t => t._dmg = calcDmg(army, t))
    targets.forEach(t => {
      console.log(`${army.army} group ${army.group} would deal defending group ${t.group} ${t._dmg} damage`)
    })
    targets.sort((a,b) => b.attSort - a.attSort)
    const tgt = targets.find(t => t._dmg && !t._def)
    if(!tgt) continue
    army._att = tgt
    tgt._def = true
  }
  console.log()
  // Attack Phase
  armies.sort((a,b) => b.init - a.init)
  for(const army of armies) {
    if(!army._att) continue
    if(army.units <= 0) continue
    const tgt = army._att
    const dmg = calcDmg(army, tgt)
    const killed = Math.floor(dmg / tgt.hits)
    tgt.units -= killed
    console.log(`${army.army} group ${army.group} attacks defending group ${tgt.group}, killing ${killed} units`)
  }
  console.log('-------------------------')
  // console.log(armies)
}
{
  console.log(`Immune System: ${teams.immune.reduce((r,u) => r + u.units, 0)}`)
  if(!teams.immune.length) {
    console.log(`No Groups Remain.`)
  }
  teams.immune.forEach(u => {
    console.log(`Group ${u.group} contains ${u.units}`)
  })
  console.log(`Infection: ${teams.infection.reduce((r,u) => r + u.units, 0)}`)
  if(!teams.infection.length) {
    console.log(`No Groups Remain.`)
  }
  teams.infection.forEach(u => {
    console.log(`Group ${u.group} contains ${u.units}`)
  })
  const count = teams.immune.reduce((a,b) => a + b.units, 0) + teams.infection.reduce((a,b) => a + b.units, 0)
  console.log(count)
}

function  enemies({ army }) {
  if(army === 'Infection') return teams.immune
  if(army === 'Immune System') return teams.infection
}


function calcDmg(att,tar) {
  if(tar.immuneTo.includes(att.dmgType)) return 0
  return att.power * (tar.weakTo.includes(att.dmgType) ? 2 : 1)
}

function int(v) { return +v }