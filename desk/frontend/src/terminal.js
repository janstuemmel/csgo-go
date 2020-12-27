import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'

import LocalEchoController from 'local-echo'

import 'xterm/css/xterm.css'

const defaultOptions = {
  ps1: '> '
}

export default class {

  constructor(elem, options = {}, commands = []) {
    
    this.elem = elem
    this.options = Object.assign({}, defaultOptions, options)
    this.commands = commands
    this.ps1 = this.options.ps1

    // input listeners
    this.inputListeners = []
    
    // init term
    this.term = new Terminal(options)
    
    // fit addon
    this._fitAddon = new FitAddon()
    this.term.loadAddon(this._fitAddon);
    
    // add local echo
    this.localEcho = new LocalEchoController(this.term)
    this.localEcho.addAutocompleteHandler(index => {
      if (index === 0) return this.commands
      return []
    })

    // init
    this.term.open(elem)
    this.fit()

    // add resize listener
    window.addEventListener('resize', () => this.fit())
  }

  fit() {
    this._fitAddon.fit()
  }

  write(val) {
    this.term.write(val)
  }

  writeln(val) {
    this.term.writeln(val)
  }

  onInput(fun) {
    this.inputListeners.push(fun)
  }

  _emitInput(inp) {
    this.inputListeners.forEach(fun => fun(inp))
  }

  async read() {
    
    const inp = await this.localEcho.read(this.ps1)
    
    this._emitInput(inp)
  }
}