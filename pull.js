const axios = require('axios')
const fs = require('fs').promises
const argv = process.argv.slice(2)

async function run() {
  axios.defaults.headers.common.cookie = `session=${process.env.COOKIE}`
  const { data: input } = await axios.get(`https://adventofcode.com/2018/day/${argv[0]}/input`)
  await fs.writeFile(`input/day${argv[0]}`, input)
}

run().catch(console.error)