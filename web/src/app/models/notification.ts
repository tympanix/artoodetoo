import { Input, Output } from '../model'

export interface Notification {
  error: string
}

export class EditorError implements Notification {
  error: string
}

export class SourceWarning extends EditorError {
  input: Input
  source: Output
  error: string

  constructor(input: Input, source: Output) {
    super()
    this.input = input
    this.source = source
    this.error = `Prossible type incompatability of ${this.input.type} with ${this.source.type}`
  }
}