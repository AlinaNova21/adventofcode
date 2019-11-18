const fs = require('fs')
const path = require('path')
const argv = process.argv.slice(2)
const day = path.basename(process.argv[1]).slice(0,-3)
const file = argv[0] || `input/${day}`
const input = fs.readFileSync(file, 'utf8').split("\n").filter(Boolean)

input.asNumbers = function() { return input.map(v => parseFloat(v)) }

module.exports = input