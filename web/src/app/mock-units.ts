import { Unit } from './unit';

export const UNITS: Unit[] = [
  { id: "example.EmailAction",
    input: [{name:"Receiver",type:"string"},{name:"Subject",type:"string"},{name:"Message",type:"string"}],
    output: []},
  { id:"example.StringConverter",
    input:[{name:"Format",type:"string"},{name:"Placeholder",type:"interface {}"}],
    output:[{name:"Formatted",type:"string"}]},
  { id:"example.PersonEvent",
    input:[],
    output:[{name:"Name",type:"string"},{name:"Age",type:"int"},{name:"Heigth",type:"float32"},{name:"Married",type:"bool"}]}
];
