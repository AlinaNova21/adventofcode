const axios = require('axios')
const fs = require('fs')
const argv = process.argv.slice(2)

axios.get(`https://adventofcode.com/2018/day/${argv[0]}/input`, { 
  headers: {
    cookie: `session=${process.env.COOKIE}`
  }
}).then(({ data }) => {
  fs.writeFile(`input.day${argv[0]}`, data, () => {})
})