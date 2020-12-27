const fetch = require('node-fetch')
const fs = require('fs')

const URL = 'https://raw.githubusercontent.com/funeralchris/csgo-list-of-cvars/master/cvars_all.log'

// https://regex101.com/r/NOL90b/1/
const pattern = /^(.*) +\| (.+) +\| ([\w,\" ]+) \|(.*)/

async function get() {

  const response = await fetch(URL)

  let cvars = await response.text()

  cvars = cvars.split('\n');

  let result = []
  let errs = []
  let parsingErr = 0

  cvars.forEach((cvar, idx) => {

    if(idx == 0) return 

    const match = cvar.match(pattern)

    if (match) {
      result.push({
        command: match[1].trim(),
        default: match[2].trim(),
        groups: match[3].trim().replace(/\"/gi, ''),
        description: match[4].trim(),
      })
    } else {
      errs.push(cvar)
    }
  })

  console.log(errs)
  console.log(errs.length)

  return result
}


(async function() {
  const result = await get()

  fs.writeFileSync('./commands.json', JSON.stringify(result))

  // console.log(result) 
})()