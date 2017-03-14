class Input {
  name: string;
  type: string;
}

class Output {
  name: string;
  type: string;
}

export class Unit {
  id: string;
  description: string;
  input: Input[];
  output: Output[];
}
