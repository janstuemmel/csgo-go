import Terminal from '../../src/terminal'

describe('keyPress', () => {

  let mock, _keyPress = Terminal.prototype._keyPress

  beforeEach(() => mock = createTerminalMock()) 

  it('should write key to buffer', () => {
  
    // given
    const keyPress = _keyPress.bind(mock)
  
    // when
    keyPress({ key: 'h', domEvent: { keyCode: 72 } })
  
    // then
    expect(mock.write).toHaveBeenCalledWith('h')
  })
})


// helper
function createTerminalMock() {
  return {
    write: jest.fn(),
    curser: { x: 0, y: 0 },
  }
}