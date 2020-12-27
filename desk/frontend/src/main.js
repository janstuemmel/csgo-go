import Terminal from './terminal'
import cvars from './cvars.json'

const elemAddress = document.getElementById('address'),
      elemPassword = document.getElementById('password')

let address = elemAddress.value, 
    password = elemPassword.value

elemAddress.addEventListener('input', e => address = e.target.value)
elemPassword.addEventListener('input', e => password = e.target.value)

// init term
const term = new Terminal(
  document.getElementById('terminal'), 
  { theme: { background: '#073642' } }, 
  cvars.map(cvar => cvar.command)
)

// term.fit()

term.read()

term.onInput(inp => {
  
  if (inp === '') {
    return term.read()
  }

  sendCommand(inp).then(resp => {

    term.localEcho.println(resp)

    term.read()

  }).catch(console.log)
})

// helper

async function sendCommand(command) {
  
  const options = {
    method: 'POST',
    body: JSON.stringify({ address, password, command })
  }

  const response = await fetch('/message', options)
  
  return await response.text()
}

