const charTable = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567'
const byteTable = [
  0xff, 0xff, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
  0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
  0xff, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
  0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
  0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16,
  0x17, 0x18, 0x19, 0xff, 0xff, 0xff, 0xff, 0xff,
  0xff, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
  0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
  0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16,
  0x17, 0x18, 0x19, 0xff, 0xff, 0xff, 0xff, 0xff
]

export default class {
  encode (plain) {
    if (!Buffer.isBuffer(plain) && typeof plain !== 'string') {
      throw new TypeError('base32.encode only takes Buffer or string as parameter')
    }
    if (!Buffer.isBuffer(plain)) {
      plain = Buffer.from(plain)
    }
    let i = 0
    let j = 0
    let shiftIndex = 0
    let digit = 0
    const encoded = Buffer.alloc(this._quintetCount(plain) * 8)
    while (i < plain.length) {
      const current = plain[i]
      if (shiftIndex > 3) {
        digit = current & (0xff >> shiftIndex)
        shiftIndex = (shiftIndex + 5) % 8
        digit = (digit << shiftIndex) | ((i + 1 < plain.length) ? plain[i + 1] : 0) >> (8 - shiftIndex)
        i++
      } else {
        digit = (current >> (8 - (shiftIndex + 5))) & 0x1f
        shiftIndex = (shiftIndex + 5) % 8
        if (shiftIndex === 0) {
          i++
        }
      }
      encoded[j] = charTable.charCodeAt(digit)
      j++
    }
    for (i = j; i < encoded.length; i++) {
      encoded[i] = 0x3d // '='.charCodeAt(0)
    }
    return encoded
  }

  decode (encoded) {
    if (!Buffer.isBuffer(encoded) && typeof encoded !== 'string') {
      throw new TypeError('base32.decode only takes Buffer or string as parameter')
    }
    let shiftIndex = 0
    let plainDigit = 0
    let plainChar = null
    let plainPos = 0
    if (!Buffer.isBuffer(encoded)) {
      encoded = Buffer.from(encoded)
    }
    const decoded = Buffer.alloc(Math.ceil(encoded.length * 5 / 8))
    /* byte by byte isn't as pretty as octet by octet but tests a bit
        faster. will have to revisit. */
    for (let i = 0; i < encoded.length; i++) {
      if (encoded[i] === 0x3d) {
        break
      }
      const encodedByte = encoded[i] - 0x30
      if (encodedByte < byteTable.length) {
        plainDigit = byteTable[encodedByte]
        if (shiftIndex <= 3) {
          shiftIndex = (shiftIndex + 5) % 8
          if (shiftIndex === 0) {
            plainChar |= plainDigit
            decoded[plainPos] = plainChar
            plainPos++
            plainChar = 0
          } else {
            plainChar |= 0xff & (plainDigit << (8 - shiftIndex))
          }
        } else {
          shiftIndex = (shiftIndex + 5) % 8
          plainChar |= 0xff & (plainDigit >>> shiftIndex)
          decoded[plainPos] = plainChar
          plainPos++
          plainChar = 0xff & (plainDigit << (8 - shiftIndex))
        }
      } else {
        throw new Error('Invalid input - it is not base32 encoded string')
      }
    }
    return decoded.slice(0, plainPos)
  }

  _quintetCount (buff) {
    const quintets = Math.floor(buff.length / 5)
    return buff.length % 5 === 0 ? quintets : quintets + 1
  }
}
